package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestCalculator(t *testing.T) {
	t.Run("Should add the 2 numbers", func(t *testing.T) {
		a := 1
		b := 2

		c := &Calculator{var1: a, var2: b}
		res := c.add()

		if res != 3 {
			t.Fatal("Expected 3, got %v", res)
		}
	})
}
