package ex03

func PopCountLowest(x uint64) int {
	s := 0
	for x != 0 {
		x = x & (x - 1)
		s++
	}
	return int(s)
}
