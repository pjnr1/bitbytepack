package bitbytepack

import (
	"bytes"
	"testing"
)

func TestReadFromArray(t *testing.T) {
	array := []byte{0x01, 0x02}
	mask := []byte{0x0F, 0x0F}
	want := uint(0x12)
	if got := ReadFromArray(array, mask); got != want {
		t.Errorf("ReadFromArray(%x, %x) = %x, want %x", array, mask, got, want)
	}
	array = []byte{0x10, 0x02}
	mask = []byte{0xF0, 0x0F}
	if got := ReadFromArray(array, mask); got != want {
		t.Errorf("ReadFromArray(%x, %x) = %x, want %x", array, mask, got, want)
	}
}

func TestReadFromArrayTypeSpecifics(t *testing.T) {
	array := []byte{0x01, 0x02}
	mask := []byte{0x0F, 0x0F}
	want := uint(0x12)
	if got := ReadFromArray(array, mask); got != want {
		t.Errorf("ReadFromArray(%x, %x) = %x, want %x", array, mask, got, want)
	}

	if got := ReadFromArray64(array, mask); got != uint64(want) {
		t.Errorf("ReadFromArray64(%x, %x) = %x, want %x", array, mask, got, uint64(want))
	}
	if got := ReadFromArray32(array, mask); got != uint32(want) {
		t.Errorf("ReadFromArray64(%x, %x) = %x, want %x", array, mask, got, uint32(want))
	}
	if got := ReadFromArray16(array, mask); got != uint16(want) {
		t.Errorf("ReadFromArray64(%x, %x) = %x, want %x", array, mask, got, uint16(want))
	}
	if got := ReadFromArray8(array, mask); got != uint8(want) {
		t.Errorf("ReadFromArray64(%x, %x) = %x, want %x", array, mask, got, uint8(want))
	}
}

func TestWriteToArray(t *testing.T) {
	array := []byte{0x00, 0x00}
	mask := []byte{0x0F, 0x0F}
	value := uint(0x12)
	want := []byte{0x01, 0x02}
	if got := WriteToArray(array, mask, value); !bytes.Equal(got, want) {
		t.Errorf("ReadFromArray(%x, %x) = %x, want %x", array, mask, got, want)
	}
	mask = []byte{0xF0, 0x0F}
	want = []byte{0x10, 0x02}
	if got := WriteToArray(array, mask, value); !bytes.Equal(got, want) {
		t.Errorf("ReadFromArray(%x, %x) = %x, want %x", array, mask, got, want)
	}
}
