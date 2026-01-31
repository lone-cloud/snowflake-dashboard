Bun.serve({
	port: 3001,
	async fetch(req) {
		const url = new URL(req.url);
		const path = url.pathname;

		const proc = Bun.spawn([
			"docker",
			"logs",
			"--tail",
			"500",
			"snowflake-proxy",
		]);
		const output = await new Response(proc.stdout).text();

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
	},
});

console.log("Logs server running on port 3001");
