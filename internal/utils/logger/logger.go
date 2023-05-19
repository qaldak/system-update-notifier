package logger

import (
	"log"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Sugar *zap.SugaredLogger
var logFile *os.File

func InitLogger(logPath string, debug bool) {
	logDir := filepath.Dir(logPath)

	if (logPath != "none") {
		if _, err := os.Stat(logDir); os.IsNotExist(err){
			errDir := os.MkdirAll(logDir, 0755)
			if errDir != nil {
				log.Printf("Error creating log directory. %v", errDir)
			}
		} 
	
		var err error
		logFile, err = os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatalf("Error initializing logfile. %v", err)
		}		
	}

	var loggerConfig zap.Config
	if debug {
		loggerConfig = zap.NewDevelopmentConfig()
	} else {
		loggerConfig = zap.NewProductionConfig()
		loggerConfig.Encoding = "console"
		loggerConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		loggerConfig.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	}

	if (logPath != "none") {
		loggerConfig.OutputPaths = []string{logFile.Name()}
		loggerConfig.ErrorOutputPaths = []string{"stderr", logFile.Name()}	
	} else {
		loggerConfig.OutputPaths = []string{"stdout"}
		loggerConfig.ErrorOutputPaths = []string{"stderr"}
	
	}


	logger, err := loggerConfig.Build()
	if err != nil {
		log.Fatalf("Error building zap logger. %v", err)
	}

	defer logger.Sync()

	Sugar = logger.Sugar()
}

func Debug(msg string, args ...interface{}) {
	if args != nil {
		Sugar.Debugf(msg, args)
	} else {
		Sugar.Debug(msg)
	}
}

func Info(msg string, args ...interface{}) {
	if args != nil {
		Sugar.Infof(msg, args)
	} else {
		Sugar.Info(msg)
	}
}

func Warn(msg string, args ...interface{}) {
	if args != nil {
		Sugar.Warnf(msg, args)
	} else {
		Sugar.Warn(msg)
	}
}

func Error(msg string, args ...interface{}) {
	if args != nil {
		Sugar.Errorf(msg, args)
	} else {
		Sugar.Error(msg)
	}
}

func Fatal(msg string, args ...interface{}) {
	if args != nil {
		Sugar.Fatalf(msg, args)
	} else {
		Sugar.Fatal(msg)
	}
}
