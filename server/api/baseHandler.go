package api

import (
	"io"
	"net/http"
)

var helloMsg = "Hello student!\n"

func BaseHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, helloMsg)
}
