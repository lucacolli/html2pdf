# Stage 1: Build del binario Go
FROM golang:1.22 AS builder

WORKDIR /app

# Copia i moduli Go e scarica le dipendenze
COPY go.mod go.sum ./
RUN go mod download

# Copia tutto il codice sorgente
COPY . .

# Compila il binario
RUN go build -o htmltopdf main.go

# Stage 2: Immagine finale
FROM debian:bookworm

ENV DEBIAN_FRONTEND=noninteractive

# Installa wkhtmltopdf e dipendenze necessarie
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        wkhtmltopdf \
        ca-certificates \
        fontconfig \
        fonts-dejavu \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copia il binario compilato dallo stage builder
COPY --from=builder /app/htmltopdf .

# Espone la porta del servizio
EXPOSE 7879

# Comando di default
CMD ["./htmltopdf"]
