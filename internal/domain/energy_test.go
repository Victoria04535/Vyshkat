package domain

import "testing"

func TestNewEnergyProfile(t *testing.T) {
	t.Run("сохраняет корректный расход энергии", func(t *testing.T) {
		profile, err := NewEnergyProfile(2.5)
		if err != nil {
			t.Fatalf("не ожидали ошибку: %v", err)
		}

		if got := profile.DailyEnergyConsumption(); got != 2.5 {
			t.Errorf(
				"DailyEnergyConsumption() = %v, ожидали %v",
				got,
				2.5,
			)
		}
	})

	t.Run("допускает нулевой расход энергии", func(t *testing.T) {
		profile, err := NewEnergyProfile(0)
		if err != nil {
			t.Fatalf("не ожидали ошибку для нулевого расхода: %v", err)
		}

		if got := profile.DailyEnergyConsumption(); got != 0 {
			t.Errorf(
				"DailyEnergyConsumption() = %v, ожидали 0",
				got,
			)
		}
	})

	t.Run("отклоняет отрицательный расход энергии", func(t *testing.T) {
		_, err := NewEnergyProfile(-1)
		if err == nil {
			t.Fatal("ожидали ошибку для отрицательного расхода")
		}
	})
}
