package main

import (
	"alexedwards.net/snippetbox/pkg/domain"
	"alexedwards.net/snippetbox/pkg/forms"
	"errors"
	"net/http"
)

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

	//var user domain.User
	//json.NewDecoder(r.Body).Decode(&user)
	//spew.Dump(user) //print

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

	err = app.users.Insert(form.Get("name"), form.Get("email"), form.Get("password"))
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

	//hash, err := bcrypt.GenerateFromPassword(user.Password, 10)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//user.Password = hash
	//
	//id, err := app.snippets.Insert(
	//	form.Get("title"), form.Get("content"), form.Get("expires"))
	//if err != nil {
	//	app.serverError(w, err)
	//	return
	//}
	//
	//app.session.Put(r, "flash", "Snippet successfully created!")
	//
	//http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
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
	user, err = app.users.Authenticate(form.Get("email"), form.Get("password"))
	if err != nil {
		if errors.Is(err, domain.ErrInvalidCredentials) {
			form.Errors.Add("generic", "Email or Password is incorrect")
			app.render(w, r, "login.page.tmpl", &templateData{Form: form})
		} else {
			app.serverError(w, err)
		}
		return
	}
	err = app.createSession(r, user)
	if err != nil {
		app.serverError(w, err)
	}
	http.Redirect(w, r, "/snippet/create", http.StatusSeeOther)
}

func (app *application) logout(w http.ResponseWriter, r *http.Request) {
	app.session.Remove(r, "accessToken")
	app.session.Put(r, "flash", "You've been logged out successfully!")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
