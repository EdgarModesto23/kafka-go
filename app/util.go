package main

import "errors"

func GetSize() {}

func ParseUnsignedVarint(data []byte) (uint64, int, error) {
	var result uint64
	var shift uint
	var i int

	// Iterate over the bytes of the data
	for i = 0; i < len(data); i++ {
		// Get the current byte
		b := data[i]

		// Add the lower 7 bits to the result
		result |= uint64(b&0x7F) << shift
		shift += 7

		// If the continuation bit is not set, we've reached the last byte
		if b&0x80 == 0 {
			return result, i + 1, nil
		}
	}

	// If we exit the loop, it means we did not find the end of the varint
	return 0, 0, errors.New("varint is too long or malformed")
}

func Int8ToBigEndianBytes(value int8) []byte {
	return []byte{byte(value)} 
}
