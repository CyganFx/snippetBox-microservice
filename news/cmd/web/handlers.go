package main

import (
	"errors"
	"fmt"
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

func (app *application) createNews(c *gin.Context) {
	var news domain.News
	err := c.BindJSON(&news)
	if err != nil {
		app.clientErrorWithDescription(c, http.StatusBadRequest, "couldn't bind json news")
		return
	}

	v := validator.New()
	v.Required("title", "content", "expires")
	v.MaxLength("title", 100)
	v.PermittedValues("expires", "365", "7", "1")
	if !v.Valid() {
		app.validationError(c, http.StatusBadRequest, v.Errors)
		return
	}

	id, err := app.newsService.Save(
		news.Title, news.Content, news.Expires)

	if err != nil {
		app.serverError(c, err)
		return
	}

	c.Data(http.StatusOK, "application/json", []byte(fmt.Sprintf("News with id %d successfully created!", id)))
	c.JSON(http.StatusOK, news)
}
