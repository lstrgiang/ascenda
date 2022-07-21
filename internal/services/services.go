package services

import (
	"github.com/lstrgiang/ascenda/internal/domain"
	"github.com/lstrgiang/ascenda/internal/infra/db"
	"github.com/lstrgiang/ascenda/internal/services/hotel"
	"github.com/lstrgiang/ascenda/internal/services/merger"
	"github.com/lstrgiang/ascenda/internal/services/parser"
)

type GetMergerFn func() merger.Merger
type GetParserfn func() parser.Parser
type GetHotelServiceFn func(db *db.Storage) hotel.HotelService

func GetMerger() merger.Merger {
	return merger.NewMerger()
}
func GetParser() parser.Parser {
	return parser.NewParser()
}

func GetHotelService(db *db.Storage) hotel.HotelService {
	return hotel.NewHotelService(db)
}

func LoadData(
	suppliers []string,
	apiCaller func(string) ([]byte, error),
	merger merger.Merger,
	parser parser.Parser,
	hotelService hotel.HotelService,
) error {
	// get data from suppliers
	parsedData := make(map[string][]*domain.Hotel)
	for _, supplier := range suppliers {
		byteData, err := apiCaller(supplier)
		if err != nil {
			continue
		}
		supplierData, err := parser.ParseData(byteData)
		if err != nil {
			return err
		}
		//merge into main arr
		for key, dataArr := range supplierData {
			if _, ok := parsedData[key]; !ok {
				parsedData[key] = make([]*domain.Hotel, 0)
			}
			parsedData[key] = append(parsedData[key], dataArr...)
		}
	}

	newData := merger.MergeData(parsedData)

	if err := hotelService.InsertMany(newData); err != nil {
		return err
	}
	return nil
}
