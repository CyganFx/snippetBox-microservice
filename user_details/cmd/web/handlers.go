package main

import (
	"errors"
	"github.com/CyganFx/snippetBox-microservice/user_details/pkg/domain"
	"github.com/CyganFx/snippetBox-microservice/user_details/pkg/forms"
	"github.com/markbates/goth/gothic"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

func (app *application) signupForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "signup.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) signup(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("name", "email", "password")
	form.MinLength("password", 10)
	form.MatchesPattern("email", forms.EmailRX)
	form.MaxLength("name", 255)
	form.MaxLength("email", 255)

	if !form.Valid() {
		app.render(w, r, "signup.page.tmpl", &templateData{
			Form: form,
		})
		return
	}

	err = app.userService.Save(form.Get("name"), form.Get("email"), form.Get("password"))
	if err != nil {
		if errors.Is(err, domain.ErrDuplicateEmail) {
			form.Errors.Add("email", "Address is already in use")
			app.render(w, r, "signup.page.tmpl", &templateData{
				Form: form,
			})
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.session.Put(r, "flash", "Your signup was successful. Please log in.")
	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

func (app *application) loginForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	var user *domain.User
	form := forms.New(r.PostForm)
	user, err = app.userService.Authenticate(form.Get("email"), form.Get("password"))
	if err != nil {
		if errors.Is(err, domain.ErrInvalidCredentials) {
			form.Errors.Add("generic", "Email or Password is incorrect")
			app.render(w, r, "login.page.tmpl", &templateData{Form: form})
		} else {
			app.serverError(w, err)
		}
		return
	}
	err = app.generateTokenAndPutInSession(r, user.Email)
	if err != nil {
		app.serverError(w, err)
	}
	http.Redirect(w, r, "/news/create", http.StatusSeeOther)
}

func (app *application) logout(w http.ResponseWriter, r *http.Request) {
	app.session.Remove(r, "accessToken")
	app.session.Put(r, "flash", "You've been logged out successfully!")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) oauthCallback(w http.ResponseWriter, r *http.Request) {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.infoLog.Printf("oauth completed: %s", user)
	err = app.generateTokenAndPutInSession(r, user.Email)
	if err != nil {
		app.serverError(w, err)
	}
	http.Redirect(w, r, "/news/create", http.StatusSeeOther)
}

func (app *application) oauthCheck(w http.ResponseWriter, r *http.Request) {
	gothic.BeginAuthHandler(w, r)
}
