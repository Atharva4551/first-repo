package funding

type Fund struct {
	balance iNT
}

func NewFund(initialBalance int) *Fund {

	return &Fund{
		balance: initialBalance,
	}
}

func (f *Fund) Balance() int {
	return f.balance

}

func (f *Fund) Withdraw(amount int) {
	f.balance -= amount
}
