#!/bin/sh
set -e

# Start Python API on localhost:8001 (internal only)
uv run uvicorn main:app --host 127.0.0.1 --port 8001 &

# Start Go backend on Render-assigned port ($PORT)
./app --port ${PORT:-8080}