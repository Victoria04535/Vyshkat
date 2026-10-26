package ui

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"

	"hw1/internal/domain"
	"hw1/internal/service"
)

type Console struct {
	input   *bufio.Scanner
	output  io.Writer
	park    *service.ParkService
	reports *service.ReportService
}

func NewConsole(
	input io.Reader,
	output io.Writer,
	park *service.ParkService,
	reports *service.ReportService,
) *Console {
	return &Console{
		input:   bufio.NewScanner(input),
		output:  output,
		park:    park,
		reports: reports,
	}
}

func (c *Console) Run() error {
	for {
		c.showMenu()

		choice, err := c.readLine("Выберите действие: ")
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}

		switch choice {
		case "1":
			c.addTransport()
		case "2":
			c.addInventoryItem()
		case "3":
			c.showInventory()
		case "4":
			c.showEnergyConsumption()
		case "5":
			c.showBeginnerFriendlyTransport()
		case "6":
			c.showInventoryCounts()
		case "0":
			fmt.Fprintln(c.output, "Завершение работы.")
			return nil
		default:
			fmt.Fprintln(c.output, "Неизвестная команда.")
		}
	}
}

func (c *Console) showMenu() {
	fmt.Fprintln(c.output, "\n=== ВышКат ===")
	fmt.Fprintln(c.output, "1. Добавить транспорт")
	fmt.Fprintln(c.output, "2. Добавить инвентарный предмет")
	fmt.Fprintln(c.output, "3. Показать всю технику и инвентарь")
	fmt.Fprintln(c.output, "4. Показать суточное энергопотребление")
	fmt.Fprintln(c.output, "5. Показать транспорт для новичков")
	fmt.Fprintln(c.output, "6. Показать количество транспортных средств и инвентаря")
	fmt.Fprintln(c.output, "0. Выход")
}

func (c *Console) showInventoryCounts() {
	transports := c.reports.TotalTransportCount()
	inventory := len(c.reports.AllInventoryItems())

	fmt.Fprintf(
		c.output,
		"Количество транспортных средств: %d. Количество инвентаря: %d\n",
		transports,
		inventory-transports,
	)
}

func (c *Console) addTransport() {
	kind, err := c.readLine(
		"Тип транспорта (1 — самокат, 2 — электросамокат, 3 — велосипед, 4 — электровелосипед): ",
	)
	if err != nil {
		c.printError(err)
		return
	}

	name, err := c.readLine("Название: ")
	if err != nil {
		c.printError(err)
		return
	}

	number, err := c.readLine("Инвентарный номер: ")
	if err != nil {
		c.printError(err)
		return
	}

	ease, err := c.readInt("Удобство для новичка (1–10): ")
	if err != nil {
		c.printError(err)
		return
	}

	var transport domain.Transport

	switch kind {
	case "1":
		transport, err = domain.NewScooter(name, number, ease)

	case "2":
		var consumption float64
		consumption, err = c.readFloat(
			"Расход энергии за день (кВт·ч): ",
		)
		if err == nil {
			transport, err = domain.NewElectricScooter(
				name, number, ease, consumption,
			)
		}

	case "3":
		transport, err = domain.NewBicycle(name, number, ease)

	case "4":
		var consumption float64
		consumption, err = c.readFloat(
			"Расход энергии за день (кВт·ч): ",
		)
		if err == nil {
			transport, err = domain.NewElectricBicycle(
				name, number, ease, consumption,
			)
		}

	default:
		fmt.Fprintln(c.output, "Неизвестный тип транспорта.")
		return
	}

	if err != nil {
		c.printError(err)
		return
	}

	if err := c.park.AddTransport(transport); err != nil {
		c.printError(err)
		return
	}

	fmt.Fprintln(c.output, "Транспорт успешно добавлен.")
}

func (c *Console) addInventoryItem() {
	kind, err := c.readLine(
		"Тип предмета (1 — шлем, 2 — зарядная станция, 3 — станция парковки): ",
	)
	if err != nil {
		c.printError(err)
		return
	}

	name, err := c.readLine("Название: ")
	if err != nil {
		c.printError(err)
		return
	}

	number, err := c.readLine("Инвентарный номер: ")
	if err != nil {
		c.printError(err)
		return
	}

	var item domain.InventoryItem

	switch kind {
	case "1":
		size, readErr := c.readLine("Размер шлема: ")
		if readErr != nil {
			c.printError(readErr)
			return
		}

		item, err = domain.NewHelmet(name, number, size)

	case "2", "3":
		consumption, readErr := c.readFloat(
			"Расход энергии за день (кВт·ч): ",
		)
		if readErr != nil {
			c.printError(readErr)
			return
		}

		if kind == "2" {
			item, err = domain.NewChargingStation(
				name, number, consumption,
			)
		} else {
			item, err = domain.NewDockingStation(
				name, number, consumption,
			)
		}

	default:
		fmt.Fprintln(c.output, "Неизвестный тип предмета.")
		return
	}

	if err != nil {
		c.printError(err)
		return
	}

	if err := c.park.AddItem(item); err != nil {
		c.printError(err)
		return
	}

	fmt.Fprintln(c.output, "Инвентарный предмет успешно добавлен.")
}

func (c *Console) showInventory() {
	items := c.reports.AllInventoryItems()

	if len(items) == 0 {
		fmt.Fprintln(c.output, "Инвентаря и техники нет.")
		return
	}

	fmt.Fprintln(c.output, "\n=== Вся техника и инвентарь ===")

	for _, item := range items {
		fmt.Fprintf(
			c.output,
			"- %s (инв. № %s)\n",
			item.Name(),
			item.InventoryNumber(),
		)
	}
}

func (c *Console) showEnergyConsumption() {
	total := c.reports.TotalDailyEnergyConsumption()

	fmt.Fprintf(
		c.output,
		"Суммарный расход энергии за день: %.2f кВт·ч\n",
		total,
	)
}

func (c *Console) showBeginnerFriendlyTransport() {
	transports := c.reports.BeginnerFriendlyTransport()

	if len(transports) == 0 {
		fmt.Fprintln(c.output, "Подходящий транспорт не найден.")
		return
	}

	fmt.Fprintln(c.output, "\n=== Транспорт для новичков ===")

	for _, transport := range transports {
		fmt.Fprintf(
			c.output,
			"- %s (инв. № %s, удобство: %d/10)\n",
			transport.Name(),
			transport.InventoryNumber(),
			transport.EaseOfUse(),
		)
	}
}

func (c *Console) readLine(prompt string) (string, error) {
	fmt.Fprint(c.output, prompt)

	if !c.input.Scan() {
		if err := c.input.Err(); err != nil {
			return "", err
		}
		return "", io.EOF
	}

	return strings.TrimSpace(c.input.Text()), nil
}

func (c *Console) readInt(prompt string) (int, error) {
	value, err := c.readLine(prompt)
	if err != nil {
		return 0, err
	}

	number, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf(
			"ожидалось целое число: %w",
			err,
		)
	}

	return number, nil
}

func (c *Console) readFloat(prompt string) (float64, error) {
	value, err := c.readLine(prompt)
	if err != nil {
		return 0, err
	}

	value = strings.ReplaceAll(value, ",", ".")

	number, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, fmt.Errorf(
			"ожидалось число: %w",
			err,
		)
	}
	if math.IsNaN(number) || math.IsInf(number, 0) {
		return 0, fmt.Errorf("некорректное число: %q", value)
	}

	return number, nil
}

func (c *Console) printError(err error) {
	fmt.Fprintf(c.output, "Ошибка: %v\n", err)
}
