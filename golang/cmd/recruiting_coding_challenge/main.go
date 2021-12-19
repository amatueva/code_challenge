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
	enableGracefulShutdown(ctx)

	query, err := ui.PromptUser(ctx)
	if err != nil {
		fmt.Println("Goodbye")
		return
	}

	searchResult := search.Search(ctx, query)
	searchResult.PrintRecord()
}

func enableGracefulShutdown(ctx context.Context) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		ctx.Done()
		fmt.Println("got a signal, context cancelled - ", sig)
	}()
	fmt.Println("awaiting signal")
}
