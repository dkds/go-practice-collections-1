package main

import "fmt"

func Maps() {
	websites := map[string]string{
		"Google": "https://google.com",
		"AWS":    "https://aws.com",
	}
	fmt.Println("websites: ", websites)
	fmt.Println("AWS address: ", websites["AWS"])

	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println("websites: ", websites)

	delete(websites, "Google")
	fmt.Println("websites: ", websites)

	// -------------------------------- //
	testMake := make(map[string]float64, 3)

	testMake["key1"] = 1.1
	testMake["key2"] = 2.2
	testMake["key3"] = 3.3

	fmt.Println("Test make map: ", testMake)

	testMake["key4"] = 4.4

	fmt.Println("Test make map: ", testMake)
}
