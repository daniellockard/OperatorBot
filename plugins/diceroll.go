package main

import (
	"fmt"
	"strings"

	"github.com/justinian/dice"
	"github.com/nlopes/slack"
)

func main() {}

// nolint: deadcode, U1000
func ProcessMessage(slackRTMClient *slack.RTM, slackAPIClient *slack.Client, message string, channel string) {
	if strings.HasPrefix(message, "!roll") {
		splitString := strings.Split(message, " ")
		diceRoll, _, err := dice.Roll(splitString[1])
		if err != nil {
			slackRTMClient.SendMessage(slackRTMClient.NewOutgoingMessage(fmt.Sprintf("Error rolling: %s", err), channel))
		}
		resultString := fmt.Sprintf("Roll Result: %d\n", diceRoll.Int())
		slackRTMClient.SendMessage(slackRTMClient.NewOutgoingMessage(resultString, channel))

	}
}
