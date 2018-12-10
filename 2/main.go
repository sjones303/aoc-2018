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
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	defer f.Close()

	var (
		s      = bufio.NewScanner(f)
		t2, t3 int
	)
	for s.Scan() {
		t := s.Text()
		c2, c3 := letters.Count(t)
		if c2 {
			t2++
		}
		if c3 {
			t3++
		}
	}
	if err = s.Err(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("checksum: %d\n", t2*t3)
}
