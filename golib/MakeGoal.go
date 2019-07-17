package golib

// MakeGoal returns a spiral target matrix of size s
func MakeGoal(s int) []int {
	ts := s * s
	puzzle := make([]int, ts)
	for i := 0; i < ts; i++ {
		puzzle[i] = -1
	}
	cur, x, ix, y, iy := 1, 0, 1, 0, 0
	for {
		puzzle[x+y*s] = cur
		if cur == 0 {
			break
		}
		cur++
		if x+ix == s || x+ix < 0 || (ix != 0 && puzzle[x+ix+y*s] != -1) {
			iy = ix
			ix = 0
		} else if y+iy == s || y+iy < 0 || (iy != 0 && puzzle[x+(y+iy)*s] != -1) {
			ix = -iy
			iy = 0
		}
		x += ix
		y += iy
		if cur == s*s {
			cur = 0
		}
	}
	return puzzle
}
