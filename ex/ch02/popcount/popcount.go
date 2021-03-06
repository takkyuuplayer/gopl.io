package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(
		pc[byte(x>>(0*8))] +
			pc[byte(x>>(1*8))] +
			pc[byte(x>>(2*8))] +
			pc[byte(x>>(3*8))] +
			pc[byte(x>>(4*8))] +
			pc[byte(x>>(5*8))] +
			pc[byte(x>>(6*8))] +
			pc[byte(x>>(7*8))])
}

// ex 2.3
func PopCountByLoop(x uint64) int {
	ret := pc[byte(x>>(0*8))]

	for i := uint64(1); i < 8; i++ {
		ret += pc[byte(x>>(i*8))]
	}

	return int(ret)
}

// ex 2.4
func PopCountByShift(x uint64) int {
	ret := uint64(0)

	for i := uint64(0); i < 64; i++ {
		ret += (x >> i) & 1
	}

	return int(ret)
}

// ex 2.5
func PopCountByAndOperator(x uint64) int {
	ret := 0

	for x > 0 {
		x = x & (x - 1)
		ret++
	}

	return ret
}
