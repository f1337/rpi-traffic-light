## Build

```
docker build -t pilight-http .
docker run -it --rm -v "$PWD":/go/src/app -e GOARM=6 -e GOARCH=arm -e GOOS=linux pilight-http go build pilight-http.go
```

## Install onto RPi

```
scp pilight-http pi@192.168.1.xxx:/home/pi/
scp -r public pi@192.168.1.xxx:/home/pi/
```

## Run on the Pi

```
ssh pi@192.168.1.xxx
sudo ./pilight-http
```

## Use the web UI to control the pins/lights

Open http://192.168.1.xxx:8000/ in your browser.

## Get the value of pin #1 using cURL

`curl -H 'Content-Type: application/json'  http://192.168.1.xxx:8000/gpio/1`

## Set the value of pin #2 to 1 (high/on) using cURL

`curl -H 'Content-Type: application/json' -X PUT -d '{"value": "1"}' http://192.168.1.xxx:8000/gpio/2`

## TODO

* docs: screenshots, hardware sketches, video?
* split embd http server into separate repo?
