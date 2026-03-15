#!/usr/bin/env bash
set -euo pipefail

# Development
npx @tailwindcss/cli -i ./public/css/input.css -o ./public/css/output.css

templ generate
go run ./cmd/**/*.go