package popcount

// pc[i] is the population count of i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns population count (number of set bits) of x
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

func PopCountLoop(x uint64) int {
	var r byte
	for i := 0; i < 8; i++ {
		r += pc[byte(x>>(i*8))]
	}
	return int(r)
}

func PopCountShift(x uint64) int {
	var r uint64
	for i := 0; i < 64; i++ {
		r += (x >> i) & 1
	}
	return int(r)
}

func PopCountClear(x uint64) int {
	var r int
	for x > 0 {
		r++
		x = x & (x - 1)
	}
	return r
}
