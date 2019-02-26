package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/sjones303/aoc-2018/4/event"
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
		guards = map[int][]event.Event{}
		s      = bufio.NewScanner(file)
		l      = 0
		g      = 0
	)
	for s.Scan() {
		l++
		e, err := event.Parse(s.Text())
		if err != nil {
			fmt.Fprintf(os.Stderr, "line %d: unable to parse: %v\n",
				l, err)
			os.Exit(1)
		}
		if e.A == event.Start {
			g = e.G
		}

		guards[g] = append(guards[g], e)
	}

	if err = s.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for g, e := range guards {
		fmt.Printf("%d: %d\n", g, len(e))
	}
}
