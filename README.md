# html2pdf

Un servizio REST scritto in Go che converte documenti HTML in PDF utilizzando `weasyprint`.

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

---

## Endpoint API

### **POST /v0/pdfgen/fromhtml**

Accetta un JSON con un campo `html` contenente il documento HTML in **base64**.

**Esempio con `curl`:**

```bash
b64=$(base64 -w0 document.html)   # su macOS: base64 document.html | tr -d '\n'
curl -X POST http://localhost:7979/v0/pdfgen/fromhtml   -H "Content-Type: application/json"   -d "{\"html\":\"$b64\"}"   -o output.pdf
```

---

### **POST /v0/pdfgen/fromhtml-multipart**

Accetta un form-data con il file HTML (`file`).

**Esempio con `curl`:**

```bash
curl -X POST http://localhost:7979/v0/pdfgen/fromhtml-multipart   -F "file=@./document.html"   -o output.pdf
```

---

## Build locale (senza Docker)

Se vuoi compilare localmente:

```bash
go build -o htmltopdf main.go
./htmltopdf
```

---

## Dipendenze principali

- [Go](https://go.dev/)
- [weasyprint](https://weasyprint.org/)

---

## Note

- Assicurati che `weasyprint` sia installato e disponibile nel container o nel tuo ambiente locale.
- La porta predefinita è **7979**, ma può essere sovrascritta tramite la variabile d’ambiente `PORT`.
