# Hướng dẫn cài đặt KomeTimeLap

Dự án này sử dụng ngôn ngữ Go (Golang) và không yêu cầu thư viện bên ngoài (zero dependencies).

## 1. Yêu cầu hệ thống
- Đã cài đặt **Go** (phiên bản 1.16 trở lên). Tải tại: https://go.dev/dl/

## 2. Cách cài đặt trên máy mới
1. Tải toàn bộ thư mục này về máy.
2. Mở terminal tại thư mục dự án.
3. Chạy lệnh sau để đảm bảo môi trường Go đã sẵn sàng:
   ```bash
   go mod tidy
   ```

## 3. Cách chạy ứng dụng
### Trên Linux / macOS / Android (Termux):
Sử dụng script tự động để có tính năng tự khởi động lại sau 12h:
```bash
chmod +x run.sh
./run.sh
```

### Trên Windows:
Bạn có thể chạy trực tiếp:
```powershell
go run main.go
```
*(Lưu ý: Trên Windows nếu muốn tự restart sau 12h, bạn cần tạo file `.bat` tương đương hoặc chạy thủ công).*

## 4. Cấu trúc thư mục
- `main.go`: Mã nguồn server backend (Go).
- `static/`: Chứa giao diện frontend (HTML, CSS, JS, Favicon).
- `run.sh`: Script quản lý vòng lặp chạy server.
- `go.mod`: File quản lý module của Go.
