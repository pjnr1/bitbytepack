package bitbytepack

import (
	"bytes"
	"reflect"
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

	array = []byte{0x04, 0x02}
	mask = []byte{0x3C, 0x0F}
	want = uint(0x12)

	if got := ReadFromArray(array, mask); got != want {
		t.Errorf("ReadFromArray(%x, %x) = %x, want %x", array, mask, got, want)
	}

	// Test early return
	want = 0
	if got := ReadFromArray([]byte{}, []byte{0x0F}); got != want {
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

	if got := ReadFromArray8(array, mask); got != uint8(want) {
		t.Errorf("ReadFromArray8(%x, %x) = %x, want %x", array, mask, got, uint8(want))
	}
	if got := ReadFromArray16(array, mask); got != uint16(want) {
		t.Errorf("ReadFromArray16(%x, %x) = %x, want %x", array, mask, got, uint16(want))
	}
	if got := ReadFromArray32(array, mask); got != uint32(want) {
		t.Errorf("ReadFromArray32(%x, %x) = %x, want %x", array, mask, got, uint32(want))
	}
	if got := ReadFromArray64(array, mask); got != uint64(want) {
		t.Errorf("ReadFromArray64(%x, %x) = %x, want %x", array, mask, got, uint64(want))
	}
}

func TestWriteToArray(t *testing.T) {
	array := []byte{0x00, 0x00}
	mask := []byte{0x0F, 0x0F}
	value := uint(0x12)
	want := []byte{0x01, 0x02}

	if got, e := WriteToArray(array, mask, value); e != nil || !bytes.Equal(got, want) {
		t.Errorf("ReadFromArray(%x, %x) = %x, want %x", array, mask, got, want)
	}

	array = []byte{0x0a, 0x00}
	mask = []byte{0xF0, 0x0F}
	want = []byte{0x1a, 0x02}

	if got, e := WriteToArray(array, mask, value); e != nil || !bytes.Equal(got, want) {
		t.Errorf("WriteToArray(%x, %x) = %x, want %x", array, mask, got, want)
	}

	// Test early return
	array = []byte{0x00}
	mask = []byte{0x0F, 0x0F}
	want = []byte{}
	if _, e := WriteToArray(array, mask, value); e != ErrArrayShorterThanMask {
		t.Errorf("WriteToArray(%x, %x, %x) didn't throw '%s', but '%s'",
			array, mask, value, ErrNotEnoughBitsToEmbedValue, e)
	}

	// Test early return
	array = []byte{0x00, 0x00}
	mask = []byte{0x00, 0x0F}
	want = []byte{}
	if _, e := WriteToArray(array, mask, value); e != ErrNotEnoughBitsToEmbedValue {
		t.Errorf("WriteToArray(%x, %x, %x) didn't throw '%s', but '%s'",
			array, mask, value, ErrNotEnoughBitsToEmbedValue, e)
	}
}

func TestWriteToArrayTypeSpecifics(t *testing.T) {
	array := []byte{0x00, 0x00}
	mask := []byte{0x0F, 0x0F}
	value := uint(0x12)
	want := []byte{0x01, 0x02}

	if got, e := WriteToArray8(array, mask, uint8(value)); e != nil || !bytes.Equal(got, want) {
		t.Errorf("WriteToArray8(%x, %x, %x) = %x, want %x", array, mask, uint8(value), got, want)
	}
	if got, e := WriteToArray16(array, mask, uint16(value)); e != nil || !bytes.Equal(got, want) {
		t.Errorf("WriteToArray16(%x, %x, %x) = %x, want %x", array, mask, uint16(value), got, want)
	}
	if got, e := WriteToArray32(array, mask, uint32(value)); e != nil || !bytes.Equal(got, want) {
		t.Errorf("WriteToArray32(%x, %x, %x) = %x, want %x", array, mask, uint32(value), got, want)
	}
	if got, e := WriteToArray64(array, mask, uint64(value)); e != nil || !bytes.Equal(got, want) {
		t.Errorf("WriteToArray64(%x, %x, %x) = %x, want %x", array, mask, uint64(value), got, want)
	}
}

func TestMultReadFromArray(t *testing.T) {
	array := []byte{0x12, 0x34, 0x56, 0x78}
	masks := []MaskTypePair{
		{[]byte{0xF0, 0xF0, 0x00, 0x00}, reflect.Uint},
		{[]byte{0x0F, 0x0F, 0x00, 0x00}, reflect.Uint},
		{[]byte{0xFF, 0x00, 0xFF, 0x00}, reflect.Uint},
		{[]byte{0x00, 0xFF, 0x00, 0x0F}, reflect.Uint}}
	want := []interface{}{
		uint(0x13),
		uint(0x24),
		uint(0x1256),
		uint(0x0348)}

	if got := MultReadFromArray(array, masks...); !reflect.DeepEqual(want, got) {
		t.Errorf("MultReadFromArray(%x, %x) = %x, want %x", array, masks, got, want)

	}
}

func TestMultReadFromArrayTypeSpecifics(t *testing.T) {
	array := []byte{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0}
	masks := []MaskTypePair{
		{[]byte{0xF0, 0xF0, 0x00, 0x00}, reflect.Uint},
		{[]byte{0x0F, 0x0F, 0x00, 0x00}, reflect.Uint},
		{[]byte{0xFF, 0x00, 0xFF, 0x00}, reflect.Uint},
		{[]byte{0x00, 0xFF, 0x00, 0x0F}, reflect.Uint},
		{[]byte{0xFF, 0xFF, 0xFF, 0xFF}, reflect.Uint},
		{[]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF}, reflect.Uint}}
	want := []interface{}{
		uint(0x13),
		uint(0x24),
		uint(0x1256),
		uint(0x348),
		uint(0x12345678),
		uint(0x123456789A),
	}
	if got := MultReadFromArray(array, masks...); !reflect.DeepEqual(want, got) {
		t.Errorf("MultReadFromArray(%x, %x) = %x, want %x", array, masks, got, want)
	}

	masksB := [][]byte{
		{0xF0, 0xF0, 0x00, 0x00},
		{0x0F, 0x0F, 0x00, 0x00},
		{0xFF, 0x00, 0xFF, 0x00},
		{0x00, 0xFF, 0x00, 0x0F},
		{0xFF, 0xFF, 0xFF, 0xFF},
		{0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
		{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}}

	want8 := []uint8{
		0x13,
		0x24,
		0x56,
		0x48,
		0x78,
		0x9A,
		0xBC}
	if got := MultReadFromArray8(array, masksB...); !reflect.DeepEqual(want8, got) {
		t.Errorf("MultReadFromArray8(%x, %x) = %x, want %x", array, masks, got, want)
	}

	want16 := []uint16{
		0x13,
		0x24,
		0x1256,
		0x348,
		0x5678,
		0x789A,
		0x9ABC}
	if got := MultReadFromArray16(array, masksB...); !reflect.DeepEqual(want16, got) {
		t.Errorf("MultReadFromArray16(%x, %x) = %x, want %x", array, masks, got, want)
	}

	want32 := []uint32{
		0x13,
		0x24,
		0x1256,
		0x348,
		0x12345678,
		0x3456789A,
		0x56789ABC}
	if got := MultReadFromArray32(array, masksB...); !reflect.DeepEqual(want32, got) {
		t.Errorf("MultReadFromArray32(%x, %x) = %x, want %x", array, masks, got, want)
	}

	want64 := []uint64{
		0x13,
		0x24,
		0x1256,
		0x348,
		0x12345678,
		0x123456789A,
		0x123456789ABC}
	if got := MultReadFromArray64(array, masksB...); !reflect.DeepEqual(want64, got) {
		t.Errorf("MultReadFromArray64(%x, %x) = %x, want %x", array, masks, got, want)
	}
}

func TestMultWriteToArray(t *testing.T) {
	// Unsigned values
	array := []byte{0x00, 0x00, 0x00, 0x00, 0x00}
	maskValuePairs := []interface{}{
		MaskValuePair{[]byte{0x00, 0x00, 0x00, 0x00, 0xFF}, 0x9A},
		MaskValuePair8{[]byte{0x00, 0x00, 0x00, 0xFF}, 0x78},
		MaskValuePair16{[]byte{0x00, 0x0F, 0x0F, 0x00}, 0x46},
		MaskValuePair32{[]byte{0xF0, 0x00, 0xF0, 0x00}, 0x15},
		MaskValuePair64{[]byte{0x0F, 0xF0, 0x00, 0x00}, 0x23},
	}
	want := []byte{0x12, 0x34, 0x56, 0x78, 0x9A}
	if got, e := MultWriteToArray(array, maskValuePairs...); e != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("MultWriteToArray(%x, %x) = %x, want %x", array, maskValuePairs, got, want)
	}

	// Signed values
	array = make([]byte, 24)
	maskValuePairs = []interface{}{
		MaskValuePairS{[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}, -1},
		MaskValuePair8S{[]byte{0x00, 0x00, 0x00, 0xFF}, -43},
		MaskValuePair16S{[]byte{0x00, 0xFF, 0xFF, 0x00}, -1345},
		MaskValuePair32S{[]byte{0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF, 0xFF, 0xFF}, -705422},
		MaskValuePair64S{[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}, -12134142},
	}
	want = []byte{0x00, 0xFA, 0xBF, 0xD5, 0xFF, 0xF5, 0x3C, 0x72, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x46, 0xD9, 0x02}
	if got, e := MultWriteToArray(array, maskValuePairs...); e != nil || !reflect.DeepEqual(got, want) {
		if e != nil {
			t.Errorf("Got the error: %s", e)
		}
		t.Errorf("MultWriteToArray(%x, %x) = %x, want %x", array, maskValuePairs, got, want)
	}

	// Float values
	array = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	maskValuePairs = []interface{}{
		MaskValuePair32F{[]byte{0xF0, 0x0F, 0xF0, 0x0F, 0xF0, 0x0F, 0xF0, 0x0F, 0x00, 0x00}, 1.00},
		MaskValuePair32F{[]byte{0x0F, 0xF0, 0x0F, 0xF0, 0x0F, 0xF0, 0x0F, 0xF0, 0x00, 0x00}, 1.50},
	}
	want = []byte{0x33, 0xFF, 0x8C, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	if got, e := MultWriteToArray(array, maskValuePairs...); e != nil || !reflect.DeepEqual(got, want) {
		if e != nil {
			t.Errorf("Got the error: %s", e)
		}
		t.Errorf("MultWriteToArray(%x, %x) = %x, want %x", array, maskValuePairs, got, want)
	}

	wantErr := ErrNotEnoughBitsToEmbedValue
	mask := make([]byte, 2)
	value := 125.0
	if _, got := WriteToArray32F(array, mask, float32(value)); got != wantErr {
		t.Errorf("WriteToArray32F(%x, %x, flaot32(%f)) => %x, want %x", array, mask, float32(value), got, wantErr)
	}
	if _, got := WriteToArray64F(array, mask, value); got != wantErr {
		t.Errorf("WriteToArray32F(%x, %x, %f) => %x, want %x", array, mask, value, got, wantErr)
	}

	array = make([]byte, 12)
	maskValuePairs = []interface{}{
		MaskValuePair64F{[]byte{0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}, 0.25},
		MaskValuePair32F{[]byte{0xFF, 0xFF, 0xFF, 0xFF}, 0.50},
	}
	want = []byte{0x3f, 0x00, 0x00, 0x00, 0x3f, 0xd0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	if got, e := MultWriteToArray(array, maskValuePairs...); e != nil || !reflect.DeepEqual(got, want) {
		if e != nil {
			t.Errorf("Got the error: %s", e)
		}
		t.Errorf("MultWriteToArray(%x, %x) = %x, want %x", array, maskValuePairs, got, want)
	}

}

func BenchmarkReadFromArray(b *testing.B) {
	array := []byte{0x81, 0x09, 0x04, 0x4A, 0x00, 0x00, 0x05, 0x01, 0xFF}
	mask := []byte{0x00, 0x00, 0x00, 0x00, 0x0F, 0x0F, 0x0F, 0x0F, 0x00}

	for i := 0; i < b.N; i++ {
		ReadFromArray(array, mask)
	}
}

func BenchmarkWriteToArray(b *testing.B) {
	array := []byte{0x81, 0x09, 0x04, 0x4A, 0x00, 0x00, 0x00, 0x00, 0xFF}
	mask := []byte{0x00, 0x00, 0x00, 0x00, 0x0F, 0x0F, 0x0F, 0x0F, 0x00}
	value := uint(0x1234)

	for i := 0; i < b.N; i++ {
		WriteToArray(array, mask, value)
	}
}
