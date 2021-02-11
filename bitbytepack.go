// Package bitbytepacket provides functions to embed values across multiple
// bytes in a byte array.
package bitbytepack

import (
	"errors"
	"math"
	"math/bits"
	"reflect"
)

// Errors
var (
	ErrNotEnoughBitsToEmbedValue = errors.New("not enough values to embed value")
	ErrArrayShorterThanMask      = errors.New("array is shorter than the mask")
	ErrInterfaceTypeNotSupported = errors.New("deducted interface type is not supported")
)

// Constants
const (
	MaxNumberOfValuesToRead = 128
)

// Struct type to contain both a mask array and value
type MaskValuePair struct {
	Mask  []byte // mask array
	Value uint   // value to be embedded
}

// Struct type to contain both a mask array and value
type MaskValuePair8 struct {
	Mask  []byte // mask array
	Value uint8  // value to be embedded
}

// Struct type to contain both a mask array and value
type MaskValuePair16 struct {
	Mask  []byte // mask array
	Value uint16 // value to be embedded
}

// Struct type to contain both a mask array and value
type MaskValuePair32 struct {
	Mask  []byte // mask array
	Value uint32 // value to be embedded
}

// Struct type to contain both a mask array and value
type MaskValuePair64 struct {
	Mask  []byte // mask array
	Value uint64 // value to be embedded
}

// Struct type to contain both a mask array and value
type MaskValuePairS struct {
	Mask  []byte // mask array
	Value int    // value to be embedded
}

// Struct type to contain both a mask array and value
type MaskValuePair8S struct {
	Mask  []byte // mask array
	Value int8   // value to be embedded
}

// Struct type to contain both a mask array and value
type MaskValuePair16S struct {
	Mask  []byte // mask array
	Value int16  // value to be embedded
}

// Struct type to contain both a mask array and value
type MaskValuePair32S struct {
	Mask  []byte // mask array
	Value int32  // value to be embedded
}

// Struct type to contain both a mask array and value
type MaskValuePair64S struct {
	Mask  []byte // mask array
	Value int64  // value to be embedded
}

// Struct type to contain both a mask array and value
type MaskValuePair32F struct {
	Mask  []byte  // mask array
	Value float32 // float32 value to be embedded
}

// Struct type to contain both a mask array and value
type MaskValuePair64F struct {
	Mask  []byte  // mask array
	Value float64 // float64 value to be embedded
}

// Struct type to contain both a mask array and the value type to read
type MaskTypePair struct {
	Mask []byte       // mask array
	Type reflect.Kind // type to read out
}

// Accumulative count ones in every byte of an []byte
func CountOnes(array []byte) int {
	count := 0
	for _, m := range array {
		count += bits.OnesCount8(m)
	}
	return count
}

// Base function for reading a value of an array
func ReadFromArray(array []byte, mask []byte) uint {
	if len(array) < len(mask) {
		return 0
	}

	var finalValue uint = 0
	var b = 0

	for i, m := range mask {

		// Extract byte with mask
		var maskedValue = uint(array[i] & m)

		// Shift all the way to the left
		maskedValue <<= (bits.UintSize - 8) + bits.LeadingZeros8(m)

		// Shift to fit end of the right
		maskedValue >>= b

		// Add to final value
		finalValue += maskedValue

		// Update number of bytes used
		b += bits.OnesCount8(m)
	}

	// Shift all the way to the left
	finalValue >>= bits.UintSize - b

	return finalValue
}

// Base function for writing an unsigned integer value
func WriteToArray(array []byte, mask []byte, value uint) ([]byte, error) {
	if len(array) < len(mask) {
		return []byte{}, ErrArrayShorterThanMask
	}

	if CountOnes(mask) < bits.Len(value) {
		return array, ErrNotEnoughBitsToEmbedValue
	}

	for i := range mask {

		// Reverse iteration
		j := len(mask) - i - 1

		// Shift value if mask is shifted
		valueByte := byte(value) << bits.TrailingZeros8(mask[j])

		// Apply mask
		valueByte &= mask[j]

		// Add to array
		array[j] |= valueByte

		// Shift value by the number of bits that was written to the array
		value >>= bits.OnesCount8(mask[j])
	}
	return array, nil
}

// Overload for uint8
func ReadFromArray8(array []byte, mask []byte) uint8 {
	return uint8(ReadFromArray(array, mask))
}

// Overload for uint16
func ReadFromArray16(array []byte, mask []byte) uint16 {
	return uint16(ReadFromArray(array, mask))
}

// Overload for uint32
func ReadFromArray32(array []byte, mask []byte) uint32 {
	return uint32(ReadFromArray(array, mask))
}

// Overload for uint64
func ReadFromArray64(array []byte, mask []byte) uint64 {
	return uint64(ReadFromArray(array, mask))
}

// Overload for int
func ReadFromArrayS(array []byte, mask []byte) int {
	return int(ReadFromArray(array, mask))
}
// Overload for uint8
func ReadFromArray8S(array []byte, mask []byte) int8 {
	return int8(ReadFromArray(array, mask))
}

// Overload for uint16
func ReadFromArray16S(array []byte, mask []byte) int16 {
	return int16(ReadFromArray(array, mask))
}
// Overload for uint32
func ReadFromArray32S(array []byte, mask []byte) int32 {
	return int32(ReadFromArray(array, mask))
}
// Overload for uint64
func ReadFromArray64S(array []byte, mask []byte) int64 {
	return int64(ReadFromArray(array, mask))
}

// Overload to read embedded float32 values of []byte array
func ReadFromArray32F(array []byte, mask []byte) float32 {
	return math.Float32frombits(ReadFromArray32(array, mask))
}

// Overload to read embedded float64 values of []byte array
func ReadFromArray64F(array []byte, mask []byte) float64 {
	return math.Float64frombits(ReadFromArray64(array, mask))
}

func WriteToArrayS(array []byte, mask []byte, value int) ([]byte, error) {
	return WriteToArray(array, mask, uint(value))
}
func WriteToArray8S(array []byte, mask []byte, value int8) ([]byte, error) {
	// Using double conversion to ensure correct representation
	return WriteToArray(array, mask, uint(uint8(value)))
}
func WriteToArray16S(array []byte, mask []byte, value int16) ([]byte, error) {
	// Using double conversion to ensure correct representation
	return WriteToArray(array, mask, uint(uint16(value)))
}
func WriteToArray32S(array []byte, mask []byte, value int32) ([]byte, error) {
	// Using double conversion to ensure correct representation
	return WriteToArray(array, mask, uint(uint32(value)))
}
func WriteToArray64S(array []byte, mask []byte, value int64) ([]byte, error) {
	// Using double conversion to ensure correct representation
	return WriteToArray(array, mask, uint(uint64(value)))
}

func WriteToArray8(array []byte, mask []byte, value uint8) ([]byte, error) {
	return WriteToArray(array, mask, uint(value))
}
func WriteToArray16(array []byte, mask []byte, value uint16) ([]byte, error) {
	return WriteToArray(array, mask, uint(value))
}
func WriteToArray32(array []byte, mask []byte, value uint32) ([]byte, error) {
	return WriteToArray(array, mask, uint(value))
}
func WriteToArray64(array []byte, mask []byte, value uint64) ([]byte, error) {
	return WriteToArray(array, mask, uint(value))
}

// "Overload" to embed float32 value in a []byte array
func WriteToArray32F(array []byte, mask []byte, value float32) ([]byte, error) {
	if CountOnes(mask) < 32 {
		return array, ErrNotEnoughBitsToEmbedValue
	}
	return WriteToArray32(array, mask, math.Float32bits(value))
}

// "Overload" to embed float64 value in a []byte array
func WriteToArray64F(array []byte, mask []byte, value float64) ([]byte, error) {
	if CountOnes(mask) < 64 {
		return array, ErrNotEnoughBitsToEmbedValue
	}
	return WriteToArray64(array, mask, math.Float64bits(value))
}

// Read multiple values from array using an array of masks
func MultReadFromArray(array []byte, mask ...MaskTypePair) []interface{} {
	output := make([]interface{}, 0, MaxNumberOfValuesToRead)

	for _, m := range mask {
		switch m.Type {
		case reflect.Uint:
			output = append(output, ReadFromArray(array, m.Mask))
		case reflect.Uint8:
			output = append(output, ReadFromArray8(array, m.Mask))
		case reflect.Uint16:
			output = append(output, ReadFromArray16(array, m.Mask))
		case reflect.Uint32:
			output = append(output, ReadFromArray32(array, m.Mask))
		case reflect.Uint64:
			output = append(output, ReadFromArray64(array, m.Mask))
		case reflect.Int:
			output = append(output, ReadFromArrayS(array, m.Mask))
		case reflect.Int8:
			output = append(output, ReadFromArray8S(array, m.Mask))
		case reflect.Int16:
			output = append(output, ReadFromArray16S(array, m.Mask))
		case reflect.Int32:
			output = append(output, ReadFromArray32S(array, m.Mask))
		case reflect.Int64:
			output = append(output, ReadFromArray64S(array, m.Mask))
		case reflect.Float32:
			output = append(output, ReadFromArray32F(array, m.Mask))
		case reflect.Float64:
			output = append(output, ReadFromArray64F(array, m.Mask))
		}
	}

	return output
}

// Read multiple values (uint8) from array using an array of masks
func MultReadFromArray8(array []byte, mask ...[]byte) []uint8 {
	output := make([]uint8, 0, MaxNumberOfValuesToRead)

	for _, m := range mask {
		output = append(output, ReadFromArray8(array, m))
	}

	return output
}

// Read multiple values (uint16) from array using an array of masks
func MultReadFromArray16(array []byte, mask ...[]byte) []uint16 {
	output := make([]uint16, 0, MaxNumberOfValuesToRead)

	for _, m := range mask {
		output = append(output, ReadFromArray16(array, m))
	}

	return output
}

// Read multiple values (uint32) from array using an array of masks
func MultReadFromArray32(array []byte, mask ...[]byte) []uint32 {
	output := make([]uint32, 0, MaxNumberOfValuesToRead)

	for _, m := range mask {
		output = append(output, ReadFromArray32(array, m))
	}

	return output
}

// Read multiple values (uint64) from array using an array of masks
func MultReadFromArray64(array []byte, mask ...[]byte) []uint64 {
	output := make([]uint64, 0, MaxNumberOfValuesToRead)

	for _, m := range mask {
		output = append(output, ReadFromArray64(array, m))
	}

	return output
}

// Read multiple float32 values from array using an array of masks
func MultReadFromArray32F(array []byte, mask ...[]byte) []float32 {
	output := make([]float32, 0, MaxNumberOfValuesToRead)

	for _, m := range mask {
		output = append(output, ReadFromArray32F(array, m))
	}

	return output
}

// Read multiple float64 values from array using an array of masks
func MultReadFromArray64F(array []byte, mask ...[]byte) []float64 {
	output := make([]float64, 0, MaxNumberOfValuesToRead)

	for _, m := range mask {
		output = append(output, ReadFromArray64F(array, m))
	}

	return output
}

func MultWriteToArray(array []byte, mvp ...interface{}) ([]byte, error) {
	var err error = nil
	// Iterate over all Mask-Value pairs
	for _, m := range mvp {
		switch m.(type) {
		case MaskValuePair:
			array, err = WriteToArray(array, m.(MaskValuePair).Mask, m.(MaskValuePair).Value)
		case MaskValuePair8:
			array, err = WriteToArray8(array, m.(MaskValuePair8).Mask, m.(MaskValuePair8).Value)
		case MaskValuePair16:
			array, err = WriteToArray16(array, m.(MaskValuePair16).Mask, m.(MaskValuePair16).Value)
		case MaskValuePair32:
			array, err = WriteToArray32(array, m.(MaskValuePair32).Mask, m.(MaskValuePair32).Value)
		case MaskValuePair64:
			array, err = WriteToArray64(array, m.(MaskValuePair64).Mask, m.(MaskValuePair64).Value)
		case MaskValuePairS:
			array, err = WriteToArrayS(array, m.(MaskValuePairS).Mask, m.(MaskValuePairS).Value)
		case MaskValuePair8S:
			array, err = WriteToArray8S(array, m.(MaskValuePair8S).Mask, m.(MaskValuePair8S).Value)
		case MaskValuePair16S:
			array, err = WriteToArray16S(array, m.(MaskValuePair16S).Mask, m.(MaskValuePair16S).Value)
		case MaskValuePair32S:
			array, err = WriteToArray32S(array, m.(MaskValuePair32S).Mask, m.(MaskValuePair32S).Value)
		case MaskValuePair64S:
			array, err = WriteToArray64S(array, m.(MaskValuePair64S).Mask, m.(MaskValuePair64S).Value)
		case MaskValuePair32F:
			array, err = WriteToArray32F(array, m.(MaskValuePair32F).Mask, m.(MaskValuePair32F).Value)
		case MaskValuePair64F:
			array, err = WriteToArray64F(array, m.(MaskValuePair64F).Mask, m.(MaskValuePair64F).Value)
		default:
			return array, ErrInterfaceTypeNotSupported
		}

		if err != nil {
			return array, err
		}
	}

	return array, err
}
