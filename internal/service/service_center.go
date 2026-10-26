package service

import "hw1/internal/domain"

// ServiceCenter определяет, можно ли добавить транспорт в парк
type ServiceCenter interface {
	Accept(transport domain.Transport) bool
}

// SampleServiceCenter — временная реализация сервисного центра. Для полноценной реализации необходимо знать критерии
// отбора подходящего транспорта
type SampleServiceCenter struct{}

// Accept для SampleServiceCenter разрешает добавление любого транспорта
func (SampleServiceCenter) Accept(
	_ domain.Transport,
) bool {
	return true
}
