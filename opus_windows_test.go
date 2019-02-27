// +build windows

package opus

import (
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
