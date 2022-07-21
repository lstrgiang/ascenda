package repository

import (
	"errors"

	"github.com/lstrgiang/ascenda/internal/domain"
	"github.com/lstrgiang/ascenda/internal/infra/db"
)

func NewHotelRepository(db *db.Storage) HotelRepository {
	return &hotelRepository{
		db: db,
	}
}

type HotelRepository interface {
	AddNewData(data *domain.Hotel) error
	UpdateData(d *domain.Hotel) error
	GetDataByHotelId(id string) *domain.Hotel
	GetDataByDestinationId(id int) []*domain.Hotel
	GetDataByDestinationIdAndHotelId(hotelId string, destinationId int) *domain.Hotel
	All() []*domain.Hotel
}

type hotelRepository struct {
	db *db.Storage
}

func (r *hotelRepository) InsertMany(d []*domain.Hotel) error {
	for _, data := range d {
		if err := r.AddNewData(data); err != nil {
			return err
		}
	}
	return nil
}
func (r *hotelRepository) AddNewData(d *domain.Hotel) error {
	if r.db.IsHotelIdExist(d.HotelID) {
		return errors.New("data existed")
	}
	// other validation will not be done with assumption that data will be consistent
	r.db.SetHotelIdMap(d.HotelID, d)
	r.db.SetDestinationIdMap(d.DestinationID, d.HotelID, d)
	r.db.SetDestinationHotelMap(d.HotelID, d.DestinationID, d)
	return nil
}

func (r *hotelRepository) UpdateData(d *domain.Hotel) error {
	if !r.db.IsHotelIdExist(d.HotelID) {
		return errors.New("data not existed")
	}

	// other validation will not be done with assumption that data will be consistent
	r.db.SetHotelIdMap(d.HotelID, d)
	r.db.SetDestinationIdMap(d.DestinationID, d.HotelID, d)
	r.db.SetDestinationHotelMap(d.HotelID, d.DestinationID, d)
	return nil
}

func (r hotelRepository) GetDataByHotelId(id string) *domain.Hotel {
	return r.db.GetHotelIdMap()[id]

}
func (r hotelRepository) GetDataByDestinationId(id int) []*domain.Hotel {
	result := make([]*domain.Hotel, 0)
	for _, data := range r.db.GetDestinationIdMap()[id] {
		result = append(result, data)
	}
	return result
}

func (r hotelRepository) GetDataByDestinationIdAndHotelId(hotelId string, destinationId int) *domain.Hotel {
	return r.db.GetDestinationHotelMap(hotelId, destinationId)

}
func (r hotelRepository) All() []*domain.Hotel {
	result := make([]*domain.Hotel, 0)
	for _, data := range r.db.GetHotelIdMap() {
		result = append(result, data)
	}
	return result
}
