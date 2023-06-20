package pub

import (
	"fmt"
	"os"
	"strings"

	"sysup-notifier/internal/utils/logger"

	"github.com/slack-go/slack"
)

// Generates and sends message to Slack.
func SlackMsg(newVersion string, cntAptPacks string) {
	msg := generateMsg(newVersion, cntAptPacks)

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
func generateMsg(newVersion string, cntAptPacks string) (msg string) {
	h := getHostname()
	v := ""
	a := ""

	if newVersion != "" {
		v = (" New version: " + newVersion + ",")
		logger.Debug(v)
	}

	if cntAptPacks != "" {
		a = (" APT package updates: " + cntAptPacks)
		logger.Debug(a)
	}

	msg = fmt.Sprintf("[%s] System updates available.%s%s", h, v, a)
	logger.Debug("Notification message created. %v", msg)
	return
}

// Get hostname from operating system.
func getHostname() string {
	h, err := os.Hostname()
	if err != nil {
		logger.Fatal("Error determine hostname. %v", err)
	}
	return strings.ToLower(h)
}
