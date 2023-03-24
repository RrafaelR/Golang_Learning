package main

import (
	"fmt"
	"sort"
)

type person struct {
	name string
	age  int
}

func main() {
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

	println("***********************************")
	people := map[string]person{
		"Rafael": {"Rafael", 21},
		"Maria":  {"Maria", 21},
		"Julia":  {"Julia", 26},
		"Kely":   {"Kely", 31},
	}
	var keys []string
	for key := range people {
		keys = append(keys, key)
	}
	sort.Strings(keys) // it put in alphabetic order the keys
	for _, key := range keys {
		fmt.Printf("[%s]->{%v}\n", key, people[key])
	}
}

func printMap(c map[int]string) {
	for key, value := range c {
		fmt.Println(key, ":", value) // print keys in alphabetic order
	}
}
