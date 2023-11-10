package groupieTrecker
// custom structure
type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Relations    string   `json:"-"`
	DatesLocation  map[string][]string `json:"-"`
}

type Relations struct {
	Index []struct {
		Id int `json:"id"`
		DatesLocation map[string][]string `json:"datesLocations"` 
	} 
}

type ParsedLocations struct {
	city string
	country string
	date []string
}