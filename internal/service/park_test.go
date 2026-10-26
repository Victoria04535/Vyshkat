package service

import (
	"strings"
	"testing"

	"hw1/internal/domain"
	"hw1/internal/repository"
)

// fakeServiceCenter позволяет подменить техосмотр в тестах.
type fakeServiceCenter struct {
	accepted bool
	calls    int
}

func (f *fakeServiceCenter) Accept(transport domain.Transport) bool {
	f.calls++
	return f.accepted
}

func TestParkService_AddTransport(t *testing.T) {
	t.Run("добавляет транспорт после успешного техосмотра", func(t *testing.T) {
		repo := repository.NewInMemoryStorage()
		center := &fakeServiceCenter{accepted: true}
		park := NewParkService(repo, center)

		scooter, err := domain.NewScooter("Самокат", "TR-001", 8)
		if err != nil {
			t.Fatalf("не удалось создать самокат: %v", err)
		}

		if err := park.AddTransport(scooter); err != nil {
			t.Fatalf("ожидалось успешное добавление, получена ошибка: %v", err)
		}

		if center.calls != 1 {
			t.Errorf("ожидался 1 вызов техосмотра, получено: %d", center.calls)
		}

		if _, exists := repo.FindByInventoryNumber("TR-001"); !exists {
			t.Error("транспорт не найден в репозитории после добавления")
		}
	})

	t.Run("не добавляет транспорт после неуспешного техосмотра", func(t *testing.T) {
		repo := repository.NewInMemoryStorage()
		center := &fakeServiceCenter{accepted: false}
		park := NewParkService(repo, center)

		scooter, err := domain.NewScooter("Самокат", "TR-002", 8)
		if err != nil {
			t.Fatalf("не удалось создать самокат: %v", err)
		}

		err = park.AddTransport(scooter)
		if err == nil {
			t.Fatal("ожидалась ошибка при провале техосмотра")
		}

		if !strings.Contains(err.Error(), "не прошёл технический осмотр") {
			t.Errorf("неожиданный текст ошибки: %v", err)
		}

		if center.calls != 1 {
			t.Errorf("ожидался 1 вызов техосмотра, получено: %d", center.calls)
		}

		if _, exists := repo.FindByInventoryNumber("TR-002"); exists {
			t.Error("транспорт не должен добавляться после провала техосмотра")
		}
	})

}

func TestParkService_AddItem(t *testing.T) {
	t.Run("добавляет инвентарный объект без техосмотра", func(t *testing.T) {
		repo := repository.NewInMemoryStorage()
		center := &fakeServiceCenter{accepted: false}
		park := NewParkService(repo, center)

		helmet, err := domain.NewHelmet("Шлем", "IT-001", "M")
		if err != nil {
			t.Fatalf("не удалось создать шлем: %v", err)
		}

		if err := park.AddItem(helmet); err != nil {
			t.Fatalf("не удалось добавить инвентарный объект: %v", err)
		}

		if center.calls != 0 {
			t.Errorf(
				"техосмотр не должен вызываться при добавлении инвентарного объекта, вызовов: %d",
				center.calls,
			)
		}

		if _, exists := repo.FindByInventoryNumber("IT-001"); !exists {
			t.Error("инвентарный объект не найден в репозитории")
		}
	})
}
