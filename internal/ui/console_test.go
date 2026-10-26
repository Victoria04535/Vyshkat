package ui

import (
	"bytes"
	"strings"
	"testing"

	"hw1/internal/domain"
	"hw1/internal/repository"
	"hw1/internal/service"
)

type testServiceCenter struct {
	accepted bool
}

func (f testServiceCenter) Accept(_ domain.Transport) bool {
	return f.accepted
}

func newTestConsole(input string) (*Console, *repository.InMemoryStorage, *bytes.Buffer) {
	storage := repository.NewInMemoryStorage()
	center := testServiceCenter{accepted: true}
	park := service.NewParkService(storage, center)
	reports := service.NewReportService(storage)
	output := &bytes.Buffer{}

	console := NewConsole(strings.NewReader(input), output, park, reports)

	return console, storage, output
}

func TestConsole_Run_Exit(t *testing.T) {
	console, _, output := newTestConsole("0\n")

	if err := console.Run(); err != nil {
		t.Fatalf("Run() вернул ошибку: %v", err)
	}

	if !strings.Contains(output.String(), "Завершение работы.") {
		t.Error("ожидалось сообщение о завершении работы")
	}
}

func TestConsole_Run_UnknownCommand(t *testing.T) {
	console, _, output := newTestConsole("9\n0\n")

	if err := console.Run(); err != nil {
		t.Fatalf("Run() вернул ошибку: %v", err)
	}

	if !strings.Contains(output.String(), "Неизвестная команда.") {
		t.Error("ожидалось сообщение о неизвестной команде")
	}
}

func TestConsole_AddTransport(t *testing.T) {
	console, storage, output := newTestConsole(
		"1\n" + // добавить транспорт
			"1\n" + // самокат
			"Самокат\n" +
			"TR-001\n" +
			"7\n" +
			"0\n",
	)

	if err := console.Run(); err != nil {
		t.Fatalf("Run() вернул ошибку: %v", err)
	}

	item, found := storage.FindByInventoryNumber("TR-001")
	if !found {
		t.Fatal("транспорт не найден в хранилище")
	}

	if item.Name() != "Самокат" {
		t.Errorf("имя = %q, ожидалось %q", item.Name(), "Самокат")
	}

	if !strings.Contains(output.String(), "Транспорт успешно добавлен.") {
		t.Error("ожидалось сообщение об успешном добавлении транспорта")
	}
}

func TestConsole_AddTransport_InvalidInput(t *testing.T) {
	console, storage, output := newTestConsole(
		"1\n" +
			"1\n" +
			"Самокат\n" +
			"TR-001\n" +
			"abc\n" +
			"0\n",
	)

	if err := console.Run(); err != nil {
		t.Fatalf("Run() вернул ошибку: %v", err)
	}

	if _, found := storage.FindByInventoryNumber("TR-001"); found {
		t.Error("транспорт не должен добавляться с некорректным значением удобства")
	}

	if !strings.Contains(output.String(), "Ошибка:") {
		t.Error("ожидалось сообщение об ошибке")
	}
}

func TestConsole_AddTransport_ServiceError(t *testing.T) {
	storage := repository.NewInMemoryStorage()
	center := testServiceCenter{accepted: false}
	park := service.NewParkService(storage, center)
	reports := service.NewReportService(storage)
	output := &bytes.Buffer{}

	console := NewConsole(
		strings.NewReader("1\n1\nСамокат\nTR-001\n7\n0\n"),
		output,
		park,
		reports,
	)

	if err := console.Run(); err != nil {
		t.Fatalf("Run() вернул ошибку: %v", err)
	}

	if _, found := storage.FindByInventoryNumber("TR-001"); found {
		t.Error("транспорт, не прошедший техосмотр, не должен добавляться")
	}

	if !strings.Contains(output.String(), "не прошёл технический осмотр") {
		t.Error("ожидалось сообщение о непрохождении техосмотра")
	}
}

func TestConsole_AddInventoryItem(t *testing.T) {
	console, storage, output := newTestConsole(
		"2\n" +
			"1\n" +
			"Шлем\n" +
			"HL-001\n" +
			"M\n" +
			"0\n",
	)

	if err := console.Run(); err != nil {
		t.Fatalf("Run() вернул ошибку: %v", err)
	}

	item, found := storage.FindByInventoryNumber("HL-001")
	if !found {
		t.Fatal("шлем не найден в хранилище")
	}

	if item.Name() != "Шлем" {
		t.Errorf("имя = %q, ожидалось %q", item.Name(), "Шлем")
	}

	if !strings.Contains(output.String(), "Инвентарный предмет успешно добавлен.") {
		t.Error("ожидалось сообщение об успешном добавлении предмета")
	}
}

func TestConsole_ShowInventory(t *testing.T) {
	console, _, output := newTestConsole(
		"2\n1\nШлем\nHL-001\nM\n" +
			"3\n0\n",
	)

	if err := console.Run(); err != nil {
		t.Fatalf("Run() вернул ошибку: %v", err)
	}

	result := output.String()
	if !strings.Contains(result, "Шлем") {
		t.Error("в инвентарном объекте должен отображаться шлем")
	}
	if !strings.Contains(result, "HL-001") {
		t.Error("в инвентарном объекте должен отображаться инвентарный номер")
	}
}

func TestConsole_ShowInventory_Empty(t *testing.T) {
	console, _, output := newTestConsole("3\n0\n")

	if err := console.Run(); err != nil {
		t.Fatalf("Run() вернул ошибку: %v", err)
	}

	if !strings.Contains(output.String(), "Инвентаря и техники нет.") {
		t.Error("ожидалось сообщение об отсутствии инвентаря и техники")
	}
}

func TestConsole_ShowEnergyConsumption(t *testing.T) {
	console, _, output := newTestConsole(
		"2\n2\nЗарядная станция\nCH-001\n3,5\n" +
			"4\n0\n",
	)

	if err := console.Run(); err != nil {
		t.Fatalf("Run() вернул ошибку: %v", err)
	}

	if !strings.Contains(output.String(), "3.50 кВт·ч") {
		t.Error("ожидалось суммарное энергопотребление 3.50 кВт·ч")
	}
}

func TestConsole_ShowBeginnerFriendlyTransport(t *testing.T) {
	console, _, output := newTestConsole(
		"1\n1\nСамокат для новичка\nTR-001\n8\n" +
			"1\n3\nВелосипед\nTR-002\n4\n" +
			"5\n0\n",
	)

	if err := console.Run(); err != nil {
		t.Fatalf("Run() вернул ошибку: %v", err)
	}

	result := output.String()

	if !strings.Contains(result, "Самокат для новичка") {
		t.Error("подходящий для новичка транспорт должен отображаться")
	}

	if strings.Contains(result, "Велосипед (инв. № TR-002") {
		t.Error("транспорт с удобством 4 не должен отображаться в списке для новичков")
	}
}

func TestConsole_ShowBeginnerFriendlyTransport_Empty(t *testing.T) {
	console, _, output := newTestConsole("5\n0\n")

	if err := console.Run(); err != nil {
		t.Fatalf("Run() вернул ошибку: %v", err)
	}

	if !strings.Contains(output.String(), "Подходящий транспорт не найден.") {
		t.Error("ожидалось сообщение об отсутствии подходящего транспорта")
	}
}
