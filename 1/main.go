package main

import (
	"fmt"
	"strconv"
)

func parse(in []string) ([]int, error) {
	out := make([]int, len(in))
	for i, v := range in {
		var err error
		out[i], err = strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("arg %d: %v", i, err)
		}
	}
	return out, nil
}
