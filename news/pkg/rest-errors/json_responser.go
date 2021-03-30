package rest_errors

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime/debug"
)

type jsonResponser struct {
	errorLog *log.Logger
}

func NewJsonResponser(errorLog *log.Logger) Responser {
	return &jsonResponser{errorLog: errorLog}
}

func (h *jsonResponser) ServerError(c *gin.Context, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	h.errorLog.Output(2, trace)

	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
}

func (h *jsonResponser) ClientError(c *gin.Context, status int) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": http.StatusText(status),
	})
}

func (h *jsonResponser) ClientErrorWithDescription(c *gin.Context, status int, description string) {
	c.JSON(status, gin.H{
		"error":       http.StatusText(status),
		"description": description,
	})
}

func (h *jsonResponser) NotFound(c *gin.Context) {
	h.ClientError(c, http.StatusNotFound)
}

//TODO
//func (app *application) isAuthenticated(r *http.Request) bool {
//	return app.session.Exists(r, "accessToken")
//}

//func (app *application) ExtractToken(r *http.Request) string {
//	bearToken := app.session.GetById(r, "accessToken")
//	str := fmt.Sprintf("%v", bearToken)
//	return str
//}
//
//func (app *application) generateTokenAndPutInSession(r *http.Request, email string) error {
//	tokenManager := domain.NewManager(os.Getenv("signingKey"))
//	token, err := tokenManager.NewJWT(email)
//	if err != nil {
//		return errors.NewJsonResponser("JWT token creation problems")
//	}
//	app.infoLog.Printf("Access token created: %s", token)
//	app.session.Put(r, "accessToken", token)
//	return nil
//}
