package domain

import (
	"fmt"
	"strings"
)

type InventoryData struct {
	name            string
	inventoryNumber string
}

func NewInventoryData(name, inventoryNumber string) (InventoryData, error) {
	data := InventoryData{
		name:            strings.TrimSpace(name),
		inventoryNumber: strings.TrimSpace(inventoryNumber),
	}

	if err := data.Validate(); err != nil {
		return InventoryData{}, err
	}

	return data, nil
}

func (d InventoryData) Name() string {
	return d.name
}

func (d InventoryData) InventoryNumber() string {
	return d.inventoryNumber
}

func (d InventoryData) Validate() error {
	if d.name == "" {
		return fmt.Errorf("название не должно быть пустым")
	}

	if d.inventoryNumber == "" {
		return fmt.Errorf("инвентарный номер не должен быть пустым")
	}

	return nil
}
