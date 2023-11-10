package groupieTrecker

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	models "groupieTrecker/internal/models"
)

var (
	artistsUrl   = "https://groupietrackers.herokuapp.com/api/artists"
	relationsUrl = "https://groupietrackers.herokuapp.com/api/relation"
)

// functions for getting json files and unmarshal(parsing json info to custom structure called Artist). With this function we can get list of artists or info of one artist by id
func GetJson(url string, object interface{}) error {
	// client := &http.Client{Timeout: 10 * time.Second}
	resp, err := http.Get(url) // GET request to url which contain jsonfiles. Imagine that we are client
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, object) // object - it is array of Artists([]Artist) or just Artist(struct type Artist{})
}

// It is function that gives list of all artists for GET request from mainpage(or just from "/" url)
func GetArtistsList() ([]models.Artist, error) {
	artistsList := []models.Artist{}
	err := GetJson(artistsUrl, &artistsList)
	if err != nil {
		return nil, err
	}
	return artistsList, nil
}

// It is function that gives information about artist that was clicked by user(as you can see its use Id that retrived from userside)

func GetRelations() (models.Relations, error) {
	var relation models.Relations
	err := GetJson(relationsUrl, &relation)
	if err != nil {
		return relation, err
	}
	return relation, nil
}
