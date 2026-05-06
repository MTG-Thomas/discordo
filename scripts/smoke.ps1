$ErrorActionPreference = "Stop"

$root = Split-Path -Parent $PSScriptRoot
$outDir = Join-Path ([System.IO.Path]::GetTempPath()) ("discordo-smoke-" + [System.Guid]::NewGuid().ToString("N"))
New-Item -ItemType Directory -Path $outDir | Out-Null

try {
    $exe = Join-Path $outDir "discordo.exe"
    $version = "smoke-mtg"
    $commit = "smoke"
    $date = "2000-01-01T00:00:00Z"
    $ldflags = "-s -X github.com/ayn2op/discordo/internal/version.Version=$version -X github.com/ayn2op/discordo/internal/version.Commit=$commit -X github.com/ayn2op/discordo/internal/version.Date=$date -X github.com/ayn2op/discordo/internal/version.Distribution=mtg-fork"

    Push-Location $root
    try {
        go build -trimpath -ldflags="$ldflags" -o $exe .

        $got = & $exe -version
        $want = "discordo distribution=mtg-fork version=$version commit=$commit date=$date"
        if ($got -ne $want) {
            throw "unexpected version output: got '$got', want '$want'"
        }

        go test ./internal/config -run TestLoad
    }
    finally {
        Pop-Location
    }
}
finally {
    Remove-Item -Recurse -Force $outDir
}
