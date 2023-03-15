// this is my handler
package main

import (
	"log"
	"net/http"

	"github.com/MejiaFrancis/quotes/helpers"
)

func (app *application) quoteCreateShow(w http.ResponseWriter, r *http.Request) {

	helpers.RenderTemplates(w, "./static/html/quote.page.tmpl")
	//RenderTemplate(w, "home.page.tmpl", nil)
	// w.Write([]byte("Welcome to my page."))
	//question, err := app.question.Get()
	//if err != nil {
	//return
	//}
	//w.Write([]byte(question.Body))
}

// create handler for about
func (app *application) quoteCreateSubmit(w http.ResponseWriter, r *http.Request) {

	//get the data from the form
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	// insert the quote and author into the database
	quote := r.PostForm.Get("quote")
	author := r.PostForm.Get("author_name")
	log.Printf("%s %s\n", quote, author)
	id, err := app.quote.Insert(quote, author)
	log.Printf("%s %s %d\n", quote, author, id)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}

}
