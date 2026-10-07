#!/bin/bash
set -e
cd "$(dirname "$0")"
echo "=== DEMOSHIELD BUILD ==="
go mod init demoshield 2>/dev/null || true
go build -ldflags="-s -w" -o demoshield demoshield.go
echo "[OK] $(ls -lh demoshield | awk '{print $5}')"
echo "Usage: ./demoshield --encrypt | --decrypt"
