package main

import (
	"context"
	"sync"

	"github.com/Yusuf1901-cloud/message_sender/api"
	"github.com/Yusuf1901-cloud/message_sender/config"
	"github.com/Yusuf1901-cloud/message_sender/daemon"
)

func main() {
	cfg := config.New()

	api := api.New(cfg)

	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer cancel()
		api.Run()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		daemon.SendMessages(ctx)
	}()

	wg.Wait()
}
