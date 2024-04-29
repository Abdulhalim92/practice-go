package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/status", status)
	err := http.ListenAndServe(":8091", nil)
	if err != nil {
		log.Fatal(err)
	}
}

type key int

const (
	requestID key = iota
	jwt
)

func status(w http.ResponseWriter, r *http.Request) {
	// Add request ID to the context
	ctx := context.WithValue(r.Context(), requestID, uuid.NewV4().String())
	// Add credentials to the context
	ctx = context.WithValue(ctx, jwt, r.Header.Get("Authorization"))

	upDB, err := isDatabaseUp(ctx)
	if err != nil {
		log.Printf("failed to check if database is up: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	upAuth, err := isMonitoringUp(ctx)
	if err != nil {
		fmt.Printf("failed to check if monitoring is up: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	_, err = fmt.Fprintf(w, "DB up: %t, Auth up: %t", upDB, upAuth)
	if err != nil {
		fmt.Printf("failed to write response: %v", err)
		return
	}
}

func isDatabaseUp(ctx context.Context) (bool, error) {
	// retrieve the request ID from the context
	reqID, ok := ctx.Value(requestID).(string)
	if !ok {
		return false, fmt.Errorf("request ID in context does not have the expected type")
	}
	log.Printf("req %s - checking db status", reqID)

	return true, nil
}

func isMonitoringUp(ctx context.Context) (bool, error) {
	// retrieve the request ID from the context
	reqID, ok := ctx.Value(requestID).(string)
	if !ok {
		return false, fmt.Errorf("requestID in context does not have the expected type")
	}
	log.Printf("req %s - checking monitoring status", reqID)
	return true, nil
}
