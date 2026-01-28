# Snowflake Proxy Stats Dashboard

A beautiful, real-time stats dashboard for your Tor Snowflake proxy. Monitor connections, success rates, bandwidth usage, and geographic distribution of users you're helping bypass censorship.

![Dashboard Preview](https://via.placeholder.com/800x400?text=Snowflake+Stats+Dashboard)

## Features

- 🌍 **Geographic Distribution** - See which countries you're helping with flag emojis
- 📊 **Real-time Metrics** - Total connections, timeouts, and success rates
- 📈 **Network Stats** - Monitor bandwidth usage (inbound/outbound traffic)
- 💾 **Resource Monitoring** - Memory usage and uptime tracking
- 🎨 **Dark Theme** - Easy on the eyes, modern GitHub-style interface
- 🔄 **Auto-refresh** - Updates every 60 seconds

## Quick Start

### Prerequisites

- Docker
- Docker Compose

### Installation

1. Clone this repository:
```bash
git clone https://github.com/lone-cloud/snowflake-dashboard.git
cd snowflake-dashboard
```

2. Start the services:
```bash
docker compose up -d
```

3. Open the dashboard in your browser:
```
http://localhost:8888
```

That's it! Your Snowflake proxy is now running and helping people bypass censorship.

## Configuration

### Change Dashboard Port

Edit `docker-compose.yml` and modify the port mapping:
```yaml
ports:
  - "YOUR_PORT:8888"  # Change YOUR_PORT to desired port
```

### Metrics Port

The Snowflake proxy exposes Prometheus metrics on port 9999. If you need to change this, update both `docker-compose.yml` and `nginx.conf`.

## What is Snowflake?

Snowflake is a pluggable transport from the Tor Project that helps people bypass internet censorship. When you run a Snowflake proxy, you're providing an entry point for censored users to access the Tor network.

Your proxy acts as a bridge, making it harder for censors to block access. The more Snowflake proxies running, the harder it becomes to block Tor.

Learn more: https://snowflake.torproject.org/

## Architecture

This project runs two containers:

1. **snowflake-proxy** - The actual Tor Snowflake proxy with Prometheus metrics enabled
2. **stats-dashboard** - Nginx serving the web dashboard and proxying metrics requests

The Snowflake proxy runs in `host` network mode for optimal NAT traversal, while the stats dashboard runs on port 8888.

## Network Requirements

For best results:
- **Open UDP ports** - Forward UDP ports 32768-60999 on your router to your server
- **Static IP** - Helpful but not required
- **Stable connection** - 24/7 uptime helps the most users

Even with restricted NAT or no port forwarding, your proxy will still help users - just fewer of them.

## Success Rates

Don't be alarmed by success rates around 30-60%. This is normal for Snowflake proxies, especially with restricted NAT types. Failed connections can occur due to:

- Client-side NAT/firewall issues
- Network instability
- Client disconnections
- Geographic routing challenges

Every successful connection helps someone access the free internet.

## Monitoring

The dashboard displays:

- **Total Connections** - Lifetime successful connections
- **Connection Timeouts** - Failed connection attempts
- **Success Rate** - Percentage of successful connections
- **Inbound/Outbound Traffic** - Total bandwidth used
- **Memory Usage** - Current proxy memory consumption
- **Uptime** - How long the proxy has been running
- **Country Breakdown** - Geographic distribution of users

## Updating

Pull the latest Snowflake proxy image:
```bash
docker compose pull
docker compose up -d
```

## Troubleshooting

### Dashboard shows "Error loading"
- Check that both containers are running: `docker compose ps`
- Verify the Snowflake proxy is exposing metrics: `curl http://localhost:9999/internal/metrics`

### Low connection counts
- This is normal, especially during off-peak hours
- Connection rates vary by time of day and geographic demand
- You may see 20-150+ connections per hour depending on various factors

### Stats show 0 B network traffic
- The traffic metrics track Snowflake process network usage, not relayed Tor traffic
- This is expected if you just started the proxy

## Contributing

Contributions welcome! Please feel free to submit a Pull Request.

## License

MIT License - see LICENSE file for details

## Acknowledgments

- [Tor Project](https://www.torproject.org/) for Snowflake
- All Snowflake proxy operators helping fight censorship

## Privacy

The Snowflake proxy logs only aggregate statistics. No user traffic is logged or inspectable. The country codes shown are derived from connection metadata and don't identify specific users.

---

**Run a Snowflake proxy. Help people access the free internet.** 🌐
