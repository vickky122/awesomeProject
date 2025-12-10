package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Student struct {
	Name  string `json:"name"`
	Marks []int  `json:"marks"`
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
	switch {
	case avg >= 90:
		return "A"
	case avg >= 75:
		return "B"
	case avg >= 60:
		return "C"
	case avg >= 40:
		return "D"
	}
	return "F"
}

func (s *Student) AddMark(m int) {
	s.Marks = append(s.Marks, m)
}

func (s Student) SaveToFile(filename string) error {
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func LoadFromFile(filename string) (Student, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return Student{}, err
	}

	var s Student
	err = json.Unmarshal(data, &s)
	if err != nil {
		return Student{}, err
	}

	return s, nil
}

func main() {
	var choice int

	fmt.Println("Choose an option:")
	fmt.Println("1. Save student data to JSON")
	fmt.Println("2. Load student data from JSON")
	fmt.Print("Enter choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		s := Student{Name: "Vikrant", Marks: []int{80, 85, 90}}
		s.AddMark(95)

		fmt.Println("Saving Data:")
		fmt.Println("Name:", s.Name)
		fmt.Println("Average:", s.Average())
		fmt.Println("Grade:", s.Grade())

		err := s.SaveToFile("student.json")
		if err != nil {
			fmt.Println("Error writing file:", err)
		} else {
			fmt.Println("student.json file created successfully.")
		}
	} else if choice == 2 {
		loaded, err := LoadFromFile("student.json")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("\nLoaded Data:")
		fmt.Println("Name:", loaded.Name)
		fmt.Println("Marks:", loaded.Marks)
		fmt.Println("Average:", loaded.Average())
		fmt.Println("Grade:", loaded.Grade())
	} else {
		fmt.Println("Invalid choice!")
	}
}
