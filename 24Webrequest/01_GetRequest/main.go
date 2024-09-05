package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

const serverPort = 1212

func main() {
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("server: %s/\n", r.Method)
			fmt.Fprintf(w, `{"message": "This is our first request!"}`)
		})

		server := http.Server{
			Addr:    fmt.Sprintf(":%d", serverPort),
			Handler: mux,
		}

		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				fmt.Printf("error running http server: %s\n", err)
			}
		}
	}()

	time.Sleep(100 * time.Millisecond)

	requestURL := fmt.Sprintf("http://localhost:%d", serverPort)
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	fmt.Println("client: got response!")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	select {}
}
