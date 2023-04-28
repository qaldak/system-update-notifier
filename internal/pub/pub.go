package pub

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/slack-go/slack"
)

// Generates and sends message to Slack.
func SlackMsg() {
	msg := generateMsg()

	token := os.Getenv("SLACK_AUTH_TOKEN")
	channelId := os.Getenv("SLACK_CHANNEL_ID")
	log.Println(token, channelId)

	bot := slack.New(token)
	_, timestamp, err := bot.PostMessage(
		channelId,
		slack.MsgOptionText(msg, false),
	)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("Message sent successfully on Channel %s at %s", channelId, timestamp)
}

// Generates message about available system updates.
func generateMsg() string {
	h := getHostname()
	msg := fmt.Sprintf("[%s] System updates available", h)
	log.Println(msg)
	return msg
}

// Get hostname from operating system.
func getHostname() string {
	h, err := os.Hostname()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return strings.ToLower(h)
}