# nhohkstracker

A professional real-time "nho hks" tracking application featuring a modern Glassmorphism UI and a robust Go backend.

## 🚀 Features

- **Modern Glassmorphism UI**: High-end aesthetic with frosted glass effects, smooth transitions, and a clean layout.
- **Go Backend**: High-performance server built with Go (Golang) using only the standard library for maximum efficiency.
- **Real-time Tracking**: Dynamic countdown/count-up timers synced directly from the server via API.
- **Security & Protection**:
    - **Rate Limiting**: Built-in protection (10 req/30s) to prevent API abuse.
    - **IP Banlist**: Comprehensive blacklisting system to block malicious actors.
    - **IPv4 Optimized**: Clear IPv4 logging and server binding.
- **Reliability**:
    - **Auto-restart System**: Includes a dedicated shell script for continuous uptime with 12-hour periodic restarts.
- **Responsive Design**: Fully optimized for mobile, tablet, and desktop viewing.
- **API Documentation**: Built-in, easy-to-read technical guide for developers.

## 🛠️ Technologies

- **Backend**: Go (Golang)
- **Frontend**: HTML5, CSS3 (Vanilla), JavaScript (Vanilla)
- **Typography**: Montserrat (Google Fonts)
- **Server**: Custom HTTP Multiplexer

## 📁 Project Structure

```text
.
├── main.go             # Go backend server logic
├── go.mod              # Go module definition
├── run.sh              # Auto-restart & management script
├── README.md           # Project documentation
├── static/
│   ├── index.html      # Main application entry point
│   ├── style.css       # Modern Glassmorphism styles
│   ├── script.js       # Frontend real-time logic
│   ├── assets/
│   │   └── tuyu1.png   # Background artwork
│   └── misc/
│       └── docs.html   # API technical documentation
```

## ⚙️ Setup & Installation

### Prerequisites
- **Go** (1.16 or higher installed)
- **Bash** (for the auto-restart script)

### Running the Application

1. **Clone the repository**:
   ```bash
   git clone https://github.com/akikohatsune/nhohkstracker.git
   cd nhohkstracker
   ```

2. **Run with Auto-restart (Recommended)**:
   ```bash
   chmod +x run.sh
   ./run.sh
   ```

3. **Run manually**:
   ```bash
   go run main.go
   ```

The application will be available at `http://127.0.0.1:3939`.

## 📍 API Endpoints

- `GET /api/timer1`: Returns data for the first milestone (Label & Start Date).
- `GET /api/timer2`: Returns data for the second milestone (Label & Start Date).

## 📄 License

This project is licensed under the CC0 1.0 Universal - see the [LICENSE](LICENSE) file for details.

---
*dcm lcf suot ngay nho hks*
