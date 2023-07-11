.PHONY: test cover-test cover-test-html

test:
	go test ./...

cover-test:
	go test -v -coverpkg=./... -coverprofile=coverage.out ./...

cover-test-html: cover-test
	go tool cover -html=coverage.out
