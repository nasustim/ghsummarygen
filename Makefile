.PHONY: build_darwin_arm64 build_linux_amd64

build_darwin_arm64:
	env GOOS=darwin GOARCH=arm64 go build -o dist/github_all_summary_macos_arm64 main.go

build_linux_amd64:
	env GOOS=linux GOARCH=amd64 go build -o dist/github_all_summary_linux_amd64 main.go