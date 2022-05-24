package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
)

func serveStart(s *http.Server) error {
	fmt.Println("HTTP Server Start!")
	return s.ListenAndServe()

}

func serveStop(s *http.Server, errctx context.Context) error {
	return s.Shutdown(errctx)
}

func main() {
	s1 := &http.Server{Addr: ":7788"}
	ctx, cancel := context.WithCancel(context.Background())
	g, errctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		err := serveStart(s1)
		return err
	})

	g.Go(func() error {
		<-errctx.Done()
		return serveStop(s1, errctx)
	})

	cn := make(chan os.Signal, 1)
	signal.Notify(cn)

	g.Go(func() error {
		for {
			select {
			case <-errctx.Done():
				return errctx.Err()
			case <-cn:
				cancel()

			}
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Println("HTTP Server Stop!")
	}

}
