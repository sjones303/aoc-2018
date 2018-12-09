package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/sjones303/aoc-2018/1/freq"
)

func main() {
	flag.Parse()
	in, err := parse(flag.Args())
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("sum: %d\n", freq.Sum(in))
}

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
