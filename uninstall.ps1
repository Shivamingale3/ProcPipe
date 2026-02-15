$ErrorActionPreference = "Stop"

$InstallDir = "$env:LOCALAPPDATA\ProcPipe"
Write-Host "Uninstalling ProcPipe..." -ForegroundColor Cyan

if (Test-Path $InstallDir) {
    Remove-Item -Recurse -Force $InstallDir
    Write-Host "✅ Removed $InstallDir" -ForegroundColor Green
} else {
    Write-Host "⚠️  Directory not found" -ForegroundColor Yellow
}

Write-Host "Note: You may need to remove ProcPipe from your PATH manually." -ForegroundColor Yellow
