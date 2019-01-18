package event

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Action int

const (
	Start Action = iota + 1
	FallAsleep
	WakeUp
)

type Event struct {
	T time.Time
	G int
	A Action
}

func Parse(in string) (Event, error) {
	e := Event{}

	n := strings.LastIndex(in, "]")
	if n < 0 {
		return e, errors.New("unable to find end of time")
	}

	var err error
	e.T, err = time.Parse("[2006-01-02 15:04]", in[:n+1])
	if err != nil {
		return e, fmt.Errorf("malformed time: %v", err)
	}

	switch {
	case strings.Contains(in[n+2:], "falls asleep"):
		e.A = FallAsleep
	case strings.Contains(in[n+2:], "wakes up"):
		e.A = WakeUp
	default:
		p := 0
		p, err := fmt.Sscanf(in[n+2:], "Guard #%d begins shift", &e.G)
		switch {
		case err != nil:
			return e, fmt.Errorf("bad start action: %v", err)
		case p != 1:
			return e, fmt.Errorf(
				"bad start action: %d/1 items found", p)
		}
		e.A = Start
	}

	return e, nil
}
