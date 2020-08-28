package popcount

// pc[i] is the population count of i
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

// PCLoop returns the population count (number of set bits) of x. Uses loop
func PCLoop(x uint64) int {
	sum := 0
	for i := 0; i < 8; i++ {
		sum += int(pc[byte(x>>(i*8))])
	}
	return sum
}

// PCShift returns the population count (number of set bits) of x. Uses shift
func PCShift(x uint64) int {
	sum := 0
	for x > 0 {
		x = x >> 1
		if x&1 == 1 {
			sum++
		}
	}
	return sum
}

// PCMask returns the population count (number of set bits) of x. Uses mask
func PCMask(x uint64) int {
	sum := 0
	for x > 0 {
		sum++
		x = x & (x - 1)
	}
	return sum
}
