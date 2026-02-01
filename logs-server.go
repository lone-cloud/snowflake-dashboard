//go:build !dev

package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func main() {
	fs := http.FileServer(http.Dir("/usr/share/nginx/html"))
	http.Handle("/", addSecurityHeaders(fs))

	http.HandleFunc("/internal/nat", handleNAT)
	http.HandleFunc("/internal/logs", handleLogs)
	http.HandleFunc("/internal/metrics", handleMetrics)

	log.Println("Server running on port 8888")
	server := &http.Server{
		Addr:         ":8888",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}

func getDockerLogs() (string, error) {
	client := &http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", "/var/run/docker.sock")
			},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost/containers/snowflake-proxy/logs?stdout=true&stderr=true&tail=500", nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			log.Printf("Failed to close response body: %v", closeErr)
		}
	}()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Docker API error: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	return string(body), err
}

func handleNAT(w http.ResponseWriter, _ *http.Request) {
	output, err := getDockerLogs()
	if err != nil {
		log.Printf("Failed to fetch logs: %v", err)
		http.Error(w, "Logs unavailable", 500)
		return
	}

	lines := strings.Split(output, "\n")
	natRegex := regexp.MustCompile(`\bNAT type:\s*([^\r\n]+)\s*$`)

	natType := "Unknown"
	for i := len(lines) - 1; i >= 0; i-- {
		if match := natRegex.FindStringSubmatch(lines[i]); match != nil {
			natType = strings.TrimSpace(match[1])
			break
		}
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if _, err := fmt.Fprint(w, natType); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func handleLogs(w http.ResponseWriter, _ *http.Request) {
	output, err := getDockerLogs()
	if err != nil {
		log.Printf("Failed to fetch logs: %v", err)
		http.Error(w, "Logs unavailable", 500)
		return
	}

	var filtered []string
	for _, line := range strings.Split(output, "\n") {
		if strings.Contains(line, "In the last") {
			filtered = append(filtered, line)
		}
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if _, err := fmt.Fprint(w, strings.Join(filtered, "\n")); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func handleMetrics(w http.ResponseWriter, _ *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:9999/internal/metrics", nil)
	if err != nil {
		log.Printf("Failed to create request: %v", err)
		http.Error(w, "Metrics unavailable", 500)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Failed to fetch metrics: %v", err)
		http.Error(w, "Metrics unavailable", 500)
		return
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			log.Printf("Failed to close response body: %v", closeErr)
		}
	}()

	if resp.StatusCode != 200 {
		http.Error(w, "Metrics unavailable", resp.StatusCode)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if _, err := io.Copy(w, resp.Body); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func addSecurityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Referrer-Policy", "no-referrer")
		w.Header().Set("Permissions-Policy", "geolocation=(), microphone=(), camera=()")
		w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'; style-src 'self'; connect-src 'self'")
		next.ServeHTTP(w, r)
	})
}
