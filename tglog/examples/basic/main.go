package main

import (
	"fmt"
	"os"
	"time"

	"github.com/nshiffer/tglog"
)

func main() {
	// You must set these environment variables before running
	// DO NOT hardcode tokens or credentials in your code
	// os.Setenv("TELEGRAM_BOT_TOKEN", "your_bot_token")
	// os.Setenv("TELEGRAM_CHAT_ID", "your_chat_id")
	os.Setenv("TELEGRAM_LOG_LEVEL", "debug")
	os.Setenv("TELEGRAM_APP_NAME", "tglog-demo")
	os.Setenv("TELEGRAM_ASYNC", "true")
	os.Setenv("TELEGRAM_DISABLE_COLORS", "false")

	// Check if Telegram bot token and chat ID are set
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	chatID := os.Getenv("TELEGRAM_CHAT_ID")

	if botToken == "" || chatID == "" {
		fmt.Println("Please set TELEGRAM_BOT_TOKEN and TELEGRAM_CHAT_ID environment variables")
		fmt.Println("You can create a bot using @BotFather on Telegram")
		fmt.Println("To get your chat ID, send a message to @userinfobot on Telegram")
		os.Exit(1)
	}

	fmt.Println("Starting Telegram Logger Demo")
	fmt.Println("-----------------------------")
	fmt.Println("📊 Demonstrating all log levels with emojis and timestamps")

	// Create a logger with a manual configuration
	config := tglog.DefaultConfig()
	config.BotToken = botToken
	config.ChatID = chatID
	config.AppName = "Demo App"
	config.MinLevel = tglog.Debug // Log all messages including debug
	config.DisableColors = false
	config.Async = true
	// You can customize the time format if needed
	config.TimeFormat = "2006-01-02 15:04:05"

	logger, err := tglog.New(config)
	if err != nil {
		fmt.Printf("Failed to create logger: %v\n", err)
		os.Exit(1)
	}
	// Make sure to close the logger at the end to flush any pending messages
	defer logger.Close()

	// Log messages with different severity levels
	logger.Debug("🔍 This is a debug message with emoji and timestamp")
	logger.Info("ℹ️ Application started successfully - normal operation")
	logger.Warning("⚠️ This is a warning - something might need attention")
	logger.Error("❌ An error occurred - operation failed but program continues")
	logger.Fatal("💀 Fatal error - application will exit")

	// Add some structured information with formatting
	type User struct {
		ID   int
		Name string
	}
	user := User{ID: 123, Name: "John Doe"}
	logger.Info("User logged in: ID=%d, Name=%s", user.ID, user.Name)

	// Demonstrate error with additional context
	err = fmt.Errorf("database connection failed")
	logger.Error("System error occurred: %v", err)

	// Wait a moment to make sure async messages are processed
	// This isn't necessary in a real application as long as you call logger.Close()
	time.Sleep(2 * time.Second)

	fmt.Println("✅ All log messages sent. Check your Telegram!")
	fmt.Println("Messages should include emojis, timestamps, and proper formatting")
}
