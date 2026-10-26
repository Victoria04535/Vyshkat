package repository

import (
	"hw1/internal/domain"
	"testing"
)

func TestNewInMemoryStorage(t *testing.T) {
	storage := NewInMemoryStorage()

	if storage == nil {
		t.Fatal("ожидали созданный репозиторий")
	}

	if got := len(storage.InventoryItems()); got != 0 {
		t.Errorf("количество объектов = %d, ожидали 0", got)
	}

	if got := len(storage.Transports()); got != 0 {
		t.Errorf("количество транспортных средств = %d, ожидали 0", got)
	}
}

func TestAddInventoryItem(t *testing.T) {
	t.Run("добавляет инвентарный объект", func(t *testing.T) {
		storage := NewInMemoryStorage()

		helmet, err := domain.NewHelmet("Шлем", "HM-001", "M")
		if err != nil {
			t.Fatalf("не удалось создать шлем: %v", err)
		}

		if err := storage.AddInventory(helmet); err != nil {
			t.Fatalf("не ожидали ошибку при добавлении: %v", err)
		}

		items := storage.InventoryItems()
		if len(items) != 1 {
			t.Fatalf("количество объектов = %d, ожидали 1", len(items))
		}

		got, exists := storage.FindByInventoryNumber("HM-001")
		if !exists {
			t.Fatal("добавленный объект не найден")
		}

		if got.InventoryNumber() != "HM-001" {
			t.Errorf(
				"InventoryNumber() = %q, ожидали %q",
				got.InventoryNumber(),
				"HM-001",
			)
		}
	})

	t.Run("отклоняет повторяющийся инвентарный номер", func(t *testing.T) {
		storage := NewInMemoryStorage()

		first, err := domain.NewHelmet("Шлем первый", "HM-001", "M")
		if err != nil {
			t.Fatalf("не удалось создать первый шлем: %v", err)
		}

		second, err := domain.NewHelmet("Шлем второй", "HM-001", "L")
		if err != nil {
			t.Fatalf("не удалось создать второй шлем: %v", err)
		}

		if err := storage.AddInventory(first); err != nil {
			t.Fatalf("не удалось добавить первый шлем: %v", err)
		}

		if err := storage.AddInventory(second); err == nil {
			t.Fatal("ожидали ошибку при повторном инвентарном номере")
		}

		if got := len(storage.InventoryItems()); got != 1 {
			t.Errorf(
				"количество объектов = %d, ожидали 1",
				got,
			)
		}
	})

	t.Run("разрешает разные инвентарные номера", func(t *testing.T) {
		storage := NewInMemoryStorage()

		first, err := domain.NewHelmet("Шлем", "HM-001", "M")
		if err != nil {
			t.Fatalf("не удалось создать первый шлем: %v", err)
		}

		second, err := domain.NewHelmet("Шлем", "HM-002", "L")
		if err != nil {
			t.Fatalf("не удалось создать второй шлем: %v", err)
		}

		if err := storage.AddInventory(first); err != nil {
			t.Fatalf("не удалось добавить первый шлем: %v", err)
		}

		if err := storage.AddInventory(second); err != nil {
			t.Fatalf("не удалось добавить второй шлем: %v", err)
		}

		if got := len(storage.InventoryItems()); got != 2 {
			t.Errorf("количество объектов = %d, ожидали 2", got)
		}
	})
}

func TestTransports(t *testing.T) {
	storage := NewInMemoryStorage()

	scooter, err := domain.NewScooter("Самокат", "SC-001", 7)
	if err != nil {
		t.Fatalf("не удалось создать самокат: %v", err)
	}

	helmet, err := domain.NewHelmet("Шлем", "HM-001", "M")
	if err != nil {
		t.Fatalf("не удалось создать шлем: %v", err)
	}

	electricScooter, err := domain.NewElectricScooter(
		"Электросамокат",
		"ES-001",
		8,
		2.5,
	)
	if err != nil {
		t.Fatalf("не удалось создать электросамокат: %v", err)
	}

	for _, item := range []domain.InventoryItem{scooter, helmet, electricScooter} {
		if err := storage.AddInventory(item); err != nil {
			t.Fatalf("не удалось добавить объект: %v", err)
		}
	}

	transports := storage.Transports()
	if len(transports) != 2 {
		t.Fatalf("количество транспортных средств = %d, ожидали 2", len(transports))
	}

	for _, transport := range transports {
		if transport.InventoryNumber() == "HM-001" {
			t.Error("шлем не должен попадать в список транспорта")
		}
	}
}

func TestInventoryItems(t *testing.T) {
	storage := NewInMemoryStorage()

	helmet, err := domain.NewHelmet("Шлем", "HM-001", "M")
	if err != nil {
		t.Fatalf("не удалось создать шлем: %v", err)
	}

	if err := storage.AddInventory(helmet); err != nil {
		t.Fatalf("не удалось добавить шлем: %v", err)
	}

	items := storage.InventoryItems()
	if len(items) != 1 {
		t.Fatalf("количество объектов = %d, ожидали 1", len(items))
	}

	// Изменение возвращённого среза не должно менять внутренний список.
	items[0] = nil

	remaining := storage.InventoryItems()
	if len(remaining) != 1 {
		t.Fatalf("количество объектов после изменения копии = %d, ожидали 1", len(remaining))
	}

	if remaining[0] == nil {
		t.Error("изменение возвращённого среза затронуло внутренний список")
	}
}

func TestFindByInventoryNumber(t *testing.T) {
	storage := NewInMemoryStorage()

	helmet, err := domain.NewHelmet("Шлем", "HM-001", "M")
	if err != nil {
		t.Fatalf("не удалось создать шлем: %v", err)
	}

	if err := storage.AddInventory(helmet); err != nil {
		t.Fatalf("не удалось добавить шлем: %v", err)
	}

	t.Run("находит существующий объект", func(t *testing.T) {
		item, exists := storage.FindByInventoryNumber("HM-001")
		if !exists {
			t.Fatal("ожидали найти объект")
		}

		if item.InventoryNumber() != "HM-001" {
			t.Errorf(
				"InventoryNumber() = %q, ожидали %q",
				item.InventoryNumber(),
				"HM-001",
			)
		}
	})

	t.Run("возвращает отсутствие для неизвестного номера", func(t *testing.T) {
		item, exists := storage.FindByInventoryNumber("UNKNOWN")
		if exists {
			t.Fatal("не ожидали найти объект с неизвестным номером")
		}

		if item != nil {
			t.Errorf("ожидали nil, получили %v", item)
		}
	})
}
