# Build

`docker build -t pilight-http .`

`docker run -it --rm -v "$PWD":/go/src/app -e GOARM=6 -e GOARCH=arm -e GOOS=linux pilight-http go build pilight-http.go`

# Install onto RPi

`scp pilight-http pi@192.168.1.xxx:/home/pi/`

# Run on the Pi

`sudo ./pilight-http`

# Get the value of pin #1

GET http://192.168.1.xxx:8000/gpio/1

# Set the value of pin #2 to 1 (high/on)

GET http://192.168.1.xxx:8000/gpio/2/1
