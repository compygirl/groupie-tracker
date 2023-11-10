package groupieTrecker

import (
	"log"

	models "groupieTrecker/internal/models"
)



func UnmarshalArtistsAndRelations() ([]models.Artist,error) {
	artistsList, err :=GetArtistsList()
	if err != nil {
		log.Fatalln(err)
		return nil,err
	}
	relationlist, err := GetRelations()
	if err != nil {
		log.Fatalln(err)
		return nil,err
	}
	for i := 0 ; i < len(artistsList); i++ {
		artistsList[i].DatesLocation = relationlist.Index[i].DatesLocation
	}
	return artistsList,nil
}
