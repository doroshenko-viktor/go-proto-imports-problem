package main

import (
	y "go-program/company"
	x "go-program/company/nested1/nested2"
)

func main() {
	_ = x.CompanyMessage{
		CompEnum: y.CompanyEnum_VAR1,
	}
}
