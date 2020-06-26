## Build, Compile, Install & Run

`make`

## Build

`make build`

## Compile

`make compile`

## Install onto RPi

`make install`

## Run on the Pi

`make run`

## Use the web UI to control the pins/lights

Open http://pilight.local:8000/ in your browser.

## Get the value of pin #1 using cURL

`curl -H 'Content-Type: application/json'  http://pilight.local:8000/gpio/1`

## Set the value of pin #2 to 1 (high/on) using cURL

`curl -H 'Content-Type: application/json' -X PUT -d '{"value": "1"}' http://pilight.local:8000/gpio/2`

## TODO

* docs: screenshots, hardware sketches, video?
* split embd http server into separate repo?
