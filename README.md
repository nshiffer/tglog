# Telegram Alert

This project provides a simple, secure, and efficient Golang package for sending log messages to Telegram.

## Features

- **Easy Integration**: Import into any Go application with minimal setup
- **Multiple Log Levels**: Debug, Info, Warning, Error, and Fatal with distinctive emojis
- **Secure by Default**: Uses HTTPS, modern TLS settings, and environment variables
- **Asynchronous Logging**: Non-blocking log delivery (with synchronous option)
- **Customizable**: Flexible configuration options to suit your needs
- **Visual Distinction**: Log levels include emojis and formatting for clear visibility
- **Timestamps**: Each log message includes the date and time it was generated

## Getting Started

### Setting Up a Telegram Bot

1. Contact [@BotFather](https://t.me/botfather) on Telegram
2. Create a new bot with `/newbot` command
3. Copy the API token provided by BotFather
4. Start a conversation with your new bot
5. Get your chat ID by messaging [@userinfobot](https://t.me/userinfobot)

### Installation

```bash
go get github.com/nickshiffer/tglog
```

### Running Examples

This repository contains example applications in the `tglog/examples` directory:

#### Basic Example

```bash
cd tglog/examples/basic
export TELEGRAM_BOT_TOKEN=your_bot_token
export TELEGRAM_CHAT_ID=your_chat_id
go run main.go
```

#### Secure Example

```bash
cd tglog/examples/secure
# Edit main.go to add your bot token and chat ID
go run main.go
```

## Usage

```go
package main

import (
    "github.com/nickshiffer/tglog"
)

func main() {
    // Create logger with default configuration from environment variables
    logger, err := tglog.Simple("My Application")
    if err != nil {
        panic(err)
    }
    defer logger.Close()
    
    // Log messages at different levels
    logger.Debug("Debug message with %s formatting", "string")  // 🔍 [DEBUG]
    logger.Info("Application started")                          // ℹ️ [INFO]
    logger.Warning("This is a warning")                         // ⚠️ [WARNING]
    logger.Error("An error occurred: %v", err)                  // ❌ [ERROR]
    logger.Fatal("Fatal error, application will exit")          // 💀 [FATAL]
}
```

## Log Levels with Emojis

TGLog uses the following emojis for different log levels to make them visually distinct:

- **Debug** 🔍: Detailed information for developers
- **Info** ℹ️: General operational information
- **Warning** ⚠️: Potential issues that don't cause errors
- **Error** ❌: Error conditions that allow continued operation
- **Fatal** 💀: Critical errors that stop execution

For more detailed usage and configuration options, see the [package documentation](tglog/README.md).

## License

[MIT License](tglog/LICENSE)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. 