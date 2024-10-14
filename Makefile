.PHONY: build_darwin_arm64 build_linux_amd64

build_all: clean build_darwin_arm64 build_linux_amd64

build_darwin_arm64:
	env GOOS=darwin GOARCH=arm64 go build -o dist/github_all_summary_macos_arm64 cmd/main.go

build_linux_amd64:
	env GOOS=linux GOARCH=amd64 go build -o dist/github_all_summary_linux_amd64 cmd/main.go

clean:
	rm -fr ./dist