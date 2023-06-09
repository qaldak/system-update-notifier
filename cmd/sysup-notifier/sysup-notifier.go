package main

import (
	"flag"
	"log"

	"sysup-notifier/internal/pub"
	"sysup-notifier/internal/syschk"
	"sysup-notifier/internal/utils/logger"

	"github.com/joho/godotenv"
)

func main() {
	hasUpdates, newVersion, cntAptPacks := syschk.SearchForUpdates()
	logger.Debug("Updates available: %v", hasUpdates, newVersion, cntAptPacks)

	if hasUpdates {
		logger.Debug("Updates available. Publish Slack message...")
		pub.SlackMsg(newVersion, cntAptPacks)
	}
}

func init() {
	// input args declaration
	var debug bool
	var logfile string

	// flags declaration using flag package
	flag.CommandLine.StringVar(&logfile, "log", "log/sysup-notifier.log", "--log: set path/filename for log or 'none' to disable logfile and use Stdout instead.")
	flag.CommandLine.BoolVar(&debug, "debug", false, "--debug: set loglevel to 'debug'.")

	flag.Parse() // after declaring flags we need to call it

	if isFlagPassed("log") {
		if logfile == "none" {
			log.Println("Logfile disabled. Stdout used instead.")
		}
	}

	if isFlagPassed("debug") {
		debug = true
	}

	// initialize logger
	logger.InitLogger(logfile, debug)

	// load env
	err := godotenv.Load(".env")
	if err != nil {
		logger.Fatal("Error initializing environment variables from '.env'. %v", err)
	}
}

func isFlagPassed(name string) bool {
	flagFound := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			flagFound = true
		}
	})
	return flagFound
}
