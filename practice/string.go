package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	s1 := "hello"
	fmt.Println(s1);
	fmt.Println(strings.Contains(s1, "hel"))

	s2 := "1,2,3"
	fmt.Println(strings.Split(s2, ","))
//	a3 := strings.Split(s2,",")
//	fmt.Println(strings.co)
	i3 := strings.Index(s1,"l")
	fmt.Println(i3)

	i4 := 4

	s,err := strconv.Atoi(s1))
	fmt.Println(strconv.Itoa(i4))

}
