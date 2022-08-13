package main

import "fmt"

func main() {
	var Name, Age = "Bear", 12
	var c = fmt.Sprintf(Name, Age)
	fmt.Println("Hellow Orld#3")
	fmt.Println(c)
	fmt.Println("Print your name plz:")
	fmt.Scanln(&Name)
	fmt.Println("You entered:" + Name)
}
