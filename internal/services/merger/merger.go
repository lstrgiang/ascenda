package merger

import (
	"github.com/lstrgiang/ascenda/internal/domain"
)

func NewMerger() Merger {
	return &merger{}
}

type Merger interface {
	// merge data with newly parsed data and list of existed data retrieve from database
	MergeData(map[string][]*domain.Hotel) []*domain.Hotel
}

type merger struct {
}

// merge between hotels with the same id
func (p merger) MergeData(
	data map[string][]*domain.Hotel,
) (result []*domain.Hotel) {
	if result == nil {
		result = make([]*domain.Hotel, 0)
	}

	for _, arr := range data {
		// data not existed, creat e new
		result = append(result, p.MergeHotelArray(arr...))
	}

	return result
}

func (p merger) MergeHotelArray(srcList ...*domain.Hotel) *domain.Hotel {
	dst := &domain.Hotel{}
	for _, hotel := range srcList {
		dst = p.MergeHotel(dst, hotel)
	}
	return dst
}

// merge hotel object, which will trigger deeper layer merge
func (p merger) MergeHotel(src *domain.Hotel, dst *domain.Hotel) *domain.Hotel {
	return &domain.Hotel{
		HotelID:           p.mergeID(src.HotelID, dst.HotelID), // must be the same
		Name:              p.mergeName(src.Name, dst.Name),
		DestinationID:     p.mergeDestinationID(src.DestinationID, dst.DestinationID),
		Description:       p.mergeDescription(src.Description, dst.Description),
		Location:          p.mergeLocation(src.Location, dst.Location),
		Amenities:         p.mergeAmenities(src.Amenities, dst.Amenities),
		BookingConditions: p.mergeBookingConditions(src.BookingConditions, dst.BookingConditions),
		Images:            p.mergeImages(src.Images, dst.Images),
	}
}

// merge id string
func (p merger) mergeID(ids ...string) string {
	return nonEmptyString(ids...)
}

// merge destination ids
func (p merger) mergeDestinationID(ids ...int) int {
	return nonZeroInt(ids...)
}

// merge name
func (p merger) mergeName(names ...string) string {
	return longestString(names...)
}

// merge description
func (p merger) mergeDescription(descriptions ...string) string {
	return longestString(descriptions...)
}

// merge location
func (p merger) mergeLocation(src domain.Location, dst domain.Location) domain.Location {
	return domain.Location{
		Lat:        nonZeroFloat(src.Lat, dst.Lat),
		Lng:        nonZeroFloat(src.Lat, dst.Lat),
		Address:    longestString(src.Address, dst.Address),
		City:       longestString(src.City, dst.City),
		Country:    longestString(src.Country, dst.Country),
		PostalCode: longestString(src.PostalCode, dst.PostalCode),
	}
}

// merge images
func (p merger) mergeImages(src domain.Images, dst domain.Images) domain.Images {
	return domain.Images{
		Room:      append(src.Room, dst.Room...),
		Site:      append(src.Site, dst.Site...),
		Amenities: append(src.Amenities, dst.Amenities...),
	}
}

// merge booking conditions
func (p merger) mergeBookingConditions(stringSlices ...[]string) []string {
	return uniqueString(stringSlices...)
}

// merge amenities object
func (p merger) mergeAmenities(src domain.Amenities, dst domain.Amenities) domain.Amenities {
	return domain.Amenities{
		General: uniqueString(src.General, dst.General),
		Room:    uniqueString(src.Room, dst.Room),
	}
}
