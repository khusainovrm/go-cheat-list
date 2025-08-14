package main

import "fmt"

func Create() map[string]string {
	fmt.Println("Start Create")
	return map[string]string{
		"Perfect": "World",
	}
}

func Exposed() {
	fmt.Println("Exposed")
}

func private() {
	fmt.Println("private")
}
