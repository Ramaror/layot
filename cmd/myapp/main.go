package main

import (
	"errors"
	"fmt"
	"layout/internal"
)

const DEFAULT_DISCOUNT = 50

func main() {
	r := internal.NewCustomer("Ramzan", 21, 1000, 10, true)
	r.CalcDiscount = func() (int, error) {
		if !r.Discount {
			return 0, errors.New("Discount not available")
		}
		result := DEFAULT_DISCOUNT - r.Debt
		if result < 0 {
			return 0, nil
		}
		return result, nil
	}
	fmt.Println(internal.CalcPrice(*r, 500))
}
