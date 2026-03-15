#!/usr/bin/env bash
set -euo pipefail

docker build -t gui-tester . 

docker run -p 3000:3000 gui-tester