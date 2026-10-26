package app

import (
	"hw1/internal/repository"
	"hw1/internal/service"
	"hw1/internal/ui"
	"io"
)

type Container struct {
	Runner ui.Runner
}

func NewContainer(input io.Reader, output io.Writer) *Container {
	storage := repository.NewInMemoryStorage()
	serviceCenter := service.SampleServiceCenter{}

	parkService := service.NewParkService(
		storage,
		serviceCenter,
	)
	reportService := service.NewReportService(storage)

	console := ui.NewConsole(
		input,
		output,
		parkService,
		reportService,
	)

	return &Container{
		Runner: console,
	}
}
