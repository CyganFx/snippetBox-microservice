package main

import (
	"alexedwards.net/snippetbox/pkg/domain"
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"os"
	"runtime/debug"
	"time"
)

func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}
	td.CurrentYear = time.Now().Year()
	td.Flash = app.session.PopString(r, "flash")
	td.IsAuthenticated = app.isAuthenticated(r)
	return td
}

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("The template %s does not exist", name))
		return
	}
	buf := new(bytes.Buffer)

	err := ts.Execute(buf, app.addDefaultData(td, r))
	if err != nil {
		app.serverError(w, err)
		return
	}
	buf.WriteTo(w)
}

func (app *application) isAuthenticated(r *http.Request) bool {
	return app.session.Exists(r, "accessToken")
}

func (app *application) ExtractToken(r *http.Request) string {
	//bearToken := r.Header.Get("Authorization")
	bearToken := app.session.Get(r, "accessToken")
	str := fmt.Sprintf("%v", bearToken)
	return str
	//strArr := strings.Split(str, " ")
	//if len(strArr) == 2 {
	//	return strArr[1]
	//}
	//return ""
}

func (app *application) createSession(r *http.Request, user *domain.User) error {
	tokenManager := domain.NewManager(os.Getenv("signingKey"))
	token, err := tokenManager.NewJWT(user)
	if err != nil {
		return errors.New("JWT token creation problems")
	}
	app.infoLog.Printf("Access token created: %s", token)
	app.session.Put(r, "accessToken", token)
	return nil
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
