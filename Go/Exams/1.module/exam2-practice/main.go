package main

import (
	"github.com/fatih/color"	
)

func main(){
	color.Green("sum of 3 + 2 : %d\n", Add(3, 2))
	color.Blue("sum of 3 - 2 : %d\n", Subtract(3, 2))
}

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}