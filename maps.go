package main

import "fmt"

type floatMap map[string]float64

func (f floatMap) output() {
	fmt.Println(f)
}

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
	testMake := make(floatMap, 3)

	testMake["key1"] = 1.1
	testMake["key2"] = 2.2
	testMake["key3"] = 3.3

	testMake.output()

	testMake["key4"] = 4.4

	testMake.output()

	// -------------------------------- //
	for name, url := range websites {
		fmt.Println("Name: ", name, "URL:", url)
	}
}
