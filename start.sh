#! /bin/sh
# Start Go backend
./app &

# TODO: See if uv run python main.py is faster
# Start Python API
uvicorn main:app --host 0.0.0.0 --port 8001