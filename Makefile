# Meta
VERSION=`git describe --always`
DATE=`date`
LDFLAGS := -X 'main.version=$(VERSION)' -X 'main.date=$(DATE)'

# Setup
setup:
	  go get github.com/golang/lint/golint
	  go get golang.org/x/tools/cmd/goimports
	  go get github.com/Songmu/make2help/cmd/make2help
	  go get github.com/golang/dep

# Install dependencies
deps: setup
	  dep ensure

# Build
build: deps
	  GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o bin/linux64/text2img app/cli/app.go
		GOOS=linux GOARCH=386 go build -ldflags "$(LDFLAGS)" -o bin/linux386/text2img app/cli/app.go
		GOOS=windows GOARCH=386 go build -ldflags "$(LDFLAGS)" -o bin/windows386/text2img.exe app/cli/app.go
		GOOS=windows GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o bin/windows64/text2img.exe app/cli/app.go
		GOOS=darwin GOARCH=386 go build -ldflags "$(LDFLAGS)" -o bin/darwin386/text2img app/cli/app.go
		GOOS=darwin GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o bin/darwin64/text2img app/cli/app.go

# Show help
help:
	  @make2help $(MAKEFILE_LIST)

.PHONY: setup deps build help
