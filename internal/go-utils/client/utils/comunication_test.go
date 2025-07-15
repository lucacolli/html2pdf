package utils

import (
	//"github.com/stretchr/testify/assert"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var testHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bar"))
})

func Testrequest(t *testing.T) {

	server := httptest.NewServer(testHandler)
	fmt.Println(server)
	fmt.Println(server.URL)
}
