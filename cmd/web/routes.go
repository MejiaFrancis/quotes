// this is my route
package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	//create multiplexer
	router := httprouter.New()
	// create file server
	fileServer := http.FileServer(http.Dir("./static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer)) //exclude resource and go to static

	router.HandlerFunc(http.MethodGet, "/quote/create", app.quoteCreateShow) //passing in pointer, say where to find handler func
	// callback - above shows passing of the address not the func itself
	router.HandlerFunc(http.MethodPost, "quote/create", app.quoteCreateSubmit)
	//router.HandlerFunc(http.MethodGet, "/", app.Home)
	return router
}
