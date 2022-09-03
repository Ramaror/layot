package main

import (
	"errors"
	"fmt"
	"layout/internal"
)

const DEFAULT_DISCOUNT = 50000

func main() {
	r := internal.Customer{Name: "Ramzan", Age: 23, Discount: true, Debt: 5000, Balance: 10000}
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
	fmt.Println(internal.CalcPrice(r, 500))
}
