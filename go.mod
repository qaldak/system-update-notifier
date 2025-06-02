module sysup-notifier

go 1.23.0

toolchain go1.24.1

require github.com/slack-go/slack v0.17.0

require go.uber.org/multierr v1.10.0 // indirect

require (
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/joho/godotenv v1.5.1
	go.uber.org/zap v1.27.0
	golang.org/x/text v0.25.0
)
