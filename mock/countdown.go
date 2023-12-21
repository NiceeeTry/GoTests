package mock

import (
	"fmt"
	"io"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// type SpySleeper struct {
// 	Calls int
// }

// func (s *SpySleeper) Sleep() {
// 	s.Calls++
// }

type SpyCountdownOoperations struct {
	Calls []string
}

func (s *SpyCountdownOoperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOoperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type ConfigurableSleeper struct {
	Duration time.Duration
	Sleeper  func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.Sleeper(c.Duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

const write = "write"
const sleep = "sleep"

func Countdown(out io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		s.Sleep()
	}
	fmt.Fprintln(out, finalWord)
}
