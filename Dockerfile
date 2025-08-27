# Build Go application
FROM golang:1.23.3-alpine AS go-builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app ./backend

# Final stage with Python and uv
FROM python:3.11-slim
WORKDIR /app

# Install uv
RUN pip install uv

# Copy Go binary from builder stage
COPY --from=go-builder /app/app ./

# Copy Python dependency files
COPY pyproject.toml uv.lock ./

# Install Python dependencies with uv
RUN uv sync --frozen

# Copy Python files
COPY main.py models.py ./

# Create and copy startup script
COPY start.sh ./
RUN chmod +x start.sh

# Expose the port your Go app uses
EXPOSE 8080

# Use uv to run Python
CMD ["./start.sh"]