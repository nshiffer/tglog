package main

import (
	"fmt"
	"os"
	"time"

	"github.com/nshiffer/tglog"
)

func main() {
	// SECURITY NOTICE: You should set these environment variables outside your code
	// DO NOT hardcode tokens or credentials in your source code
	// These environment variables should be set before running the application:
	// export TELEGRAM_BOT_TOKEN=your_bot_token
	// export TELEGRAM_CHAT_ID=your_chat_id

	// For this example only, we set non-sensitive variables directly
	os.Setenv("TELEGRAM_APP_NAME", "Secure Example")
	os.Setenv("TELEGRAM_LOG_LEVEL", "info") // Only send info and above

	// Check if required environment variables are set
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	chatID := os.Getenv("TELEGRAM_CHAT_ID")

	if botToken == "" || chatID == "" {
		fmt.Println("ERROR: Missing required environment variables")
		fmt.Println("Please set these environment variables before running:")
		fmt.Println("  export TELEGRAM_BOT_TOKEN=your_bot_token")
		fmt.Println("  export TELEGRAM_CHAT_ID=your_chat_id")
		fmt.Println("")
		fmt.Println("You can create a bot using @BotFather on Telegram")
		fmt.Println("To get your chat ID, send a message to @userinfobot on Telegram")
		os.Exit(1)
	}

	// Create a secure logger using environment variables
	logger, err := tglog.SecureWithEnv()
	if err != nil {
		fmt.Printf("Failed to create secure logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Close()

	// Log some messages
	logger.Debug("This debug message will NOT be sent because log level is set to info")
	logger.Info("Application started with secure configuration")
	logger.Warning("Security notice: remember to never commit API tokens to version control")
	logger.Error("Critical: this is a demonstration error message")

	// Wait for async messages to be sent
	time.Sleep(2 * time.Second)

	fmt.Println("All log messages sent securely. Check your Telegram!")

	// SECURITY BEST PRACTICES:
	// 1. Never hardcode tokens or sensitive data in your source code
	// 2. Use environment variables or a secure vault service
	// 3. Implement proper access controls for your bot
	// 4. Regularly rotate your tokens
	// 5. Monitor for suspicious activity
}
