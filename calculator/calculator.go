package main

type Calculator struct {
	var1 int
	var2 int
}

func (c *Calculator) add() int {
	return c.var1 + c.var2
}

func NewCalculator(v1, v2 int) *Calculator {
	return &Calculator{var1: v1, var2: v2}
}
