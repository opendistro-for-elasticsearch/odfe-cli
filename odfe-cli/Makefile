# we will put our integration testing in this path
INTEGRATION_TEST_PATH=./it

# this command will start a docker components that we set in docker-compose.yml
docker.start.components:
	docker-compose up;

# shutting down docker components
docker.stop:
	docker-compose down;

# clean up docker
docker.clean:
	docker system prune --volumes;

# this command will trigger integration test
# INTEGRATION_TEST_SUITE_PATH is used for run specific test in Golang, if it's not specified
# it will run all tests under ./it directory
test.integration:
	go clean -testcache;
	go test -tags=integration $(INTEGRATION_TEST_PATH);

test.unit:
	go clean -testcache;
	go test ./...;

# format project using goimports tool
format:
	goimports -w .;

# generate odfe-cli file in current directory
# update GOOS / GOARCH if you want to build for other os and architecture
build:
	go build .