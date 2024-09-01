package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

type Config struct {
	Target      string
	Charset     string
	MaxLength   int
	Concurrency int
}

type BruteForceState struct {
	Found        bool
	Attempts     int
	FoundAttempt string
	Mutex        sync.Mutex
}

func main() {
	target := flag.String("target", "!@#%$%^&&******DSJKLDSLJKSDLKJDSKJLSDJKLDSJKLSEUIWUIO", "Target password to brute force")
	charset := flag.String("charset", "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_.,", "Character set for brute force")
	maxLength := flag.Int("length", 42, "Maximum password length to attempt")
	concurrency := flag.Int("concurrency", 10, "Number of goroutines for parallel attempts")
	flag.Parse()

	config := Config{
		Target:      *target,
		Charset:     *charset,
		MaxLength:   *maxLength,
		Concurrency: *concurrency,
	}

	state := BruteForceState{}
	startTime := time.Now()

	runBruteForce(config, &state)

	duration := time.Since(startTime)
	if state.Found {
		fmt.Printf("Password found: %s\nAttempts: %d\nTime taken: %v\n", state.FoundAttempt, state.Attempts, duration)
	} else {
		fmt.Println("Password not found within the given length constraints.")
	}
}

func runBruteForce(config Config, state *BruteForceState) {
	var wg sync.WaitGroup

	for i := 0; i < config.Concurrency; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for length := 1; length <= config.MaxLength && !state.Found; length++ {
				generateCombinations(config, length, "", state)
			}
		}(i)
	}

	wg.Wait()
}

func generateCombinations(config Config, length int, prefix string, state *BruteForceState) {
	if state.Found {
		return
	}

	if length == 0 {
		state.Mutex.Lock()
		state.Attempts++
		if prefix == config.Target {
			state.Found = true
			state.FoundAttempt = prefix
		}
		state.Mutex.Unlock()
		return
	}

	for _, char := range config.Charset {
		newPrefix := prefix + string(char)
		generateCombinations(config, length-1, newPrefix, state)
	}
}

// go run main.go -target "Pa$$123" -charset "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()" -length 6 -concurrency 8
