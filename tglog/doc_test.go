package tglog_test

import (
	"fmt"
	"os"

	"github.com/nickshiffer/tglog"
)

// This example demonstrates how to create a logger with manual configuration
func Example_basic() {
	// In a real application, get these from environment variables or secure configuration
	// For example only - these are placeholder values
	botToken := "YOUR_BOT_TOKEN"
	chatID := "YOUR_CHAT_ID"

	if botToken == "YOUR_BOT_TOKEN" {
		// This is just for documentation, not for actual execution
		fmt.Println("Logger created")
		fmt.Println("Info message sent")
		fmt.Println("Warning message sent")
		return
	}

	// Create configuration
	config := tglog.DefaultConfig()
	config.BotToken = botToken
	config.ChatID = chatID
	config.AppName = "Example App"

	// Create the logger
	logger, err := tglog.New(config)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer logger.Close()

	// Log messages
	logger.Info("Application started")
	logger.Warning("Resource usage is high")

	// Output:
	// Logger created
	// Info message sent
	// Warning message sent
}

// This example demonstrates using environment variables for configuration
func Example_environmentVariables() {
	// In a real application, set these through your environment
	// For example: export TELEGRAM_BOT_TOKEN=your_token
	if os.Getenv("TELEGRAM_BOT_TOKEN") == "" {
		// This is just for documentation, not for actual execution
		fmt.Println("Logger created from environment")
		fmt.Println("Message sent")
		return
	}

	// Create logger from environment variables
	logger, err := tglog.WithEnv()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer logger.Close()

	// Log a message
	logger.Info("Using environment variables for configuration")

	// Output:
	// Logger created from environment
	// Message sent
}

// This example demonstrates using the secure configuration
func Example_secureConfiguration() {
	// In a real application, get these securely
	// For example only - these are placeholder values
	botToken := "YOUR_BOT_TOKEN"
	chatID := "YOUR_CHAT_ID"

	if botToken == "YOUR_BOT_TOKEN" {
		// This is just for documentation, not for actual execution
		fmt.Println("Secure logger created")
		fmt.Println("Secure message sent")
		return
	}

	// Create logger configuration
	logConfig := tglog.DefaultConfig()
	logConfig.BotToken = botToken
	logConfig.ChatID = chatID

	// Create secure TLS configuration
	secureConfig := tglog.DefaultSecureConfig()

	// Create the secure logger
	logger, err := tglog.NewSecure(logConfig, secureConfig)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer logger.Close()

	// Log a message securely
	logger.Info("Secure logging initialized")

	// Output:
	// Secure logger created
	// Secure message sent
}
