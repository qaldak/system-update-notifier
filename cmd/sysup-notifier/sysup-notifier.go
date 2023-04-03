package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	hasUpdates := checkUpdates()

	if hasUpdates {
		log.Println("Updates available.")
		sendSlackMsg()
	}
}

func sendSlackMsg() {

	h := getHostname()
	msg := fmt.Sprintf("[%s] System updates available", h)
	log.Println(h)
	log.Println(msg)

	godotenv.Load(".env")
	token := os.Getenv("SLACK_AUTH_TOKEN")
	channel := os.Getenv("SLACK_CHANNEL_ID")
	log.Println(token, channel)

	bot := slack.New(token)
	chl , timestamp, err := bot.PostMessage(
		channel,
		slack.MsgOptionText(msg, false),
	)

	if err != nil {
		log.Println(err)
		return 
	}

	log.Printf("Message sent successfully on Channel %s at %s", chl, timestamp)
}

func getHostname() string {
	h, err := os.Hostname()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return strings.ToLower(h)
}

func checkUpdates() bool {
	// check update files on Dietpi
	updateFiles := []string{"testdata/.update_available", "testdata/.apt_updates"}

	for _, file := range updateFiles {
		fileExists := determineUpdateFiles(file)
		if fileExists {
			log.Println("Update file found:", file)
			return true
		}
	}
	log.Println("No update files found")
	return false
}

func determineUpdateFiles(file string) bool {
	_, err := os.Stat(file)
	if err != nil {
		log.Println("Update file not found:", file)
		return false
	}
	return true
}
