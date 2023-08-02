package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Send a GET request to retrieve the third part of the file
	resp, err := http.Get("http://localhost:8090/part3")
	if err != nil {
		log.Fatal("Error sending request:", err)
	}
	defer resp.Body.Close()

	// Read the response body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response:", err)
	}

	// Print the third part of the file
	fmt.Println("Part 3:")
	fmt.Println(string(data))
}
