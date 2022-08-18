.SILENT:
.DEFAULT_GOAL := help

export USRID := $(shell id -u)
export GRPID := $(shell id -g)

.PHONY: help
help: ; $(info Usage:)
	echo "\033[0;32mmake config\033[0;0m             Make some pre-build configuration"
	echo ""

	echo "\033[0;32mmake build\033[0;0m              Build app"
	echo "\033[0;32mmake run\033[0;0m                Run app"
	echo "\033[0;32mmake test\033[0;0m               Run tests"
	echo ""

	echo "\033[0;32mmake docker-build\033[0;0m       Build app through docker"
	echo "\033[0;32mmake docker-run\033[0;0m         Run app through docker"
	echo "\033[0;32mmake docker-test\033[0;0m        Run tests through docker"
	echo ""

.PHONY: build
build: ; $(info building app...)
	go build -o ./pwd ./cmd/cli
	echo "Built successfully!"

.PHONY: run
run: ;
	go run ./cmd/cli

.PHONY: test
test: ; $(info running tests...)
	go test ./...

.PHONY: godog
godog: ; $(info running tests...)
	cp -f .env.test .env
	godog run

.PHONY: config
config: ; $(info Preparing configuration...)
	cp -i .env.dist .env
	echo "Done!"

.PHONY: docker-build
docker-build: ; $(info building app through docker...)
	docker build -f Dockerfile --target builder -t pwd-builder .
	docker build -f Dockerfile -t pwd .
	echo "Built successfully!"

.PHONY: docker-run
docker-run: ;
	docker run -it pwd ./pwd

.PHONY: docker-test
docker-test: ; $(info running tests through docker...)
	docker run --rm pwd-builder godog run
