package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if err := write("readme.txt", "This is a readme file"); err != nil {
		log.Fatal("Failed to write file", err)
	}
	readFile("./readme.txt")
}

func write(fileName string, text string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.WriteString(file, text)
	if err != nil {
		return err
	}
	return file.Close()
}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	fmt.Println("Text data inside the file is: ", dataByte)
}
