package domain

import "testing"

func TestNewInventoryData(t *testing.T) {
	t.Run("сохраняет корректные данные", func(t *testing.T) {
		item, err := NewInventoryData("Самокат", "TR-001")
		if err != nil {
			t.Fatalf("не ожидали ошибку: %v", err)
		}

		if got := item.Name(); got != "Самокат" {
			t.Errorf("Name() = %q, ожидали %q", got, "Самокат")
		}

		if got := item.InventoryNumber(); got != "TR-001" {
			t.Errorf(
				"InventoryNumber() = %q, ожидали %q",
				got,
				"TR-001",
			)
		}
	})

	t.Run("обрезает пробелы по краям", func(t *testing.T) {
		item, err := NewInventoryData("  Самокат  ", " TR-001 ")
		if err != nil {
			t.Fatalf("не ожидали ошибку: %v", err)
		}

		if got := item.Name(); got != "Самокат" {
			t.Errorf("Name() = %q, ожидали %q", got, "Самокат")
		}

		if got := item.InventoryNumber(); got != "TR-001" {
			t.Errorf(
				"InventoryNumber() = %q, ожидали %q",
				got,
				"TR-001",
			)
		}
	})

	t.Run("отклоняет пустое название", func(t *testing.T) {
		_, err := NewInventoryData("", "TR-001")
		if err == nil {
			t.Fatal("ожидали ошибку для пустого названия")
		}
	})

	t.Run("отклоняет название из пробелов", func(t *testing.T) {
		_, err := NewInventoryData("   ", "TR-001")
		if err == nil {
			t.Fatal("ожидали ошибку для названия из пробелов")
		}
	})

	t.Run("отклоняет пустой инвентарный номер", func(t *testing.T) {
		_, err := NewInventoryData("Самокат", "")
		if err == nil {
			t.Fatal("ожидали ошибку для пустого номера")
		}
	})

	t.Run("отклоняет номер из пробелов", func(t *testing.T) {
		_, err := NewInventoryData("Самокат", "   ")
		if err == nil {
			t.Fatal("ожидали ошибку для номера из пробелов")
		}
	})
}

func TestInventoryDataValidate(t *testing.T) {
	t.Run("корректные данные проходят проверку", func(t *testing.T) {
		item, err := NewInventoryData("Самокат", "TR-001")
		if err != nil {
			t.Fatalf("не удалось создать объект: %v", err)
		}

		if err := item.Validate(); err != nil {
			t.Errorf("Validate() вернул неожиданную ошибку: %v", err)
		}
	})
}
