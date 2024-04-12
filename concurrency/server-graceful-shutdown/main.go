package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// create the notification channel
	bye := make(chan os.Signal)
	signal.Notify(bye, os.Interrupt, syscall.SIGTERM)

	mux := http.NewServeMux()
	mux.Handle("/status", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			_, _ = fmt.Fprint(w, "OK")
		},
	))
	srv := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}
	// launch the server in another goroutine
	go func() {
		// launch the server
		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server: %q\n", err)
		}
	}()
	// wait for os signal
	sig := <-bye
	// the code below is executed when we receive an os.Signal
	log.Printf("detected os signal %s \n", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	err := srv.Shutdown(ctx)
	cancel()
	if err != nil {
		log.Fatal(err)
	}
}
