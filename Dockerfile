# Stage 1: Build del binario Go
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o html2pdf main.go

# Stage 2: Runtime minimale con WeasyPrint installato via pip (con override)
FROM debian:bookworm

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        python3 \
        python3-pip \
        python3-cffi \
        libcairo2 \
        pango1.0-tools \
        libpango-1.0-0 \
        fonts-liberation \
        fonts-dejavu \
        fontconfig \
        ca-certificates && \
    pip install --break-system-packages --no-cache-dir weasyprint && \
    fc-cache -fv && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/html2pdf /usr/local/bin/html2pdf

EXPOSE 8080
CMD ["/usr/local/bin/html2pdf"]
