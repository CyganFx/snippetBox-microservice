package main

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable)

	r := mux.NewRouter()
	getRouter := r.Methods("GET").Subrouter()
	postRouter := r.Methods("POST").Subrouter()

	getRouter.Handle("/", dynamicMiddleware.ThenFunc(app.home))
	getRouter.Handle("/news/create",
		dynamicMiddleware.ThenFunc(app.createNewsForm))
	postRouter.Handle("/news/create",
		dynamicMiddleware.ThenFunc(app.createNews))
	getRouter.Handle("/news/{id}", dynamicMiddleware.ThenFunc(app.showNews))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	return standardMiddleware.Then(r)
}
