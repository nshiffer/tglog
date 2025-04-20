// Package tglog provides a simple logging system that sends log messages to Telegram
package tglog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

// LogLevel represents the severity level of a log message
type LogLevel int

const (
	// Debug is the lowest severity level, used for detailed debugging information
	Debug LogLevel = iota
	// Info is for general information about normal operation
	Info
	// Warning is for potentially problematic situations that don't cause errors
	Warning
	// Error is for error conditions that might still allow the application to continue running
	Error
	// Fatal is the highest severity level, used for critical errors that prevent further execution
	Fatal
)

// Config holds the configuration for the logger
type Config struct {
	// BotToken is the Telegram bot token
	BotToken string
	// ChatID is the Telegram chat ID where messages will be sent
	ChatID string
	// MinLevel is the minimum log level to send to Telegram (defaults to Info)
	MinLevel LogLevel
	// AppName is the name of the application, prepended to log messages
	AppName string
	// Async determines whether messages are sent asynchronously (defaults to true)
	Async bool
	// DisableColors disables colored messages in Telegram (defaults to false)
	DisableColors bool
	// TimeFormat determines the format for timestamps (defaults to "2006-01-02 15:04:05")
	TimeFormat string
	// HTTPClient allows setting a custom HTTP client
	HTTPClient *http.Client
}

// Logger represents a Telegram logger instance
type Logger struct {
	config     Config
	httpClient *http.Client
	msgQueue   chan message
	wg         sync.WaitGroup
	mu         sync.Mutex
	closed     bool
}

// message represents a log message to be sent to Telegram
type message struct {
	level   LogLevel
	content string
	time    time.Time
}

// telegramMessage represents the JSON structure for a Telegram message
type telegramMessage struct {
	ChatID    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

// DefaultConfig returns a default configuration
func DefaultConfig() Config {
	return Config{
		MinLevel:      Info,
		Async:         true,
		DisableColors: false,
		TimeFormat:    "2006-01-02 15:04:05",
		HTTPClient:    &http.Client{Timeout: 10 * time.Second},
	}
}

// New creates a new Logger with the given configuration
func New(config Config) (*Logger, error) {
	if config.BotToken == "" {
		return nil, fmt.Errorf("bot token is required")
	}
	if config.ChatID == "" {
		return nil, fmt.Errorf("chat ID is required")
	}

	if config.HTTPClient == nil {
		config.HTTPClient = &http.Client{Timeout: 10 * time.Second}
	}

	if config.TimeFormat == "" {
		config.TimeFormat = "2006-01-02 15:04:05"
	}

	logger := &Logger{
		config:     config,
		httpClient: config.HTTPClient,
	}

	if config.Async {
		logger.msgQueue = make(chan message, 100)
		go logger.processQueue()
	}

	return logger, nil
}

// Close waits for all async messages to be sent and closes the logger
func (l *Logger) Close() {
	l.mu.Lock()
	defer l.mu.Unlock()

	if !l.config.Async || l.closed {
		return
	}

	l.closed = true
	close(l.msgQueue)
	l.wg.Wait()
}

// processQueue handles asynchronous message sending
func (l *Logger) processQueue() {
	for msg := range l.msgQueue {
		l.sendMessage(msg.level, msg.content)
		l.wg.Done()
	}
}

// log sends a log message to Telegram
func (l *Logger) log(level LogLevel, format string, args ...interface{}) {
	if level < l.config.MinLevel {
		return
	}

	content := fmt.Sprintf(format, args...)
	now := time.Now()

	// Format message with app name, timestamp, and level prefix
	prefix := getLogLevelPrefix(level, !l.config.DisableColors)
	timestamp := now.Format(l.config.TimeFormat)

	appName := ""
	if l.config.AppName != "" {
		appName = fmt.Sprintf("<b>[%s]</b> ", l.config.AppName)
	}

	formattedMsg := fmt.Sprintf("%s%s %s - %s", appName, prefix, timestamp, content)

	if l.config.Async && !l.closed {
		l.mu.Lock()
		defer l.mu.Unlock()

		if !l.closed {
			l.wg.Add(1)
			l.msgQueue <- message{
				level:   level,
				content: formattedMsg,
				time:    now,
			}
			return
		}
	}

	l.sendMessage(level, formattedMsg)
}

// sendMessage sends a message to Telegram
func (l *Logger) sendMessage(level LogLevel, content string) {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", l.config.BotToken)

	msg := telegramMessage{
		ChatID:    l.config.ChatID,
		Text:      content,
		ParseMode: "HTML",
	}

	jsonData, err := json.Marshal(msg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "tglog: failed to marshal message: %v\n", err)
		return
	}

	resp, err := l.httpClient.Post(apiURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Fprintf(os.Stderr, "tglog: failed to send message: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		fmt.Fprintf(os.Stderr, "tglog: telegram API returned error status: %d\n", resp.StatusCode)
	}

	// If Fatal, exit the program
	if level == Fatal {
		os.Exit(1)
	}
}

// getLogLevelPrefix returns the prefix for a log level with emoji
func getLogLevelPrefix(level LogLevel, colored bool) string {
	var prefix, emoji string

	switch level {
	case Debug:
		prefix = "DEBUG"
		emoji = "🔍" // Magnifying glass
		if colored {
			prefix = fmt.Sprintf("<code>%s</code>", prefix)
		}
	case Info:
		prefix = "INFO"
		emoji = "ℹ️" // Information
		if colored {
			prefix = fmt.Sprintf("<code>%s</code>", prefix)
		}
	case Warning:
		prefix = "WARNING"
		emoji = "⚠️" // Warning
		if colored {
			prefix = fmt.Sprintf("<b><code>%s</code></b>", prefix)
		}
	case Error:
		prefix = "ERROR"
		emoji = "❌" // Cross mark
		if colored {
			// Using HTML that Telegram supports for better visibility
			prefix = fmt.Sprintf("<b><code>%s</code></b>", prefix)
		}
	case Fatal:
		prefix = "FATAL"
		emoji = "💀" // Skull
		if colored {
			prefix = fmt.Sprintf("<b><code>%s</code></b>", prefix)
		}
	default:
		prefix = "UNKNOWN"
		emoji = "❓" // Question mark
	}

	// Return emoji and prefix
	return fmt.Sprintf("%s [%s]", emoji, prefix)
}

// Debug logs a debug message
func (l *Logger) Debug(format string, args ...interface{}) {
	l.log(Debug, format, args...)
}

// Info logs an info message
func (l *Logger) Info(format string, args ...interface{}) {
	l.log(Info, format, args...)
}

// Warning logs a warning message
func (l *Logger) Warning(format string, args ...interface{}) {
	l.log(Warning, format, args...)
}

// Warn is an alias for Warning
func (l *Logger) Warn(format string, args ...interface{}) {
	l.Warning(format, args...)
}

// Error logs an error message
func (l *Logger) Error(format string, args ...interface{}) {
	l.log(Error, format, args...)
}

// Fatal logs a fatal message and exits the program
func (l *Logger) Fatal(format string, args ...interface{}) {
	l.log(Fatal, format, args...)
}
