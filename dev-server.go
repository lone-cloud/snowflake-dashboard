//go:build dev

package main

import (
	"log"
	"net/http"
	"os"
)

const mockMetrics = `# HELP tor_snowflake_proxy_connection_timeouts_total The total number of client connection attempts that failed after successful rendezvous
# TYPE tor_snowflake_proxy_connection_timeouts_total counter
tor_snowflake_proxy_connection_timeouts_total 765
# HELP tor_snowflake_proxy_connections_total The total number of successful connections handled by the snowflake proxy
# TYPE tor_snowflake_proxy_connections_total counter
tor_snowflake_proxy_connections_total{country=""} 46
tor_snowflake_proxy_connections_total{country="??"} 1
tor_snowflake_proxy_connections_total{country="AE"} 4
tor_snowflake_proxy_connections_total{country="AU"} 3
tor_snowflake_proxy_connections_total{country="BF"} 3
tor_snowflake_proxy_connections_total{country="BR"} 1
tor_snowflake_proxy_connections_total{country="BY"} 1
tor_snowflake_proxy_connections_total{country="CA"} 4
tor_snowflake_proxy_connections_total{country="CH"} 2
tor_snowflake_proxy_connections_total{country="CI"} 5
tor_snowflake_proxy_connections_total{country="CM"} 1
tor_snowflake_proxy_connections_total{country="CN"} 8
tor_snowflake_proxy_connections_total{country="CV"} 1
tor_snowflake_proxy_connections_total{country="DE"} 8
tor_snowflake_proxy_connections_total{country="DK"} 3
tor_snowflake_proxy_connections_total{country="EG"} 8
tor_snowflake_proxy_connections_total{country="ES"} 6
tor_snowflake_proxy_connections_total{country="FI"} 1
tor_snowflake_proxy_connections_total{country="FR"} 7
tor_snowflake_proxy_connections_total{country="GA"} 1
tor_snowflake_proxy_connections_total{country="GB"} 7
tor_snowflake_proxy_connections_total{country="GM"} 1
tor_snowflake_proxy_connections_total{country="IE"} 1
tor_snowflake_proxy_connections_total{country="IL"} 1
tor_snowflake_proxy_connections_total{country="IN"} 4
tor_snowflake_proxy_connections_total{country="IR"} 84
tor_snowflake_proxy_connections_total{country="KE"} 3
tor_snowflake_proxy_connections_total{country="LT"} 1
tor_snowflake_proxy_connections_total{country="LY"} 1
tor_snowflake_proxy_connections_total{country="MA"} 9
tor_snowflake_proxy_connections_total{country="MU"} 33
tor_snowflake_proxy_connections_total{country="MW"} 1
tor_snowflake_proxy_connections_total{country="NG"} 6
tor_snowflake_proxy_connections_total{country="NL"} 4
tor_snowflake_proxy_connections_total{country="PK"} 1
tor_snowflake_proxy_connections_total{country="PL"} 2
tor_snowflake_proxy_connections_total{country="RU"} 140
tor_snowflake_proxy_connections_total{country="RW"} 1
tor_snowflake_proxy_connections_total{country="SD"} 5
tor_snowflake_proxy_connections_total{country="SG"} 1
tor_snowflake_proxy_connections_total{country="SO"} 1
tor_snowflake_proxy_connections_total{country="TG"} 3
tor_snowflake_proxy_connections_total{country="TM"} 4
tor_snowflake_proxy_connections_total{country="TN"} 21
tor_snowflake_proxy_connections_total{country="TZ"} 1
tor_snowflake_proxy_connections_total{country="UA"} 2
tor_snowflake_proxy_connections_total{country="UG"} 5
tor_snowflake_proxy_connections_total{country="US"} 445
tor_snowflake_proxy_connections_total{country="ZA"} 8
tor_snowflake_proxy_connections_total{country="ZM"} 7
# HELP tor_snowflake_proxy_traffic_inbound_bytes_total The total in bound traffic by the snowflake proxy (KB)
# TYPE tor_snowflake_proxy_traffic_inbound_bytes_total counter
tor_snowflake_proxy_traffic_inbound_bytes_total 3.226803e+06
# HELP tor_snowflake_proxy_traffic_outbound_bytes_total The total out bound traffic by the snowflake proxy (KB)
# TYPE tor_snowflake_proxy_traffic_outbound_bytes_total counter
tor_snowflake_proxy_traffic_outbound_bytes_total 672663
# HELP process_network_receive_bytes_total Number of bytes received by the process over the network
# TYPE process_network_receive_bytes_total counter
process_network_receive_bytes_total 6.066517003e+09
# HELP process_network_transmit_bytes_total Number of bytes sent by the process over the network
# TYPE process_network_transmit_bytes_total counter
process_network_transmit_bytes_total 6.091533673e+09
# HELP process_resident_memory_bytes Resident memory size in bytes
# TYPE process_resident_memory_bytes gauge
process_resident_memory_bytes 1.02674432e+08
# HELP process_start_time_seconds Start time of the process since unix epoch in seconds
# TYPE process_start_time_seconds gauge
process_start_time_seconds 1.76976251641e+09
# HELP go_goroutines Number of goroutines that currently exist
# TYPE go_goroutines gauge
go_goroutines 416
# HELP process_open_fds Number of open file descriptors
# TYPE process_open_fds gauge
process_open_fds 47
# HELP process_cpu_seconds_total Total user and system CPU time spent in seconds
# TYPE process_cpu_seconds_total counter
process_cpu_seconds_total 5244.22
`

const mockLogs = `2026/01/29 19:36:17 In the last 1h0m0s, there were 150 completed successful connections. Traffic Relayed ↓ 358643 KB (99.62 KB/s), ↑ 87018 KB (24.17 KB/s).
2026/01/29 20:36:17 In the last 1h0m0s, there were 137 completed successful connections. Traffic Relayed ↓ 705340 KB (195.93 KB/s), ↑ 137061 KB (38.07 KB/s).
2026/01/29 21:36:17 In the last 1h0m0s, there were 101 completed successful connections. Traffic Relayed ↓ 489758 KB (136.04 KB/s), ↑ 106948 KB (29.71 KB/s).
2026/01/29 22:36:17 In the last 1h0m0s, there were 34 completed successful connections. Traffic Relayed ↓ 571290 KB (158.69 KB/s), ↑ 113590 KB (31.55 KB/s).
2026/01/29 23:36:17 In the last 1h0m0s, there were 31 completed successful connections. Traffic Relayed ↓ 565185 KB (157.00 KB/s), ↑ 43956 KB (12.21 KB/s).
2026/01/30 00:36:17 In the last 1h0m0s, there were 8 completed successful connections. Traffic Relayed ↓ 209866 KB (58.30 KB/s), ↑ 86463 KB (24.02 KB/s).
2026/01/30 01:36:17 In the last 1h0m0s, there were 6 completed successful connections. Traffic Relayed ↓ 150602 KB (41.83 KB/s), ↑ 72347 KB (20.10 KB/s).
2026/01/30 02:36:17 In the last 1h0m0s, there were 5 completed successful connections. Traffic Relayed ↓ 111978 KB (31.11 KB/s), ↑ 62467 KB (17.35 KB/s).
2026/01/30 03:36:17 In the last 1h0m0s, there were 8 completed successful connections. Traffic Relayed ↓ 66912 KB (18.59 KB/s), ↑ 50987 KB (14.16 KB/s).
2026/01/30 04:36:17 In the last 1h0m0s, there were 11 completed successful connections. Traffic Relayed ↓ 307124 KB (85.31 KB/s), ↑ 39730 KB (11.04 KB/s).
2026/01/30 05:36:17 In the last 1h0m0s, there were 18 completed successful connections. Traffic Relayed ↓ 219477 KB (60.97 KB/s), ↑ 23241 KB (6.46 KB/s).
2026/01/30 06:36:17 In the last 1h0m0s, there were 13 completed successful connections. Traffic Relayed ↓ 483875 KB (134.41 KB/s), ↑ 52759 KB (14.66 KB/s).
2026/01/30 07:36:17 In the last 1h0m0s, there were 18 completed successful connections. Traffic Relayed ↓ 608490 KB (169.03 KB/s), ↑ 75063 KB (20.85 KB/s).
2026/01/30 08:36:17 In the last 1h0m0s, there were 22 completed successful connections. Traffic Relayed ↓ 585820 KB (162.73 KB/s), ↑ 95537 KB (26.54 KB/s).
2026/01/30 09:41:59 In the last 1h0m0s, there were 51 completed successful connections. Traffic Relayed ↓ 528797 KB (146.89 KB/s), ↑ 90265 KB (25.07 KB/s).
2026/01/30 10:41:59 In the last 1h0m0s, there were 31 completed successful connections. Traffic Relayed ↓ 585162 KB (162.54 KB/s), ↑ 92028 KB (25.56 KB/s).
2026/01/30 11:41:59 In the last 1h0m0s, there were 97 completed successful connections. Traffic Relayed ↓ 287205 KB (79.78 KB/s), ↑ 64850 KB (18.01 KB/s).
2026/01/30 12:41:59 In the last 1h0m0s, there were 105 completed successful connections. Traffic Relayed ↓ 154081 KB (42.80 KB/s), ↑ 34662 KB (9.63 KB/s).
2026/01/30 13:41:59 In the last 1h0m0s, there were 54 completed successful connections. Traffic Relayed ↓ 330043 KB (91.68 KB/s), ↑ 58548 KB (16.26 KB/s).
2026/01/30 14:41:59 In the last 1h0m0s, there were 61 completed successful connections. Traffic Relayed ↓ 283171 KB (78.66 KB/s), ↑ 51404 KB (14.28 KB/s).
2026/01/30 15:41:59 In the last 1h0m0s, there were 75 completed successful connections. Traffic Relayed ↓ 232058 KB (64.46 KB/s), ↑ 58970 KB (16.38 KB/s).
2026/01/30 16:41:59 In the last 1h0m0s, there were 95 completed successful connections. Traffic Relayed ↓ 283368 KB (78.71 KB/s), ↑ 82115 KB (22.81 KB/s).
2026/01/30 17:41:59 In the last 1h0m0s, there were 126 completed successful connections. Traffic Relayed ↓ 251105 KB (69.75 KB/s), ↑ 74885 KB (20.80 KB/s).
2026/01/30 18:41:59 In the last 1h0m0s, there were 110 completed successful connections. Traffic Relayed ↓ 291813 KB (81.06 KB/s), ↑ 64936 KB (18.04 KB/s).`

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	http.HandleFunc("/internal/metrics", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write([]byte(mockMetrics))
	})

	http.HandleFunc("/internal/logs", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write([]byte(mockLogs))
	})

	http.HandleFunc("/internal/nat", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write([]byte("restricted"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Dev server running at http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
