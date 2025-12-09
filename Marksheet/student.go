package main

import "fmt"

type Student struct {
	Name  string
	Marks []int
}

func (s Student) Average() float64 {
	total := 0
	for _, v := range s.Marks {
		total += v
	}
	return float64(total) / float64(len(s.Marks))
}

func (s Student) Grade() string {
	avg := s.Average()

	if avg >= 90 {
		return "A"
	} else if avg >= 75 {
		return "B"
	} else if avg >= 60 {
		return "C"
	} else if avg >= 40 {
		return "D"
	}
	return "F"
}

func main() {
	s := Student{
		Name:  "Vikrant",
		Marks: []int{85, 90, 80, 75},
	}

	fmt.Println("Name:", s.Name)
	fmt.Println("Average:", s.Average())
	fmt.Println("Grade:", s.Grade())
}
