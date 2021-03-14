package main

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, app.secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable)

	r := mux.NewRouter()
	getRouter := r.Methods("GET").Subrouter()
	postRouter := r.Methods("POST").Subrouter()

	getRouter.HandleFunc("/", app.home)
	getRouter.Handle("/user/signup", dynamicMiddleware.ThenFunc(app.signupForm))
	postRouter.Handle("/user/signup", dynamicMiddleware.ThenFunc(app.signup))
	getRouter.Handle("/user/login", dynamicMiddleware.ThenFunc(app.loginForm))
	postRouter.Handle("/user/login", dynamicMiddleware.ThenFunc(app.login))
	postRouter.Handle("/user/logout",
		dynamicMiddleware.Append(app.TokenVerify).ThenFunc(app.logout))

	getRouter.Handle("/auth/{provider}",
		dynamicMiddleware.ThenFunc(app.oauthCheck))

	getRouter.Handle("/auth/{provider}/callback",
		dynamicMiddleware.ThenFunc(app.oauthCallback))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	return standardMiddleware.Then(r)
}
