// +build windows

package opus

import (
	"fmt"
	"unsafe"
)

// BandWidth is type of Opus BandWidth
type BandWidth int

// Application is type of opus application
type Application int

const (
	setBitrateRequest      = 4002
	getBitrateRequest      = 4003
	setMaxBandwidthRequest = 4004
	getMaxBandWidthRequest = 4005
	setBandwidthRequest    = 4008
	getBandWidthRequest    = 4009

	// ApplicationVoip is ...
	ApplicationVoip Application = 2048
	// ApplicationAudio is ...
	ApplicationAudio Application = 2049
	// ApplcationRestrictedLowdelay is ...
	ApplcationRestrictedLowdelay Application = 2050
	// BandwidthNarrowband is set bandwidth to narrowband
	BandwidthNarrowband BandWidth = 1101
	// BandwidthMediumband is set bandwidth to mediumband
	BandwidthMediumband BandWidth = 1102
	// BandwidthWideband is set bandwidth to wideband
	BandwidthWideband BandWidth = 1103
	// BandwidthSuperwideband is set bandwidth to superwideband
	BandwidthSuperwideband BandWidth = 1104
	// BandwidthFullband is set bandwidth to fullband
	BandwidthFullband BandWidth = 1105
)

// Encoder is Encoder interface for opus.dll
type Encoder struct {
	mem        []byte
	p          unsafe.Pointer
	sampleRate int
	channels   int
}

// NewEncoder creates Opus Encoder
func NewEncoder(sampleRate int, channels int, application Application) (*Encoder, error) {
	size := opus_encoder_get_size(channels)
	mem := make([]byte, size)
	p := unsafe.Pointer(&mem[0])
	n := opus_encoder_init(p, int32(sampleRate), channels, int(application))
	if n != 0 {
		return nil, Error(n)
	}
	return &Encoder{mem, p, sampleRate, channels}, nil
}

// SetBandwidth sets bandwidth to encoder
func (e *Encoder) SetBandwidth(bandwidth BandWidth) error {
	n := opus_encoder_ctl(e.p, setBandwidthRequest, int32(bandwidth))
	if n != 0 {
		return Error(n)
	}
	return nil
}

// SetMaxBandwidth sets max bandwidth
func (e *Encoder) SetMaxBandwidth(bandwidth BandWidth) error {
	n := opus_encoder_ctl(e.p, setMaxBandwidthRequest, int32(bandwidth))
	if n != 0 {
		return Error(n)
	}
	return nil
}

// MaxBandwidth gets max bandwidth
func (e *Encoder) MaxBandwidth() (BandWidth, error) {
	var out int32
	res := opus_encoder_get_ctl(e.p, getMaxBandWidthRequest, &out)
	if res != 0 {
		fmt.Println(res)
		return 0, Error(res)
	}
	return BandWidth(out), nil
}

// SetBitrate sets bitrate to encoder
func (e *Encoder) SetBitrate(bitrate int) error {
	res := opus_encoder_ctl(e.p, setBitrateRequest, int32(bitrate))
	if res != 0 {
		return Error(res)
	}
	return nil
}

// Bitrate returns bitrate of encoder
func (e *Encoder) Bitrate() (int, error) {
	var out int32
	res := opus_encoder_get_ctl(e.p, getBitrateRequest, &out)
	if res != 0 {
		return 0, Error(res)
	}
	return int(res), nil
}

// Encode encodes int16 pcm to encoded byte array
func (e *Encoder) Encode(in []int16, out []byte) (int, error) {
	n := opus_encode(e.p, in, out, e.channels)
	if n < 0 {
		return 0, Error(n)
	}
	return n, nil
}

// EncodeFloat32 encodes float32 pcm to encoded byte array
func (e *Encoder) EncodeFloat32(in []float32, out []byte) (int, error) {
	n := opus_encode_float(e.p, in, out, e.channels)
	if n < 0 {
		return 0, Error(n)
	}
	return n, nil
}
