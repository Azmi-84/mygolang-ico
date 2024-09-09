package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

const serverPort = 3333

func main() {
	// Start the server in a separate goroutine
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("server: received a %s request\n", r.Method)
			w.WriteHeader(http.StatusOK)
		})

		server := http.Server{
			Addr:    fmt.Sprintf(":%d", serverPort),
			Handler: mux,
		}

		// Start the server
		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				fmt.Printf("error running http server: %s\n", err)
			}
		}
	}()

	// Give the server a moment to start
	time.Sleep(100 * time.Millisecond)

	// Send a POST request to the server
	requestURL := fmt.Sprintf("http://localhost:%d", serverPort)
	res, err := http.Post(requestURL, "application/json", nil)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	// Print the response status code
	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	// Keep the program running
	select {}
}
