package iota_test

import "testing"

const (
	_ = 256 << (10 * iota)
	KB
	MB
	GB
)

func TestKBMBGB(t *testing.T) {
	if KB != 1024*256 {
		t.Errorf(`KB = %#v, want %#v`, KB, 1024*256)
	}

	if GB != 1024*1024*1024*256 {
		t.Errorf(`GB = %#v, want %#v`, GB, 1024*1024*1024*256)
	}
}
