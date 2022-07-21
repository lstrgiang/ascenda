package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lstrgiang/ascenda/internal/app/server/config"
	"github.com/lstrgiang/ascenda/internal/infra/db"
	"github.com/lstrgiang/ascenda/internal/infra/utils/request"
	"github.com/lstrgiang/ascenda/internal/services"
)

func NewServer(cfg config.Config) Server {
	return &server{
		db:              db.NewStorage(),
		cfg:             cfg,
		e:               gin.Default(),
		NewMerger:       services.GetMerger,
		NewParser:       services.GetParser,
		NewHotelService: services.GetHotelService,
	}
}

type Server interface {
	Run()
	ProcessData()
	RegisterHandler()
}

type server struct {
	e               *gin.Engine
	db              *db.Storage
	cfg             config.Config
	NewMerger       services.GetMergerFn
	NewParser       services.GetParserfn
	NewHotelService services.GetHotelServiceFn
}

func (s *server) ProcessData() {
	suppliers := []string{}
	file, err := ioutil.ReadFile(s.cfg.SupplierFilePath)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal([]byte(file), &suppliers); err != nil {
		panic(err)
	}

	if err := services.LoadData(
		suppliers,
		request.GetJsonBody,
		s.NewMerger(),
		s.NewParser(),
		s.NewHotelService(s.db),
	); err != nil {
		panic(err)
	}
}

func (s server) Run() {
	s.e.Run(fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port))
}

func (s *server) path(p string) string {
	return path.Join(s.cfg.Path, p)
}

func (s *server) RegisterHandler() {
	api := s.e
	api.GET(s.path("/hotels"), func(c *gin.Context) {
		hotelId := c.Query("hotel_id")
		destinationId := c.Query("destination_id")
		if hotelId != "" && destinationId != "" {
			destinationIdVal, err := strconv.Atoi(destinationId)
			if err != nil {
				//response error
				c.JSON(http.StatusBadRequest, Response{
					Error: err.Error(),
				})
				return
			}
			hotelService := s.NewHotelService(s.db)
			data := hotelService.GetHotelByDestinationAndHotelId(hotelId, destinationIdVal)
			c.JSON(http.StatusOK, Response{
				Data:    data,
				Message: "Data list retrieved successfully",
			})
			return
		}
		if hotelId != "" {
			hotelService := s.NewHotelService(s.db)
			data := hotelService.GetHotelById(hotelId)
			c.JSON(http.StatusOK, Response{
				Data:    data,
				Message: "Data list retrieved successfully",
			})
			return

		}
		if destinationId != "" {
			destinationIdVal, err := strconv.Atoi(destinationId)
			if err != nil {
				//response error
				c.JSON(http.StatusBadRequest, Response{
					Error: err.Error(),
				})
				return
			}
			hotelService := s.NewHotelService(s.db)
			data := hotelService.GetHotelByDestinationId(destinationIdVal)
			c.JSON(http.StatusOK, Response{
				Data:    data,
				Message: "Data list retrieved successfully",
			})
			return
		}

		// find all
		hotelService := s.NewHotelService(s.db)
		data := hotelService.FindAll()
		c.JSON(http.StatusOK, Response{
			Data:    data,
			Message: "Data list retrieved successfully",
		})
	})
}
