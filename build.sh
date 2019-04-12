#!/bin/bash 


go build -buildmode=plugin -o echo.so plugins/echo.go
go build -buildmode=plugin -o diceroll.so plugins/diceroll.go
go build

echo "run \`SLACK_TOKEN="SLACK_TOKEN_HERE" ./OperatorBot\`"
