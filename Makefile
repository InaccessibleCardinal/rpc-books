dev1:
	go run ./cmd/main.go
test:
	go test ./... -coverprofile cover.out
cov:
	go tool cover -func cover.out