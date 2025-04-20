/*
Package tglog provides a secure and efficient logging system that sends log messages to Telegram.

This package allows applications to send log messages with different severity levels
(Debug, Info, Warning, Error, Fatal) to a Telegram chat. It's designed to be simple
to use while also providing robust security features.

Basic Usage:

	// Create a logger with environment variables
	logger, err := tglog.Simple("My App")
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	// Log messages at different levels
	logger.Debug("Debug message")
	logger.Info("Application started")
	logger.Warning("Resource usage high")
	logger.Error("Failed to process: %v", err)
	logger.Fatal("Critical error, shutting down")

Features:

  - Multiple log levels with distinctive emojis
  - Secure environment variable configuration
  - Asynchronous or synchronous operation
  - Timestamps on all messages
  - HTML formatting for better readability in Telegram

Security Considerations:

This package is designed with security in mind. It uses HTTPS for all
communications with the Telegram API, enforces modern TLS settings, and
avoids storing sensitive credentials in code.

See the package documentation for more detailed information on configuration
and usage options.
*/
package tglog
