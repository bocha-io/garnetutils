.PHONY: garnetutils build

build:
	@go build -o ./build/garnetutils ./main.go

lint:
	golangci-lint run --fix --out-format=line-number --issues-exit-code=0 --config .golangci.yml --color always ./...

test:
	@go test -v ./...

clean:
	@rm x/garnethelpers/*.go

fix-lines:
	@golines -w .

release-dry:
	@goreleaser release --snapshot --clean

release:
	@goreleaser release
