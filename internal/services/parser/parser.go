package parser

import (
	"encoding/json"

	"github.com/lstrgiang/ascenda/internal/domain"
)

// 1. call api, retrieve data
// 2. map data into object for validation if necessary
// 3. merge data from different sources
// 4. find and merge to update existed data
// 5. insert new data
func NewParser() Parser {
	return &parser{}
}

type Parser interface {
	ParseData(rawData []byte) (map[string][]*domain.Hotel, error)
}

type parser struct {
}

func (s parser) ParseData(rawData []byte) (map[string][]*domain.Hotel, error) {
	parsedData := make([]interface{}, 0)
	dataList := make(map[string][]*domain.Hotel)
	if err := json.Unmarshal(rawData, &parsedData); err != nil {
		return nil, err
	}

	for _, object := range parsedData {
		result := &domain.Hotel{}
		s.handleObject(object, result)
		dataList[result.HotelID] = append(dataList[result.HotelID], result)
	}
	return dataList, nil
}

func (s parser) handleObject(object interface{}, result *domain.Hotel) {
	mapObject, _ := object.(map[string]interface{})
	for key, value := range mapObject {
		s.handleNested(key, value, result)
	}
}

func (s parser) handleNested(key string, data interface{}, result *domain.Hotel) {
	switch keyMapper[key] {
	case domain.HotelCols.HotelID:
		result.HotelID = ParseString(data)
	case domain.HotelCols.Name:
		result.Name = ParseString(data)
	case domain.HotelCols.DestinationID:
		value := ParseFloat64(data)
		result.DestinationID = int(value)
	case domain.HotelCols.Description:
		result.Description = ParseString(data)
	case domain.HotelCols.Location:
		ParseLocation(result, locationMapper[key], data)
	case domain.HotelCols.Amenities:
		result.Amenities = ParseAmenities(data)
	case domain.HotelCols.BookingConditions:
		condition := ParseBookingCondition(data)
		result.BookingConditions = condition
	case domain.HotelCols.Images:
		parseImages(data, result)
	}
}
