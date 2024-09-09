package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	const URL = "http://localhost:8080"

	data := url.Values{}
	data.Add("name", "azmi-84")
	data.Add("name", "azmi-84")
	data.Add("name", "azmi-84")
	data.Add("name", "azmi-84")
	data.Add("name", "azmi-84")
	data.Add("name", "azmi-84")

	response, err := http.PostForm(URL, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
