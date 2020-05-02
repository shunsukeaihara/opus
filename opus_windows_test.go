// +build windows

package opus

import (
	"math"
	"testing"
)

func TestCreateDecoder(t *testing.T) {
	_, err := NewDecoder(48000, 2)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateEncoder(t *testing.T) {
	enc, err := NewEncoder(48000, 2, ApplicationVoip)
	if err != nil {
		t.Fatal(err)
	}

	err = enc.SetMaxBandwidth(BandwidthWideband)
	if err != nil {
		t.Fatal(err)
	}

	b, err := enc.MaxBandwidth()
	if err != nil {
		t.Fatal(err)
	}
	if b != BandwidthWideband {
		t.Fatal("max bandwidth did not matche")
	}
}

func absDiff(a, b []int16, t *testing.T) (diff float64) {
	for i := 0; i < len(a); i++ {
		d := math.Abs(float64(a[i] - b[i]))
		diff += d
	}
	return diff / float64(len(a))
}

func TestEncodeDecode(t *testing.T) {
	sampleRate := 16000
	samples := 640 * 20
	frequency := 440.0
	volume := 5000
	var phase float64
	pcm := make([]int16, samples)
	frequencyRadiansPerSample := frequency * 2 * math.Pi / float64(sampleRate)

	for sample := 0; sample < samples; sample++ {
		phase += frequencyRadiansPerSample
		sampleValue := float64(volume) * math.Sin(phase)
		pcm[sample] = int16(sampleValue)
	}

	enc, err := NewEncoder(sampleRate, 1, ApplicationVoip)
	if err != nil {
		t.Fatal(err)
	}
	err = enc.SetMaxBandwidth(BandwidthWideband)
	if err != nil {
		t.Fatal(err)
	}

	dec, err := NewDecoder(sampleRate, 1)
	if err != nil {
		t.Fatal(err)
	}

	out := make([]int16, 640)
	buf := make([]byte, 640)
	for i := 0; i < samples; i += 640 {
		n, err := enc.Encode(pcm[i:i+640], buf)
		if err != nil {
			t.Fatal(err)
		}
		n, err = dec.Decode(buf[:n], out)
		if err != nil {
			t.Fatal(err)
		}
		if n != 640 {
			t.Fatal("decoding error")
		}
		//if diff := absDiff(out, frame, t); diff > 0.1 {
		//	t.Fatal("encoding or decoding error")
		//}
	}
}
