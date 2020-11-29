package main

import (
	"fmt"

	"github.com/hooklift/gowsdl/soap"

	"github.com/hiwyw/go-test/soap/calc"
)

func main() {
	// soapCli := soap.NewClient("http://www.dneonline.com/calculator.asmx?wsdl")
	soapCli := soap.NewClient("http://localhost:8080")
	soap := calc.NewCalculatorSoap(soapCli)

	res, err := soap.Add(&calc.Add{
		IntA: 1,
		IntB: 1,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res.AddResult)
}
