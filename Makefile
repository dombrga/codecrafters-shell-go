run:
	go run ./cmd/myshell/main.go
test:
	go test -coverprofile cover.out -v ./...