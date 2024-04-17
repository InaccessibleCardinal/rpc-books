dev1:
	go run ./cmd/main.go
dev2:
	go run ./cmd/server/...
test:
	go test ./... -coverprofile cover.out | grep -v /gen
test_int:
	go test -tags=integration ./internal/integrations/... 
cov:
	go tool cover -func cover.out

covh:
	go tool cover -html=cover.out

inter:
	go run ./cmd/inter/...

buildui:
	(cd ./cmd/webui; GOARCH=wasm GOOS=js go build -o web/app.wasm; go build)

runui: buildui
	./cmd/webui/webui.exe
