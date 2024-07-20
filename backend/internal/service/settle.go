package service

import "slices"

type Settle struct {
	// The members involved with their overall balance.
	// Negative values are debt.
	members  map[string]int
	expenses []Expense
}
type Payment struct {
	From   string
	To     string
	Amount int
}
type Expense struct {
	Amount  int
	Members []string
	Source  string
}

func NewSettle(members []string, expenses []Expense) *Settle {
	membersMap := make(map[string]int)
	for _, m := range members {
		membersMap[m] = 0
	}

	return &Settle{
		members:  membersMap,
		expenses: expenses,
	}
}

// Populates Settle.Members with the calculated values. Assumes that those are all 0.
func (s *Settle) calculateMembersBalance() {
	for _, expense := range s.expenses {
		s.members[expense.Source] += expense.Amount

		amountDebt := expense.Amount / len(expense.Members)
		for _, member := range expense.Members {
			s.members[member] -= amountDebt
		}
	}
}

// Create Payments from Settle.
func (s *Settle) CreatePayments() []Payment {
	s.calculateMembersBalance()

	payments := []Payment{}

	memberKeysSorted := make([]string, len(s.members))
	i := 0
	for k := range s.members {
		memberKeysSorted[i] = k
		i++
	}
	slices.SortStableFunc(memberKeysSorted, func(a, b string) int {
		return s.members[a] - s.members[b]
	})

	right := len(memberKeysSorted) - 1
	for _, member := range memberKeysSorted {
		if s.members[member] >= 0 {
			// we are finished if all debts have been calculated
			break
		}

		for s.members[member] < 0 {
			to := memberKeysSorted[right]
			amount := min(-s.members[member], s.members[to])

			payments = append(payments, Payment{
				To:     to,
				From:   member,
				Amount: amount,
			})

			// update balance
			s.members[member] += amount
			s.members[to] -= amount

			if s.members[to] == 0 {
				right--
			}
		}
	}

	return payments
}
