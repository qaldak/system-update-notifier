package logger

import (
	"log"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Sugar *zap.SugaredLogger

func InitLogger(logpath string, debug bool) {
	logDir := filepath.Dir(logpath)

	if _, err := os.Stat(logDir); os.IsNotExist(err){
		errDir := os.MkdirAll(logDir, 0755)
		if errDir != nil {
			log.Printf("Error creating log directory. %v", errDir)
		}
	} 

	logfile, err := os.OpenFile(logpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Error initializing logfile. %v", err)
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

	loggerConfig.OutputPaths = []string{logfile.Name()}
	loggerConfig.ErrorOutputPaths = []string{"stderr", logfile.Name()}

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
