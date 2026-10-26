package domain

import (
	"fmt"
	"strings"
)

// Helmet — шлем с размером
type Helmet struct {
	InventoryData
	size string
}

func NewHelmet(
	name, inventoryNumber, size string,
) (Helmet, error) {
	data, err := NewInventoryData(name, inventoryNumber)
	if err != nil {
		return Helmet{}, err
	}

	size = strings.TrimSpace(size)
	if size == "" {
		return Helmet{}, fmt.Errorf("размер шлема не должен быть пустым")
	}

	return Helmet{
		InventoryData: data,
		size:          size,
	}, nil
}

func (h Helmet) Size() string {
	return h.size
}

// ChargingStation — зарядная станция
type ChargingStation struct {
	InventoryData
	EnergyProfile
}

func NewChargingStation(
	name, inventoryNumber string,
	dailyConsumption float64,
) (ChargingStation, error) {
	data, err := NewInventoryData(name, inventoryNumber)
	if err != nil {
		return ChargingStation{}, err
	}

	energy, err := NewEnergyProfile(dailyConsumption)
	if err != nil {
		return ChargingStation{}, err
	}

	return ChargingStation{
		InventoryData: data,
		EnergyProfile: energy,
	}, nil
}

// DockingStation — док-станция
type DockingStation struct {
	InventoryData
	EnergyProfile
}

func NewDockingStation(
	name, inventoryNumber string,
	dailyConsumption float64,
) (DockingStation, error) {
	data, err := NewInventoryData(name, inventoryNumber)
	if err != nil {
		return DockingStation{}, err
	}

	energy, err := NewEnergyProfile(dailyConsumption)
	if err != nil {
		return DockingStation{}, err
	}

	return DockingStation{
		InventoryData: data,
		EnergyProfile: energy,
	}, nil
}
