package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)

// BanList manages blocked IP addresses
type BanList struct {
	ips map[string]bool
	mu  sync.RWMutex
}

func (bl *BanList) IsBanned(ip string) bool {
	bl.mu.RLock()
	defer bl.mu.RUnlock()
	host, _, _ := net.SplitHostPort(ip)
	if host == "::1" {
		host = "127.0.0.1"
	}
	return bl.ips[host]
}

func (bl *BanList) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		host, _, _ := net.SplitHostPort(ip)
		if host == "::1" {
			host = "127.0.0.1"
		}
		
		if bl.IsBanned(r.RemoteAddr) {
			now := time.Now().Format("15:04:05")
			fmt.Printf("[%s] 🚫 BANNED ATTEMPT: %s was denied access\n", now, host)
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprint(w, "Access Denied: Your IP is blacklisted.")
			return
		}
		next.ServeHTTP(w, r)
	})
}

// RateLimiter stores request timestamps for limiting
type RateLimiter struct {
	ips map[string][]time.Time
	mu  sync.Mutex
}

func (rl *RateLimiter) Limit(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		if ip == "::1" {
			ip = "127.0.0.1"
		}

		rl.mu.Lock()
		defer rl.mu.Unlock()

		now := time.Now()
		var recent []time.Time
		for _, t := range rl.ips[ip] {
			if now.Sub(t) < 30*time.Second {
				recent = append(recent, t)
			}
		}

		if len(recent) >= 10 {
			fmt.Printf("[%s] ⚠️  RATE LIMIT: %s is calling too fast (%d req/30s)\n", now.Format("15:04:05"), ip, len(recent))
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Too many requests. Please try again in 30 seconds.",
			})
			return
		}

		recent = append(recent, now)
		rl.ips[ip] = recent

		fmt.Printf("[%s] ✅ API CALL: %s -> %s\n", now.Format("15:04:05"), ip, r.URL.Path)
		next(w, r)
	}
}

type TimerData struct {
	Label       string `json:"label"`
	StartDate   string `json:"start_date"`
	DaysElapsed int    `json:"days_elapsed"`
}

func calculateDays(startDateStr string) int {
	startDate, _ := time.Parse(time.RFC3339, startDateStr)
	diff := time.Since(startDate)
	return int(diff.Hours() / 24)
}

func main() {
	banList := &BanList{
		ips: map[string]bool{
			"1.2.3.4": true,
			"127.0.0.2": true,
		},
	}

	limiter := &RateLimiter{
		ips: make(map[string][]time.Time),
	}

	mux := http.NewServeMux()

	// API 1: Milestone 1
	mux.HandleFunc("/api/timer1", limiter.Limit(func(w http.ResponseWriter, r *http.Request) {
		date := "2025-11-24T00:00:00Z"
		data := TimerData{
			Label:       "Break up with lcf",
			StartDate:   date,
			DaysElapsed: calculateDays(date),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}))

	// API 2: Milestone 2
	mux.HandleFunc("/api/timer2", limiter.Limit(func(w http.ResponseWriter, r *http.Request) {
		date := "2025-11-08T00:00:00Z"
		data := TimerData{
			Label:       "First time lcf said \"nho hks\"",
			StartDate:   date,
			DaysElapsed: calculateDays(date),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}))

	mux.Handle("/", http.FileServer(http.Dir("./static")))

	finalHandler := banList.Middleware(mux)

	go func() {
		restartAfter := 12 * time.Hour
		timer := time.NewTimer(restartAfter)
		<-timer.C
		fmt.Printf("[%s] 🔄 12 hours reached. Auto-exiting for restart...\n", time.Now().Format("15:04:05"))
		time.Sleep(2 * time.Second)
		os.Exit(0)
	}()

	fmt.Println("Server running at http://127.0.0.1:3939")
	fmt.Println("Security: Banlist & Rate Limit (10 req/30s) enabled")
	fmt.Println("Schedule: Auto-restart every 12 hours.")
	
	if err := http.ListenAndServe("0.0.0.0:3939", finalHandler); err != nil {
		fmt.Printf("Server Error: %v\n", err)
	}
}
