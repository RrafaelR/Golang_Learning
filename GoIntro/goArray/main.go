package main

import "fmt"

// array had assigned lenght
// slice is dinamic
// select specific elements in slice nameSlice[startIndexIncluding : upToNotIncluding]
func main() {
	friends := []string{"Annie", "Betty", "Carlos", "Julia", "Roxan"}
	fmt.Println(friends[1])
	// for i := range friends {
	// 	friends[1] = "Jack" // here you are changing the value directly in the array
	// 	if i == 1 {
	// 		fmt.Println(friends[1]) //-> Jack
	// 	}
	// }
	// for i, v := range friends {
	// 	friends[1] = "Jack" // it didn't change the value of v, because v took a copy of the array friends
	// 	if i == 1 {
	// 		fmt.Printf("v[%s]\n", v) // -> v[Jack]
	// 	}
	// }
	var friends2 []string
	friends2 = friends
	friends2[1] = "Jack"
	fmt.Printf("friends[%s] - friends2[%s]", friends[1], friends2[1]) //-> friends[Jack] - friends2[Jack]
}
