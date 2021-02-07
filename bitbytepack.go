package bitbytepack

import (
	"errors"
	"math"
	"math/bits"
)

var (
	ErrNotEnoughBitsToEmbedValue = errors.New("the mask array doesn't contain enough bits to embed the value object")
)

const (
	MaxNumberOfValuesToRead = 128
)

type MaskValuePair struct {
	mask  []byte
	value uint
}

func CountOnes(array []byte) int {
	count := 0
	for _, m := range array {
		count += bits.OnesCount8(m)
	}
	return count
}

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

// "Overload" to read embedded float32 values of []byte array
func ReadFromArray32F(array []byte, mask []byte) float32 {
	return math.Float32frombits(ReadFromArray32(array, mask))
}

// "Overload" to read embedded float64 values of []byte array
func ReadFromArray64F(array []byte, mask []byte) float64 {
	return math.Float64frombits(ReadFromArray64(array, mask))
}

func WriteToArray(array []byte, mask []byte, value uint) []byte {
	if len(array) < len(mask) {
		return []byte{}
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
	return array
}

func WriteToArray8(array []byte, mask []byte, value uint8) ([]byte, error) {
	if CountOnes(mask) < bits.Len8(value) {
		return array, ErrNotEnoughBitsToEmbedValue
	}
	return WriteToArray(array, mask, uint(value)), nil
}
func WriteToArray16(array []byte, mask []byte, value uint16) ([]byte, error) {
	if CountOnes(mask) < bits.Len16(value) {
		return array, ErrNotEnoughBitsToEmbedValue
	}
	return WriteToArray(array, mask, uint(value)), nil
}
func WriteToArray32(array []byte, mask []byte, value uint32) ([]byte, error) {
	if CountOnes(mask) < bits.Len32(value) {
		return array, ErrNotEnoughBitsToEmbedValue
	}
	return WriteToArray(array, mask, uint(value)), nil
}
func WriteToArray64(array []byte, mask []byte, value uint64) ([]byte, error) {
	if CountOnes(mask) < bits.Len64(value) {
		return array, ErrNotEnoughBitsToEmbedValue
	}
	return WriteToArray(array, mask, uint(value)), nil
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
func MultReadFromArray(array []byte, mask ...[]byte) []uint {
	output := make([]uint, 0, MaxNumberOfValuesToRead)

	for _, m := range mask {
		output = append(output, ReadFromArray(array, m))
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

func MultWriteToArray(array []byte, mvp ...MaskValuePair) []byte {
	// Iterate over all Mask-Value pairs
	for _, m := range mvp {
		array = WriteToArray(array, m.mask, m.value)
	}
	return array
}
