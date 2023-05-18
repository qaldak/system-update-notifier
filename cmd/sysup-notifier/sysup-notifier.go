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
	hasUpdates := syschk.SearchForUpdates()

	if hasUpdates {
		logger.Debug("Updates available.")
		pub.SlackMsg()
	}
}

func init() {
	// input args declaration
	var nolog bool
	var debug bool

	// flags declaration using flag package
	flag.CommandLine.BoolVar(&nolog, "nl", false, "")
	flag.CommandLine.BoolVar(&nolog, "nolog", false, "--nolog disables logging")
	flag.CommandLine.BoolVar(&debug, "d", false, "Specify pass. Default is password")
	flag.CommandLine.BoolVar(&debug, "debug", false, "Specify pass. Default is password")

	flag.Parse() // after declaring flags we need to call it

	flag.Visit(func(f *flag.Flag) {
		if f.Name == "nolog" || f.Name == "nl" {
			nolog = true
		}
	})

	flag.VisitAll(func(f *flag.Flag) {
		if f.Name == "debug" || f.Name == "d" {
			nolog = true
		}
	})



	log.Println()
	log.Printf("nolog: %v", nolog)
	log.Printf("debug: %v", debug)

	// Todo: check and map input args.
	// debug := true

	// initialize logger
	logger.InitLogger(debug)

	// load env
	err := godotenv.Load(".env")
	if err != nil {
		logger.Fatal("Error initializing environment variables from '.env'. %v", err)
	}
}
