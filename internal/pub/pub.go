package pub

import (
	"fmt"
	"os"
	"strings"

	"sysup-notifier/internal/utils/logger"

	"github.com/slack-go/slack"
)

// Generates and sends message to Slack.
func SlackMsg() {
	msg := generateMsg()

	token := os.Getenv("SLACK_AUTH_TOKEN")
	channelId := os.Getenv("SLACK_CHANNEL_ID")

	bot := slack.New(token)
	_, timestamp, err := bot.PostMessage(
		channelId,
		slack.MsgOptionText(msg, false),
	)
	if err != nil {
		logger.Warn("Error sending Slack message. %v", err)
		return
	}

	logger.Debug("Message sent successfully on Channel '%v' at '%v'.", channelId, timestamp)
}

// Generates message about available system updates.
func generateMsg() string {
	h := getHostname()
	msg := fmt.Sprintf("[%v] System updates available", h)
	logger.Debug("Notification message created. %v.", msg)
	return msg
}

// Get hostname from operating system.
func getHostname() string {
	h, err := os.Hostname()
	if err != nil {
		logger.Fatal("Error determine hostname. %v", err)
	}
	return strings.ToLower(h)
}
