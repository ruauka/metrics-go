// Package response - ответ сервиса.
package response

import (
	"time"

	"metrics-service-2/internal/logic"
	"metrics-service-2/internal/utils/dictionary"
)

// Response - Структура ответа.
type Response struct {
	Version     string `json:"version"`
	ExecuteDate string `json:"execute_date"`
	Resp        []Resp `json:"resp"`
}

// Resp - блок из основной логики.
type Resp struct {
	logic.Result
}

// NewResponse - конструктор объекта ответа.
func NewResponse(data *logic.Data) *Response {
	return &Response{
		Version:     "v.1.0.0",
		ExecuteDate: time.Now().Format(dictionary.LayoutToString),
		Resp: []Resp{
			{
				data.Result,
			},
		},
	}
}
