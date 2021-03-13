package main

import (
	"errors"
	"fmt"
	"github.com/CyganFx/snippetBox-microservice/news/pkg/domain"
	"github.com/CyganFx/snippetBox-microservice/news/pkg/forms"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	s, err := app.newsService.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.render(w, r, "home.page.tmpl", &templateData{
		NewsArray: s,
	})
}

func (app *application) showNews(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.newsService.Find(id)
	if err != nil {
		if errors.Is(err, domain.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.render(w, r, "show.page.tmpl", &templateData{
		News: s,
	})
}

func (app *application) createNewsForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) createNews(w http.ResponseWriter, r *http.Request) {
	// Limit the request body size to 4096 bytes, default: 10MB
	r.Body = http.MaxBytesReader(w, r.Body, 4096)
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("title", "content", "expires")
	form.MaxLength("title", 100)
	form.PermittedValues("expires", "365", "7", "1")
	if !form.Valid() {
		app.render(w, r, "create.page.tmpl", &templateData{
			Form: form,
		})
		return
	}

	id, err := app.newsService.Save(
		form.Get("title"), form.Get("content"), form.Get("expires"))
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.session.Put(r, "flash", "News successfully created!")

	http.Redirect(w, r, fmt.Sprintf("/news/%d", id), http.StatusSeeOther)
}
