const mockMetrics = `# HELP tor_snowflake_proxy_connections_total The total number of successful connections handled by the snowflake proxy
# TYPE tor_snowflake_proxy_connections_total counter
tor_snowflake_proxy_connections_total{country=""} 3
tor_snowflake_proxy_connections_total{country="CI"} 1
tor_snowflake_proxy_connections_total{country="ES"} 1
tor_snowflake_proxy_connections_total{country="FR"} 1
tor_snowflake_proxy_connections_total{country="GB"} 2
tor_snowflake_proxy_connections_total{country="IR"} 9
tor_snowflake_proxy_connections_total{country="IT"} 1
tor_snowflake_proxy_connections_total{country="JP"} 1
tor_snowflake_proxy_connections_total{country="KR"} 1
tor_snowflake_proxy_connections_total{country="MA"} 1
tor_snowflake_proxy_connections_total{country="MU"} 2
tor_snowflake_proxy_connections_total{country="NG"} 2
tor_snowflake_proxy_connections_total{country="PG"} 1
tor_snowflake_proxy_connections_total{country="PT"} 1
tor_snowflake_proxy_connections_total{country="RU"} 7
tor_snowflake_proxy_connections_total{country="SG"} 1
tor_snowflake_proxy_connections_total{country="TN"} 3
tor_snowflake_proxy_connections_total{country="US"} 41
# HELP tor_snowflake_proxy_connection_timeouts_total The total number of client connection attempts that failed after successful rendezvous
# TYPE tor_snowflake_proxy_connection_timeouts_total counter
tor_snowflake_proxy_connection_timeouts_total 91
# HELP tor_snowflake_proxy_traffic_inbound_bytes_total The total in bound traffic by the snowflake proxy (KB)
# TYPE tor_snowflake_proxy_traffic_inbound_bytes_total counter
tor_snowflake_proxy_traffic_inbound_bytes_total 490539
# HELP tor_snowflake_proxy_traffic_outbound_bytes_total The total out bound traffic by the snowflake proxy (KB)
# TYPE tor_snowflake_proxy_traffic_outbound_bytes_total counter
tor_snowflake_proxy_traffic_outbound_bytes_total 110937
# HELP process_network_receive_bytes_total Number of bytes received by the process over the network
# TYPE process_network_receive_bytes_total counter
process_network_receive_bytes_total 1.78163309e+09
# HELP process_network_transmit_bytes_total Number of bytes sent by the process over the network
# TYPE process_network_transmit_bytes_total counter
process_network_transmit_bytes_total 1.780049424e+09
# HELP process_resident_memory_bytes Resident memory size in bytes
# TYPE process_resident_memory_bytes gauge
process_resident_memory_bytes 9.8537472e+07
# HELP process_start_time_seconds Start time of the process since unix epoch in seconds
# TYPE process_start_time_seconds gauge
process_start_time_seconds 1.7695821806e+09
# HELP go_goroutines Number of goroutines that currently exist
# TYPE go_goroutines gauge
go_goroutines 571
# HELP process_open_fds Number of open file descriptors
# TYPE process_open_fds gauge
process_open_fds 65
# HELP process_cpu_seconds_total Total user and system CPU time spent in seconds
# TYPE process_cpu_seconds_total counter
process_cpu_seconds_total 1395.02
`;

const server = Bun.serve({
	port: 3000,
	async fetch(req) {
		const url = new URL(req.url);

		if (url.pathname === "/internal/metrics") {
			return new Response(mockMetrics, {
				headers: { "Content-Type": "text/plain" },
			});
		}

		if (url.pathname === "/internal/logs") {
			const mockLogs = `2026/01/29 06:36:23 In the last 1h0m0s, there were 28 completed successful connections. Traffic Relayed ↓ 198456 KB (55.13 KB/s), ↑ 38921 KB (10.81 KB/s).
2026/01/29 05:36:23 In the last 1h0m0s, there were 35 completed successful connections. Traffic Relayed ↓ 251032 KB (69.73 KB/s), ↑ 47852 KB (13.29 KB/s).
2026/01/29 04:36:23 In the last 1h0m0s, there were 42 completed successful connections. Traffic Relayed ↓ 312847 KB (86.90 KB/s), ↑ 59234 KB (16.45 KB/s).
2026/01/29 03:36:23 In the last 1h0m0s, there were 31 completed successful connections. Traffic Relayed ↓ 221543 KB (61.54 KB/s), ↑ 41876 KB (11.63 KB/s).
2026/01/29 02:36:23 In the last 1h0m0s, there were 19 completed successful connections. Traffic Relayed ↓ 134219 KB (37.28 KB/s), ↑ 25183 KB (6.99 KB/s).
2026/01/29 01:36:23 In the last 1h0m0s, there were 47 completed successful connections. Traffic Relayed ↓ 341876 KB (94.97 KB/s), ↑ 64392 KB (17.89 KB/s).
2026/01/29 00:36:23 In the last 1h0m0s, there were 38 completed successful connections. Traffic Relayed ↓ 267543 KB (74.32 KB/s), ↑ 50124 KB (13.92 KB/s).
2026/01/28 23:36:23 In the last 1h0m0s, there were 52 completed successful connections. Traffic Relayed ↓ 389234 KB (108.12 KB/s), ↑ 72841 KB (20.23 KB/s).
2026/01/28 22:36:23 In the last 1h0m0s, there were 44 completed successful connections. Traffic Relayed ↓ 298712 KB (82.98 KB/s), ↑ 56234 KB (15.62 KB/s).
2026/01/28 21:36:23 In the last 1h0m0s, there were 29 completed successful connections. Traffic Relayed ↓ 203487 KB (56.52 KB/s), ↑ 38291 KB (10.64 KB/s).
2026/01/28 20:36:23 In the last 1h0m0s, there were 36 completed successful connections. Traffic Relayed ↓ 254098 KB (70.58 KB/s), ↑ 47659 KB (13.24 KB/s).
2026/01/28 19:36:23 In the last 1h0m0s, there were 41 completed successful connections. Traffic Relayed ↓ 289543 KB (80.43 KB/s), ↑ 54327 KB (15.09 KB/s).
2026/01/28 18:36:23 In the last 1h0m0s, there were 33 completed successful connections. Traffic Relayed ↓ 232109 KB (64.47 KB/s), ↑ 43652 KB (12.13 KB/s).
2026/01/28 17:36:23 In the last 1h0m0s, there were 27 completed successful connections. Traffic Relayed ↓ 189321 KB (52.59 KB/s), ↑ 35614 KB (9.89 KB/s).
2026/01/28 16:36:23 In the last 1h0m0s, there were 49 completed successful connections. Traffic Relayed ↓ 356712 KB (99.09 KB/s), ↑ 67089 KB (18.64 KB/s).
2026/01/28 15:36:23 In the last 1h0m0s, there were 39 completed successful connections. Traffic Relayed ↓ 275432 KB (76.51 KB/s), ↑ 51743 KB (14.37 KB/s).
2026/01/28 14:36:23 In the last 1h0m0s, there were 45 completed successful connections. Traffic Relayed ↓ 318765 KB (88.55 KB/s), ↑ 59876 KB (16.63 KB/s).
2026/01/28 13:36:23 In the last 1h0m0s, there were 34 completed successful connections. Traffic Relayed ↓ 239654 KB (66.57 KB/s), ↑ 45021 KB (12.51 KB/s).
2026/01/28 12:36:23 In the last 1h0m0s, there were 26 completed successful connections. Traffic Relayed ↓ 182347 KB (50.65 KB/s), ↑ 34298 KB (9.53 KB/s).
2026/01/28 11:36:23 In the last 1h0m0s, there were 51 completed successful connections. Traffic Relayed ↓ 367891 KB (102.19 KB/s), ↑ 69123 KB (19.20 KB/s).`;
			return new Response(mockLogs, {
				headers: { "Content-Type": "text/plain" },
			});
		}

		const filePath = url.pathname === "/" ? "./index.html" : `.${url.pathname}`;
		const file = Bun.file(filePath);

		if (await file.exists()) {
			return new Response(file);
		}

		return new Response("Not Found", { status: 404 });
	},
});

console.log(`Dev server running at http://localhost:${server.port}`);
