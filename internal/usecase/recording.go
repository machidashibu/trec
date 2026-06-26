package usecase

import (
	"context"
	"log/slog"
	"time"
	"trec/internal/core/logger"
	"trec/internal/domain"
)

type resoringRepository interface {
	Add(name string, start time.Time, end time.Time, result string) (domain.Test, error)
}

type recordingTicker interface {
	Start(ctx context.Context)
	Tick() <-chan time.Time
}

type recordingInput interface {
	Get(prompt string) (string, error)
}

type recordingReporter interface {
	Start()
	Stop()
	Ticked(d time.Duration)
	Completed()
}

type Recording struct {
	repo     resoringRepository
	ticker   recordingTicker
	in       recordingInput
	reporter recordingReporter
}

func NewRecording(repo resoringRepository, ticker recordingTicker, in recordingInput, reporter recordingReporter) *Recording {
	return &Recording{
		repo:     repo,
		ticker:   ticker,
		in:       in,
		reporter: reporter,
	}
}

func (uc *Recording) Recording(ctx context.Context, testname string) error {
	slog.Debug("Execute Recording")

	// start recording
	start := time.Now()
	uc.reporter.Start()
	for {
		select {
		case <-ctx.Done():
			// stop
			stop := time.Now()
			uc.reporter.Stop()
			slog.Debug("Stop recording", "testname", testname, "start", start, "stop", stop)

			// input memo
			result, err := uc.in.Get("Input result: ")
			if err != nil {
				logger.Error("Recording", "input result error", err)
				result = ""
			}
			slog.Debug("Inputted result", "result", result)

			// add to DB
			record, err := uc.repo.Add(testname, start, stop, result)
			if err != nil {
				return logger.Error("Recording", "Failed to add record", err, "testname", testname, "start", start, "stop", stop, "result", result)
			}
			slog.Debug("Recorded to DB", "record", record)

			uc.reporter.Completed()
			slog.Debug("Finished Recording")
			return nil
		case <-uc.ticker.Tick():
			uc.reporter.Ticked(time.Since(start).Truncate(time.Second))
		}
	}
}
