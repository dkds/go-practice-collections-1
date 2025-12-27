package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"A book"}
	fmt.Println(productNames)

	productNames[2] = "A Carpet"

	prices := [4]float64{1, 2, 3, 4}
	fmt.Println(prices)

	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices)

	featuredPrices = prices[:3]
	fmt.Println(featuredPrices)

	featuredPrices = prices[1:]
	fmt.Println(featuredPrices)

	featuredPrices[0] = 199.99
	fmt.Println(featuredPrices)
	fmt.Println(prices)

	pricesLength := len(prices)
	featuredPricesLength := len(featuredPrices)
	fmt.Println(pricesLength)
	fmt.Println(featuredPricesLength)

	pricesCapacity := cap(prices)
	featuredPricesCapacity := cap(featuredPrices)
	fmt.Println("Prices: ", prices, pricesLength, pricesCapacity)
	fmt.Println("Featured Prices: ", featuredPrices, featuredPricesLength, featuredPricesCapacity)

	highlightedPrices := featuredPrices[:1]
	highlightedPricesLength := len(highlightedPrices)
	highlightedPricesCapacity := cap(highlightedPrices)
	fmt.Println("Highlighted Prices: ", highlightedPrices, highlightedPricesLength, highlightedPricesCapacity)
}
