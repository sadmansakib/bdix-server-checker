SHELL := bash
.SHELLFLAGS := -euo pipefail -c

packager_name = bdix-tester

.PHONY: dist-dir
dist-dir:
	mkdir -p dist
	
PHONY: binary-darwin-amd64
binary-darwin-amd64: dist-dir
	GOOS=darwin GOARCH=amd64 go build -o dist/"${packager_name}"-darwin-amd64 ./main.go

.PHONY: binary-darwin-arm64
binary-darwin-arm64: dist-dir
	GOOS=darwin GOARCH=arm64 go build -o dist/"${packager_name}"-darwin-arm64 ./main.go

.PHONY: binary-linux-amd64
binary-linux-amd64: dist-dir
	GOOS=linux GOARCH=amd64 go build -o dist/"${packager_name}"-linux-amd64 ./main.go

.PHONY: binary-windows-amd64
binary-windows-amd64: dist-dir
	GOOS=windows GOARCH=amd64 go build -o dist/"${packager_name}"-windows-amd64.exe ./main.go

.PHONY: binary-all
binary-all: binary-darwin-amd64 binary-darwin-arm64 binary-linux-amd64 binary-windows-amd64

.PHONY: release
release: binary-all
	./script/release.sh