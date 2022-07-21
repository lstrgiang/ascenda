package parser

import "github.com/lstrgiang/ascenda/internal/domain"

var keyMapper = map[string]string{
	"id":                 domain.HotelCols.HotelID,
	"Id":                 domain.HotelCols.HotelID,
	"hotel_id":           domain.HotelCols.HotelID,
	"Destinations":       domain.HotelCols.DestinationID,
	"DestinationId":      domain.HotelCols.DestinationID,
	"destination_id":     domain.HotelCols.DestinationID,
	"Name":               domain.HotelCols.Name,
	"name":               domain.HotelCols.Name,
	"hotel_name":         domain.HotelCols.Name,
	"lat":                domain.HotelCols.Location,
	"lng":                domain.HotelCols.Location,
	"address":            domain.HotelCols.Location,
	"Latitude":           domain.HotelCols.Location,
	"Longitude":          domain.HotelCols.Location,
	"City":               domain.HotelCols.Location,
	"Address":            domain.HotelCols.Location,
	"Country":            domain.HotelCols.Location,
	"PostalCode":         domain.HotelCols.Location,
	"location":           domain.HotelCols.Location,
	"Description":        domain.HotelCols.Description,
	"info":               domain.HotelCols.Description,
	"details":            domain.HotelCols.Description,
	"amenities":          domain.HotelCols.Amenities,
	"Amenities":          domain.HotelCols.Amenities,
	"Facilities":         domain.HotelCols.Amenities,
	"images":             domain.HotelCols.Images,
	"booking_conditions": domain.HotelCols.BookingConditions,
}

var amenitiesMapper = map[string]string{
	"rooms":   domain.AmenityCols.Room,
	"room":    domain.AmenityCols.Room,
	"Rooms":   domain.AmenityCols.Room,
	"General": domain.AmenityCols.General,
	"general": domain.AmenityCols.General,
}

var locationMapper = map[string]string{
	"Latitude":    domain.LocationCols.Lat,
	"Longitude":   domain.LocationCols.Lng,
	"lat":         domain.LocationCols.Lat,
	"lng":         domain.LocationCols.Lng,
	"Address":     domain.LocationCols.Address,
	"address":     domain.LocationCols.Address,
	"city":        domain.LocationCols.City,
	"City":        domain.LocationCols.City,
	"country":     domain.LocationCols.Country,
	"Country":     domain.LocationCols.Country,
	"postal_code": domain.LocationCols.PostalCode,
	"PostalCode":  domain.LocationCols.PostalCode,
}
var imageMapper = map[string]string{
	"rooms":     domain.ImageCols.Room,
	"amenities": domain.ImageCols.Amenities,
	"site":      domain.ImageCols.Site,
}
var imageDetailMapper = map[string]string{
	"url":         domain.ImageDetailCols.Link,
	"link":        domain.ImageDetailCols.Link,
	"description": domain.ImageDetailCols.Description,
	"caption":     domain.ImageDetailCols.Description,
	"Description": domain.ImageDetailCols.Description,
}
