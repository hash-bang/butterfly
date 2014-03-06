#!/bin/bash
echo "Killing existing server..."
pkill -f 'python ./butterfly.server.py'

echo "Starting server..."
chmod +x butterfly.server.py
./butterfly.server.py &

echo "Opening terminal..."
google-chrome --app=http://localhost:57575
