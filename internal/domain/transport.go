package domain

import "fmt"

type TransportBase struct {
	InventoryData
	easeOfUse int
}

func NewTransportBase(
	name, inventoryNumber string,
	easeOfUse int,
) (TransportBase, error) {
	data, err := NewInventoryData(name, inventoryNumber)
	if err != nil {
		return TransportBase{}, err
	}

	transport := TransportBase{
		InventoryData: data,
		easeOfUse:     easeOfUse,
	}

	if err := transport.Validate(); err != nil {
		return TransportBase{}, err
	}

	return transport, nil
}

func (t TransportBase) EaseOfUse() int {
	return t.easeOfUse
}

func (t TransportBase) IsSuitableForBeginners() bool {
	return t.easeOfUse >= 6
}

func (t TransportBase) Validate() error {
	if err := t.InventoryData.Validate(); err != nil {
		return err
	}

	if t.easeOfUse < 1 || t.easeOfUse > 10 {
		return fmt.Errorf(
			"оценка простоты управления должна быть от 1 до 10",
		)
	}

	return nil
}

// Scooter — обычный самокат
type Scooter struct {
	TransportBase
}

func NewScooter(
	name, inventoryNumber string,
	easeOfUse int,
) (Scooter, error) {
	base, err := NewTransportBase(name, inventoryNumber, easeOfUse)
	if err != nil {
		return Scooter{}, err
	}

	return Scooter{TransportBase: base}, nil
}

// ElectricScooter — электросамокат
type ElectricScooter struct {
	TransportBase
	EnergyProfile
}

func NewElectricScooter(
	name, inventoryNumber string,
	easeOfUse int,
	dailyConsumption float64,
) (ElectricScooter, error) {
	base, err := NewTransportBase(name, inventoryNumber, easeOfUse)
	if err != nil {
		return ElectricScooter{}, err
	}

	energy, err := NewEnergyProfile(dailyConsumption)
	if err != nil {
		return ElectricScooter{}, err
	}

	return ElectricScooter{
		TransportBase: base,
		EnergyProfile: energy,
	}, nil
}

// Bicycle — обычный велосипед
type Bicycle struct {
	TransportBase
}

func NewBicycle(
	name, inventoryNumber string,
	easeOfUse int,
) (Bicycle, error) {
	base, err := NewTransportBase(name, inventoryNumber, easeOfUse)
	if err != nil {
		return Bicycle{}, err
	}

	return Bicycle{TransportBase: base}, nil
}

// ElectricBicycle — электровелосипед
type ElectricBicycle struct {
	TransportBase
	EnergyProfile
}

func NewElectricBicycle(
	name, inventoryNumber string,
	easeOfUse int,
	dailyConsumption float64,
) (ElectricBicycle, error) {
	base, err := NewTransportBase(name, inventoryNumber, easeOfUse)
	if err != nil {
		return ElectricBicycle{}, err
	}

	energy, err := NewEnergyProfile(dailyConsumption)
	if err != nil {
		return ElectricBicycle{}, err
	}

	return ElectricBicycle{
		TransportBase: base,
		EnergyProfile: energy,
	}, nil
}
