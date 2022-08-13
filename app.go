package main

import "fmt"

func main() {
	var Ali string = "Here"
	var Name, Age = "Bear", 12
	var c = fmt.Sprintf(Name, Age)
	fmt.Println("Hellow Orld#3")
	fmt.Println(c)
	fmt.Println("Print your name plz:")
	fmt.Scanln(&Name)
	fmt.Println("You entered:" + Name)

	var Foo int = 123
	var result string
	result = fmt.Sprintf("Ёбаные якоря %d %s", Foo, Ali)
	fmt.Println(result)
}
