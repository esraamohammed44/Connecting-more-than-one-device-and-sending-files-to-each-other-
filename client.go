package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

const (
	ipAddress = "localhost"
	baseURL   = "http://%s:8090"
)

func main() {
	client := &http.Client{}

	// Fetch URL
	url, err := fetchData(client, "/URL")
	if err != nil {
		log.Fatal("Error fetching URL:", err)
	}

	log.Printf("Received URL: %s", url)

	fmt.Println("Please enter the part you need ('part1', 'part2', 'part3', 'part4') or 'file' if you need all parts:")
	var index string
	fmt.Scanln(&index)

	switch index {
	case "part1", "part2", "part3", "part4":
		part, err := fetchData(client, fmt.Sprintf("/%s", index))
		if err != nil {
			log.Fatalf("Error fetching %s: %s", index, err)
		}

		log.Printf("Received %s: %s", index, part)

		wordCounts := countWords(part)
		printWordCounts(wordCounts)

	case "file":
		file, err := fetchData(client, "/file")
		if err != nil {
			log.Fatal("Error fetching file:", err)
		}

		log.Printf("Received file: %s", file)

		parts := strings.Split(file, "#####")

		var wg sync.WaitGroup
		wg.Add(len(parts))

		resultsCh := make(chan map[string]int, len(parts))

		for _, part := range parts {
			go func(content string) {
				defer wg.Done()
				wordCounts := countWords(content)
				resultsCh <- wordCounts
				close(resultsCh)
			}(part)
		}

		wg.Wait()

		finalWordCounts := reduceWordCounts(resultsCh)
		printWordCounts(finalWordCounts)

	default:
		log.Fatal("Invalid index")
	}
}

func fetchData(client *http.Client, path string) (string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(baseURL+path, ipAddress), nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func countWords(text string) map[string]int {
	words := strings.Fields(text)
	wordCounts := make(map[string]int)

	for _, word := range words {
		wordCounts[word]++
	}

	return wordCounts
}

func reduceWordCounts(resultsCh <-chan map[string]int) map[string]int {
	finalWordCounts := make(map[string]int)

	for wordCounts := range resultsCh {
		for word, count := range wordCounts {
			finalWordCounts[word] += count
		}
	}

	return finalWordCounts
}

func printWordCounts(wordCounts map[string]int) {
	fmt.Println("Word counts:")
	for word, count := range wordCounts {
		fmt.Printf("%s : %d\n", word, count)
	}

	totalWordCount := calculateTotalWordCount(wordCounts)
	fmt.Printf("Total word count: %d\n", totalWordCount)
}

func calculateTotalWordCount(wordCounts map[string]int) int {
	total := 0
	for _, count := range wordCounts {
		total += count
	}
	return total
}
