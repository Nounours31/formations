package main

import (
	"fmt"
)

func exo0() {
	var alias string = "Go"
	alice := "no go"
	fmt.Printf("Hello, World! %s - %s\n", alias, alice)
	fmt.Println("Go est correctement installé !")
}

func exo1() {
	var alias string = "Go"
	fmt.Printf("Hello, [%s]! \n", alias)
}

func main() {
	exo0()
	exo1()
}
