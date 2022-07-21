package parser

import (
	"encoding/json"

	"github.com/lstrgiang/ascenda/internal/domain"
)

type ParseFn func(interface{}) interface{}

func ParseInt(data interface{}) int {
	val, _ := data.(int)
	return val
}

func ParseFloat64(data interface{}) float64 {
	val, _ := data.(float64)
	return val
}

func ParseString(data interface{}) string {
	val, _ := data.(string)
	return val
}

func ParseStringArray(data interface{}) []string {
	val, _ := data.([]string)
	return val
}

func ArrayToStringArray(data []interface{}) []string {
	result := make([]string, 0)
	for _, val := range data {
		result = append(result, val.(string))
	}
	return result
}

func marshalAmenities(data map[string][]string) domain.Amenities {
	jsonObj, _ := json.Marshal(data)
	result := domain.Amenities{}
	json.Unmarshal(jsonObj, &result)
	return result
}

func ParseAmenities(data interface{}) domain.Amenities {
	switch t := data.(type) {
	case []interface{}:
		return domain.Amenities{
			General: ArrayToStringArray(t),
		}
	case map[string]interface{}:
		tempData := make(map[string][]string)
		for key, val := range t {
			targetKey := amenitiesMapper[key]
			tempData[targetKey] = ArrayToStringArray(val.([]interface{}))
		}
		return marshalAmenities(tempData)
	}
	return domain.Amenities{}
}

func parseLocationRow(target *domain.Hotel, key string, data interface{}) {
	switch key {
	case domain.LocationCols.PostalCode:
		target.Location.PostalCode = data.(string)
	case domain.LocationCols.Address:
		target.Location.Address = data.(string)
	case domain.LocationCols.City:
		target.Location.City = data.(string)
	case domain.LocationCols.Country:
		target.Location.Country = data.(string)
	case domain.LocationCols.Lat:
		val, _ := data.(float64)
		target.Location.Lat = val
	case domain.LocationCols.Lng:
		val, _ := data.(float64)
		target.Location.Lng = val
	}
}
func ParseLocation(target *domain.Hotel, key string, data interface{}) {
	switch t := data.(type) {
	case float64:
		parseLocationRow(target, key, t)
	case string:
		parseLocationRow(target, key, t)
	case map[string]interface{}:
		for key, val := range t {
			parseLocationRow(target, key, val)
		}
	}
}

func ParseBookingCondition(data interface{}) []string {
	dataArr := data.([]interface{})
	return ArrayToStringArray(dataArr)
}

func parseImageDetailList(data interface{}) []domain.ImageDetail {
	result := make([]domain.ImageDetail, 0)
	switch t := data.(type) {
	case []interface{}:
		for _, dataRow := range t {
			mapDataRow := dataRow.(map[string]interface{})
			image := domain.ImageDetail{}
			for key, val := range mapDataRow {
				switch imageDetailMapper[key] {
				case domain.ImageDetailCols.Link:
					image.Link = val.(string)
				case domain.ImageDetailCols.Description:
					image.Description = val.(string)
				}
			}
			result = append(result, image)
		}
	}
	return result
}

func parseImageDetail(key string, data interface{}, result *domain.Hotel) {
	switch key {
	case domain.ImageCols.Amenities:
		result.Images.Amenities = append(result.Images.Amenities, parseImageDetailList(data)...)
	case domain.ImageCols.Room:
		result.Images.Room = append(result.Images.Room, parseImageDetailList(data)...)
	case domain.ImageCols.Site:
		result.Images.Site = append(result.Images.Site, parseImageDetailList(data)...)
	}
}

func parseImages(data interface{}, result *domain.Hotel) {
	switch t := data.(type) {
	case map[string]interface{}:
		for key, val := range t {
			parseImageDetail(imageMapper[key], val, result)
		}
	}
}
