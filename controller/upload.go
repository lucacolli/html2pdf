package controller

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

	"html2pdf/internal/go-utils/httpw"
	"html2pdf/internal/db"
)

// FromHtml gestisce una richiesta JSON con l'HTML codificato in base64.
func FromHtml(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var item db.HtmlDoc
	json.Unmarshal(body, &item)

	// Crea cartella temporanea
	folder, err := ioutil.TempDir(os.TempDir(), "pdfgen")
	if err != nil {
		log.Println("Impossible to create the temp folder")
		httpw.Respond(w, r, http.StatusInternalServerError, httpw.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	defer os.RemoveAll(folder)

	// Decodifica HTML
	unbased, err := base64.StdEncoding.DecodeString(item.HTML)
	if err != nil {
		log.Println("Impossible to decode the HTML content")
		httpw.Respond(w, r, http.StatusInternalServerError, httpw.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	// Scrive file HTML
	htmlPath := folder + "/source.html"
	if err := ioutil.WriteFile(htmlPath, unbased, 0600); err != nil {
		log.Println("Impossible to create the HTML file")
		httpw.Respond(w, r, http.StatusInternalServerError, httpw.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	// Converte in PDF
	command := "wkhtmltopdf " + htmlPath + " " + folder + "/output.pdf"
	if _, err = exec.Command("sh", "-c", command).CombinedOutput(); err != nil {
		log.Println("Error while running wkhtmltopdf")
		httpw.Respond(w, r, http.StatusInternalServerError, httpw.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	http.ServeFile(w, r, folder+"/output.pdf")
}

// FromHtmlMultipart gestisce una richiesta multipart/form-data con un file HTML.
func FromHtmlMultipart(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		httpw.Respond(w, r, http.StatusBadRequest, httpw.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	f, _, err := r.FormFile("file")
	if err != nil {
		httpw.Respond(w, r, http.StatusBadRequest, httpw.Error{
			Code:    http.StatusBadRequest,
			Message: "file missing",
		})
		return
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		httpw.Respond(w, r, http.StatusInternalServerError, httpw.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	// Crea cartella temporanea
	folder, err := ioutil.TempDir(os.TempDir(), "pdfgen")
	if err != nil {
		httpw.Respond(w, r, http.StatusInternalServerError, httpw.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	defer os.RemoveAll(folder)

	// Scrive file HTML
	htmlPath := folder + "/source.html"
	if err := ioutil.WriteFile(htmlPath, data, 0600); err != nil {
		httpw.Respond(w, r, http.StatusInternalServerError, httpw.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	// Converte in PDF
	command := "wkhtmltopdf " + htmlPath + " " + folder + "/output.pdf"
	if _, err = exec.Command("sh", "-c", command).CombinedOutput(); err != nil {
		httpw.Respond(w, r, http.StatusInternalServerError, httpw.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	http.ServeFile(w, r, folder+"/output.pdf")
}
