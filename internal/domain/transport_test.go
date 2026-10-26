package domain

import "testing"

func TestNewTransportBase(t *testing.T) {
	t.Run("отклоняет оценку ниже допустимого диапазона", func(t *testing.T) {
		_, err := NewTransportBase("Самокат", "TR-001", 0)
		if err == nil {
			t.Fatal("ожидали ошибку для оценки 0")
		}
	})

	t.Run("отклоняет оценку выше допустимого диапазона", func(t *testing.T) {
		_, err := NewTransportBase("Самокат", "TR-001", 11)
		if err == nil {
			t.Fatal("ожидали ошибку для оценки 11")
		}
	})

	t.Run("отклоняет пустое название", func(t *testing.T) {
		_, err := NewTransportBase("", "TR-001", 7)
		if err == nil {
			t.Fatal("ожидали ошибку для пустого названия")
		}
	})

	t.Run("отклоняет пустой инвентарный номер", func(t *testing.T) {
		_, err := NewTransportBase("Самокат", "", 7)
		if err == nil {
			t.Fatal("ожидали ошибку для пустого инвентарного номера")
		}
	})
}

func TestTransportBaseIsSuitableForBeginners(t *testing.T) {
	tests := []struct {
		name      string
		easeOfUse int
		want      bool
	}{
		{
			name:      "оценка ниже порога",
			easeOfUse: 5,
			want:      false,
		},
		{
			name:      "оценка равна порогу",
			easeOfUse: 6,
			want:      true,
		},
		{
			name:      "оценка выше порога",
			easeOfUse: 8,
			want:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transport, err := NewTransportBase(
				"Транспорт",
				"TR-001",
				tt.easeOfUse,
			)
			if err != nil {
				t.Fatalf("не удалось создать транспорт: %v", err)
			}

			if got := transport.IsSuitableForBeginners(); got != tt.want {
				t.Errorf(
					"IsSuitableForBeginners() = %t, ожидали %t",
					got,
					tt.want,
				)
			}
		})
	}
}

func TestNewScooter(t *testing.T) {
	t.Run("создаёт самокат с корректными данными", func(t *testing.T) {
		_, err := NewScooter("Самокат", "SC-001", 7)
		if err != nil {
			t.Fatalf("не ожидали ошибку: %v", err)
		}
	})

	t.Run("отклоняет недопустимую оценку", func(t *testing.T) {
		_, err := NewScooter("Самокат", "SC-001", 0)
		if err == nil {
			t.Fatal("ожидали ошибку для недопустимой оценки")
		}
	})
}

func TestNewElectricScooter(t *testing.T) {
	t.Run("создаёт электросамокат с корректными данными", func(t *testing.T) {
		_, err := NewElectricScooter("Электросамокат", "ES-001", 8, 2.5)
		if err != nil {
			t.Fatalf("не ожидали ошибку: %v", err)
		}
	})

	t.Run("отклоняет недопустимую оценку", func(t *testing.T) {
		_, err := NewElectricScooter("Электросамокат", "ES-001", 11, 2.5)
		if err == nil {
			t.Fatal("ожидали ошибку для недопустимой оценки")
		}
	})

	t.Run("отклоняет отрицательный расход энергии", func(t *testing.T) {
		_, err := NewElectricScooter("Электросамокат", "ES-001", 8, -1)
		if err == nil {
			t.Fatal("ожидали ошибку для отрицательного расхода энергии")
		}
	})
}

func TestNewBicycle(t *testing.T) {
	t.Run("создаёт велосипед с корректными данными", func(t *testing.T) {
		_, err := NewBicycle("Велосипед", "BI-001", 6)
		if err != nil {
			t.Fatalf("не ожидали ошибку: %v", err)
		}
	})

	t.Run("отклоняет недопустимую оценку", func(t *testing.T) {
		_, err := NewBicycle("Велосипед", "BI-001", 0)
		if err == nil {
			t.Fatal("ожидали ошибку для недопустимой оценки")
		}
	})
}

func TestNewElectricBicycle(t *testing.T) {
	t.Run("создаёт электровелосипед с корректными данными", func(t *testing.T) {
		_, err := NewElectricBicycle("Электровелосипед", "EB-001", 7, 1.8)
		if err != nil {
			t.Fatalf("не ожидали ошибку: %v", err)
		}
	})

	t.Run("отклоняет недопустимую оценку", func(t *testing.T) {
		_, err := NewElectricBicycle("Электровелосипед", "EB-001", 0, 1.8)
		if err == nil {
			t.Fatal("ожидали ошибку для недопустимой оценки")
		}
	})

	t.Run("отклоняет отрицательный расход энергии", func(t *testing.T) {
		_, err := NewElectricBicycle("Электровелосипед", "EB-001", 7, -1)
		if err == nil {
			t.Fatal("ожидали ошибку для отрицательного расхода энергии")
		}
	})
}
