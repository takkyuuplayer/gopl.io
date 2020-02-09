.PHONY: test run run-test

setup: dep

dep:
	which dep || go get -u github.com/golang/dep/cmd/dep
	dep ensure
	dep ensure -update

test:
	go vet ./...
	go test ./...

lint: golint gocyclo

golint:
	which golint || go get -u -v golang.org/x/lint/golint
	go list ./... | xargs golint

gocyclo:
	which gocyclo || go get -u -v github.com/fzipp/gocyclo
	find . -maxdepth 1 -mindepth 1 -type d -regex "\.\/[a-z].*" | grep -v vendor | xargs gocyclo -over 15

