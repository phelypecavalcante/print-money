package currency

type (
	Amount struct {
		value float64
		unit  int // currency smallest unit
	}
	Currency struct {
		amount Amount
		code   string // code following ISO 4217
		number string // number following ISO 4217
		name   string // name following ISO 4217
	}
)

// New id should be a currency Name, Code or Number following ISO 4217
// e.g: Brazilian Real	BRL	986
func New(a float64, u int, id string) (*Currency, error) {
	attr, err := newAttributes(id)
	if err != nil {
		return nil, err
	}
	return &Currency{
		amount: Amount{
			value: a,
			unit:  u,
		},
		code:   attr.code,
		number: attr.number,
		name:   attr.name,
	}, nil
}

func (c Currency) Amount() Amount {
	return c.amount
}

func (a Amount) Float64() float64 {
	return a.value
}

func (a Amount) Int64() int64 {
	return int64(a.value * float64(a.unit))
}
