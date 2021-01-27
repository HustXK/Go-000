package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func recvConn(ctx context.Context, conn net.Conn, msg chan<- string) error {
	defer conn.Close()
	readScan := bufio.NewScanner(conn)
	for readScan.Scan() {
		select {
		case <-ctx.Done():
			fmt.Printf("client read close\n")
			close(msg)
			return nil
		default:
			msg <- readScan.Text()
		}
	}
	if err := readScan.Err(); err != nil {
		return errors.Wrap(err, "read err")
	}
	return nil
}

func sendConn(ctx context.Context, conn net.Conn, msg chan string) error {
	defer conn.Close()
	w := bufio.NewWriter(conn)
	select {
	case <-ctx.Done():
		fmt.Printf("client read close\n")
		close(msg)
		return nil
	default:
		for c := range msg {
			var buf bytes.Buffer
			buf.WriteString("send" + c + "\n")
			_, err := w.WriteString(buf.String())
			err = w.Flush()
			if err != nil {
				return errors.Wrap(err, "write err")
			}
		}
	}

	fmt.Println("client write close")
	return nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	group, ctx := errgroup.WithContext(ctx)
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil{
		fmt.Printf("Listen error: %#v", err)
	}
	closeCh := make(chan os.Signal, 1)
	signal.Notify(closeCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	group.Go(func() (err error) {
		select {
		case <- closeCh:
			listen.Close()
			fmt.Println("Get single to close")
			cancel()
		}
		signal.Stop(closeCh)
		return err
	})
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}
		select {
		case <-ctx.Done():
			fmt.Println("end to stop")
			return
		default:
			go func() {
				msgCh := make(chan string, 1)
				group.Go(func() error {
					return recvConn(ctx, conn, msgCh)
				})
				group.Go(func() error {
					return sendConn(ctx, conn, msgCh)
				})
				if err := group.Wait(); err != nil {
					log.Println("group error...")
				} else {
					log.Println("done...")
				}
			}()
		}
	}
}

