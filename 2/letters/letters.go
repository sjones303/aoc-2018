package letters

func Count(in string) (c2, c3 bool) {
	l := make(map[rune]int, 26)
	for _, r := range in {
		l[r]++
	}
	for _, c := range l {
		if c == 2 {
			c2 = true
		}
		if c == 3 {
			c3 = true
		}
		if c2 && c3 {
			return c2, c3
		}
	}
	return c2, c3
}

func FirstDiff(in1, in2 string) int {
	var (
		r1 = []rune(in1)
		r2 = []rune(in2)
	)
	for i, v1 := range r1 {
		if v1 != r2[i] {
			return i
		}
	}
	return -1
}
