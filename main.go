package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"syscall/js"
)

const (
	telegramBotToken = "1491351615:AAFz-tst2LWe4YtFGCHx_HLvSwcRxkn1_7Y"
	chatID           = "-1001270836335"
)

type sendMessageRequest struct {
	ChatID    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode,omitempty"`
}

// Define a struct to represent the structure of the JSON response
type IPAddressResponse struct {
	IP string `json:"ip"`
}

func getClientIP() (string, error) {
	// Make a request to the ipify API
	resp, err := http.Get("https://api.ipify.org?format=json")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check if the response status code is OK
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Error: Unexpected status code %d", resp.StatusCode)
	}

	// Decode the JSON response into the IPAddressResponse struct
	var ipResponse IPAddressResponse
	err = json.NewDecoder(resp.Body).Decode(&ipResponse)
	if err != nil {
		return "", err
	}

	return ipResponse.IP, nil
}

func parseBrowserInfo(userAgent string) (browserName, browserVersion string) {
	// Your parsing logic here
	// This is a simple example, and you may want to use a library for more robust parsing
	// Example: extracting browser name and version using a regular expression
	// Note: This is just an illustrative example and may not cover all cases.
	if matched, _ := regexp.MatchString(`Chrome/(\d+(\.\d+)*)`, userAgent); matched {
		browserName = "Chrome"
		browserVersion = strings.Split(strings.Split(userAgent, "Chrome/")[1], " ")[0]
	} else if matched, _ := regexp.MatchString(`Firefox/(\d+(\.\d+)*)`, userAgent); matched {
		browserName = "Firefox"
		browserVersion = strings.Split(strings.Split(userAgent, "Firefox/")[1], " ")[0]
	} else {
		// Handle other browsers or return empty strings
	}
	return browserName, browserVersion
}

func escapeMarkdown(text string) string {
	// Define a regular expression pattern to match the special characters
	re := regexp.MustCompile(`[_*[\]()~` + "`" + `>#\+\-=|{}.!]`)

	// Replace the matched characters with their escaped versions
	escapedText := re.ReplaceAllString(text, "\\$0")

	return escapedText
}

func main() {

	ipAddress, err := getClientIP()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Client IP Address:", ipAddress)

	// Register the parseUserAgent function for use in JavaScript
	userAgent := js.Global().Get("navigator").Get("userAgent").String()

	// Process the User-Agent string as needed
	fmt.Println("User-Agent:", userAgent)

	// Example: Extract browser name and version
	browserName, browserVersion := parseBrowserInfo(userAgent)
	fmt.Println("Browser Name:", browserName)
	fmt.Println("Browser Version:", browserVersion)

	message := fmt.Sprintf("User-Agent: %s\nBrowser Name: %s\nBrowser Version: %s\nClient's IP address: %s",
		userAgent, browserName, browserVersion, ipAddress)

	// Create the request body
	requestBody := sendMessageRequest{
		ChatID: chatID,
		Text:   message,
	}

	// Convert struct to JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Create the HTTP request
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", telegramBotToken)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the content type header
	req.Header.Set("Content-Type", "application/json")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Check the response
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: Unexpected status code", resp.StatusCode)
		return
	}

	fmt.Println("Message sent successfully!")
}
