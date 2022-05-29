package types

type Currency string

type Card struct {
	ID       int
	PAN      string
	Balance  int
	Currency Currency
	Color    string
	Name     string
	Active   bool
}
