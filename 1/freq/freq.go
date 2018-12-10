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

func Match(in []int) int {
	f := make(map[int]struct{}, 10000)
	f[0] = struct{}{}
	t := 0
	for i := 0; ; i++ {
		v := in[i % len(in)]
		t += v
		if _, ok := f[t]; ok {
			return t
		}
		f[t] = struct{}{}
	}
}
