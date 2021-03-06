package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func main() {
	srv := &http.Server{
		Addr: ":8080",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println(run(ctx))
		cancel()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println(run(ctx))
		cancel()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	quit := make(chan os.Signal, 1)
	defer close(quit)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("start!")

	// シグナルかコンテキストキャンセルを受ける
	select {
	case <-quit:
		fmt.Println("received signal")
		cancel()
	case <-ctx.Done():
		fmt.Println("received context cancel")
	}

	// シャットダウンの為だけのcontext
	sCtx, sCancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer sCancel()
	if err := srv.Shutdown(sCtx); err != nil {
		log.Println(err)
	}

	// doneを読んだら全てのgorutineが確実に終了することが前提の実装になっている
	wg.Wait()

	fmt.Println("done")
}
