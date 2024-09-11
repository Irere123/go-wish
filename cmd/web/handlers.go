package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	cuid "github.com/irere123/bdwish/internal"
)

type CreateWishRequest struct {
	Name    string
	Age     int
	Message string
}

type WishResponse struct {
	ID      int
	Name    string
	Age     int
	Message string
	Slug    string
}

type TemplateData struct {
	Wish *Wish
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./static/html/home.page.tmpl",
		"./static/html/base.layout.tmpl",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)

	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) createWish(w http.ResponseWriter, r *http.Request) {
	data := &CreateWishRequest{}

	err := json.NewDecoder(r.Body).Decode(data)

	if err != nil {
		app.badRequest(w, err)
		return
	}

	newWish := Wish{Name: data.Name, Age: data.Age, Message: data.Message, Slug: cuid.Slug()}

	app.db.Create(&newWish)

	wish := Wish{}

	app.db.Raw("SELECT * FROM wishes WHERE ID= ?", newWish.ID).Scan(&wish)

	resp, err := json.MarshalIndent(wish, "", "  ")

	if err != nil {
		app.errorLog.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(resp))
}

func (app *application) wishCreated(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		http.Redirect(w, r, "/", http.StatusBadRequest)
		return
	}

	files := []string{
		"./static/html/created.page.tmpl",
		"./static/html/base.layout.tmpl",
	}

	wish := Wish{}

	app.db.Raw("SELECT * FROM wishes WHERE ID= ?", id).Scan(&wish)

	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, &TemplateData{Wish: &wish})

	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) retrieveWishMessage(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")

	wish := Wish{}

	app.db.Raw("SELECT * FROM wishes WHERE Slug= ?", slug).Scan(&wish)

	files := []string{
		"./static/html/retrieve.page.tmpl",
		"./static/html/base.layout.tmpl",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, &TemplateData{Wish: &wish})

	if err != nil {
		app.serverError(w, err)
		return
	}
}
