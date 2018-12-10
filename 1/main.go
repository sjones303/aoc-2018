package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/sjones303/aoc-2018/1/freq"
)

func main() {
	flag.Parse()

	f, err := os.Open(flag.Arg(0))
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	defer f.Close()

	var (
		s  = bufio.NewScanner(f)
		in []int
	)
	for s.Scan() {
		v, err := strconv.Atoi(s.Text())
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
		in = append(in, v)
	}

	fmt.Printf("sum:   %d\n", freq.Sum(in))
	fmt.Printf("match: %d\n", freq.Match(in))
}
