@echo off
echo === Compilation pour Linux depuis Windows ===

set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0

echo Compilation en cours...
go build -ldflags="-s -w" -o demoshield_linux demoshield.go

if exist demoshield_linux (
    echo ✅ demoshield_linux créé avec succès!
    dir demoshield_linux
) else (
    echo ❌ Erreur de compilation
)
pause