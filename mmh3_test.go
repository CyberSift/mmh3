package mmh3

import (
	"fmt"
	"testing"
)


// original tests should still succeed with seeds uninitialized or set to 0 - adding new tests to func TestSeed

func TestSeed(t *testing.T) {

	var seed_32_1, seed_32_2 uint32
	var seed_64_1, seed_64_2 uint64

	seed_32_1 = 2
	seed_32_2 = 128

	seed_64_1 = 2
	seed_64_2 = 128

	s := []byte("hello")
	if Hash32(s, seed_32_1) == Hash32(s, seed_32_2) {
		t.Fatal("32bit hello")
	}
	if fmt.Sprintf("%x", Hash128(s, seed_64_1)) == fmt.Sprintf("%x", Hash128(s, seed_64_2)) {
		t.Fatal("128bit hello ", Hash128(s, seed_64_1), Hash128(s, seed_64_2))
	}
	s = []byte("Winter is coming")
	if Hash32(s, seed_32_1) == Hash32(s, seed_32_2) {
		t.Fatal("32bit winter")
	}
	if fmt.Sprintf("%x", Hash128(s, seed_64_1)) == fmt.Sprintf("%x", Hash128(s, seed_64_2)) {
		t.Fatal("128bit winter", Hash128(s, seed_64_1), Hash128(s, seed_64_2))
	}

}

func TestAll(t *testing.T) {

	var seed_32 uint32
	var seed_64 uint64

	s := []byte("hello")
	if Hash32(s, seed_32) != 0x248bfa47 {
		t.Fatal("32bit hello")
	}
	if fmt.Sprintf("%x", Hash128(s, seed_64)) != "029bbd41b3a7d8cb191dae486a901e5b" {
		t.Fatal("128bit hello")
	}
	s = []byte("Winter is coming")
	if Hash32(s, seed_32) != 0x43617e8f {
		t.Fatal("32bit winter")
	}
	if fmt.Sprintf("%x", Hash128(s, seed_64)) != "95eddc615d3b376c13fb0b0cead849c5" {
		t.Fatal("128bit winter")
	}
}

// Test the x64 optimized hash against Hash128
func Testx64(t *testing.T) {
	keys := []string{
		"hello",
		"Winter is coming",
	}

	var seed_64 uint64
	seed_64 = 0

	for _, k := range keys {
		h128 := Hash128([]byte(k),seed_64)
		h128x64 := Hash128x64([]byte(k), seed_64)

		if string(h128) != string(h128x64) {
			t.Fatalf("Expected same hashes for %s, but got %x and %x", k, h128, h128x64)
		}
	}

}

func TestHashWriter128(t *testing.T) {
	s := []byte("hello")
	h := HashWriter128{}
	h.Write(s)
	res := h.Sum(nil)
	if fmt.Sprintf("%x", res) != "029bbd41b3a7d8cb191dae486a901e5b" {
		t.Fatal("128bit hello")
	}
	s = []byte("Winter is coming")
	h.Reset()
	h.Write(s)
	res = h.Sum(nil)
	if fmt.Sprintf("%x", res) != "95eddc615d3b376c13fb0b0cead849c5" {
		t.Fatal("128bit hello")
	}
	str := "Winter is coming"
	h.Reset()
	h.WriteString(str)
	res = h.Sum(nil)
	if fmt.Sprintf("%x", res) != "95eddc615d3b376c13fb0b0cead849c5" {
		t.Fatal("128bit hello")
	}

}
