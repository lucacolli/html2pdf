package httpw

import (
	"log"
	"net/http"
	"os"

	"github.com/urfave/negroni"
)

func NewRecovery() *negroni.Recovery {
	return &negroni.Recovery{
		Logger:     log.New(os.Stdout, "[negroni] ", 0),
		PrintStack: false,
		StackAll:   false,
		StackSize:  1024 * 8,
		Formatter:  &PanicFormatter{},
	}
}

type PanicFormatter struct{}

func (p *PanicFormatter) FormatPanicError(w http.ResponseWriter, r *http.Request, infos *negroni.PanicInformation) {
	Respond(w, r, http.StatusInternalServerError, Error{Code: http.StatusInternalServerError, Message: "Sorry, an unmanaged error occurred :("})
}
