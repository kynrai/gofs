#!/usr/bin/env bash

set -e
set -u
set -o pipefail
set -x

npx -y tailwindcss -i ./internal/ui/styles.css -o ./internal/server/assets/css/styles.css --minify
templ generate
go build -o ./tmp/main cmd/server/main.go
