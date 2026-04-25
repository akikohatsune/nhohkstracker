#!/bin/bash
while true; do
    echo "[$(date +%T)] 🚀 Đang khởi động Server..."
    go run main.go
    echo "[$(date +%T)] 💤 Server đã dừng. Đang chờ 3 giây trước khi khởi động lại..."
    sleep 3
done
