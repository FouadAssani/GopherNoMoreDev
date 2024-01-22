package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("GopherNoMoreDev API")

	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func run() error {
	app, err := NewApplication()
	if err != nil {
		return fmt.Errorf("failed to initialize application: %w", err)
	}

	errCh := app.Start()
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	select {
	case signalErr := <-signalCh:
		fmt.Printf("system interruption signal received: %s\n", signalErr.String())
	case err := <-errCh:
		fmt.Fprintf(os.Stderr, "error while running the application: %s\n", err)
	}

	app.Stop()

	return nil
}
