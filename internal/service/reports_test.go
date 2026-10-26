package service

import (
	"testing"

	"hw1/internal/domain"
	"hw1/internal/repository"
)

func TestReportService_TotalTransportCount(t *testing.T) {
	repo := repository.NewInMemoryStorage()
	reports := NewReportService(repo)

	if got := reports.TotalTransportCount(); got != 0 {
		t.Fatalf("ожидалось 0 транспортных средств, получено: %d", got)
	}

	scooter, err := domain.NewScooter("Самокат", "TR-001", 8)
	if err != nil {
		t.Fatalf("не удалось создать самокат: %v", err)
	}

	helmet, err := domain.NewHelmet("Шлем", "IT-001", "M")
	if err != nil {
		t.Fatalf("не удалось создать шлем: %v", err)
	}

	if err := repo.AddInventory(scooter); err != nil {
		t.Fatalf("не удалось добавить самокат: %v", err)
	}
	if err := repo.AddInventory(helmet); err != nil {
		t.Fatalf("не удалось добавить шлем: %v", err)
	}

	if got := reports.TotalTransportCount(); got != 1 {
		t.Errorf("ожидалось 1 транспортное средство, получено: %d", got)
	}
}

func TestReportService_TotalDailyEnergyConsumption(t *testing.T) {
	repo := repository.NewInMemoryStorage()
	reports := NewReportService(repo)

	if got := reports.TotalDailyEnergyConsumption(); got != 0 {
		t.Fatalf("ожидался расход 0 для пустого хранилища, получено: %v", got)
	}

	scooter, err := domain.NewElectricScooter("Электросамокат", "TR-001", 8, 2.5)
	if err != nil {
		t.Fatalf("не удалось создать электросамокат: %v", err)
	}

	charger, err := domain.NewChargingStation("Зарядная станция", "IT-001", 3.5)
	if err != nil {
		t.Fatalf("не удалось создать зарядную станцию: %v", err)
	}

	helmet, err := domain.NewHelmet("Шлем", "IT-002", "M")
	if err != nil {
		t.Fatalf("не удалось создать шлем: %v", err)
	}

	for _, item := range []domain.InventoryItem{scooter, charger, helmet} {
		if err := repo.AddInventory(item); err != nil {
			t.Fatalf("не удалось добавить объект: %v", err)
		}
	}

	if got, want := reports.TotalDailyEnergyConsumption(), 6.0; got != want {
		t.Errorf("ожидался суммарный расход %v, получено: %v", want, got)
	}
}

func TestReportService_BeginnerFriendlyTransport(t *testing.T) {
	repo := repository.NewInMemoryStorage()
	reports := NewReportService(repo)

	scooter, err := domain.NewScooter("Простой самокат", "TR-001", 6)
	if err != nil {
		t.Fatalf("не удалось создать самокат: %v", err)
	}

	bicycle, err := domain.NewBicycle("Сложный велосипед", "TR-002", 5)
	if err != nil {
		t.Fatalf("не удалось создать велосипед: %v", err)
	}

	electricScooter, err := domain.NewElectricScooter("Электросамокат", "TR-003", 9, 2.0)
	if err != nil {
		t.Fatalf("не удалось создать электросамокат: %v", err)
	}

	for _, transport := range []domain.Transport{scooter, bicycle, electricScooter} {
		if err := repo.AddInventory(transport); err != nil {
			t.Fatalf("не удалось добавить транспорт: %v", err)
		}
	}

	got := reports.BeginnerFriendlyTransport()

	if len(got) != 2 {
		t.Fatalf("ожидалось 2 транспортных средства для новичков, получено: %d", len(got))
	}

	gotNumbers := make(map[string]bool)
	for _, transport := range got {
		gotNumbers[transport.InventoryNumber()] = true
	}

	if !gotNumbers["TR-001"] {
		t.Error("самокат с оценкой 6 должен подходить новичкам")
	}
	if !gotNumbers["TR-003"] {
		t.Error("электросамокат с оценкой 9 должен подходить новичкам")
	}
	if gotNumbers["TR-002"] {
		t.Error("велосипед с оценкой 5 не должен подходить новичкам")
	}
}

func TestReportService_AllInventoryItems(t *testing.T) {
	repo := repository.NewInMemoryStorage()
	reports := NewReportService(repo)

	if got := reports.AllInventoryItems(); len(got) != 0 {
		t.Fatalf("ожидался пустой список, получено объектов: %d", len(got))
	}

	scooter, err := domain.NewScooter("Самокат", "TR-001", 8)
	if err != nil {
		t.Fatalf("не удалось создать самокат: %v", err)
	}

	helmet, err := domain.NewHelmet("Шлем", "IT-001", "M")
	if err != nil {
		t.Fatalf("не удалось создать шлем: %v", err)
	}

	if err := repo.AddInventory(scooter); err != nil {
		t.Fatalf("не удалось добавить самокат: %v", err)
	}
	if err := repo.AddInventory(helmet); err != nil {
		t.Fatalf("не удалось добавить шлем: %v", err)
	}

	got := reports.AllInventoryItems()
	if len(got) != 2 {
		t.Fatalf("ожидалось 2 инвентарных объекта, получено: %d", len(got))
	}

	gotNumbers := make(map[string]bool)
	for _, item := range got {
		gotNumbers[item.InventoryNumber()] = true
	}

	if !gotNumbers["TR-001"] || !gotNumbers["IT-001"] {
		t.Error("в отчёте отсутствует один или несколько добавленных объектов")
	}
}
