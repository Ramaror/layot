package main

import (
	"errors"
	"fmt"
	"layout/internal"
)

func main() {
	cust := internal.NewCustomer("Ramzan", 23, 10000, 1000, true)
	fmt.Println(internal.CalcPrice(cust, 500))
	cust.WrOffDebt()
	fmt.Printf("%+v\n", cust)
	StartTransactionDynamic(cust)

	cc := &internal.Overduer{
		&internal.Customer{
			Debt:    5,
			Balance: 1000,
		},
	}
	fmt.Printf("%+v\n", cc)
}
func StartTransactionDynamic(debtor interface{}) error {
	wroffdebt, ok := debtor.(internal.Debtor)
	if !ok {
		return errors.New("incorrect type")
	}

	return wroffdebt.WrOffDebt()
}
