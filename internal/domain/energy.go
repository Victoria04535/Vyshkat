package domain

import "fmt"

type EnergyProfile struct {
	dailyConsumption float64
}

func NewEnergyProfile(dailyConsumption float64) (EnergyProfile, error) {
	if dailyConsumption < 0 {
		return EnergyProfile{}, fmt.Errorf(
			"расход энергии не может быть отрицательным",
		)
	}

	return EnergyProfile{
		dailyConsumption: dailyConsumption,
	}, nil
}

func (e EnergyProfile) DailyEnergyConsumption() float64 {
	return e.dailyConsumption
}
