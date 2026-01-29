const http = require("node:http");
const { spawn } = require("node:child_process");

http
	.createServer((_req, res) => {
		const proc = spawn("docker", ["logs", "--tail", "100", "snowflake-proxy"]);

		let output = "";
		proc.stdout.on("data", (data) => {
			output += data;
		});
		proc.stderr.on("data", (data) => {
			output += data;
		});

		proc.on("close", () => {
			const logs = output
				.split("\n")
				.filter((line) => line.includes("In the last"))
				.join("\n");

			res.writeHead(200, { "Content-Type": "text/plain" });
			res.end(logs);
		});

		proc.on("error", () => {
			res.writeHead(500);
			res.end("Logs unavailable");
		});
	})
	.listen(3001, () => console.log("Logs server running on port 3001"));
