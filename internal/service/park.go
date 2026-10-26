package service

import (
	"fmt"

	"hw1/internal/domain"
	"hw1/internal/repository"
)

type ParkService struct {
	repository    repository.Storage
	serviceCenter ServiceCenter
}

func NewParkService(
	repo repository.Storage,
	serviceCenter ServiceCenter,
) *ParkService {
	return &ParkService{
		repository:    repo,
		serviceCenter: serviceCenter,
	}
}

// AddTransport добавляет транспорт после технического осмотра
func (s *ParkService) AddTransport(transport domain.Transport) error {
	if !s.serviceCenter.Accept(transport) {
		return fmt.Errorf(
			"транспорт %q не прошёл технический осмотр",
			transport.Name(),
		)
	}

	return s.repository.AddInventory(transport)
}

// AddItem добавляет инвентарный объект без технического осмотра
func (s *ParkService) AddItem(item domain.InventoryItem) error {
	return s.repository.AddInventory(item)
}
