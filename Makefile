FILEMAIN=./internal/v0/app/main.go
FILEAPP=./bin/whatsapp_chatgpt
SERVICEURL=github.com/ManyakRus/whatsapp_chatgpt/
SERVICEURL2=gitlab.aescorp.ru/dsp_dev/claim/nikitin/

run:
	clear
	go build -race -o $(FILEAPP) $(FILEMAIN)
	#	cd ./bin && \
	./bin/app_race
mod:
	clear
	go mod tidy -compat=1.17
	go mod vendor
	go fmt ./...
build:
	clear
	go build -race -o $(FILEAPP) $(FILEMAIN)
	cd ./cmd && \
	./VersionToFile.py
buildwin:
	cls
	go build -race -o bin\whatsapp_chatgpt.exe internal\v0\app\main.go
	cd ./cmd && \
	./VersionToFile.py
lint:
	clear
	go fmt ./...
	golangci-lint run ./internal/v0/...
	golangci-lint run ./pkg/v0/...
	gocyclo -over 10 ./internal/v0
	gocyclo -over 10 ./pkg/v0
	gocritic check ./internal/v0/...
	gocritic check ./pkg/v0/...
	staticcheck ./internal/v0/...
	staticcheck ./pkg/v0/...
run.test:
	clear
	go fmt ./...
	go test -coverprofile cover.out ./internal/v0/app/... ./cmd/...
	go tool cover -func=cover.out
graph:
	goda graph -f "{{.Package.Name}}" "shared($(SERVICEURL)... $(SERVICEURL2)...)" | dot -Tsvg -o graph.svg
dot:
	goda graph -f "{{.Package.Name}}" "shared($(SERVICEURL)... $(SERVICEURL2)...)" >graph.dot
