#!/bin/bash

# Determine active port
if ss -tuln | grep ':3000 ' > /dev/null; then
  active_port=3000
  next_port=3001
else
  active_port=3001
  next_port=3000
fi

# Stop existing app on next port (if any)
if ss -tuln | grep ":$next_port " > /dev/null; then
  pkill -f "app.*:$next_port"
fi

# Проверка существования app_new
if [ ! -f "app_new" ]; then
    echo "Error: app_new file not found!"
    exit 1
fi

# Replace binary and start new instance
mv -f app_new app
chmod +x app
export PORT=$next_port
nohup ./app > app.log 2>&1 &

# Update Nginx config
echo "upstream backend { server 127.0.0.1:$next_port; }" | sudo tee /etc/nginx/conf.d/upstream.conf > /dev/null
sudo nginx -s reload

# Stop old app after reload
if [ "$active_port" != "0" ]; then
  pkill -f "app.*:$active_port"
fi

echo "Deployed on port $next_port. Old port: $active_port"