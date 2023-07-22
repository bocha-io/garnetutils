.PHONY: transpiler

lint:
	golangci-lint run --fix --out-format=line-number --issues-exit-code=0 --config .golangci.yml --color always ./...

test:
	@go test -v ./...

