package logic

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"metrics-service-2/internal/request"
)

func TestNewData(t *testing.T) {
	paymentDateBalance := 0

	expected := &Data{
		Request: request.Request{
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
		},
		Delinquency: Delinquency{},
		Result:      Result{},
	}
	actual := NewData(NewTestReq())
	require.Equal(t, expected, actual)
}
