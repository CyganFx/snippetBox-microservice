package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"snippetBox-microservice/news/pkg/domain"
	"snippetBox-microservice/news/pkg/rest-errors"
	"strconv"
)

// Actually we don't need controller in microservices, just for fun

type news struct {
	service NewsServiceInterface
	errors  rest_errors.Responser
}

type NewsServiceInterface interface {
	Save(news *domain.News) (int, error)
	FindById(id int) (*domain.News, error)
	Latest() ([]*domain.News, error)
}

func New(service NewsServiceInterface, helper rest_errors.Responser) domain.NewsController {
	return &news{service: service, errors: helper}
}

func (n *news) Home(c *gin.Context) {
	news, err := n.service.Latest()
	if err != nil {
		n.errors.ServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, news)
}

func (n *news) ShowNews(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		n.errors.NotFound(c)
		return
	}
	news, err := n.service.FindById(id)
	if err != nil {
		if errors.Is(err, domain.ErrNoRecord) {
			n.errors.NotFound(c)
			return
		} else {
			n.errors.ServerError(c, err)
			return
		}
	}

	c.JSON(http.StatusOK, news)
}

// Shouldn't be in routes
func (n *news) CreateNews(news *domain.News) (int, error) {
	id, errorSlice := n.service.Save(
		news)
	if errorSlice != nil {
		return -1, errorSlice
	}
	return id, nil
}
