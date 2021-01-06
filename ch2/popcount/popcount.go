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

// Loop counts population of x using a loop instead of an unrolled loop
func Loop(x uint64) int {
	var acc int

	for i := 0; i < 8; i++ {
		acc += int(pc[byte(x>>(i*8))])
	}
	return acc
}

// Drop counts population of x by counting and shifting LMB
func Drop(x uint64) int {
	var pop = 0
	for i := 0; i < 64; i++ {
		if x != (x&x - 1) {
			pop++
		}
	}
	return pop
}
