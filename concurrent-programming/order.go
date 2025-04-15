package main

import "fmt"

type order struct {
	ProductCode int
	Quantity    float64
	Status      orderStatus
}

func (o order) String() string {
	return fmt.Sprintf("ProductCode: %d, Quantity: %f, Status: %s", 
	o.ProductCode, o.Quantity, orderStatusToText(o.Status))
}

func orderStatusToText(status orderStatus) string {
	return []string{ "none", "new", "received", "reserved", "filled" }[status]
}

type orderStatus int

const (
	none orderStatus = iota
	new
	received
	reserved
	filled
)

var orders = []order{}