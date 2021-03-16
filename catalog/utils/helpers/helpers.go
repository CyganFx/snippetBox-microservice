package helpers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

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
//		return errors.New("JWT token creation problems")
//	}
//	app.infoLog.Printf("Access token created: %s", token)
//	app.session.Put(r, "accessToken", token)
//	return nil
//}

type Helper struct {
	errorLog *log.Logger
}

func New(errorLog *log.Logger) HelperInterface {
	return &Helper{errorLog: errorLog}
}

func (h *Helper) HumanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

func (h *Helper) ServerError(c *gin.Context, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	h.errorLog.Output(2, trace)

	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
}

func (h *Helper) ClientError(c *gin.Context, status int) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": http.StatusText(status),
	})
}

func (h *Helper) ClientErrorWithDescription(c *gin.Context, status int, description string) {
	c.JSON(status, gin.H{
		"error":       http.StatusText(status),
		"description": description,
	})
}

func (h *Helper) NotFound(c *gin.Context) {
	h.ClientError(c, http.StatusNotFound)
}
