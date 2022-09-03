package internal

type Customer struct {
	Name         string
	Age          int
	Balance      int
	Debt         int
	Discount     bool
	CalcDiscount func() (int, error)
}

func CalcPrice(c Customer, sum int) (int, error) {
	discount, err := c.CalcDiscount()
	if err != nil {
		return 0, err
	}
	sum = sum - discount
	return sum, nil
}
