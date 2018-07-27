package main

import (
	"fmt"
)

func main() {
	func() int {
		var i int = 5
		fmt.Println("func1\n")
		return i
	}()

	func(age int) {
		fmt.Println(age)
	}(2)
	myfunc := func(name string, age int) string {
		fmt.Println(name)
		fmt.Println(age)
		return name
	}
	myfunc("zhangsan", 10)
}
func testfunc(name string, age int) {
	fmt.Println(name)
	fmt.Println(age)

}
