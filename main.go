package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	"html2pdf/internal/go-utils/config"
	"html2pdf/controller"
)

func main() {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("{\"code\": 404, \"message\": \"Not Found\"}"))
	})

	r.HandleFunc("/v0/pdfgen/fromhtml", controller.FromHtml).Methods("POST")

	// Health
	r.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("ok")) }).Methods("GET")
	r.HandleFunc("/readiness", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("ok")) }).Methods("GET")

	// Negroni
	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.Use(negroni.NewLogger())
	n.UseHandler(r) // cors.New() rimosso

	port := config.Get("PORT")
	if port == "" {
		port = "7979"
	}

	log.Println("Starting service on port " + port)
	http.ListenAndServe("0.0.0.0:"+port, n)
}
