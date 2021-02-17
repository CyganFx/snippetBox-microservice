package main

import (
	"alexedwards.net/snippetbox/pkg/forms"
	"alexedwards.net/snippetbox/pkg/models"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func (app *application) signupForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Display the user signup form...")

}

func (app *application) signup(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new user...")
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	spew.Dump(user) //print

	form := forms.New(r.PostForm)
	form.Required("email", "password")
	if !form.Valid() {
		//TODO
		app.render(w, r, "signup.page.tmpl", &templateData{
			Form: form,
		})
	}

	hash, err := bcrypt.GenerateFromPassword(user.Password, 10)
	//TODO
	if err != nil {
		log.Fatal(err)
	}
	user.Password = hash

	id, err := app.snippets.Insert(
		form.Get("title"), form.Get("content"), form.Get("expires"))
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.session.Put(r, "flash", "Snippet successfully created!")

	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}

func (app *application) loginForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Display the user login form...")
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Login the user...")
}

func (app *application) logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Logout the user...")
}

func (app *application) protectedEndpoint(w http.ResponseWriter, r *http.Request) {

}

func (app *application) TokenVerifyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return nil
}
