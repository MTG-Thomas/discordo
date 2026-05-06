$ErrorActionPreference = "Stop"

$changed = gofmt -l .
if ($changed) {
    $changed
    throw "Go files need gofmt"
}

go test ./...
go build -trimpath -ldflags=-s .
pwsh "$PSScriptRoot/smoke.ps1"
