package model

import "time"

type CollapsedTest struct {
	name     string
	duration time.Duration
	count    int
}

func NewCollapsedTest(name string, duration time.Duration, count int) *CollapsedTest {
	return &CollapsedTest{
		name:     name,
		duration: duration,
		count:    count,
	}
}

func (c CollapsedTest) Name() string {
	return c.name
}

func (c CollapsedTest) Duration() time.Duration {
	return c.duration
}

func (c CollapsedTest) Count() int {
	return c.count
}
