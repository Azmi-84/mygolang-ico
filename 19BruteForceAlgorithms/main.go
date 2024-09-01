package main

import (
	"fmt"
	"time"
)

func bruteForce(target string, charset string, maxLength int) (string, int) {
	var attempt string
	attemptCount := 0
	startTime := time.Now()

	for length := 1; length <= maxLength; length++ {
		if recursiveBruteForce(target, charset, length, "", &attemptCount, &attempt) {
			duration := time.Since(startTime)
			fmt.Printf("Password found: %s\nAttempts: %d\nTime taken: %v\n", attempt, attemptCount, duration)
			return attempt, attemptCount
		}
	}

	fmt.Println("Password not found within the given length constraints.")
	return "", attemptCount
}

func recursiveBruteForce(target, charset string, length int, prefix string, count *int, found *string) bool {
	if length == 0 {
		*count++
		if prefix == target {
			*found = prefix
			return true
		}
		return false
	}

	for i := 0; i < len(charset); i++ {
		newPrefix := prefix + string(charset[i])
		if recursiveBruteForce(target, charset, length-1, newPrefix, count, found) {
			return true
		}
	}
	return false
}

func main() {
	targetPassword := "a3"

	charset := "abcdefghijklmnopqrstuvwxyz0123456789"

	maxLength := 4

	foundPassword, attempts := bruteForce(targetPassword, charset, maxLength)

	if foundPassword != "" {
		fmt.Printf("Password '%s' was successfully cracked in %d attempts.\n", foundPassword, attempts)
	} else {
		fmt.Println("Failed to crack the password.")
	}
}
