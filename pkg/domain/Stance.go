package domain

type Stance int

const (
	Default Stance = iota
	Fakie   Stance = iota
	Nollie  Stance = iota
)

func (s Stance) String() string {
	return [...]string{"", "Fakie", "Nollie"}[s]
}
