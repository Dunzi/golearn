package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

/*
基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。
*/

func main() {
	ctx := context.Background()
	eg, errCtx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return http.ListenAndServe(":8080", nil)
	})
	eg.Go(
		func() error {
			<-ctx.Done()
			fmt.Println("server shutdown")
			return errors.New("server shutdown")
		})

	eg.Go(
		func() error {
			quit := make(chan os.Signal, 0)
			signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-quit:
				return ctx.Err()
			}
		})

	if err := eg.Wait(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("exit")
}
