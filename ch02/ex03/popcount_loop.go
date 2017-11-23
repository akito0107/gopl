package ex03

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCountLoop(x uint64) int {
	var s int
	for i := 0; i < 8; i++ {
		s += int(pc[byte(x>>(0*8))])
	}
	return s
}
