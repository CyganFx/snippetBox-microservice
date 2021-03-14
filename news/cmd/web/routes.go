package main

import (
	"github.com/gin-gonic/gin"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	dynamicMiddleware := alice.New(app.session.Enable)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), app.secureHeaders())

	r.GET("/", app.home)
	r.GET("/news/:id", app.showNews)

	return dynamicMiddleware.Then(r)
}
