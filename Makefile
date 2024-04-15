dev1:
	go run ./cmd/main.go
dev2:
	go run ./cmd/server/...
test:
	go test ./... -coverprofile cover.out
cov:
	go tool cover -func cover.out