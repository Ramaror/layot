package internal

import "errors"

const DEFAULT_DISCOUNT = 50

type Customer struct {
	Name     string
	Age      int
	Balance  int
	Debt     int
	discount bool
}

func (c *Customer) CalcDiscount() (int, error) {
	if !c.discount {
		return 0, errors.New("Discount not available")
	}
	result := DEFAULT_DISCOUNT - c.Debt
	if result < 0 {
		return 0, nil
	}
	return result, nil
}

func (c *Customer) WrOffDebt() error {
	if c.Debt >= c.Balance {
		return errors.New("Not possible write off")
	}
	c.Balance -= c.Debt
	c.Debt = 0
	return nil
}
func NewCustomer(name string, age int, balance int, debt int, discount bool) *Customer {
	return &Customer{
		Name:     name,
		Age:      age,
		Balance:  balance,
		Debt:     debt,
		discount: discount,
	}
}

// измените функцию так, чтобы на вход она теперь принимала объект
// реализующий интерфейс Discounter
func CalcPrice(c Discounter, sum int) (int, error) {
	discount, err := c.CalcDiscount()
	if err != nil {
		return 0, err
	}
	sum = sum - discount
	return sum, nil
}
