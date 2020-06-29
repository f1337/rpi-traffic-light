mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir := $(dir $(mkfile_path))

default: start
.PHONY: default

.PHONY: build
build:
	@docker run --rm --name yarn-build -v "$(current_dir)":/usr/src/app -w /usr/src/app node:8 yarn build

.PHONY: start
start:
	@docker run -it --rm --name yarn-start -v "$(current_dir)":/usr/src/app -w /usr/src/app -p 3000:3000 node:8 yarn start

.PHONY: test
test:
	@docker run -it --rm --name yarn-test -v "$(current_dir)":/usr/src/app -w /usr/src/app node:8 yarn test
