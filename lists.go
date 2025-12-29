package main

import "fmt"

type product struct {
	id    int32
	title string
	price float64
}

func Lists() {
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

	// -------------------------------- //
	newPrices := []float64{1, 2}
	fmt.Println(newPrices)

	newPrices[1] = 3
	newPrices = append(newPrices, 5)

	fmt.Println(newPrices)

	// -------------------------------- //
	hobbies := [3]string{"Computer Gaming", "Street Photography", "Listening to podcasts/music"}
	fmt.Println(hobbies)

	fmt.Println("First element: ", hobbies[0])
	fmt.Println("Second and third element: ", hobbies[1:3])

	slice1 := hobbies[0:2]
	slice2 := hobbies[:2]

	fmt.Println("Slice 1: ", slice1)
	fmt.Println("Slice 2: ", slice2)

	slice3 := slice1[1:3]
	fmt.Println("Slice 3: ", slice3)

	// -------------------------------- //
	goals := []string{"Learn go", "Create microservices with go"}
	fmt.Println("Goals: ", goals)

	goals[1] = "Test change"
	goals = append(goals, "Add Go to CV")

	fmt.Println("Goals: ", goals)

	// -------------------------------- //
	products := []product{
		product{id: 1, title: "Product 1", price: 11.0},
		product{id: 2, title: "Product 2", price: 12.0},
	}

	fmt.Println("Products: ", products)

	products = append(products, product{id: 3, title: "Product 3", price: 13.0})

	fmt.Println("Products: ", products)

	products2 := []product{
		{id: 4, title: "Product 4", price: 14.0},
		{id: 5, title: "Product 5", price: 15.0},
	}

	products = append(products, products2...)
	fmt.Println("Products: ", products)

	// -------------------------------- //
	testMake := make([]string, 2, 5)

	fmt.Println("Test make size: ", len(testMake))
	fmt.Println("Test make capacity: ", cap(testMake))
}
