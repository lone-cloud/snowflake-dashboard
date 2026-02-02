//go:build dev

package main

import (
	"log"
	"net/http"
	"os"
)

const mockMetrics = `# HELP tor_snowflake_proxy_connection_timeouts_total The total number of client connection attempts that failed after successful rendezvous. Note that failures can occur for reasons outside of the proxy's control, such as the client's NAT and censorship situation.
# TYPE tor_snowflake_proxy_connection_timeouts_total counter
tor_snowflake_proxy_connection_timeouts_total 1649
# HELP tor_snowflake_proxy_connections_total The total number of successful connections handled by the snowflake proxy
# TYPE tor_snowflake_proxy_connections_total counter
tor_snowflake_proxy_connections_total{country=""} 110
tor_snowflake_proxy_connections_total{country="??"} 11
tor_snowflake_proxy_connections_total{country="AE"} 4
tor_snowflake_proxy_connections_total{country="AM"} 1
tor_snowflake_proxy_connections_total{country="AO"} 1
tor_snowflake_proxy_connections_total{country="AR"} 1
tor_snowflake_proxy_connections_total{country="AT"} 3
tor_snowflake_proxy_connections_total{country="AU"} 9
tor_snowflake_proxy_connections_total{country="AZ"} 1
tor_snowflake_proxy_connections_total{country="BD"} 2
tor_snowflake_proxy_connections_total{country="BE"} 1
tor_snowflake_proxy_connections_total{country="BF"} 2
tor_snowflake_proxy_connections_total{country="BG"} 1
tor_snowflake_proxy_connections_total{country="BR"} 1
tor_snowflake_proxy_connections_total{country="BY"} 3
tor_snowflake_proxy_connections_total{country="CA"} 4
tor_snowflake_proxy_connections_total{country="CD"} 1
tor_snowflake_proxy_connections_total{country="CG"} 1
tor_snowflake_proxy_connections_total{country="CH"} 3
tor_snowflake_proxy_connections_total{country="CI"} 4
tor_snowflake_proxy_connections_total{country="CN"} 15
tor_snowflake_proxy_connections_total{country="CO"} 1
tor_snowflake_proxy_connections_total{country="CY"} 1
tor_snowflake_proxy_connections_total{country="DE"} 11
tor_snowflake_proxy_connections_total{country="DK"} 1
tor_snowflake_proxy_connections_total{country="EG"} 7
tor_snowflake_proxy_connections_total{country="EU"} 1
tor_snowflake_proxy_connections_total{country="FI"} 6
tor_snowflake_proxy_connections_total{country="FR"} 10
tor_snowflake_proxy_connections_total{country="GA"} 3
tor_snowflake_proxy_connections_total{country="GB"} 16
tor_snowflake_proxy_connections_total{country="GH"} 3
tor_snowflake_proxy_connections_total{country="GR"} 2
tor_snowflake_proxy_connections_total{country="HK"} 1
tor_snowflake_proxy_connections_total{country="HN"} 1
tor_snowflake_proxy_connections_total{country="ID"} 1
tor_snowflake_proxy_connections_total{country="IL"} 1
tor_snowflake_proxy_connections_total{country="IN"} 6
tor_snowflake_proxy_connections_total{country="IR"} 169
tor_snowflake_proxy_connections_total{country="IT"} 4
tor_snowflake_proxy_connections_total{country="JP"} 2
tor_snowflake_proxy_connections_total{country="KE"} 5
tor_snowflake_proxy_connections_total{country="KH"} 1
tor_snowflake_proxy_connections_total{country="KR"} 1
tor_snowflake_proxy_connections_total{country="LT"} 1
tor_snowflake_proxy_connections_total{country="LV"} 2
tor_snowflake_proxy_connections_total{country="LY"} 1
tor_snowflake_proxy_connections_total{country="MA"} 32
tor_snowflake_proxy_connections_total{country="MG"} 1
tor_snowflake_proxy_connections_total{country="ML"} 1
tor_snowflake_proxy_connections_total{country="MO"} 2
tor_snowflake_proxy_connections_total{country="MU"} 66
tor_snowflake_proxy_connections_total{country="MW"} 5
tor_snowflake_proxy_connections_total{country="MX"} 3
tor_snowflake_proxy_connections_total{country="MY"} 3
tor_snowflake_proxy_connections_total{country="NG"} 9
tor_snowflake_proxy_connections_total{country="NL"} 2
tor_snowflake_proxy_connections_total{country="PE"} 1
tor_snowflake_proxy_connections_total{country="PH"} 2
tor_snowflake_proxy_connections_total{country="PK"} 1
tor_snowflake_proxy_connections_total{country="PL"} 3
tor_snowflake_proxy_connections_total{country="RE"} 1
tor_snowflake_proxy_connections_total{country="RO"} 3
tor_snowflake_proxy_connections_total{country="RS"} 1
tor_snowflake_proxy_connections_total{country="RU"} 327
tor_snowflake_proxy_connections_total{country="RW"} 1
tor_snowflake_proxy_connections_total{country="SA"} 2
tor_snowflake_proxy_connections_total{country="SD"} 8
tor_snowflake_proxy_connections_total{country="SE"} 2
tor_snowflake_proxy_connections_total{country="SG"} 2
tor_snowflake_proxy_connections_total{country="SK"} 1
tor_snowflake_proxy_connections_total{country="SO"} 1
tor_snowflake_proxy_connections_total{country="TH"} 1
tor_snowflake_proxy_connections_total{country="TM"} 5
tor_snowflake_proxy_connections_total{country="TN"} 60
tor_snowflake_proxy_connections_total{country="TR"} 1
tor_snowflake_proxy_connections_total{country="TW"} 1
tor_snowflake_proxy_connections_total{country="UA"} 5
tor_snowflake_proxy_connections_total{country="UG"} 8
tor_snowflake_proxy_connections_total{country="US"} 946
tor_snowflake_proxy_connections_total{country="VN"} 1
tor_snowflake_proxy_connections_total{country="ZA"} 23
tor_snowflake_proxy_connections_total{country="ZM"} 10
# HELP tor_snowflake_proxy_traffic_inbound_bytes_total The total in bound traffic by the snowflake proxy (KB)
# TYPE tor_snowflake_proxy_traffic_inbound_bytes_total counter
tor_snowflake_proxy_traffic_inbound_bytes_total 2.7179552e+07
# HELP tor_snowflake_proxy_traffic_outbound_bytes_total The total out bound traffic by the snowflake proxy (KB)
# TYPE tor_snowflake_proxy_traffic_outbound_bytes_total counter
tor_snowflake_proxy_traffic_outbound_bytes_total 4.204604e+06
# HELP process_network_receive_bytes_total Number of bytes received by the process over the network
# TYPE process_network_receive_bytes_total counter
process_network_receive_bytes_total 4.1371771232e+10
# HELP process_network_transmit_bytes_total Number of bytes sent by the process over the network
# TYPE process_network_transmit_bytes_total counter
process_network_transmit_bytes_total 4.1382248284e+10
# HELP process_resident_memory_bytes Resident memory size in bytes
# TYPE process_resident_memory_bytes gauge
process_resident_memory_bytes 9.5772672e+07
# HELP process_start_time_seconds Start time of the process since unix epoch in seconds
# TYPE process_start_time_seconds gauge
process_start_time_seconds 1.76985915868e+09
# HELP go_goroutines Number of goroutines that currently exist
# TYPE go_goroutines gauge
go_goroutines 360
# HELP process_open_fds Number of open file descriptors
# TYPE process_open_fds gauge
process_open_fds 45
# HELP process_cpu_seconds_total Total user and system CPU time spent in seconds
# TYPE process_cpu_seconds_total counter
process_cpu_seconds_total 27268.99
`

const mockLogs = `2026/01/31 12:32:40 In the last 1h0m0s, there were 134 completed successful connections. Traffic Relayed ↓ 183457 KB (50.96 KB/s), ↑ 37294 KB (10.36 KB/s).
2026/01/31 13:32:40 In the last 1h0m0s, there were 114 completed successful connections. Traffic Relayed ↓ 352507 KB (97.92 KB/s), ↑ 101225 KB (28.12 KB/s).
2026/01/31 14:32:40 In the last 1h0m0s, there were 106 completed successful connections. Traffic Relayed ↓ 857188 KB (238.11 KB/s), ↑ 119983 KB (33.33 KB/s).
2026/01/31 15:32:40 In the last 1h0m0s, there were 93 completed successful connections. Traffic Relayed ↓ 561858 KB (156.07 KB/s), ↑ 119705 KB (33.25 KB/s).
2026/01/31 16:32:40 In the last 1h0m0s, there were 86 completed successful connections. Traffic Relayed ↓ 733619 KB (203.78 KB/s), ↑ 127004 KB (35.28 KB/s).
2026/01/31 17:32:40 In the last 1h0m0s, there were 128 completed successful connections. Traffic Relayed ↓ 597328 KB (165.92 KB/s), ↑ 121182 KB (33.66 KB/s).
2026/01/31 18:32:40 In the last 1h0m0s, there were 121 completed successful connections. Traffic Relayed ↓ 348280 KB (96.74 KB/s), ↑ 59854 KB (16.63 KB/s).
2026/01/31 19:32:40 In the last 1h0m0s, there were 120 completed successful connections. Traffic Relayed ↓ 372636 KB (103.51 KB/s), ↑ 50047 KB (13.90 KB/s).
2026/01/31 20:32:40 In the last 1h0m0s, there were 113 completed successful connections. Traffic Relayed ↓ 474974 KB (131.94 KB/s), ↑ 93429 KB (25.95 KB/s).
2026/01/31 21:32:40 In the last 1h0m0s, there were 48 completed successful connections. Traffic Relayed ↓ 729185 KB (202.55 KB/s), ↑ 92820 KB (25.78 KB/s).
2026/01/31 22:32:40 In the last 1h0m0s, there were 16 completed successful connections. Traffic Relayed ↓ 1187745 KB (329.93 KB/s), ↑ 113682 KB (31.58 KB/s).
2026/01/31 23:32:40 In the last 1h0m0s, there were 7 completed successful connections. Traffic Relayed ↓ 715292 KB (198.69 KB/s), ↑ 59830 KB (16.62 KB/s).
2026/02/01 00:32:40 In the last 1h0m0s, there were 8 completed successful connections. Traffic Relayed ↓ 517076 KB (143.63 KB/s), ↑ 32023 KB (8.90 KB/s).
2026/02/01 01:32:40 In the last 1h0m0s, there were 5 completed successful connections. Traffic Relayed ↓ 536223 KB (148.95 KB/s), ↑ 35168 KB (9.77 KB/s).
2026/02/01 02:32:40 In the last 1h0m0s, there were 9 completed successful connections. Traffic Relayed ↓ 635409 KB (176.50 KB/s), ↑ 36176 KB (10.05 KB/s).
2026/02/01 03:32:40 In the last 1h0m0s, there were 12 completed successful connections. Traffic Relayed ↓ 602633 KB (167.40 KB/s), ↑ 36340 KB (10.09 KB/s).
2026/02/01 04:32:40 In the last 1h0m0s, there were 22 completed successful connections. Traffic Relayed ↓ 588310 KB (163.42 KB/s), ↑ 42610 KB (11.84 KB/s).
2026/02/01 05:32:40 In the last 1h0m0s, there were 19 completed successful connections. Traffic Relayed ↓ 1207049 KB (335.29 KB/s), ↑ 87281 KB (24.24 KB/s).
2026/02/01 06:32:40 In the last 1h0m0s, there were 19 completed successful connections. Traffic Relayed ↓ 862143 KB (239.48 KB/s), ↑ 84283 KB (23.41 KB/s).
2026/02/01 07:32:40 In the last 1h0m0s, there were 25 completed successful connections. Traffic Relayed ↓ 994962 KB (276.38 KB/s), ↑ 87620 KB (24.34 KB/s).
2026/02/01 08:32:40 In the last 1h0m0s, there were 27 completed successful connections. Traffic Relayed ↓ 821964 KB (228.32 KB/s), ↑ 105267 KB (29.24 KB/s).
2026/02/01 09:32:40 In the last 1h0m0s, there were 28 completed successful connections. Traffic Relayed ↓ 811943 KB (225.54 KB/s), ↑ 91503 KB (25.42 KB/s).
2026/02/01 10:32:40 In the last 1h0m0s, there were 36 completed successful connections. Traffic Relayed ↓ 1523789 KB (423.27 KB/s), ↑ 108411 KB (30.11 KB/s).
2026/02/01 11:32:40 In the last 1h0m0s, there were 19 completed successful connections. Traffic Relayed ↓ 1858789 KB (516.33 KB/s), ↑ 129278 KB (35.91 KB/s).`

func main() {
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	http.HandleFunc("/api/metrics", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(mockMetrics))
	})

	http.HandleFunc("/api/logs", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(mockLogs))
	})

	http.HandleFunc("/api/nat", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("restricted"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Dev server running at http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
