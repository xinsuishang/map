package initialize

import (
	"github.com/cloudwego/kitex/pkg/klog"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"msp/common/model/config"
	"os"
	"path"
	"time"
)

func InitLogger(logConfig *config.LogConfig, logName string, isDev bool) {
	// Customizable output directory.
	logFilePath := path.Join(logConfig.LogDir, logName)
	if err := os.MkdirAll(logFilePath, 0o777); err != nil {
		panic(err)
	}

	// Set filename to date
	logFileName := time.Now().Format("2006-01-02") + ".log"
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			panic(err)
		}
	}

	logger := kitexlogrus.NewLogger(
	//kitexlogrus.WithHook(simpleutils.ContextHook{}),
	)
	// Provides compression and deletion
	lumberjackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    20,   // A file can be up to 20M.
		MaxBackups: 5,    // Save up to 5 files at the same time.
		MaxAge:     10,   // A file can exist for a maximum of 10 days.
		Compress:   true, // Compress with gzip.
	}

	if logConfig.LogDebug || isDev {
		logger.SetLevel(klog.LevelDebug)
	} else {
		logger.SetLevel(klog.LevelWarn)
	}
	logger.SetOutput(lumberjackLogger)
	klog.SetLogger(logger)
}
