package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"bot", 20})
	fmt.Println(person{name: "zhangsan", age: 20})
}
