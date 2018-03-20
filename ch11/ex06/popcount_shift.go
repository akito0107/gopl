package ex06

func PopCountShift(x uint64) int {
	s := 1 & x
	var i uint64
	for i = 1; i < 64; i++ {
		xi := x >> i
		if xi == 0 {
			break
		}
		s += 1 & xi
	}
	return int(s)
}
