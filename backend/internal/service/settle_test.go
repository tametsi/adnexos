package service

import (
	"fmt"
	"testing"

	"github.com/MarvinJWendt/testza"
)

func TestCreatePayments(t *testing.T) {
	tests := []struct {
		settle           *Settle
		paymentsExpected []Payment
	}{
		{
			NewSettle([]string{"a", "b", "c"}, []Expense{
				{Source: "b", Members: []string{"a"}, Amount: 3},
				{Source: "a", Members: []string{"b"}, Amount: 2},
				{Source: "c", Members: []string{"a"}, Amount: 5},
				{Source: "c", Members: []string{"b"}, Amount: 7},
			}), []Payment{
				{From: "a", To: "c", Amount: 6},
				{From: "b", To: "c", Amount: 6},
			},
		},
		{
			NewSettle([]string{"a", "b", "c", "d"}, []Expense{
				{Source: "a", Members: []string{"a", "b", "c"}, Amount: 6},
				{Source: "a", Members: []string{"a", "b"}, Amount: 4},
				{Source: "b", Members: []string{"a", "d"}, Amount: 2},
				{Source: "c", Members: []string{"a", "b", "c", "d"}, Amount: 16},
			}), []Payment{
				{From: "b", To: "c", Amount: 6},
				{From: "d", To: "c", Amount: 4},
				{From: "d", To: "a", Amount: 1},
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test #%d", i), func(t *testing.T) {
			payments := tt.settle.CreatePayments()
			testza.AssertSameElements(t, tt.paymentsExpected, payments)
		})
	}
}
