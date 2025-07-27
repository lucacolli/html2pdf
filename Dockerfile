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

# Stage 2: Immagine finale con wkhtmltopdf e font
FROM debian:bookworm

# Evita richieste interattive
ENV DEBIAN_FRONTEND=noninteractive

# Installa wkhtmltopdf e font aggiuntivi
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        wkhtmltopdf \
        ca-certificates \
        fonts-liberation \
        fonts-dejavu \
        fontconfig \
    && fc-cache -fv \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copia il binario compilato dallo stage builder
COPY --from=builder /app/htmltopdf .

# Espone la porta del servizio
ENV PORT=7979
EXPOSE 7979

# Comando di default
CMD ["./htmltopdf"]
