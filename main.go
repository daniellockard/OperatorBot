package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"plugin"

	"github.com/nlopes/slack"
	log "github.com/sirupsen/logrus"
)

type config struct {
	Plugins []struct {
		Name       string `json:"name"`
		ModulePath string `json:"module_path"`
	} `json:"plugins"`
}

func main() {
	pluginProcessFunctions := []plugin.Symbol{}
	slackAPIClient := slack.New(os.Getenv("SLACK_TOKEN"))
	slackRTMClient := slackAPIClient.NewRTM()
	go slackRTMClient.ManageConnection()

	configFile, err := os.Open("config.json")
	if err != nil {
		log.Panic("please create config.json like is in the repo")
	}
	defer configFile.Close()

	configBytes, err := ioutil.ReadAll(configFile)
	if err != nil {
		log.Panicf("error reading config.json: %s\n", err)
	}

	var pluginConfig config
	err = json.Unmarshal(configBytes, &pluginConfig)
	if err != nil {
		log.Panicf("error reading config.json: %s\n", err)
	}

	for _, pluginToLoad := range pluginConfig.Plugins {
		log.Infof("Loading plugin %s", pluginToLoad.Name)
		loadedPlugin, err := plugin.Open(pluginToLoad.ModulePath)
		if err != nil {
			log.Errorf("Error loading plugin %s at path %s", pluginToLoad.Name, pluginToLoad.ModulePath)
			continue
		}
		processFunction, err := loadedPlugin.Lookup("ProcessMessage")
		if err != nil {
			log.Errorf("Error looking up %s's Start function", pluginToLoad.Name)
			continue
		}
		pluginProcessFunctions = append(pluginProcessFunctions, processFunction)
	}

	for msg := range slackRTMClient.IncomingEvents {
		log.Info("Event Received: ")
		switch ev := msg.Data.(type) {
		case *slack.HelloEvent:

		case *slack.ConnectedEvent:
			log.Infof("Event Info: %#v\n", ev.Info)
			log.Infof("Connection count: %d\n", ev.ConnectionCount)

		case *slack.MessageEvent:
			channel := ev.Msg.Channel
			message := ev.Msg.Text
			log.Infof("Message: %v\n", ev)
			for _, processFunction := range pluginProcessFunctions {
				go processFunction.(func(slackRTMClient *slack.RTM, slackAPIClient *slack.Client, message string, channel string))(slackRTMClient, slackAPIClient, message, channel)
			}

		case *slack.PresenceChangeEvent:
			log.Infof("Presence Change: %v\n", ev)

		case *slack.LatencyReport:
			log.Infof("Current latency: %v\n", ev.Value)

		case *slack.RTMError:
			log.Infof("Error: %s\n", ev.Error())

		case *slack.InvalidAuthEvent:
			log.Infof("Invalid credentials\n")
			return

		default:
			log.Infof("Unexpected: %v\n", msg.Data)
		}
	}

}
