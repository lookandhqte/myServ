#!/bin/bash
set -e  # Выходить при любой ошибке

# 1. Определение активного порта через PID
active_pid=$(lsof -t -i:3000 || true)
if [ -n "$active_pid" ]; then
  active_port=3000
  next_port=3001
else
  active_port=3001
  next_port=3000
fi

# 2. Остановка приложения на следующем порту (если запущено)
next_pid=$(lsof -t -i:$next_port || true)
if [ -n "$next_pid" ]; then
  kill -9 $next_pid
  sleep 1
fi

# 3. Проверка наличия нового бинарника
if [ ! -f "app_new" ]; then
  echo "Error: app_new file not found!"
  exit 1
fi

# 4. Замена бинарника и запуск нового экземпляра
mv -f app_new app
chmod +x app
export PORT=$next_port
nohup ./app > app_$next_port.log 2>&1 &
echo "Started new app on port $next_port (PID: $!)"

# 5. Ожидание запуска приложения
sleep 3
if ! ss -tuln | grep ":$next_port " > /dev/null; then
  echo "Error: App failed to start on port $next_port!"
  exit 1
fi

# 6. Обновление Nginx конфига
echo "upstream backend { server 127.0.0.1:$next_port; }" | sudo tee /etc/nginx/conf.d/upstream.conf > /dev/null
sudo nginx -s reload || { echo "Nginx reload failed"; exit 1; }

# 7. Остановка старого приложения после переключения Nginx
if [ -n "$active_pid" ]; then
  echo "Stopping old app on port $active_port (PID: $active_pid)"
  kill -9 $active_pid
fi

echo "Deployment successful! New port: $next_port, Old port: $active_port"