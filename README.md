# OperatorBot

## Description
This does nothing but connect and log without plugins defined in config.json.

`echo` plugin is included in this repo, build it with `go build -buildmode=plugin -o echo.so plugins/echo.go`
`diceroll` plugin is included in this repo, build it with `go build -buildmode=plugin -o diceroll.so plugins/diceroll.go`


## Caveats
The way that websockets work in Slack is such that your bot will only get
messages on the rooms it has been invited to, so if it seems like nothing is
working, `/invite @botname` in Slack, in the channels you want the bot to listen
in.


## To build - This builds the main program, the echo plugin, and the diceroll plugin.
* `./build.sh` 

## How do I include my plugin?
* Create a plugin that implements:
```golang
func ProcessMessage(slackRTMClient *slack.RTM, slackAPIClient *slack.Client, message string, channel string) {

}
```
  * See `plugins/echo.go` for an example
* Build it with `go build -buildmode=plugin plugins/<name>.go -o <name>.so`
* Edit config.json
  * Add an object to the array of plugins. 
* start `./OperatorBot`

## Dockerfile
I added a Dockerfile, you can run this with `docker build .`.  It only builds
the echo module right now, but you can modify it to build yours for Linux
x86_64.

Run it with `docker run -e SLACK_TOKEN="SLACK_TOKEN_HERE" imageid`.

## :)
PR's Welcome!
