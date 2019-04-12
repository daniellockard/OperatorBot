package main

import (
	"github.com/nlopes/slack"
)

func ProcessMessage(slackRTMClient *slack.RTM, slackAPIClient *slack.Client, message string, channel string) {
	slackRTMClient.SendMessage(slackRTMClient.NewOutgoingMessage(message, channel))
}
