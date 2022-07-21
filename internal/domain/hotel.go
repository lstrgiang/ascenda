package domain

type Hotel struct {
	HotelID           string    `json:"hotel_id"`
	Name              string    `json:"name"`
	DestinationID     int       `json:"destination_id"`
	Description       string    `json:"description"`
	Location          Location  `json:"location"`
	Amenities         Amenities `json:"amenities"`
	BookingConditions []string  `json:"booking_conditions"`
	Images            Images    `json:"images"`
}

type HotelSlice []*Hotel

var HotelCols = struct {
	ID                string
	HotelID           string
	Name              string
	DestinationID     string
	Description       string
	Location          string
	Amenities         string
	BookingConditions string
	Images            string
}{
	HotelID:           "hotel_id",
	Name:              "name",
	DestinationID:     "destination_id",
	Description:       "description",
	Location:          "location",
	Amenities:         "amenities",
	BookingConditions: "booking_conditions",
	Images:            "images",
}
