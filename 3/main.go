package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/sjones303/aoc-2018/3/fabric"
)

func main() {
	flag.Parse()

	file, err := os.Open(flag.Arg(0))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	var (
		f = fabric.New(1000, 1000)
		s = bufio.NewScanner(file)
		l = 0
	)
	for s.Scan() {
		l++
		var p fabric.Piece
		p, err = fabric.ParsePiece(s.Text())
		if err != nil {
			fmt.Fprintf(os.Stderr, "line %d: unable to parse: %v\n",
				l, err)
			os.Exit(1)
		}
		err = f.Mark(p)
		if err != nil {
			fmt.Fprintf(os.Stderr, "line %d: unable to mark: %v\n",
				l, err)
			os.Exit(1)
		}
	}
	if err = s.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	count := 0
	for _, r := range f.Record() {
		for _, v := range r {
			if v > 1 {
				count++
			}
		}
	}
	fmt.Printf("overlap: %d\n", count)

	fmt.Println("=========== unique claims ==============")
	for c, m := range f.ListClaims() {
		if len(m) == 0 {
			fmt.Printf("claim: %d\n", c)
		}
	}
}
