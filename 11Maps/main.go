package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to our Maps Leaning Tutorial!")

	language := make(map[string]string)

	language["en"] = "English"
	language["fr"] = "French"
	language["es"] = "Spanish"
	language["de"] = "German"
	language["it"] = "Italian"
	language["ja"] = "Japanese"
	language["zh"] = "Chinese"
	language["pt"] = "Portuguese"
	language["ru"] = "Russian"
	language["ar"] = "Arabic"

	fmt.Printf("List of languages: %v\n", language)
	fmt.Println("English short from:", language["en"])

	delete(language, "ar")
	fmt.Println("List of languages: ", language)

	for key, value := range language {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}
