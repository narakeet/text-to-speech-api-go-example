package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	apiKey := os.Getenv("NARAKEET_API_KEY")
	if apiKey == "" {
		fmt.Println("Please set NARAKEET_API_KEY environment variable")
		os.Exit(1)
	}

	voice := "rodney"
	text := "Hi there from Go"
	url := fmt.Sprintf("https://api.narakeet.com/text-to-speech/mp3?voice=%s", voice)

	req, err := http.NewRequest("POST", url, strings.NewReader(text))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		os.Exit(1)
	}

	req.Header.Set("Accept", "application/octet-stream")
	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("x-api-key", apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Failed to generate audio: %d\n%s\n", resp.StatusCode, string(body))
		os.Exit(1)
	}

	file, err := os.Create("output.mp3")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("File saved at: output.mp3")
}
