package helpers

import (
	"html/template"
	"log"
	"net/http"
)

// See a web page
func RenderTemplates(w http.ResponseWriter, tmpl string) {
	// web page is displayed
	// template-> to write html and pass dynamic data
	ts, err := template.ParseFiles(tmpl)
	if err != nil { // parse file di not work
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)  //no data to inject so this is nil
	if err != nil {           // if parse file did now work
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
