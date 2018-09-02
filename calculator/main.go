package main

import (
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("library", "Library CLI")

	aCLI = app.Command("add", "Command for adding 2 numbers.")
	v1   = aCLI.Arg("v1", "1st value").Required().Int()
	v2   = aCLI.Arg("v2", "2nd value").Required().Int()
)

func main() {
	var calculator *Calculator
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case aCLI.FullCommand():
		calculator = NewCalculator(*v1, *v2)
		fmt.Printf("Addition is :%v\n", calculator.add())
	}
}
