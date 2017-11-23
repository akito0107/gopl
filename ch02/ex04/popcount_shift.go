package ex03

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCountShift(x uint64) int {
	s := 1 & x
	var i uint64
	for i = 1; i < 64; i++ {
		s += 1 & (x >> i)
	}
	return int(s)
}
