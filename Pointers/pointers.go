package main

import "fmt"

func main() {
	age := 32
	var agePointer *int
	agePointer = &age
	fmt.Println("Age", agePointer)  //Address of age - 0xc000012098
	fmt.Println("Age", *agePointer) //Value of age

	adultyear := getAdultYear(agePointer)
	fmt.Println("Adultyear", adultyear)

}

func getAdultYear(age *int) int {
	return *age - 18
}
