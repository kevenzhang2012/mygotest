package main

import (
	"fmt"
)

type admin struct {
	person user
	level  string
}
type user struct {
	name  string
	email string
}

func main() {

	fred := admin{
		person: user{
			name:  "lisa",
			email: "aa@aq.com",
		},

		level: "super",
	}

	fmt.Println(fred)

}
