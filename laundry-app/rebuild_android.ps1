$ErrorActionPreference = "Stop"
$adb = "C:\Users\THINKPAD X13\AppData\Local\Android\Sdk\platform-tools\adb.exe"

if (Test-Path $adb) {
    Write-Host "Setting up ADB Reverse Proxy..."
    & $adb reverse tcp:4000 tcp:4000
} else {
    Write-Host "Warning: ADB not found at $adb"
}

Write-Host "Building Frontend..."
Set-Location "$PSScriptRoot\frontend"
npm run build

Write-Host "Syncing Capacitor..."
npx cap sync

Write-Host "Done! Please Restart your Android App now."
