package main

import (
	"errors"
	"fmt"
	mathFunction "github.com/emretask1n/Calculator/math"
	"os"
	"strings"
)

type Calculator struct {
	functions []mathFunction.MathFunction
}

func (c *Calculator) addMathFunction(m mathFunction.MathFunction) {
	c.functions = append(c.functions, m)
}

func (c *Calculator) doCalculation(name string, arg float64) (float64, error) {
	var result float64
	for _, f := range c.functions {
		if strings.ToLower(name) == strings.ToLower(f.GetName()) {
			result = f.Calculate(arg)
			return result, nil
		}
	}
	return 0, errors.New("no such function exists:" + name)
}

func main() {
	startCalculator()
}

func startCalculator() {
	flag := true

	myCalculator := Calculator{}

	myCalculator.addMathFunction(mathFunction.Sin{Name: "Sine"})
	myCalculator.addMathFunction(mathFunction.Cos{Name: "Cosine"})
	myCalculator.addMathFunction(mathFunction.Log{Name: "Log"})

	fmt.Println("\nCalculator started.")
	fmt.Println("You can calculate using following functions")
	for _, f := range myCalculator.functions {
		fmt.Println(f.GetName())
	}

	for flag {
		var fName string
		fmt.Println("> Enter name of the calculation or enter x to exit:")
		_, err := fmt.Scanf("%s", &fName)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		//
		if fName == "x" {
			//os.Exit(0)
			flag = false
		} else {
			var arg float64
			fmt.Println("> Enter a value for the calculation:")
			_, err := fmt.Scan(&arg)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
			value, err2 := myCalculator.doCalculation(fName, arg)
			if err2 != nil {
				fmt.Println(err2)
			} else {
				fmt.Printf("Result of %s of %f : %f", fName, arg, value)
				os.Exit(0)
			}
		}
	}
	println("Bye!")
}
