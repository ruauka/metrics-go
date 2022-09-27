package response

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"metrics-service-1/internal/logic"
	"metrics-service-1/internal/request"
	"metrics-service-1/internal/utils/dictionary"
)

func TestNewResponse(t *testing.T) {
	data := NewTestData()
	expected := &Response{
		Version:     "v.1.0.0",
		ExecuteDate: time.Now().Format(dictionary.LayoutToString),
		Resp: []Resp{
			{
				data.Result,
			},
		},
	}
	actual := NewResponse(data)
	require.Equal(t, expected, actual)
}

func NewTestData() *logic.Data {
	return &logic.Data{Request: *NewTestReq()}
}

func NewTestReq() *request.Request {
	paymentDateBalance := 0

	return &request.Request{
		Name:    "Ivan",
		Surname: "Ivanov",
		Account: []request.Account{
			{
				PaymentDateBalance: &paymentDateBalance,
			},
		},
		Loan: []request.Loan{
			{
				Payment:           0,
				PaymentDate:       time.Time{},
				ActualPaymentDate: time.Time{},
			},
		},
	}
}
