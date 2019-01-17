package event

import (
	"errors"
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
	return Event{}, errors.New("not implemented")
}
