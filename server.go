package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	data, err := ioutil.ReadFile("file.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Split the file into four parts
	partSize := len(data) / 4
	part1 := data[:partSize]
	part2 := data[partSize : 2*partSize]
	part3 := data[2*partSize : 3*partSize]
	part4 := data[3*partSize:]
	http.HandleFunc("/URL", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)
		w.Write([]byte(fmt.Sprintf("http://%s/part1\nhttp://%s/part2\nhttp://%s/part3\nhttp://%s/part4\n", r.Host, r.Host, r.Host, r.Host)))
	})
	http.HandleFunc("/file", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)
		w.Write(data)
	})
	// Serve each part from a separate slave
	http.HandleFunc("/part1", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)
		w.Write(part1)
	})
	http.HandleFunc("/part2", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)
		w.Write(part2)
	})
	http.HandleFunc("/part3", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)
		w.Write(part3)
	})
	http.HandleFunc("/part4", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)
		w.Write(part4)
	})
	fmt.Println("start")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
