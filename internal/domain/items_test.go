package domain

import "testing"

func TestNewHelmet(t *testing.T) {
	t.Run("создаёт шлем с корректными данными", func(t *testing.T) {
		helmet, err := NewHelmet("Шлем", "HM-001", "M")
		if err != nil {
			t.Fatalf("не ожидали ошибку: %v", err)
		}

		if got := helmet.Name(); got != "Шлем" {
			t.Errorf("Name() = %q, ожидали %q", got, "Шлем")
		}

		if got := helmet.InventoryNumber(); got != "HM-001" {
			t.Errorf(
				"InventoryNumber() = %q, ожидали %q",
				got,
				"HM-001",
			)
		}

		if got := helmet.Size(); got != "M" {
			t.Errorf("Size() = %q, ожидали %q", got, "M")
		}
	})

	t.Run("отклоняет пустое название", func(t *testing.T) {
		_, err := NewHelmet("", "HM-001", "M")
		if err == nil {
			t.Fatal("ожидали ошибку для пустого названия")
		}
	})

	t.Run("отклоняет пустой инвентарный номер", func(t *testing.T) {
		_, err := NewHelmet("Шлем", "", "M")
		if err == nil {
			t.Fatal("ожидали ошибку для пустого инвентарного номера")
		}
	})

	t.Run("отклоняет пустой размер", func(t *testing.T) {
		_, err := NewHelmet("Шлем", "HM-001", "")
		if err == nil {
			t.Fatal("ожидали ошибку для пустого размера")
		}
	})
}

func TestNewChargingStation(t *testing.T) {
	t.Run("создаёт зарядную станцию с корректными данными", func(t *testing.T) {
		station, err := NewChargingStation(
			"Зарядная станция",
			"CS-001",
			2.5,
		)
		if err != nil {
			t.Fatalf("не ожидали ошибку: %v", err)
		}

		if got := station.Name(); got != "Зарядная станция" {
			t.Errorf(
				"Name() = %q, ожидали %q",
				got,
				"Зарядная станция",
			)
		}

		if got := station.InventoryNumber(); got != "CS-001" {
			t.Errorf(
				"InventoryNumber() = %q, ожидали %q",
				got,
				"CS-001",
			)
		}

		if got := station.DailyEnergyConsumption(); got != 2.5 {
			t.Errorf(
				"DailyEnergyConsumption() = %v, ожидали %v",
				got,
				2.5,
			)
		}
	})

	t.Run("отклоняет пустое название", func(t *testing.T) {
		_, err := NewChargingStation("", "CS-001", 2.5)
		if err == nil {
			t.Fatal("ожидали ошибку для пустого названия")
		}
	})

	t.Run("отклоняет пустой инвентарный номер", func(t *testing.T) {
		_, err := NewChargingStation("Зарядная станция", "", 2.5)
		if err == nil {
			t.Fatal("ожидали ошибку для пустого инвентарного номера")
		}
	})

	t.Run("отклоняет отрицательный расход энергии", func(t *testing.T) {
		_, err := NewChargingStation("Зарядная станция", "CS-001", -1)
		if err == nil {
			t.Fatal("ожидали ошибку для отрицательного расхода энергии")
		}
	})
}

func TestNewDockingStation(t *testing.T) {
	t.Run("создаёт док-станцию с корректными данными", func(t *testing.T) {
		station, err := NewDockingStation(
			"Док-станция",
			"DS-001",
			1.75,
		)
		if err != nil {
			t.Fatalf("не ожидали ошибку: %v", err)
		}

		if got := station.Name(); got != "Док-станция" {
			t.Errorf(
				"Name() = %q, ожидали %q",
				got,
				"Док-станция",
			)
		}

		if got := station.InventoryNumber(); got != "DS-001" {
			t.Errorf(
				"InventoryNumber() = %q, ожидали %q",
				got,
				"DS-001",
			)
		}

		if got := station.DailyEnergyConsumption(); got != 1.75 {
			t.Errorf(
				"DailyEnergyConsumption() = %v, ожидали %v",
				got,
				1.75,
			)
		}
	})

	t.Run("отклоняет пустое название", func(t *testing.T) {
		_, err := NewDockingStation("", "DS-001", 1.75)
		if err == nil {
			t.Fatal("ожидали ошибку для пустого названия")
		}
	})

	t.Run("отклоняет пустой инвентарный номер", func(t *testing.T) {
		_, err := NewDockingStation("Док-станция", "", 1.75)
		if err == nil {
			t.Fatal("ожидали ошибку для пустого инвентарного номера")
		}
	})

	t.Run("отклоняет отрицательный расход энергии", func(t *testing.T) {
		_, err := NewDockingStation("Док-станция", "DS-001", -1)
		if err == nil {
			t.Fatal("ожидали ошибку для отрицательного расхода энергии")
		}
	})
}
