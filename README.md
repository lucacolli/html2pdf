# html2pdf

Un servizio REST scritto in Go che converte documenti HTML in PDF utilizzando `wkhtmltopdf`.

## Build con Docker

Costruisci l'immagine Docker:

```bash
docker build --no-cache -t html2pdf .
```

## Avvio del servizio

Avvia il container esponendo la porta `7979`:

```bash
docker run -e PORT=7979 -p 7979:7979 html2pdf
```

Il servizio sarà disponibile su:

```
http://localhost:7979
```

## Endpoint API

### **POST /html2pdf/v0/convert**

Converte un documento HTML in PDF.

**Esempio con `curl`:**

```bash
curl -X POST http://localhost:7979/html2pdf/v0/convert   -F "file=@./document.html"   -o output.pdf
```

- `file`: il file HTML da convertire.
- L'output PDF viene salvato come `output.pdf`.

## Build locale (senza Docker)

Se vuoi compilare localmente:

```bash
go build -o htmltopdf main.go
./htmltopdf
```

## Dipendenze principali

- [Go](https://go.dev/)
- [wkhtmltopdf](https://wkhtmltopdf.org/)

---

## Note

- Assicurati che `wkhtmltopdf` sia installato e disponibile nel container o nel tuo ambiente locale.
- La porta predefinita è **7979**, ma può essere sovrascritta tramite la variabile d’ambiente `PORT`.
