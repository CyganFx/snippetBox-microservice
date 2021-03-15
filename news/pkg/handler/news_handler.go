package handler

import (
	"errors"
	"github.com/CyganFx/snippetBox-microservice/news/cmd/helpers"
	"github.com/CyganFx/snippetBox-microservice/news/pkg/domain"
	"github.com/CyganFx/snippetBox-microservice/news/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type newsHandler struct {
	service service.NewsServiceInterface
	helper  helpers.HelperInterface
}

func New(service service.NewsServiceInterface, helper helpers.HelperInterface) NewsHandlerInterface {
	return &newsHandler{service: service, helper: helper}
}

func (h *newsHandler) Home(c *gin.Context) {
	news, err := h.service.Latest()
	if err != nil {
		h.helper.ServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, news)
}

func (h *newsHandler) ShowNews(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		h.helper.NotFound(c)
		return
	}
	news, err := h.service.FindById(id)
	if err != nil {
		if errors.Is(err, domain.ErrNoRecord) {
			h.helper.NotFound(c)
			return
		} else {
			h.helper.ServerError(c, err)
			return
		}
	}

	c.JSON(http.StatusOK, news)
}

// Shouldn't be in routes
func (h *newsHandler) CreateNews(news *domain.News) (int, error) {
	id, errorSlice := h.service.Save(
		news)
	if errorSlice != nil {
		return -1, errorSlice
	}
	return id, nil
}
