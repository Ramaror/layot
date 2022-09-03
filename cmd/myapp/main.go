package main

import (
	"fmt"
	"layout/internal"
)

func main() {
	cust := internal.NewCustomer("Ramzan", 23, 10000, 1000, true)
	fmt.Println(internal.CalcPrice(cust, 500))
	cust.WrOffDebt()
	fmt.Printf("%+v\n", cust)
}
