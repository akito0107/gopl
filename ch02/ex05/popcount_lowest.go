package ex03

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCountLowest(x uint64) int {
	s := 0
	for x != 0 {
		x = x & (x - 1)
		s++
	}
	return int(s)
}
