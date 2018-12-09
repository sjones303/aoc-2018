package freq

func Sum(in []int) int {
	return sum(in[0], in[1:])
}

func sum(f int, r []int) int {
	v := f + r[0]
	if len(r) == 1 {
		return v
	}
	return sum(v, r[1:])
}
