package parser

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/lstrgiang/ascenda/internal/domain"
	"github.com/stretchr/testify/assert"
)

var mockData = `[
  {
    "Id": "iJhz",
    "DestinationId": 5432,
    "Name": "Beach Villas Singapore",
    "Latitude": 1.264751,
    "Longitude": 103.824006,
    "Address": " 8 Sentosa Gateway, Beach Villas ",
    "City": "Singapore",
    "Country": "SG",
    "PostalCode": "098269",
    "Description": "  This 5 star hotel is located on the coastline of Singapore.",
    "Facilities": ["Pool", "BusinessCenter", "WiFi ", "DryCleaning", " Breakfast"]
  },
  {
    "Id": "SjyX",
    "DestinationId": 5432,
    "Name": "InterContinental Singapore Robertson Quay",
    "Latitude": null,
    "Longitude": null,
    "Address": " 1 Nanson Road",
    "City": "Singapore",
    "Country": "SG",
    "PostalCode": "238909",
    "Description": "Enjoy sophisticated waterfront living at the new InterContinental® Singapore Robertson Quay, luxury's preferred address nestled in the heart of Robertson Quay along the Singapore River, with the CBD just five minutes drive away. Magnifying the comforts of home, each of our 225 studios and suites features a host of thoughtful amenities that combine modernity with elegance, whilst maintaining functional practicality. The hotel also features a chic, luxurious Club InterContinental Lounge.",
    "Facilities": ["Pool", "WiFi ", "Aircon", "BusinessCenter", "BathTub", "Breakfast", "DryCleaning", "Bar"]
  },
  {
    "Id": "f8c9",
    "DestinationId": 1122,
    "Name": "Hilton Shinjuku Tokyo",
    "Latitude": "",
    "Longitude": "",
    "Address": "160-0023, SHINJUKU-KU, 6-6-2 NISHI-SHINJUKU, JAPAN",
    "City": "Tokyo",
    "Country": "JP",
    "PostalCode": "160-0023",
    "Description": "Hilton Tokyo is located in Shinjuku, the heart of Tokyo's business, shopping and entertainment district, and is an ideal place to experience modern Japan. A complimentary shuttle operates between the hotel and Shinjuku station and the Tokyo Metro subway is connected to the hotel. Relax in one of the modern Japanese-style rooms and admire stunning city views. The hotel offers WiFi and internet access throughout all rooms and public space.",
    "Facilities": ["Pool", "WiFi ", "BusinessCenter", "DryCleaning", " Breakfast", "Bar", "BathTub"]
  }
]`
var mockKeyMap = map[string]bool{
	"SjyX": true,
	"f8c9": true,
	"iJhz": true,
}

func TestParsingFunc(t *testing.T) {
	t.Run("should parse data type 1 successfully", func(t *testing.T) {
		parser := NewParser()
		result, err := parser.ParseData([]byte(mockData))
		assert.Empty(t, err)
		assert.NotEmpty(t, result)
		for key, val := range result {
			assert.Greater(t, len(val), 0)
			assert.True(t, mockKeyMap[key])
		}
	})
}

var mockImageData = `{
      "rooms": [
        { "link": "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/2.jpg", "caption": "Double room" },
        { "link": "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/3.jpg", "caption": "Double room" }
      ],
      "site": [
        { "link": "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/1.jpg", "caption": "Front" }
      ]
    }
`

func TestImageParinsgFunc(t *testing.T) {
	t.Run("should parse image successfully", func(t *testing.T) {
		var data interface{}
		err := json.Unmarshal([]byte(mockImageData), &data)
		if err != nil {
			panic(err)
		}
		hotel := &domain.Hotel{}
		parseImages(data, hotel)
		assert.NotEmpty(t, hotel.Images.Room)
		assert.NotEmpty(t, hotel.Images.Site)
		fmt.Println(hotel)
	})
}
