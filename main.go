package main

import (
	"awesomeProject/utils"
	"fmt"
)
import _ "awesomeProject/utils"

type Logging interface {
	Log() map[string]string
}

type User struct {
	Name string
	Age  int
}

func (u User) Log() map[string]string {
	obj := map[string]string{
		"Greeting": "Hello",
	}
	return obj
}

func (u User) grownUp() {
	u.Age++
	fmt.Println(u.Age)
}

func main() {
	u := User{Name: "Rinat", Age: 38}
	fmt.Println("It's me", u)
	fmt.Printf("It's me, %v hi\n", User{"Dinara", 38})

	greet := u.Log()
	fmt.Println(greet["Greeting"])

	u.grownUp()

	fmt.Printf("Now is %s, year", utils.Now())

	//Create()
	//Exposed()
	//private()
}
