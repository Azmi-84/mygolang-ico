package main

func main() {
	// 	// Welcome message
	// 	welcome := "Welcome to the user input program"
	// 	fmt.Println(welcome)

	// 	// Create a new reader
	// 	reader := bufio.NewReader(os.Stdin)

	// 	// Prompt for the user's name
	// 	fmt.Println("Enter your name: ")
	// 	name, _ := reader.ReadString('\n')
	// 	name = strings.TrimSpace(name)
	// 	fmt.Println("Thanks for the input", name)

	// 	// Prompt for the user's age
	// 	fmt.Println("Enter your age: ")
	// 	input, _ := reader.ReadString('\n')
	// 	input = strings.TrimSpace(input)
	// 	age, err := strconv.Atoi(input)
	// 	if err != nil {
	// 		fmt.Println("Error converting input to integer:", err)
	// 		return
	// 	}
	// 	fmt.Println("Thanks for the input", age)
	// 	fmt.Printf("Type of input is %T\n", age)

	// 	// check if a number is even or odd using a function

	// 	var n int
	// 	fmt.Print("Enter a number: ")
	// 	if _, err := fmt.Scanln(&n); err != nil {
	// 		fmt.Println("Error reading input:", err)
	// 		os.Exit(1)
	// 	}

	// 	if ok, err := isEven(n); ok {
	// 		fmt.Println("it's an even number")
	// 	} else {
	// 		fmt.Println(err)
	// 	}
	// }

	// func isEven(n int) (bool, error) {
	// 	if n&1 == 1 {
	// 		return false, fmt.Errorf("it's an odd number")
	// 	}
	// 	return true, nil

	// 	// Testing if a channel is closed

	// ch := make(chan string) // Initializing a channel

	// // Launching a goroutine to send messages to the channel
	// go func() {
	// 	for i := 1; i <= 5; i++ {
	// 		msg := fmt.Sprintf("Current num is #%v", i)
	// 		ch <- msg
	// 	}
	// 	close(ch) // Closing the channel to indicate no more data will be sent
	// }()

	// // Receiving messages from the channel
	// for msg := range ch {
	// 	fmt.Println("Received message:", msg)
	// }

	// // Trying to receive from a closed channel
	// msg := <-ch
	// fmt.Printf("Message from a closed channel: %#v\n", msg)

	// // Receiving value from a closed channel with comma ok idiom
	// msg, ok := <-ch
	// fmt.Printf("Message: %#v, (was it closed? -> %#v)\n", msg, !ok)
}
