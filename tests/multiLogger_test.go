package tests

import (
	"github.com/gouef/logger"
	"github.com/gouef/standards"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestMultiLogger(t *testing.T) {
	t.Run("MultiLogger", func(t *testing.T) {
		tempFileInfo, err := os.CreateTemp("", "file_logger_test_info_*.log")
		require.NoError(t, err)
		defer os.Remove(tempFileInfo.Name())

		tempFileDebug, err := os.CreateTemp("", "file_logger_test_debug_*.log")
		require.NoError(t, err)
		defer os.Remove(tempFileDebug.Name())

		fileLoggerInfo, err := logger.NewFileLogger(tempFileInfo.Name(), standards.INFO)
		fileLoggerDebug, err := logger.NewFileLogger(tempFileDebug.Name(), standards.DEBUG)

		multiLogger := logger.NewMultiLogger(fileLoggerInfo, fileLoggerDebug)

		err = multiLogger.Info("Something", nil)
		require.NoError(t, err)
		err = multiLogger.Debug("debug Something", nil)
		require.NoError(t, err)
		err = multiLogger.Emergency("debug Something", nil)
		require.NoError(t, err)
		err = multiLogger.Alert("debug Something", nil)
		require.NoError(t, err)
		err = multiLogger.Critical("debug Something", nil)
		require.NoError(t, err)
		err = multiLogger.Error("debug Something", nil)
		require.NoError(t, err)
		err = multiLogger.Warning("debug Something", nil)
		require.NoError(t, err)
		err = multiLogger.Notice("debug Something", nil)
		require.NoError(t, err)

		require.NoError(t, err)
		defer fileLoggerInfo.Close()

		assert.NotNil(t, fileLoggerInfo)
		assert.FileExists(t, tempFileInfo.Name())

		assert.NotNil(t, fileLoggerDebug)
		assert.FileExists(t, tempFileDebug.Name())
	})

	t.Run("MultiLogger error", func(t *testing.T) {
		tempFile, err := os.CreateTemp("", "file_logger_test_*.log")
		require.NoError(t, err)
		defer os.Remove(tempFile.Name())

		fileLogger, err := logger.NewFileLogger(tempFile.Name(), standards.INFO)
		require.NoError(t, err)

		multiLogger := logger.NewMultiLogger(fileLogger)

		err = fileLogger.Close()
		assert.NoError(t, err)

		err = multiLogger.Log(standards.INFO, "Message after close", nil)
		assert.Error(t, err)
	})
}
