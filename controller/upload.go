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

func FromHtml(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var item db.HtmlDoc
	json.Unmarshal(body, &item)

	// Create temp folder
	folder, err := ioutil.TempDir(os.TempDir(), "pdfgen")
	if err != nil {
		log.Println("Impossible to create the temp folder")
		log.Println(err)
		httpw.Respond(w, r, http.StatusInternalServerError, httpw.Error{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	defer os.RemoveAll(folder)

	// Create the html file
	unbased, err := base64.StdEncoding.DecodeString(item.HTML)
	if err != nil {
		log.Println("Impossible to obtain the HTML content")
		log.Println(err)
		httpw.Respond(w, r, http.StatusInternalServerError, httpw.Error{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	err = ioutil.WriteFile(folder+"/source.html", unbased, 0600)
	if err != nil {
		log.Println("Impossible to create the HTML file")
		log.Println(err)
		httpw.Respond(w, r, http.StatusInternalServerError, httpw.Error{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	// Create the command
	command := "weasyprint " + folder + "/source.html " + folder + "/output.pdf"

	// Run command
	_, err = exec.Command("sh", "-c", command).Output()
	if err != nil {
		log.Println("Error while running weasyprint")
		log.Println(err)
		httpw.Respond(w, r, http.StatusInternalServerError, httpw.Error{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	http.ServeFile(w, r, folder+"/output.pdf")
}
