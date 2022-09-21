// (Package doc comment intentionally malformed to demonstrate golint.)
//!+
package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// exercise 2.3
func PopCount2(x uint64) int {
	var output byte
	for i := 0; i < 7; i++ {
		output += byte(x >> i * 8)
	}
	return int(output)
}

// exercise 2.4
func PopCountShift(x uint64) int {
	var output byte
	for i := 0; i < 64; i++ {
		if x&1 != 0 {
			output++
		}
		x = x >> 1
	}
	return int(output)
}

// exercise 2.5
func PopCountClearRight(x uint64) int {
	var output byte
	for x != 0 {
		x = x & (x - 1)
		output++
	}
	return int(output)
}
