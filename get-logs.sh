#!/bin/sh

docker logs --tail 100 snowflake-proxy 2>&1 | grep "In the last"
