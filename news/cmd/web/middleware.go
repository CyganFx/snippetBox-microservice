package main

import (
	"github.com/gin-gonic/gin"
)

func (app *application) secureHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
		c.Writer.Header().Set("X-Frame-Options", "deny")
		c.Next()
	}
}

// TODO
//func (app *application) TokenVerify(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		authToken := app.ExtractToken(r)
//		token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
//			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//				return nil, fmt.Errorf("There was an error")
//			}
//			return []byte(os.Getenv("signingKey")), nil
//		})
//		if err != nil || !token.Valid {
//			http.Redirect(w, r, "/user/login", http.StatusSeeOther)
//			return
//		}
//		w.Header().Add("Cache-Control", "no-store")
//		next.ServeHTTP(w, r)
//	})
//}
