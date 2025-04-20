package tglog

import (
	"fmt"
	"os"
)

// Simple creates a new logger with a minimal configuration using environment variables
// It looks for TELEGRAM_BOT_TOKEN and TELEGRAM_CHAT_ID environment variables
func Simple(appName string) (*Logger, error) {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	chatID := os.Getenv("TELEGRAM_CHAT_ID")

	if botToken == "" {
		return nil, fmt.Errorf("TELEGRAM_BOT_TOKEN environment variable not set")
	}

	if chatID == "" {
		return nil, fmt.Errorf("TELEGRAM_CHAT_ID environment variable not set")
	}

	config := DefaultConfig()
	config.BotToken = botToken
	config.ChatID = chatID
	config.AppName = appName

	return New(config)
}

// Must is a helper that wraps a call to a function returning (*Logger, error)
// and panics if the error is non-nil
func Must(logger *Logger, err error) *Logger {
	if err != nil {
		panic(err)
	}
	return logger
}

// GetLogLevelFromString converts a string to a LogLevel
func GetLogLevelFromString(level string) LogLevel {
	switch level {
	case "debug":
		return Debug
	case "info":
		return Info
	case "warning", "warn":
		return Warning
	case "error":
		return Error
	case "fatal":
		return Fatal
	default:
		return Info
	}
}

// WithEnv creates a logger with configuration from environment variables:
// - TELEGRAM_BOT_TOKEN: the Telegram bot token (required)
// - TELEGRAM_CHAT_ID: the Telegram chat ID (required)
// - TELEGRAM_LOG_LEVEL: minimum log level (default: "info")
// - TELEGRAM_APP_NAME: application name (default: "")
// - TELEGRAM_ASYNC: whether to send messages asynchronously (default: "true")
// - TELEGRAM_DISABLE_COLORS: whether to disable colors (default: "false")
// - TELEGRAM_TIME_FORMAT: format for timestamps (default: "2006-01-02 15:04:05")
func WithEnv() (*Logger, error) {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	chatID := os.Getenv("TELEGRAM_CHAT_ID")
	logLevel := os.Getenv("TELEGRAM_LOG_LEVEL")
	appName := os.Getenv("TELEGRAM_APP_NAME")
	async := os.Getenv("TELEGRAM_ASYNC")
	disableColors := os.Getenv("TELEGRAM_DISABLE_COLORS")
	timeFormat := os.Getenv("TELEGRAM_TIME_FORMAT")

	if botToken == "" {
		return nil, fmt.Errorf("TELEGRAM_BOT_TOKEN environment variable not set")
	}

	if chatID == "" {
		return nil, fmt.Errorf("TELEGRAM_CHAT_ID environment variable not set")
	}

	config := DefaultConfig()
	config.BotToken = botToken
	config.ChatID = chatID

	if logLevel != "" {
		config.MinLevel = GetLogLevelFromString(logLevel)
	}

	if appName != "" {
		config.AppName = appName
	}

	if async == "false" {
		config.Async = false
	}

	if disableColors == "true" {
		config.DisableColors = true
	}

	if timeFormat != "" {
		config.TimeFormat = timeFormat
	}

	return New(config)
}
