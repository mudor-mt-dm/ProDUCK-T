package main

import "fmt"

func totalPrice (InitProce int) func(int) int{
	sum :=InitProce
	return func(x int) int {
		sum+=x
		return sum
	}
}

func main(){
	orderPrice:= totalPrice(1)
	fmt.Println(orderPrice(3))
	fmt.Println(orderPrice(6))
}

