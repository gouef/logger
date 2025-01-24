<img align=right width="168" src="docs/gouef_logo.png">

# logger
This package provides logging utilities with different log levels and the ability to log to multiple loggers (e.g., files, remote systems). It includes the `FileLogger` for logging to a file and `MultiLoggers` for logging to multiple loggers simultaneously.

[![Static Badge](https://img.shields.io/badge/Github-gouef%2Flogger-blue?style=for-the-badge&logo=github&link=github.com%2Fgouef%2Flogger)](https://github.com/gouef/logger)

[![GoDoc](https://pkg.go.dev/badge/github.com/gouef/logger.svg)](https://pkg.go.dev/github.com/gouef/logger)
[![GitHub stars](https://img.shields.io/github/stars/gouef/logger?style=social)](https://github.com/gouef/logger/stargazers)
[![Go Report Card](https://goreportcard.com/badge/github.com/gouef/logger)](https://goreportcard.com/report/github.com/gouef/logger)
[![codecov](https://codecov.io/github/gouef/logger/branch/main/graph/badge.svg?token=YUG8EMH6Q8)](https://codecov.io/github/gouef/logger)

## Versions
![Stable Version](https://img.shields.io/github/v/release/gouef/logger?label=Stable&labelColor=green)
![GitHub Release](https://img.shields.io/github/v/release/gouef/logger?label=RC&include_prereleases&filter=*rc*&logoSize=diago)
![GitHub Release](https://img.shields.io/github/v/release/gouef/logger?label=Beta&include_prereleases&filter=*beta*&logoSize=diago)

## Installation

To use this package in your Go project, run:

```bash
go get -u github.com/gouef/logger
```

## FileLogger
The `FileLogger` logs messages to a file. You can specify the log levels you want to enable, and it will write the log entries in a JSON-like format. It's implements `github.com/gouef/standards/Logger`.

### Usage
```go
package main

import (
	"github.com/gouef/logger"
	"github.com/gouef/standards"
	"log"
)

func main() {
	// Create a new FileLogger
	fileLogger, err := logger.NewFileLogger("app.log", standards.INFO, standards.ERROR)
	if err != nil {
		log.Fatal(err)
	}
	defer fileLogger.Close()

	// Log an info message
	err = fileLogger.Info("This is an info message", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Log an error message
	err = fileLogger.Error("This is an error message", nil)
	if err != nil {
		log.Fatal(err)
	}
}

```

### Methods

- `NewFileLogger(filePath string, levels ...standards.LogLevel) (*FileLogger, error)` Creates a new instance of FileLogger. If no log levels are specified, it defaults to log levels: EMERGENCY, CRITICAL, ERROR, ALERT, and WARNING.
- `Log(level standards.LogLevel, message string, context []any) error` Logs a message with the specified level. Context is an optional array of additional data.
- `Close() error` Closes the underlying log file.

The following log level methods are provided for convenience:

- `Emergency(message string, context []any) error`
- `Alert(message string, context []any) error`
- `Critical(message string, context []any) error`
- `Error(message string, context []any) error`
- `Warning(message string, context []any) error`
- `Notice(message string, context []any) error`
- `Info(message string, context []any) error`
- `Debug(message string, context []any) error`

## MultiLogger
The `MultiLogger` allows you to log to multiple loggers at once (e.g., file, console, remote logging system). It's implements `github.com/gouef/standards/Logger`.

### Usage

```go
package main

import (
	"github.com/gouef/logger"
	"github.com/gouef/standards"
	"log"
)

func main() {
	// Create multiple loggers (e.g., FileLogger, ConsoleLogger)
	fileLogger, err := logger.NewFileLogger("app.log", standards.INFO, standards.ERROR)
	if err != nil {
		log.Fatal(err)
	}

	// Combine the loggers into MultiLoggers
	multiLogger := logger.NewMultiLogger(fileLogger)

	// Log a message to all loggers
	err = multiLogger.Info("This message will be logged to all loggers", nil)
	if err != nil {
		log.Fatal(err)
	}
}

```


## Contributing

Read [Contributing](CONTRIBUTING.md)

## Contributors

<div>
<span>
  <a href="https://github.com/JanGalek"><img src="https://raw.githubusercontent.com/gouef/logger/refs/heads/contributors-svg/.github/contributors/JanGalek.svg" alt="JanGalek" /></a>
</span>
<span>
  <a href="https://github.com/actions-user"><img src="https://raw.githubusercontent.com/gouef/logger/refs/heads/contributors-svg/.github/contributors/actions-user.svg" alt="actions-user" /></a>
</span>
</div>

