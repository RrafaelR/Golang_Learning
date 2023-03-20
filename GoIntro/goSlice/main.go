package main

import "fmt"

func main() {
	fruits := make([]string, 5, 8)
	fruits[0] = "Banana"
	fruits[1] = "Apple"
	fruits[2] = "Grape"
	fruits[3] = "Orange"
	fmt.Printf("Length[%d], Capacity[%d]\n", len(fruits), cap(fruits)) //-> Length[5], Capacity[8]
	fruits = append(fruits, "plum")                                    // Be careful because slice is already initialized, but when you use append
	//you are incresing the length, so you put this value after then the last position that is nil or empty
	fmt.Printf("Length[%d], Capacity[%d]\n", len(fruits), cap(fruits)) //->Length[6], Capacity[8]
	for _, v := range fruits {
		fmt.Printf("fruits[%s]\n", v) /* -> fruits[Banana]
		fruits[Apple]
		fruits[Grape]
		fruits[Orange]
		fruits[] //-> this position doesn't is populated due to append, because append introduce element after the last position
		fruits[plum]*/
	}
}
