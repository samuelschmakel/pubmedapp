#!/bin/bash
echo "Starting Python API with uv..."
uv run python main.py &
PYTHON_PID=$!
echo "Python API started with PID: $PYTHON_PID"

echo "Waiting for Python API to initialize..."
sleep 10

echo "Starting Go backend..."
exec ./app