package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func textToSpeech(client *http.Client, apiKey string, voice string, text string, outputPath string) error {
	url := fmt.Sprintf("https://api.narakeet.com/text-to-speech/mp3?voice=%s", voice)

	req, err := http.NewRequest("POST", url, strings.NewReader(text))
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Accept", "application/octet-stream")
	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("x-api-key", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API error %d: %s", resp.StatusCode, string(body))
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("creating file: %w", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("writing file: %w", err)
	}

	return nil
}

func main() {
	apiKey := os.Getenv("NARAKEET_API_KEY")
	if apiKey == "" {
		fmt.Println("Please set NARAKEET_API_KEY environment variable")
		os.Exit(1)
	}

	client := &http.Client{Timeout: 30 * time.Second}

	err := textToSpeech(client, apiKey, "rodney", "Hi there from Go", "output.mp3")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("File saved at: output.mp3")
}
