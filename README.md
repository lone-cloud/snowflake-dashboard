# Snowflake Dashboard

A simple [snowflake](https://snowflake.torproject.org/) dashboard to portray runtime metrics.

The dashboard is build using data from snowflake's /internal/metrics endpoint to get the totals as well as its docker logs to get its hourly totals.

![Dashboard Preview](screenshot.webp)

## Installation

Download the `docker-compose.yml` file:

```bash
curl -O https://raw.githubusercontent.com/lone-cloud/snowflake-dashboard/main/docker-compose.yml
```

Start the services:

```bash
docker compose up -d
```

Open the dashboard in your browser:

```plaintext
http://localhost:8888
```

## Updating

Pull the latest image and restart:

```bash
docker compose pull
docker compose up -d
```
