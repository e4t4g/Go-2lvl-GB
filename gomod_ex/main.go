package main

import (
	"fmt"
	"github.com/dannyvankooten/vat"
)

func main() {
	c, _ := vat.GetCountryRates("GB")
	r, _ := c.GetRate("standard")

	fmt.Printf("Standard VAT rate for Great Britain is %.2f", r)
}
