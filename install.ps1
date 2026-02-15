$ErrorActionPreference = "Stop"

$Repo = "Shivamingale3/ProcPipe"
$Arch = "amd64" # Default to amd64 for now, could detect
$Binary = "procpipe-windows.exe"
$Url = "https://github.com/$Repo/releases/latest/download/$Binary"

$InstallDir = "$env:LOCALAPPDATA\ProcPipe"
$Dest = "$InstallDir\procpipe.exe"

Write-Host "Downloading ProcPipe..." -ForegroundColor Cyan
New-Item -ItemType Directory -Force -Path $InstallDir | Out-Null
Invoke-WebRequest -Uri $Url -OutFile $Dest

Write-Host "Installing..." -ForegroundColor Cyan
& $Dest install

Write-Host "✅ Installed successfully!" -ForegroundColor Green
Write-Host "Please restart your terminal." -ForegroundColor Yellow
