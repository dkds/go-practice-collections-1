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
}
