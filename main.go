package main

import (
	"context"
	_ "embed"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"trec/internal/adapter/controller"
	"trec/internal/adapter/presenter"
	"trec/internal/adapter/repository"
	"trec/internal/core"
	"trec/internal/core/logger"
	"trec/internal/infra"
	"trec/internal/infra/database"
	"trec/internal/usecase"
)

const configPath = "trec.yaml"

//go:embed manual.txt
var helpText string

func main() {
	os.Exit(run())
}

func run() int {
	// load config
	config := new(core.Config)
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

	// run application
	switch mode {
	case core.ModeRecording:
		// parse options
		testname, opts, err := controller.ParseRecordingOptions(args, &config.Recording)
		if err != nil {
			return manual()
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
	case core.ModeLookup:
		// parse options
		opts, err := controller.ParseLookupOptions(args, &config.Lookup)
		if err != nil {
			return manual()
		}
		// prepare lookup
		formatter := presenter.NewLookupFormatter(opts.Format(), opts.TimeFormat())
		reporter := presenter.NewLookupReporter(printer, formatter)
		// Lookup records
		uc := usecase.NewLookup(repoTestResult, reporter)
		if err := uc.Lookup(ctx, opts); err != nil {
			return exit(err)
		}
	case core.ModeDelete:
		// parse options
		id, err := controller.ParseDeleteOptions(args, config)
		if err != nil {
			return manual()
		}
		// prepare delete
		reporter := presenter.NewDeleteReporter(printer)
		// Delete record
		uc := usecase.NewDelete(repoTestResult, reporter)
		if err := uc.Delete(ctx, id); err != nil {
			return exit(err)
		}
	case core.ModeHelp:
		return manual()
	default:
		return exit(fmt.Errorf("invalid mode: %s", mode))
	}

	return 0
}

func exit(err error) int {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return 1
	}
	return 0
}

func manual() int {
	fmt.Println(helpText)
	return 2
}
