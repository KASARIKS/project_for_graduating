package handlers

import (
	"net/http"
)

func Main(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/index.html")
}

func ScriptMinJs(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/js/scripts.min.js")
}

func StyleCss(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/css/style.css")
}

func FaviconIco(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/favicon.ico")
}
