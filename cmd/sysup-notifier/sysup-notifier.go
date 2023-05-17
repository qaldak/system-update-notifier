package main

import (
	"sysup-notifier/internal/pub"
	"sysup-notifier/internal/syschk"
	"sysup-notifier/internal/utils/logger"

	"github.com/joho/godotenv"
)

func main() {
	hasUpdates := syschk.SearchForUpdates()

	if hasUpdates {
		logger.Debug("Updates available.")
		pub.SlackMsg()
	}
}

func init() {
	// Todo: check and map input args.
	debug := false

	// initialize logger
	logger.InitLogger(debug)

	// load env
	err := godotenv.Load(".env")
	if err != nil {
		logger.Fatal("Error initializing environment variables from '.env'. %v", err)
	}
}
