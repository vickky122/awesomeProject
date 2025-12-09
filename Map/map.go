package main

import "fmt"

func main() {
	websites := map[string]string{
		"google":   "google.com",
		"youtube":  "youtube.com",
		"facebook": "facebook.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["google"])

	websites["twitter"] = "twitter.com"
	fmt.Println(websites)

	delete(websites, "google")
	fmt.Println(websites)
}
