package utils

import (
	"fmt"
	"io"
	"net/http"
	"github.com/joho/godotenv"
	"os"
)

var urlTemplate = "https://adventofcode.com/2025/day/%d/input"

func GetInput(day int) (input string, err error) {
	// Load .env from the current working directory (root) when running the driver.
	_ = godotenv.Load()

	// If not found, try utils/.env to remain tolerant of different run locations.
	if os.Getenv("SESSION") == "" {
		_ = godotenv.Load("utils/.env")
	}

	session := os.Getenv("SESSION")

	client := &http.Client{}

	sessCookie := &http.Cookie{
		Name:  "session",
		Value: session,
	}
    
	req, err := http.NewRequest("GET", fmt.Sprintf(urlTemplate, day), nil)
	if err != nil {
		return "", fmt.Errorf("error creating HTTP request: %v", err)
	}
	req.AddCookie(sessCookie)

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-OK HTTP status: %s", resp.Status)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	return string(bodyBytes), nil
}