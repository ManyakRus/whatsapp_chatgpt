SERVICENAME=whatsapp_chatgpt
SERVICEURL=github.com/ManyakRus/$(SERVICENAME)

FILEMAIN=./internal/main.go
FILEAPP=./bin/whatsapp_chatgpt

NEW_REPO=$(SERVICEURL)


run:
	clear
	go build -race -o $(FILEAPP) $(FILEMAIN)
	#	cd ./bin && \
	./bin/app_race
mod:
	clear
	go mod tidy -compat=1.22
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
	cd cmd && \
	VersionToFile.py
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
newrepo:
	sed -i 's+$(SERVICEURL)+$(NEW_REPO)+g' go.mod
	find -name *.go -not -path "*/vendor/*"|xargs sed -i 's+$(SERVICEURL)+$(NEW_REPO)+g'
graph:
	clear
	image_packages ./ docs/packages.graphml
conn:
	clear
	image_connections ./internal/v0/app docs/connections.graphml $(SERVICENAME)
gocyclo:
	golangci-lint run ./... --disable-all -E gocyclo -v
