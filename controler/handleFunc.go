package home

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"workspace/asciiCode"
	home "workspace/content"
)

type httpHandler struct {
	Text   string
	Banner string
	Output string
}

func Home(w http.ResponseWriter, r *http.Request) {
	text := httpHandler{
		Output: "",
	}
	tmpl, err := template.ParseFiles("templates/homePage.html")
	if r.URL.Path != "/" {
		http.NotFound(w, r) // 404
		return
	} else {
		if err != nil {
			log.Fatal(err)
		}
	}
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusInternalServerError) // 500
		fmt.Fprint(w, "500 Internal Server Error")
		return
	} else {
		tmpl.Execute(w, text)
	}
}

func AsciiArt(w http.ResponseWriter, r *http.Request) {
	text := httpHandler{
		Output: "",
	}
	for _, v := range r.FormValue("text") {
		if string(v) == "" || v < 32 || v > 126 {

			w.WriteHeader(http.StatusBadRequest) // 400
			fmt.Fprint(w, "400 Bad Request")
			return
		}
	}

	if r.Method == http.MethodPost {
		tmpl, err := template.ParseFiles("templates/homePage.html")
		if err != nil {
			log.Fatal(err)
		}
		asciiCode.Test(r.FormValue("text"), r.FormValue("banner"))
		text.Output = string(home.Content())
		tmpl.Execute(w, text)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		fmt.Fprint(w, "405 Method Not Allowed ")
		return
	}
}
