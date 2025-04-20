# TGLog - Telegram Logging for Go

A simple, secure, and efficient Golang package for sending log messages to Telegram.

## Features

- **Easy Integration**: Import into any Go application with minimal setup
- **Multiple Log Levels**: Debug, Info, Warning, Error, and Fatal with distinctive emojis
- **Secure by Default**: Uses HTTPS, modern TLS settings, and environment variables
- **Asynchronous Logging**: Non-blocking log delivery (with synchronous option)
- **Customizable**: Flexible configuration options to suit your needs
- **Visual Distinction**: Log levels include emojis and formatting for clear visibility
- **Timestamps**: Each log message includes the date and time it was generated

## Installation

```bash
go get github.com/nshiffer/tglog
```

## Getting Started

### Setting Up a Telegram Bot

1. Contact [@BotFather](https://t.me/botfather) on Telegram
2. Create a new bot with `/newbot` command
3. Copy the API token provided by BotFather
4. Start a conversation with your new bot
5. Get your chat ID by messaging [@userinfobot](https://t.me/userinfobot)

### Basic Usage

```go
package main

import (
    "github.com/nshiffer/tglog"
)

func main() {
    // Create logger with default configuration
    logger, err := tglog.Simple("My Application")
    if err != nil {
        panic(err)
    }
    defer logger.Close()
    
    // Log messages at different levels
    logger.Debug("Debug message with %s formatting", "string")
    logger.Info("Application started")
    logger.Warning("This is a warning")
    logger.Error("An error occurred: %v", err)
    logger.Fatal("Fatal error, application will exit")
}
```

### Using Environment Variables

Set these environment variables:

- `TELEGRAM_BOT_TOKEN` - Your Telegram bot token
- `TELEGRAM_CHAT_ID` - Your Telegram chat ID
- `TELEGRAM_LOG_LEVEL` - Minimum log level (debug, info, warning, error, fatal)
- `TELEGRAM_APP_NAME` - Name of your application
- `TELEGRAM_ASYNC` - Whether to send logs asynchronously (true/false)
- `TELEGRAM_DISABLE_COLORS` - Whether to disable formatted messages (true/false)

Then use:

```go
logger, err := tglog.WithEnv()
// or for enhanced security:
logger, err := tglog.SecureWithEnv()
```

### Manual Configuration

```go
config := tglog.DefaultConfig()
config.BotToken = "YOUR_BOT_TOKEN"
config.ChatID = "YOUR_CHAT_ID"
config.AppName = "My App"
config.MinLevel = tglog.Info
config.Async = true
config.DisableColors = false
config.TimeFormat = "2006-01-02 15:04:05" // Use Go time format pattern

logger, err := tglog.New(config)
```

## Log Levels with Emojis

TGLog uses the following emojis for different log levels to make them visually distinct:

- **Debug** 🔍: Detailed information for developers
- **Info** ℹ️: General operational information
- **Warning** ⚠️: Potential issues that don't cause errors
- **Error** ❌: Error conditions that allow continued operation
- **Fatal** 💀: Critical errors that stop execution

## Security Best Practices

- Never hardcode your Telegram bot token or chat ID in your source code
- Use environment variables or a secure vault service for sensitive credentials
- Regularly rotate your tokens
- Apply the principle of least privilege when configuring bots
- Keep your dependencies updated

## License

[MIT License](LICENSE)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. 