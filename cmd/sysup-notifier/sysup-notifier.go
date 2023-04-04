package main

import (
	"log"
	"os"

	"sysup-notifier/internal/pub"
	"sysup-notifier/internal/syschk"

	"github.com/joho/godotenv"
)

func main() {
	// load env
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	hasUpdates := syschk.SearchForUpdates()

	if hasUpdates {
		log.Println("Updates available.")
		pub.SlackMsg()
	}
}
