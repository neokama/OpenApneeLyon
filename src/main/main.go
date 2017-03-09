package main

import "fmt"



func main() {


    fmt.Printf("hello, world\n")
	fmt.Printf("hello, léo\n")

	fmt.Println("What is your name ?")
	var name string  // var - nom variable - type de la variable
	fmt.Scan(&name)
	fmt.Println("hello", name, "Quel age as-tu ?")
	var age int
	fmt.Scan(&age)
	fmt.Println("C'est super d'avoir",age,"ans !")
}
