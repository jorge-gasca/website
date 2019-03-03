//Handlers for the HTTP server

package handlers

import (
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
)

type homePageData struct {
	Title      string
	Background string
}

//IndexHandler handles GET requests to the index
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Select random background
	images, err := ioutil.ReadDir("./static/")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	backgroundIndex := rand.Intn(len(images))
	backgroundImagePath := "static/" + images[backgroundIndex].Name()

	t.Execute(w, homePageData{Title: "Test", Background: backgroundImagePath})
}
