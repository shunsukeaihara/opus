// +build windows

package opus

import (
	"unsafe"
)

// Decoder is Decoder interface for opus.dll
type Decoder struct {
	mem        []byte
	p          unsafe.Pointer
	sampleRate int
	channels   int
}

// NewDecoder creates Opus Decoder
func NewDecoder(sampleRate int, channels int) (*Decoder, error) {
	size := opus_decoder_get_size(channels)
	mem := make([]byte, size)
	p := unsafe.Pointer(&mem[0])
	n := opus_decoder_init(p, int32(sampleRate), channels)
	if n != 0 {
		return nil, Error(n)
	}
	return &Decoder{mem, p, sampleRate, channels}, nil
}

// Decode decodes opus encoded byte array to int16 array
func (d *Decoder) Decode(in []byte, out []int16) (int, error) {
	n := opus_decode(d.p, in, out, d.channels)
	if n < 0 {
		return 0, Error(n)
	}
	return n, nil
}

// DecodeFloat32 decodes opus encoded byte array to float32 array
func (d *Decoder) DecodeFloat32(in []byte, out []float32) (int, error) {
	n := opus_decode_float(d.p, in, out, d.channels)
	if n < 0 {
		return 0, Error(n)
	}
	return n, nil
}
