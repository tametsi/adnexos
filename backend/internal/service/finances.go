package service

type Finances struct {
	Members map[string]int
	Map     [][]int
}
type Payment struct {
	From   string
	To     string
	Amount int
}
type Expenses struct {
	Amount  int
	Members []string
	Source  string
}

func (f Finances) memberByIndex(index int) string {
	for m, i := range f.Members {
		if i == index {
			return m
		}
	}

	return ""
}

// adjust finances from expenses
func (f *Finances) AddExpenses(expenses []Expenses) {
	for _, e := range expenses {
		sourceIndex := f.Members[e.Source]
		l := len(e.Members)

		for _, member := range e.Members {
			if member == e.Source {
				continue
			}

			memberIndex := f.Members[member]
			f.Map[sourceIndex][memberIndex] += e.Amount / l // CONSIDER using float32
		}
	}
}

// merge from-to finances
func (f *Finances) Merge() {
	for toIndex, fromMap := range f.Map {
		for fromIndex, amount := range fromMap {
			if amount == 0 {
				continue
			}

			amount -= f.Map[fromIndex][toIndex]
			if amount > 0 {
				f.Map[toIndex][fromIndex] = amount
				f.Map[fromIndex][toIndex] = 0
			} else if amount < 0 {
				f.Map[toIndex][fromIndex] = 0
				f.Map[fromIndex][toIndex] = -amount
			} else {
				f.Map[toIndex][fromIndex] = 0
				f.Map[fromIndex][toIndex] = 0
				continue
			}
		}
	}
}

// create awesome payments from `Finances`
func (f *Finances) CreatePayments() []Payment {
	payments := []Payment{}

	for toIndex, fromMap := range f.Map {
		to := f.memberByIndex(toIndex)
		for fromIndex, amount := range fromMap {
			if amount == 0 {
				continue
			}

			payments = append(payments, Payment{
				To:     to,
				From:   f.memberByIndex(fromIndex),
				Amount: amount,
			})
		}
	}

	return payments
}
