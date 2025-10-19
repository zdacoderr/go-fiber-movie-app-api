# ======================================================
# Stage 1: Build
# ======================================================
FROM golang:1.25-alpine3.22 AS builder

# Set working directory
WORKDIR /app

# Copy go.mod dan go.sum untuk caching dependency
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy seluruh source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./main.go

# ======================================================
# Stage 2: Run
# ======================================================
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy binary hasil build dari stage sebelumnya
COPY --from=builder /app/server .

# Expose port 3000
EXPOSE 3000

# Jalankan server
CMD ["./server"]