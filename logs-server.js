const http = require("node:http");
const { spawn } = require("node:child_process");

function getPath(req) {
	try {
		return new URL(req.url, "http://localhost").pathname;
	} catch {
		return req.url;
	}
}

http
	.createServer((req, res) => {
		const path = getPath(req);
		const proc = spawn("docker", ["logs", "--tail", "500", "snowflake-proxy"]);

		let output = "";
		proc.stdout.on("data", (data) => {
			output += data;
		});
		proc.stderr.on("data", (data) => {
			output += data;
		});

		proc.on("close", () => {
			if (path === "/internal/nat") {
				const natMatch = output
					.split("\n")
					.reverse()
					.map((line) => line.match(/\bNAT type:\s*([^\r\n]+)\s*$/))
					.find(Boolean);

				const natType = natMatch?.[1]?.trim() || "Unknown";
				res.writeHead(200, { "Content-Type": "text/plain" });
				res.end(natType);
				return;
			}

			if (path === "/internal/logs" || path === "/") {
				const logs = output
					.split("\n")
					.filter((line) => line.includes("In the last"))
					.join("\n");

				res.writeHead(200, { "Content-Type": "text/plain" });
				res.end(logs);
				return;
			}

			res.writeHead(404, { "Content-Type": "text/plain" });
			res.end("Not Found");
		});

		proc.on("error", () => {
			res.writeHead(500);
			res.end("Logs unavailable");
		});
	})
	.listen(3001, () => console.log("Logs server running on port 3001"));
