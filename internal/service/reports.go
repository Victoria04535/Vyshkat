package service

import (
	"hw1/internal/domain"
	"hw1/internal/repository"
)

// ReportService формирует отчёты по парку.
type ReportService struct {
	storage repository.Storage
}

func NewReportService(storage repository.Storage) *ReportService {
	return &ReportService{
		storage: storage,
	}
}

// TotalTransportCount возвращает количество транспортных средств
func (s *ReportService) TotalTransportCount() int {
	return len(s.storage.Transports())
}

// TotalDailyEnergyConsumption возвращает суммарный суточный расход энергии
func (s *ReportService) TotalDailyEnergyConsumption() float64 {
	var total float64

	for _, item := range s.storage.InventoryItems() {
		if consumer, ok := item.(domain.EnergyConsumer); ok {
			total += consumer.DailyEnergyConsumption()
		}
	}

	return total
}

// BeginnerFriendlyTransport возвращает технику, подходящую для новичков (оценка простоты управления >= 6).
func (s *ReportService) BeginnerFriendlyTransport() []domain.Transport {
	result := make([]domain.Transport, 0)

	for _, transport := range s.storage.Transports() {
		if transport.IsSuitableForBeginners() {
			result = append(result, transport)
		}
	}

	return result
}

// AllInventoryItems возвращает все инвентарные объекты.
func (s *ReportService) AllInventoryItems() []domain.InventoryItem {
	return s.storage.InventoryItems()
}
