package bitbytepack

import (
	"math/bits"
)

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

func WriteToArray(array []byte, mask []byte, value uint) []byte {
	if len(array) < len(mask) {
		return []byte{}
	}

	for i := range mask {

		// Reverse iteration
		j := len(mask) - i - 1

		// Shift value if mask is shifted
		array[j] = byte(value) << bits.TrailingZeros8(mask[j])

		// Apply mask
		array[j] &= mask[j]

		// Shift value by the number of bits that was written to the array
		value >>= bits.OnesCount8(mask[j])
	}
	return array
}

func ReadFromArray8(v []byte, mask []byte) uint8 {
	return uint8(ReadFromArray(v, mask))
}
func ReadFromArray16(v []byte, mask []byte) uint16 {
	return uint16(ReadFromArray(v, mask))
}
func ReadFromArray32(v []byte, mask []byte) uint32 {
	return uint32(ReadFromArray(v, mask))
}
func ReadFromArray64(v []byte, mask []byte) uint64 {
	return uint64(ReadFromArray(v, mask))
}

func WriteToArray8(array []byte, mask []byte, value uint8) []byte {
	return WriteToArray(array, mask, uint(value))
}
func WriteToArray16(array []byte, mask []byte, value uint16) []byte {
	return WriteToArray(array, mask, uint(value))
}
func WriteToArray32(array []byte, mask []byte, value uint32) []byte {
	return WriteToArray(array, mask, uint(value))
}
func WriteToArray64(array []byte, mask []byte, value uint64) []byte {
	return WriteToArray(array, mask, uint(value))
}
