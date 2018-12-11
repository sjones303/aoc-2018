package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/sjones303/aoc-2018/2/letters"
)

func main() {
	flag.Parse()

	f, err := os.Open(flag.Arg(0))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer f.Close()

	var (
		s      = bufio.NewScanner(f)
		t2, t3 int
		in     []string
	)
	for s.Scan() {
		t := s.Text()
		in = append(in, t)
		c2, c3 := letters.Count(t)
		if c2 {
			t2++
		}
		if c3 {
			t3++
		}
	}
	if err = s.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("checksum: %d\n", t2*t3)

	for i, v1 := range in {
		for _, v2 := range in[i+1:] {
			m := letters.FirstDiff(v1, v2)
			if m < 0 {
				continue
			}
			if letters.FirstDiff(v1[m+1:], v2[m+1:]) >= 0 {
				continue
			}
			fmt.Printf("\"%s{%c,%c}%s\"\n", v1[:m], v1[m], v2[m], v1[m+1:])
			os.Exit(0)
		}
	}

	fmt.Fprintln(os.Stderr, "no similar IDs found")
	os.Exit(1)
}
