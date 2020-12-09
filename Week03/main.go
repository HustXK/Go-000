package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func serverWork(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "The server is working")
}

//Week03 作业题目：
//1.基于 errgroup 实现一个 http server 的启动和关闭 ，
// 及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出。
func main() {
	group := new(errgroup.Group)
	server := &http.Server{Addr: ":9090"}
	closeCh := make(chan error)
	group.Go(func() error {
		http.HandleFunc("/", serverWork)
		//服务正常会阻塞
		err := server.ListenAndServe()
		closeCh <- err
		return err
	})

	singleCh := make(chan os.Signal, 1)
	signal.Notify(singleCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	group.Go(func() (err error) {
		select {
		case <- closeCh:
			fmt.Println("Get single to close")
			err = errors.New("server to close")
		case <- singleCh:
			fmt.Println("Get single to close")
			err = errors.New("get single to close")
		}
		signal.Stop(singleCh)
		_ = server.Shutdown(context.TODO())
		return err
	})

	err := group.Wait()
	fmt.Printf("End err %+v \n", err)
}
