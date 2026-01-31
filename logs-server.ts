Bun.serve({
	port: 3001,
	async fetch(req) {
		const url = new URL(req.url);
		const path = url.pathname;

		const socketPath = "/var/run/docker.sock";
		const containerName = "snowflake-proxy";

		try {
			const response = await fetch(
				`http://localhost/containers/${containerName}/logs?stdout=true&stderr=true&tail=500`,
				{
					unix: socketPath,
				},
			);

			if (!response.ok) {
				throw new Error(`Docker API error: ${response.status}`);
			}

			const output = await response.text();

			if (path === "/internal/nat") {
				const natMatch = output
					.split("\n")
					.reverse()
					.map((line) => line.match(/\bNAT type:\s*([^\r\n]+)\s*$/))
					.find(Boolean);

				const natType = natMatch?.[1]?.trim() || "Unknown";
				return new Response(natType, {
					headers: { "Content-Type": "text/plain" },
				});
			}

			if (path === "/internal/logs" || path === "/") {
				const logs = output
					.split("\n")
					.filter((line) => line.includes("In the last"))
					.join("\n");

				return new Response(logs, {
					headers: { "Content-Type": "text/plain" },
				});
			}

			return new Response("Not Found", {
				status: 404,
				headers: { "Content-Type": "text/plain" },
			});
		} catch (error) {
			console.error("Failed to fetch logs:", error);
			return new Response("Logs unavailable", {
				status: 500,
				headers: { "Content-Type": "text/plain" },
			});
		}
	},
});

console.log("Logs server running on port 3001");
