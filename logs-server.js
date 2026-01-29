Bun.serve({
	port: 3001,
	async fetch() {
		try {
			const proc = Bun.spawn(
				["docker", "logs", "--tail", "100", "snowflake-proxy"],
				{
					stdout: "pipe",
					stderr: "pipe",
				},
			);
			const output = await new Response(proc.stdout).text();
			const errors = await new Response(proc.stderr).text();
			const logs = (output + errors)
				.split("\n")
				.filter((line) => line.includes("In the last"))
				.join("\n");

			return new Response(logs, {
				headers: { "Content-Type": "text/plain" },
			});
		} catch {
			return new Response("Logs unavailable", { status: 500 });
		}
	},
});

console.log("Logs server running on port 3001");
