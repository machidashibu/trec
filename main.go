package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"trec/internal/adapter/controller"
	"trec/internal/adapter/model"
	"trec/internal/adapter/presenter"
	"trec/internal/adapter/repository"
	"trec/internal/core/logger"
	"trec/internal/infra"
	"trec/internal/infra/database"
	"trec/internal/usecase"
	"trec/manual"
)

const configPath = "trec.yaml"

func exit(err error) int {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return 1
	}
	return 0
}

func main() {
	os.Exit(run())
}

func run() int {
	// load config
	config := new(repository.Config)
	if err := config.Read(configPath); err != nil {
		// used default setting if file is not exists. (do not exit)
		slog.Error("config read error", "path", configPath)
	}

	// configure logger
	if err := logger.ApplyConfig(config); err != nil {
		return exit(err)
	}

	// parse command line arguments (Get mode and remaining arguments for options)
	mode, args := controller.ParseArgs(os.Args[1:])

	// create application context
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// prepare database
	db, err := database.OpenSQLite(config)
	if err != nil {
		return exit(err)
	}
	factory := repository.NewFactory(db)
	repoTestResult, err := factory.CreateTestResultRepository()
	if err != nil {
		return exit(err)
	}

	// prepare input/output
	inputter := infra.NewConsoleReader()
	printer := infra.NewConsolePrinter()

	man := manual.NewManual(printer)

	// run application
	switch mode {
	case model.ModeRecording:
		// parse options
		testname, opts, err := controller.ParseRecordingOptions(args, &config.Recording)
		if err != nil {
			return man.Show(mode)
		}
		// prepare recording
		ticker := infra.NewTicker(opts.Interval())
		formatter := presenter.NewDurationFormatter(opts.TimeFormat())
		reporter := presenter.NewRecordingReporter(printer, formatter)
		// Start recording
		go ticker.Start(ctx)
		uc := usecase.NewRecording(repoTestResult, ticker, inputter, reporter)
		if err := uc.Recording(ctx, testname); err != nil {
			return exit(err)
		}
		// wait termination signal
		<-ctx.Done()
	case model.ModeLookup:
		// parse options
		opts, err := controller.ParseLookupOptions(args, &config.Lookup)
		if err != nil {
			return man.Show(mode)
		}
		// prepare lookup
		formatter := presenter.NewLookupFormatter(opts.Format(), opts.TimeFormat())
		reporter := presenter.NewLookupReporter(printer, formatter)
		// Lookup records
		uc := usecase.NewLookup(repoTestResult, reporter)
		if err := uc.Lookup(ctx, opts); err != nil {
			return exit(err)
		}
	case model.ModeDelete:
		// parse options
		id, err := controller.ParseDeleteOptions(args, config)
		if err != nil {
			return man.Show(mode)
		}
		// prepare delete
		reporter := presenter.NewDeleteReporter(printer)
		// Delete record
		uc := usecase.NewDelete(repoTestResult, reporter)
		if err := uc.Delete(ctx, id); err != nil {
			return exit(err)
		}
	case model.ModeHelp:
		if len(args) > 0 {
			return man.Show(controller.MapToMode(args[0]))
		} else {
			return man.Show(model.ModeHelp)
		}
	default:
		return man.Show(model.ModeUnknown)
	}

	return 0
}
