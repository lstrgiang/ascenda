package db

import (
	"fmt"

	"github.com/lstrgiang/ascenda/internal/domain"
)

const (
	DestinationHotelKeyPattern = "%d-%s"
)

func NewStorage() *Storage {
	return &Storage{
		hotelIdMap:            make(map[string]*domain.Hotel),
		destinationIdMap:      make(map[int]map[string]*domain.Hotel),
		destinationHotelIdMap: make(map[string]*domain.Hotel),
	}

}

type Storage struct {
	hotelIdMap            map[string]*domain.Hotel         // like an indexing version
	destinationIdMap      map[int]map[string]*domain.Hotel // like an indexing version
	destinationHotelIdMap map[string]*domain.Hotel
}

func (s Storage) GetHotelIdMap() map[string]*domain.Hotel {
	return s.hotelIdMap
}

func (s Storage) GetDestinationIdMap() map[int]map[string]*domain.Hotel {
	return s.destinationIdMap
}

func (s *Storage) SetHotelIdMap(key string, data *domain.Hotel) {
	s.hotelIdMap[key] = data
}

func (s *Storage) IsHotelIdExist(key string) bool {
	_, ok := s.hotelIdMap[key]
	return ok
}

func (s *Storage) IsDestinationIdExist(key int) bool {
	_, ok := s.destinationIdMap[key]
	return ok
}

func (s *Storage) SetDestinationIdMap(destinationId int, hotelId string, data *domain.Hotel) {
	if _, ok := s.destinationIdMap[destinationId]; !ok {
		s.destinationIdMap[destinationId] = make(map[string]*domain.Hotel)
	}
	s.destinationIdMap[destinationId][hotelId] = data
}

func (s *Storage) SetDestinationHotelMap(hotelId string, destinationId int, data *domain.Hotel) {
	s.destinationHotelIdMap[s.destinationHotelIdKey(hotelId, destinationId)] = data
}

func (s *Storage) GetDestinationHotelMap(hotelId string, destinationId int) *domain.Hotel {
	return s.destinationHotelIdMap[s.destinationHotelIdKey(hotelId, destinationId)]
}

func (s Storage) destinationHotelIdKey(hotelId string, destinationId int) string {
	return fmt.Sprintf(DestinationHotelKeyPattern, destinationId, hotelId)
}
