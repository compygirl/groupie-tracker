package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"

	getJson "groupieTrecker/internal/getJson"
	models "groupieTrecker/internal/models"
)

func GetMainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusNotFound, errors.New(" "))
		return
	}
	switch r.Method {
	case "GET":
		indexTemplate := "internal/templates/html_templates/index.html"
		// artistList, err := GetArtistsList()
		RenderTemplate(w, indexTemplate, ArtistsList)
		return
	default:
		ErrorHandler(w, http.StatusMethodNotAllowed, errors.New(" "))
		return
	}
}

// artist pages
func GetArtist(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artist" {
		ErrorHandler(w, http.StatusNotFound, errors.New(" "))
		return
	}
	switch r.Method {
	case "GET":
		id := r.FormValue("id") // we get id from <a href="/artist?id={{.Id}}"
		if !r.Form.Has("id") {
			ErrorHandler(w, http.StatusBadRequest, errors.New("you hould specify more details"))
			return
		}
		if len(id) > 2 || id[0] == '0'  {
			ErrorHandler(w, http.StatusNotFound, errors.New("there is no artist with this ID"))
			return
		}
		indexTemplate := "internal/templates/html_templates/artistPage.html"
		index, err := strconv.Atoi(id)
		if err != nil {
			ErrorHandler(w, http.StatusNotFound, errors.New("there is no artist with this ID"))
			fmt.Println(err)
			return
		}
		if index <= 0 || index > len(ArtistsList) {
			ErrorHandler(w, http.StatusBadRequest, errors.New("there is no artist with this ID"))
			return
		}
		artistInfo := ArtistsList[index-1]
		fmt.Println(artistInfo.DatesLocation)
		RenderTemplate(w, indexTemplate, artistInfo)
		return
	default:
		ErrorHandler(w, http.StatusMethodNotAllowed, errors.New(" "))
		return
	}
}

func RenderTemplate(w http.ResponseWriter, htmltemplate string, resp interface{}) {
	temp, err := template.ParseFiles(htmltemplate)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, errors.New("problems with parsing"))
		fmt.Println(err)
		return
	}

	err = temp.Execute(w, resp)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, errors.New("problems with executing"))
		fmt.Println(err)
		return
	}
}
func ErrorHandler(w http.ResponseWriter, errorNum int, errDetails error) {
	var resp models.ErrorResp
	resp.ErrorNum = errorNum
	resp.ErrorMessage = http.StatusText(errorNum) + "\n" + errDetails.Error()
	w.WriteHeader(errorNum)
	temp, err := template.ParseFiles("internal/templates/html_templates/errors.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = temp.Execute(w, resp)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func ErrorCheck(err error) bool {
	if err != nil {
		log.Fatalln(err)
		return false
	}
	return true
}

var ArtistsList, err = getJson.UnmarshalArtistsAndRelations()

func main() {
	if !ErrorCheck(err) {
		return
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	styles := http.FileServer(http.Dir("internal/templates/css"))

	router := http.NewServeMux()
	router.Handle("/css/", http.StripPrefix("/css/", styles))
	router.HandleFunc("/", GetMainPage)
	router.HandleFunc("/artist", GetArtist)

	fmt.Println("listening on: http://localhost:" + port + "/")

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Println("Error: several connecrtions")
		return
	}
}
