package repository

import (
	"fmt"
	"hw1/internal/domain"
)

type Storage interface {
	AddInventory(item domain.InventoryItem) error
	Transports() []domain.Transport
	InventoryItems() []domain.InventoryItem
	FindByInventoryNumber(number string) (domain.InventoryItem, bool)
}

// InMemoryStorage хранит объекты в оперативной памяти
type InMemoryStorage struct {
	items             []domain.InventoryItem
	inventoryByNumber map[string]domain.InventoryItem
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		items:             make([]domain.InventoryItem, 0),
		inventoryByNumber: make(map[string]domain.InventoryItem),
	}
}

// Add добавляет инвентарный объект, если его номер уникален.
func (r *InMemoryStorage) AddInventory(
	item domain.InventoryItem,
) error {
	number := item.InventoryNumber()

	if _, exists := r.inventoryByNumber[number]; exists {
		return fmt.Errorf(
			"инвентарный номер %q уже используется",
			number,
		)
	}

	r.items = append(r.items, item)
	r.inventoryByNumber[number] = item

	return nil
}

// Transports возвращает все объекты, реализующие интерфейс Transport
func (r *InMemoryStorage) Transports() []domain.Transport {
	transports := make([]domain.Transport, 0)

	for _, item := range r.items {
		if transport, ok := item.(domain.Transport); ok {
			transports = append(transports, transport)
		}
	}

	return transports
}

// InventoryItems возвращает копию списка инвентарных объектов
func (r *InMemoryStorage) InventoryItems() []domain.InventoryItem {
	result := make([]domain.InventoryItem, len(r.items))
	copy(result, r.items)

	return result
}

// FindByInventoryNumber ищет объект по инвентарному номеру
func (r *InMemoryStorage) FindByInventoryNumber(
	number string,
) (domain.InventoryItem, bool) {
	item, exists := r.inventoryByNumber[number]
	return item, exists
}
