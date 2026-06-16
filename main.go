package main

import (
	"context"
	_ "embed"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"trec/internal/adapter/persistence"
	"trec/internal/adapter/presenter"
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

	// parse and overwrite configuration by command line arguments
	if err := config.ParseArgs(os.Args[1:]); err != nil {
		return exit(err)
	}

	// create application context
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// prepare database
	db, err := database.OpenSQLite(config)
	if err != nil {
		return exit(err)
	}
	factory := persistence.NewFactory(db)
	repoRecord, err := factory.CreateRecordRepository()
	if err != nil {
		return exit(err)
	}

	// prepare input/output
	inputter := infra.NewConsoleReader()
	printer := infra.NewConsolePrinter()

	// run application
	switch config.AppMode() {
	case core.ModeRecording:
		// Start recording
		ticker := infra.NewTicker(config.Interval())
		uc := usecase.NewRecording(repoRecord, ticker, inputter, printer)
		if err := uc.Recording(ctx, config.Label()); err != nil {
			return exit(err)
		}
	case core.ModeLookup:
		// Lookup records
		formatter := presenter.NewLookupFormatter(config.LookupFormat())
		reporter := presenter.NewLookupReporter(printer, formatter)
		uc := usecase.NewLookup(repoRecord, reporter)
		if err := uc.Lookup(config); err != nil {
			return exit(err)
		}
		cancel() // normal terminate
	case core.ModeHelp:
		return manual()
	default:
		return exit(fmt.Errorf("invalid mode: %s", config.AppMode()))
	}

	// wait termination signal
	<-ctx.Done()

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
