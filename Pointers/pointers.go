package main

import "fmt"

func main() {
	age := 32
	var agePointer *int
	agePointer = &age
	fmt.Println("Age", agePointer)  //Address of age - 0xc000012098
	fmt.Println("Age", *agePointer) //Value of age

	editAdultYear(agePointer)
	fmt.Println("Age", *agePointer)
	fmt.Println("Age", age)
	//fmt.Println("Adultyear", adultyear)

}

func editAdultYear(age *int) {
	*age = *age - 18
}
