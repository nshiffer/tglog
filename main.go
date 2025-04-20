package main

import (
	"fmt"
	"os"

	"github.com/nshiffer/tglog"
)

func main() {
	fmt.Println("Telegram Alert Package Demo")
	fmt.Println("---------------------------")

	// Check for environment variables
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	chatID := os.Getenv("TELEGRAM_CHAT_ID")

	if botToken == "" || chatID == "" {
		fmt.Println("Please set these environment variables before running:")
		fmt.Println("  export TELEGRAM_BOT_TOKEN=your_bot_token")
		fmt.Println("  export TELEGRAM_CHAT_ID=your_chat_id")
		fmt.Println("")
		fmt.Println("Instructions:")
		fmt.Println("1. Create a bot using @BotFather on Telegram")
		fmt.Println("2. Get your chat ID by messaging @userinfobot on Telegram")
		os.Exit(1)
	}

	// Create a logger with manual configuration
	config := tglog.DefaultConfig()
	config.BotToken = botToken
	config.ChatID = chatID
	config.AppName = "TGLog Demo"

	// Create the logger
	logger, err := tglog.New(config)
	if err != nil {
		fmt.Printf("Failed to create logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Close()

	// Send a test message
	logger.Info("This is a test message from the Telegram Alert package")
	logger.Warning("Various log levels are available: Debug, Info, Warning, Error, and Fatal")
	logger.Info("Check out the examples directory for more complete demonstrations")

	fmt.Println("Messages sent! Check your Telegram.")
}
