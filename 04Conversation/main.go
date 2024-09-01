package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please rate our service between 1 to 5: ")
	inputRating, _ := reader.ReadString('\n')

	// // Trim the newline characters from the input
	// inputRating = strings.TrimSpace(inputRating)

	rating, err := strconv.ParseFloat(strings.TrimSpace(inputRating), 64)
	if err != nil {
		fmt.Println("Error parsing rating:", err)
	} else {
		fmt.Println("Thank you for rating us:", rating+1)
	}
}
