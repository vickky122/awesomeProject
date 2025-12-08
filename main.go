package main

import (
	"fmt"
	"math"
)

func main() {
	//	fmt.Printf("hi")
	//
	//	//solve the error what is coming here
	//	//compute() is not defined
	//	//so define it
	//	func compute() (int, error) {
	//		return 0, nil
	//	}
	//
	//	if v, err := compute(); err != nil {
	//		return err
	//	}
	//	fmt.Println(v)
	//

	//if x := 10; x > 5 {
	//	fmt.Println(x)
	//}
	//x := 20
	//fmt.Println(x)
	//
	//

	////for style
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}
	//
	////for in while style
	//x := 0
	//for x < 5 {
	//	fmt.Println(x)
	//	x++
	//}
	//
	//s := "नमस्ते"
	//for idx, r := range s {
	//	fmt.Println(idx, string(r))
	//}

	//array and slices
	//a := [5]int{1, 2, 3, 4, 5}
	//b := a
	//b[0] = 10
	//fmt.Println(a)
	//fmt.Println(b)

	//var a []int  // nil slice
	//b := []int{} // empty slice, not nil
	//fmt.Println(a)
	//fmt.Println(b)

	//pointers
	//x := 10
	//p := &x
	//fmt.Println(*p)

	Hello()

	const inflationRate = 2.5
	var price float64
	rate := 5.0
	year := 10.0

	fmt.Printf("Enter the price: ")
	fmt.Scan(&price)

	fmt.Printf("Enter the rate: ")
	fmt.Scan(&rate)

	fmt.Printf("Enter the year: ")
	fmt.Scan(&year)

	amount := price * math.Pow(1+rate/100, year)
	futureRealValue := amount / math.Pow(1+inflationRate/100, year)
	fmt.Println(amount, futureRealValue)
	fmt.Println(futureRealValue)
}
