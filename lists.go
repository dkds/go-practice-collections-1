package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"A book"}
	fmt.Println(productNames)

	productNames[2] = "A Carpet"

	prices := [4]float64{11.99, 9.99, 49.99, 19.99}
	fmt.Println(prices)
}
