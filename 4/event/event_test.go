package event_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/sjones303/aoc-2018/4/event"
)

func TestParse(t *testing.T) {
	for i, td := range []struct {
		in  string
		exp event.Event
	}{
		{
			"[1518-11-01 00:00] Guard #1 begins shift",
			event.Event{
				T: time.Date(1518, 11, 01, 0, 0, 0, 0, time.UTC),
				G: 1,
				A: event.Start,
			},
		},
		{
			"[1518-11-01 23:58] Guard #99 begins shift",
			event.Event{
				T: time.Date(1518, 11, 01, 23, 58, 0, 0, time.UTC),
				G: 99,
				A: event.Start,
			},
		},
		{
			"[1518-11-04 00:36] falls asleep",
			event.Event{
				T: time.Date(1518, 11, 04, 00, 36, 0, 0, time.UTC),
				G: 0,
				A: event.FallAsleep,
			},
		},
		{
			"[1518-11-04 00:46] wakes up",
			event.Event{
				T: time.Date(1518, 11, 04, 00, 46, 0, 0, time.UTC),
				G: 0,
				A: event.WakeUp,
			},
		},
	} {
		t.Run(fmt.Sprintf("Input%d", i), func(t *testing.T) {
			act, err := event.Parse(td.in)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if act.T != td.exp.T {
				t.Errorf("T = %v, want %v", act.T, td.exp.T)
			}
			if act.G != td.exp.G {
				t.Errorf("G = %d, want %d", act.G, td.exp.G)
			}
			if act.A != td.exp.A {
				t.Errorf("A = %v, want %v", act.A, td.exp.A)
			}
		})
	}
}
