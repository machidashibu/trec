package usecase

import (
	"context"
	"fmt"
	"log/slog"
	"time"
	"trec/internal/core/logger"
	"trec/internal/domain"
)

type recordingTicker interface {
	Start(ctx context.Context)
	Tick() <-chan time.Time
}

type recordingInput interface {
	Get(prompt string) (string, error)
}

type recordingPrinter interface {
	PrintLine(text string)
	Print(text string)
}

type recordingTimeFormatter interface {
	String(d time.Duration) string
}

type Recording struct {
	repo      domain.TestResultRepository
	ticker    recordingTicker
	inputter  recordingInput
	printer   recordingPrinter
	formatter recordingTimeFormatter
}

func NewRecording(repo domain.TestResultRepository, ticker recordingTicker, inputter recordingInput, printer recordingPrinter, formatter recordingTimeFormatter) *Recording {
	return &Recording{
		repo:      repo,
		ticker:    ticker,
		inputter:  inputter,
		printer:   printer,
		formatter: formatter,
	}
}

func (uc *Recording) Recording(ctx context.Context, testname string) error {
	slog.Debug("Execute Recording")

	// start recording
	start := time.Now()
	uc.printer.PrintLine(fmt.Sprintf("Recording... %s", uc.formatter.String(0)))
	for {
		select {
		case <-ctx.Done():
			// stop
			stop := time.Now()
			uc.printer.Print("")
			slog.Debug("Stop recording", "testname", testname, "start", start, "stop", stop)

			// input memo
			result, err := uc.inputter.Get("Input result: ")
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

			uc.printer.Print("Recorded.")

			slog.Debug("Finished Recording")
			return nil
		case <-uc.ticker.Tick():
			uc.printer.PrintLine(fmt.Sprintf("Recording... %s", uc.formatter.String(time.Since(start).Truncate(time.Second))))
		}
	}
}
