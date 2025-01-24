package tests

import (
	"github.com/gouef/logger"
	"github.com/gouef/standards"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestFileLogger(t *testing.T) {
	t.Run("Test NewFileLogger", func(t *testing.T) {
		tempFile, err := os.CreateTemp("", "file_logger_test_*.log")
		require.NoError(t, err)
		defer os.Remove(tempFile.Name())

		fileLogger, err := logger.NewFileLogger(tempFile.Name(), standards.INFO, standards.ERROR)
		require.NoError(t, err)
		defer fileLogger.Close()

		assert.NotNil(t, fileLogger)
		assert.FileExists(t, tempFile.Name())
	})

	t.Run("Test Log with enabled and disabled levels", func(t *testing.T) {
		tempFile, err := os.CreateTemp("", "file_logger_test_*.log")
		require.NoError(t, err)
		defer os.Remove(tempFile.Name())

		fileLogger, err := logger.NewFileLogger(tempFile.Name(), standards.INFO, standards.ERROR)
		require.NoError(t, err)
		defer fileLogger.Close()

		err = fileLogger.Log(standards.INFO, "Info message", nil)
		assert.NoError(t, err)

		err = fileLogger.Log(standards.DEBUG, "Debug message", nil)
		assert.NoError(t, err)

		content, err := os.ReadFile(tempFile.Name())
		require.NoError(t, err)

		assert.Contains(t, string(content), "Info message")
		assert.NotContains(t, string(content), "Debug message")
	})

	t.Run("test log with nested context", func(t *testing.T) {
		fileName := "test.log"
		os.Remove(fileName)
		log, _ := logger.NewFileLogger(fileName, standards.INFO)

		context := []any{
			map[string]any{"user": "john_doe"},
			map[string]any{
				"details": map[string]any{
					"age": 30,
					"location": map[string]any{
						"city":    "Brno",
						"country": "Czech Republic",
					},
				},
			},
			map[string]any{"tags": []any{"important", "urgent"}},
		}

		err := log.Info("User logged in", context)
		assert.NoError(t, err)

		fileContent, err := os.ReadFile("test.log")
		assert.NoError(t, err)

		assert.Contains(t, string(fileContent), "\"user\":\"john_doe\"")
		assert.Contains(t, string(fileContent), "{\"age\":30,\"location\":{\"city\":\"Brno\",\"country\":\"Czech Republic\"}}")
		assert.Contains(t, string(fileContent), "[\"important\",\"urgent\"]")

	})

	t.Run("test log levels", func(t *testing.T) {
		fileName := "test.log"
		os.Remove(fileName)
		log, _ := logger.NewFileLogger(fileName,
			standards.INFO,
			standards.ERROR,
			standards.ALERT,
			standards.WARNING,
			standards.CRITICAL,
			standards.EMERGENCY,
			standards.DEBUG,
			standards.NOTICE,
		)

		context := []any{
			map[string]any{"user": "john_doe"},
			map[string]any{
				"details": map[string]any{
					"age": 30,
					"location": map[string]any{
						"city":    "Brno",
						"country": "Czech Republic",
					},
				},
			},
			map[string]any{"tags": []any{"important", "urgent"}},
		}

		err := log.Info("User logged in", context)
		assert.NoError(t, err)

		fileContent, err := os.ReadFile("test.log")
		assert.NoError(t, err)

		assert.Contains(t, string(fileContent), "[info]")
		assert.Contains(t, string(fileContent), "\"user\":\"john_doe\"")
		assert.Contains(t, string(fileContent), "{\"age\":30,\"location\":{\"city\":\"Brno\",\"country\":\"Czech Republic\"}}")
		assert.Contains(t, string(fileContent), "[\"important\",\"urgent\"]")

		err = log.Debug("User logged in", context)
		assert.NoError(t, err)
		fileContent, err = os.ReadFile("test.log")
		assert.NoError(t, err)
		assert.Contains(t, string(fileContent), "[debug]")

		err = log.Warning("User logged in", context)
		assert.NoError(t, err)
		fileContent, err = os.ReadFile("test.log")
		assert.NoError(t, err)
		assert.Contains(t, string(fileContent), "[warning]")

		err = log.Emergency("User logged in", context)
		assert.NoError(t, err)
		fileContent, err = os.ReadFile("test.log")
		assert.NoError(t, err)
		assert.Contains(t, string(fileContent), "[emergency]")

		err = log.Critical("User logged in", context)
		assert.NoError(t, err)
		fileContent, err = os.ReadFile("test.log")
		assert.NoError(t, err)
		assert.Contains(t, string(fileContent), "[critical]")

		err = log.Alert("User logged in", context)
		assert.NoError(t, err)
		fileContent, err = os.ReadFile("test.log")
		assert.NoError(t, err)
		assert.Contains(t, string(fileContent), "[alert]")

		err = log.Notice("User logged in", context)
		assert.NoError(t, err)
		fileContent, err = os.ReadFile("test.log")
		assert.NoError(t, err)
		assert.Contains(t, string(fileContent), "[notice]")

		err = log.Error("User logged in", context)
		assert.NoError(t, err)
		fileContent, err = os.ReadFile("test.log")
		assert.NoError(t, err)
		assert.Contains(t, string(fileContent), "[error]")

		os.Remove(fileName)
	})

	t.Run("test NewFileLogger with default levels", func(t *testing.T) {
		fileName := "test.log"
		log, _ := logger.NewFileLogger(fileName)

		context := []any{
			map[string]any{"user": "john_doe"},
			map[string]any{
				"details": map[string]any{
					"age": 30,
					"location": map[string]any{
						"city":    "Brno",
						"country": "Czech Republic",
					},
				},
			},
			map[string]any{"tags": []any{"important", "urgent"}},
		}

		err := log.Critical("User logged in", context)
		assert.NoError(t, err)

		fileContent, err := os.ReadFile(fileName)
		assert.NoError(t, err)

		assert.Contains(t, string(fileContent), "[critical]")
		assert.Contains(t, string(fileContent), "\"user\":\"john_doe\"")
		assert.Contains(t, string(fileContent), "{\"age\":30,\"location\":{\"city\":\"Brno\",\"country\":\"Czech Republic\"}}")
		assert.Contains(t, string(fileContent), "[\"important\",\"urgent\"]")

		err = log.Debug("User logged in", context)
		assert.NoError(t, err)
		fileContent, err = os.ReadFile(fileName)
		assert.NoError(t, err)
		assert.NotContains(t, string(fileContent), "[debug]")

		err = log.Warning("User logged in", context)
		assert.NoError(t, err)
		fileContent, err = os.ReadFile(fileName)
		assert.NoError(t, err)
		assert.Contains(t, string(fileContent), "[warning]")

		err = log.Emergency("User logged in", context)
		assert.NoError(t, err)
		fileContent, err = os.ReadFile(fileName)
		assert.NoError(t, err)
		assert.Contains(t, string(fileContent), "[emergency]")

		err = log.Info("User logged in", context)
		assert.NoError(t, err)
		fileContent, err = os.ReadFile(fileName)
		assert.NoError(t, err)
		assert.NotContains(t, string(fileContent), "[info]")

		err = log.Alert("User logged in", context)
		assert.NoError(t, err)
		fileContent, err = os.ReadFile(fileName)
		assert.NoError(t, err)
		assert.Contains(t, string(fileContent), "[alert]")

		err = log.Notice("User logged in", context)
		assert.NoError(t, err)
		fileContent, err = os.ReadFile(fileName)
		assert.NoError(t, err)
		assert.NotContains(t, string(fileContent), "[notice]")

		err = log.Error("User logged in", context)
		assert.NoError(t, err)
		fileContent, err = os.ReadFile(fileName)
		assert.NoError(t, err)
		assert.Contains(t, string(fileContent), "[error]")

		os.Remove(fileName)
	})

	t.Run("Test Close and write after close", func(t *testing.T) {
		tempFile, err := os.CreateTemp("", "file_logger_test_*.log")
		require.NoError(t, err)
		defer os.Remove(tempFile.Name())

		fileLogger, err := logger.NewFileLogger(tempFile.Name(), standards.INFO)
		require.NoError(t, err)

		err = fileLogger.Close()
		assert.NoError(t, err)

		err = fileLogger.Log(standards.INFO, "Message after close", nil)
		assert.Error(t, err)
	})

	t.Run("NewFileLogger error opening file", func(t *testing.T) {
		nonExistentFile := "/path/to/nonexistent/file.log"

		log, err := logger.NewFileLogger(nonExistentFile, standards.INFO)

		assert.Nil(t, log, "Expected logger to be nil when file can't be opened")
		assert.NotNil(t, err, "Expected an error when file can't be opened")
		assert.Contains(t, err.Error(), "failed to open log file", "Error message should contain 'failed to open log file'")
	})

	t.Run("NewFileLogger error opening file", func(t *testing.T) {
		context := []any{func() {}}
		fileName := "test.log"

		log, err := logger.NewFileLogger(fileName, standards.INFO)
		assert.NoError(t, err)

		err = log.Info("Test log message with function", context)

		assert.NotNil(t, err, "Expected error when marshaling context with a function")
		assert.Contains(t, err.Error(), "json: unsupported type: func()", "Expected error message to contain 'json: unsupported type: func()'")
	})

}
