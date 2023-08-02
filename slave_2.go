package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Send a GET request to retrieve the second part of the file
	resp, err := http.Get("http://localhost:8090/part2")
	if err != nil {
		log.Fatal("Error sending request:", err)
	}
	defer resp.Body.Close()

	// Read the response body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response:", err)
	}

	// Print the second part of the file
	fmt.Println("Part 2:")
	fmt.Println(string(data))
}
