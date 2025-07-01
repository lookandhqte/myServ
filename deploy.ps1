param([string]$version = "new")

# 1. Сборка приложения
go build -o app.exe

# 2. Определение активного порта
$activePort = 3000
$nextPort = 3001

if (Get-Process -Name "app" -ErrorAction SilentlyContinue) {
    $activeProcess = Get-Process app | Where-Object { $_.MainWindowTitle -match ":3000" }
    if ($activeProcess) {
        $activePort = 3000
        $nextPort = 3001
    }
    else {
        $activePort = 3001
        $nextPort = 3000
    }
}

# 3. Запуск новой версии на неактивном порту
$env:PORT = $nextPort
Start-Process -FilePath ".\app.exe" -WindowStyle Hidden -PassThru
Start-Sleep -Seconds 5  # Ожидаем инициализации

# 4. Переключение Nginx
nginx -s reload

# 5. Остановка старой версии
if ($activePort -ne 0) {
    Get-Process -Name "app" | Where-Object { 
        $_.MainWindowTitle -match ":${activePort}" 
    } | Stop-Process -Force
}

Write-Host "Deployment complete! New version on port $nextPort"