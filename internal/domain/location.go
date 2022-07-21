package domain

type Location struct {
	Lat        float64 `json:"lat"`
	Lng        float64 `json:"lng"`
	Address    string  `json:"address"`
	City       string  `json:"city"`
	Country    string  `json:"country"`
	PostalCode string  `json:"postal_code"`
}

var LocationCols = struct {
	Lat        string
	Lng        string
	Address    string
	City       string
	Country    string
	PostalCode string
}{
	Lat:        "lat",
	Lng:        "lng",
	Address:    "address",
	City:       "city",
	Country:    "country",
	PostalCode: "postal_code",
}
