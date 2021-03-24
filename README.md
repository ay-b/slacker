# Slacker

Another mini-tool for sending messages to Slack or another compatible service from pipelines using binary or calling a docker image

## Requirements
System variables:
* API_URL — e.g. https://hooks.slack.com/services/xxxxxxxx/xxxxxxxx/xxxxxxxxxxxxxxxxxxxxxxxx
* MESSAGE — e.g. {"channel": "#devops", "username": "test", "text": "This is a test."}

## How to use

`docker run --rm -e MESSAGE=$MESSAGE -e API_URL=$API_URL suncheez/slacker`

## Why

Because curl isn't available everywhere and just for fun.
This tool replaces a command: `curl -s -X POST --data-urlencode "${MESSAGE}" ${API_URL}`
