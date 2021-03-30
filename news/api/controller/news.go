package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"snippetBox-microservice/news/internal/service"
	"snippetBox-microservice/news/pkg/domain"
	"snippetBox-microservice/news/pkg/rest-errors"
	"strconv"
)

type news struct {
	service service.NewsInterface
	errors  rest_errors.Responser
}

func New(service service.NewsInterface, helper rest_errors.Responser) domain.NewsController {
	return &news{service: service, errors: helper}
}

func (h *news) Home(c *gin.Context) {
	news, err := h.service.Latest()
	if err != nil {
		h.errors.ServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, news)
}

func (h *news) ShowNews(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		h.errors.NotFound(c)
		return
	}
	news, err := h.service.FindById(id)
	if err != nil {
		if errors.Is(err, domain.ErrNoRecord) {
			h.errors.NotFound(c)
			return
		} else {
			h.errors.ServerError(c, err)
			return
		}
	}

	c.JSON(http.StatusOK, news)
}

// Shouldn't be in routes
func (h *news) CreateNews(news *domain.News) (int, error) {
	id, errorSlice := h.service.Save(
		news)
	if errorSlice != nil {
		return -1, errorSlice
	}
	return id, nil
}
