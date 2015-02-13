/*
    Package bitops provide a set of bit operations for uint32/uint64. They are similar to 
    what you can find in conventional C library.
*/
package bitops

import "fmt"

// real implementation for Extract32 and GetField32
func extract32(value uint32, start uint, length uint) (uint32, error) {
    return (value >> start) & ( ^uint32(0) >> (32 - length)), nil
}

// real implementation for Extract64 and GetField64
func extract64(value uint64, start uint, length uint) (uint64, error) {
    return (value >> start) & ( ^uint64(0) >> (64 - length)), nil
}

// Extract32 specify field from uint32 by staring position and length
// LSB/MSB are 0/31 and return original value if error occurs
func Extract32(value uint32, start uint, length uint) (uint32, error) {
    if (start > 31  ||  length >  32 - start ) {
        return value, fmt.Errorf("invalid start(%v) or length(%v)", start, length);
    }

    return extract32(value, start, length)
}

// Extract64 specify field from uint64 by starting position and length
// LSB/MSB are 0/63 and return original value if error occurs
func Extract64(value uint64, start uint, length uint) (uint64, error) {
    if (start > 63  ||  length >  64 - start ) {
        return value, fmt.Errorf("invalid start(%v) or length(%v)", start, length);
    }

    return extract64(value, start, length)
}

// GetField32 specify field between high and low bit from uint32 
// LSB/MSB are 0/31 and return original value if error occurs
func GetField32(value uint32, high uint, low uint) (uint32, error) {
    if (high > 31  || low > 31 || high < low) {
        return value, fmt.Errorf("invalid high(%v) or low(%v)", high, low);
    }

    return extract32(value, low, high - low + 1)
}

// GetField64 specify field between high and low bit from uint64 
// LSB/MSB are 0/63 and return original value if error occurs
func GetField64(value uint64, high uint, low uint) (uint64, error) {
    if (high > 63  || low > 63 || high < low) {
        return value, fmt.Errorf("invalid high(%v) or low(%v)", high, low);
    }

    return extract64(value, low, high - low + 1)
}
