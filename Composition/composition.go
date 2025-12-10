package main

import "fmt"

type Engine struct {
	HP int
}

type Car struct {
	Engine
	name string
}

func main() {
	c := Car{Engine{300}, "BMW"}
	fmt.Println(c.HP)
	fmt.Println(c.Engine)
	fmt.Println(c.name)
	fmt.Println(c)
	fmt.Println(c.Engine.HP)

}
