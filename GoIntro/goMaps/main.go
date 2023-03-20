package main

import "fmt"

func runMaps() {
	// lockers := map[int32]string{
	// 	1: "Books and notebooks",
	// 	2: "Laptops and ipads",
	// } -> first way yo declare map
	//var lockers map[int]string //-> second way to declare a map
	lockers := make(map[int]string) //-> third way to declare map
	lockers[1] = "Books and notebooks"
	lockers[2] = "Laptops and ipads"
	//delete(lockers, 1) // -> for delete a pair key - value of the map
	printMap(lockers)
}

func printMap(c map[int]string) {
	for key, value := range c {
		fmt.Println(key, ":", value)
	}
}
