// Package logic - основная логика.
package logic

import (
	"metrics-service-2/internal/request"
)

// Data - словарь данных.
type Data struct {
	// входящий JSON
	request.Request
	// Основная логика
	Delinquency
	// Логика формирования выходных агрегатов
	Result
}

// NewData - конструктор словаря данных.
func NewData(inputMessage *request.Request) *Data {
	return &Data{
		Request:     *inputMessage,
		Delinquency: *NewLocal(),
	}
}

// LocalCount - вызов Delinquency логики.
func (d *Data) LocalCount() {
	d.Delinquency.EnoughMoneyCount(d)
	d.Delinquency.DelinquencyCount(d)
	d.Delinquency.DelinquencyDurationCount(d)
	d.Delinquency.DelinquencySumCount(d)
}

// ResultCount - вызов Result логики.
func (d *Data) ResultCount() {
	d.Result.DelinquencyFinal(d)
	d.Result.EnoughMoneyFinal(d)
	d.Result.DelinquencyDurationFinal(d)
	d.Result.DelinquencySumTotalCount(d)
}
