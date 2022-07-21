package hotel

import (
	"github.com/lstrgiang/ascenda/internal/domain"
	"github.com/lstrgiang/ascenda/internal/infra/db"
	"github.com/lstrgiang/ascenda/internal/infra/repository"
)

func NewHotelService(db *db.Storage) HotelService {
	return &hotelService{
		hotelRepository: repository.NewHotelRepository(db),
	}
}

type HotelService interface {
	NewHotel(data *domain.Hotel) error
	UpdateHotel(data *domain.Hotel) error
	GetHotelById(hotelId string) *domain.Hotel
	GetHotelByDestinationId(destinationId int) []*domain.Hotel
	GetHotelByDestinationAndHotelId(hotelId string, destinationId int) *domain.Hotel
	FindAll() []*domain.Hotel
	InsertMany(d []*domain.Hotel) error
}

type hotelService struct {
	hotelRepository repository.HotelRepository
}

func (r hotelService) NewHotel(data *domain.Hotel) error {
	return r.hotelRepository.AddNewData(data)
}

func (r hotelService) InsertMany(d []*domain.Hotel) error {
	for _, data := range d {
		if err := r.hotelRepository.AddNewData(data); err != nil {
			return err
		}
	}
	return nil
}

func (r hotelService) UpdateHotel(data *domain.Hotel) error {
	return r.hotelRepository.UpdateData(data)
}

func (r hotelService) GetHotelById(hotelId string) *domain.Hotel {
	return r.hotelRepository.GetDataByHotelId(hotelId)
}

func (r hotelService) GetHotelByDestinationId(destinationId int) []*domain.Hotel {
	return r.hotelRepository.GetDataByDestinationId(destinationId)
}

func (r hotelService) GetHotelByDestinationAndHotelId(hotelId string, destinationId int) *domain.Hotel {
	return r.hotelRepository.GetDataByDestinationIdAndHotelId(hotelId, destinationId)
}

func (r hotelService) FindAll() []*domain.Hotel {
	return r.hotelRepository.All()
}
