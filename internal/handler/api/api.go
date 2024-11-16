package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

const TIMEOUT = 30 * time.Second

func Start(port string, handler http.Handler) error {
	srv := &http.Server{
		ReadTimeout:  TIMEOUT,
		WriteTimeout: TIMEOUT,
		Addr:         ":" + port,
		Handler:      handler,
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT,
	)

	defer stop()

	errShutDown := make(chan error, 1)

	go shutdown(srv, ctx, errShutDown)

	log.Printf("Current service listening on port %s \n", port)

	err := srv.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		return err
	}

	err = <-errShutDown

	if err != nil {
		return err
	}

	return nil
}

func shutdown(server *http.Server, ctxShutDown context.Context, errShutDown chan error) {
	<-ctxShutDown.Done()

	ctxTimeOut, stop := context.WithTimeout(context.Background(), TIMEOUT)
	defer stop()

	err := server.Shutdown(ctxTimeOut)

	switch err {
	case nil:
		errShutDown <- nil
	case context.DeadlineExceeded:
		errShutDown <- fmt.Errorf("forcing closing the server")
	default:
		errShutDown <- fmt.Errorf("forcing closing the server")
	}
}
