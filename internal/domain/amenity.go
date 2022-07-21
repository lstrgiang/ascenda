package domain

type Amenities struct {
	General []string `json:"general"`
	Room    []string `json:"room"`
}

var AmenityCols = struct {
	Room    string
	General string
}{
	Room:    "room",
	General: "general",
}
