mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir := $(dir $(mkfile_path))
docker_run  := docker run --rm -v "$(current_dir)":/usr/src/app -w /usr/src/app

default: install build
.PHONY: default

.PHONY: build
build:
	$(docker_run) --name yarn-build node:8 yarn build

.PHONY: install
install:
	$(docker_run) --name yarn-install node:8 yarn install

.PHONY: start
start:
	$(docker_run) -i -p 3000:3000 --name yarn-start node:8 yarn start

.PHONY: test
test:
	$(docker_run) --name yarn-test node:8 yarn test
