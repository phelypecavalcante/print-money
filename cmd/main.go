package main

import (
	"fmt"
	"github.com/phelypecavalcante/print-money/pkg/currency"
)

func main() {
	c, _ := currency.New(55, 100, "BRL")
	fmt.Println(c)
	fmt.Println(c.Amount())
	fmt.Println(c.Amount().Float64())
	fmt.Println(c.Amount().Int64())
}
