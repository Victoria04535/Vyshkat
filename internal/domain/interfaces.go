package domain

// InventoryItem — объект с инвентарным номером
type InventoryItem interface {
	Name() string
	InventoryNumber() string
}

// Transport — транспорт
type Transport interface {
	InventoryItem
	EaseOfUse() int
	IsSuitableForBeginners() bool
}

// EnergyConsumer — объект, потребляющий энергию
type EnergyConsumer interface {
	DailyEnergyConsumption() float64
}
