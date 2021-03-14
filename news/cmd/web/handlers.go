package main

import (
	"errors"
	"github.com/CyganFx/snippetBox-microservice/news/pkg/domain"
	"github.com/CyganFx/snippetBox-microservice/news/pkg/validator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (app *application) home(c *gin.Context) {
	news, err := app.newsService.Latest()
	if err != nil {
		app.serverError(c, err)
		return
	}
	c.JSON(http.StatusOK, news)
}

func (app *application) showNews(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		app.notFound(c)
		return
	}
	news, err := app.newsService.FindById(id)
	if err != nil {
		if errors.Is(err, domain.ErrNoRecord) {
			app.notFound(c)
			return
		} else {
			app.serverError(c, err)
			return
		}
	}

	c.JSON(http.StatusOK, news)
}

// TODO make controller class, not app
func (app *application) CreateNews(title, content, expires string) (int, []string) {
	v := validator.New()
	v.MaxLength(title, 100)
	v.PermittedValues(expires, "365", "7", "1")
	if !v.Valid() {
		return -1, v.Errors.Errors
	}

	id, err := app.newsService.Save(
		title, content, expires)
	if err != nil {
		return -1, v.Errors.Errors
	}
	return id, nil
}
