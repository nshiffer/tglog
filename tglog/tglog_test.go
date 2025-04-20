package tglog

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// TestLogger tests the basic functionality of the logger
func TestLogger(t *testing.T) {
	// Create a test server to mock Telegram API
	var receivedMessages []telegramMessage
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if !strings.Contains(r.URL.Path, "/sendMessage") {
			t.Errorf("Expected /sendMessage endpoint, got %s", r.URL.Path)
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("Failed to read request body: %v", err)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		var msg telegramMessage
		if err := json.Unmarshal(body, &msg); err != nil {
			t.Errorf("Failed to unmarshal message: %v", err)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		receivedMessages = append(receivedMessages, msg)

		// Send a successful response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := `{"ok":true,"result":{"message_id":1,"from":{"id":123,"is_bot":true,"first_name":"TestBot","username":"test_bot"},"chat":{"id":456,"first_name":"Test","type":"private"},"date":1234567890,"text":""}}`
		w.Write([]byte(response))
	}))
	defer server.Close()

	// Create a logger that uses our test server
	config := DefaultConfig()
	config.BotToken = "test_token"
	config.ChatID = "test_chat_id"
	config.AppName = "TestApp"
	config.MinLevel = Debug
	config.HTTPClient = server.Client()

	// Create a custom HTTP client that redirects to our test server
	originalClient := &http.Client{
		Transport: &testTransport{
			originalTransport: http.DefaultTransport,
			testServer:        server,
			botToken:          config.BotToken,
			t:                 t,
		},
		Timeout: 10 * time.Second,
	}
	config.HTTPClient = originalClient

	logger, err := New(config)
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}
	defer logger.Close()

	// Log messages at different levels
	testMessages := map[LogLevel]string{
		Debug:   "This is a debug message",
		Info:    "This is an info message",
		Warning: "This is a warning message",
		Error:   "This is an error message",
	}

	for level, msg := range testMessages {
		switch level {
		case Debug:
			logger.Debug(msg)
		case Info:
			logger.Info(msg)
		case Warning:
			logger.Warning(msg)
		case Error:
			logger.Error(msg)
		}
	}

	// Wait for async messages to be processed
	time.Sleep(100 * time.Millisecond)
	logger.Close()

	// Check that all messages were sent
	if len(receivedMessages) != len(testMessages) {
		t.Errorf("Expected %d messages, got %d", len(testMessages), len(receivedMessages))
	}

	// Check message content
	for _, msg := range receivedMessages {
		if msg.ChatID != config.ChatID {
			t.Errorf("Expected chat ID %s, got %s", config.ChatID, msg.ChatID)
		}
		if msg.ParseMode != "HTML" {
			t.Errorf("Expected parse mode HTML, got %s", msg.ParseMode)
		}
		// Check that the message contains the app name
		if !strings.Contains(msg.Text, config.AppName) {
			t.Errorf("Message does not contain app name: %s", msg.Text)
		}
	}
}

// Custom transport to redirect API requests to our test server
type testTransport struct {
	originalTransport http.RoundTripper
	testServer        *httptest.Server
	botToken          string
	t                 *testing.T
}

// RoundTrip implements the http.RoundTripper interface
func (tt *testTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Check if this is a request to the Telegram API
	if strings.Contains(req.URL.Host, "api.telegram.org") {
		// Create a new URL pointing to our test server
		newURL := tt.testServer.URL + req.URL.Path
		newReq := &http.Request{
			Method: req.Method,
			URL:    req.URL,
			Body:   req.Body,
			Header: req.Header,
		}
		newReq.URL, _ = req.URL.Parse(newURL)

		// Send the request to our test server
		return tt.originalTransport.RoundTrip(newReq)
	}

	// For other requests, use the original transport
	return tt.originalTransport.RoundTrip(req)
}
