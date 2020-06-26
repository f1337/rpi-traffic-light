default: build compile install run
.PHONY: default

.PHONY: build
build:
	@echo 'building image...'
	docker build -q -t pilight-http .

.PHONY: compile
compile:
	@echo 'compiling script...'
	docker run -it --rm -v "$PWD":/go/src/app -e GOARM=6 -e GOARCH=arm -e GOOS=linux pilight-http go build pilight-http.go

.PHONY: install
install:
	@echo 'copying to rpi...'
	scp pilight-http pi@pilight.local:/home/pi/

.PHONY: run
run:
	@echo 'running on rpi...'
	ssh -t pi@pilight.local 'sudo ./pilight-http'
