package infra

import (
	"context"
	"time"
)

type Ticker struct {
	interval time.Duration
	notify   chan time.Time
}

func NewTicker(interval time.Duration) *Ticker {
	return &Ticker{
		interval: interval,
		notify:   make(chan time.Time),
	}
}

func (t *Ticker) Start(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(t.interval)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case now := <-ticker.C:
				t.notify <- now
			}
		}
	}()
}

func (t *Ticker) Tick() <-chan time.Time {
	return t.notify
}
