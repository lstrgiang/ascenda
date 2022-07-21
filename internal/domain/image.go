package domain

type Images struct {
	Room      []ImageDetail `json:"room"`
	Site      []ImageDetail `json:"site"`
	Amenities []ImageDetail `json:"amenities"`
}

var ImageCols = struct {
	Room      string
	Site      string
	Amenities string
}{
	Room:      "room",
	Site:      "site",
	Amenities: "amenities",
}

type ImageDetail struct {
	Link        string `json:"link"`
	Description string `json:"description"`
}

var ImageDetailCols = struct {
	Link        string
	Description string
}{
	Link:        "link",
	Description: "description",
}
