package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"main.go/internal/search"
	"main.go/internal/ui"
)

func main() {
	ctx := context.Background()

	query, err := ui.PromptUser()
	if err != nil {
		fmt.Println("Goodbye")
		return
	}

	search.Search(ctx, query)
}

func enableGracefulShutdown() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Println("got a signal - ", sig)
	}()
	fmt.Println("awaiting signal")
}
