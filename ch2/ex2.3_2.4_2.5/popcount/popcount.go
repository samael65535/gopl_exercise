package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount1(x uint64) int {

	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCount2(x uint64) int {
	// ex2.3
	c := 0
	for i := 0; i < 8; i++ {
		s := byte(x >> uint(i*8))
		c += int(pc[s])
	}
	return c
}

func PopCount3(x uint64) int {
	// ex2.4
	c := 0
	b := x

	for {
		if b%2 == 1 {
			c++
		}

		b = b >> 1
		if b == 0 {
			return c
		}
	}
}

func PopCount4(x uint64) int {
	// ex2.5
	c := 0
	b := x
	for {
		if byte(b&(b-1)) == byte(b-1) {
			c++
		}
		b = b >> 1
		if b == 0 {
			return c
		}

	}
}
