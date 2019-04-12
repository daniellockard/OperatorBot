# OperatorBot

## Description
This does nothing but connect and log without plugins defined in config.json.

`echo` plugin is included in this repo, build it with `go build -buildmode=plugin plugins/echo.go -o echo.so`


## To build - This builds the main program, and the echo plugin.
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


## :)
PR's Welcome!
