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
	printFunc(fruits)                                                  /* -> fruits[Banana]
	// 	fruits[Apple]
	// 	fruits[Grape]
	// 	fruits[Orange]
	// 	fruits[] //-> this position doesn't is populated due to append, because append introduce element after the last position
	// 	fruits[plum]*/
	// }

	//***************

	fmt.Println("***********")
	names := []string{"Roberto", "Roberta", "Joana"}
	//names2 := names[:1] // -> it makes another reference to the value, so now we have 2 pointers, pointing to the same value that is roberto
	//names2 = append(names2, "Sol") // here you replice the value roberta, due to the the length of the names2 is 1 so after this one it will add
	//but it's the index 2 in names, so index 2 will be replaced for a new value from append of names2
	names2 := names[:1:1] // -> Here you are setting the capacity to 1 too
	fmt.Printf("%p\n", &names)
	fmt.Printf("%p\n", &names2)
	names2 = append(names2, "Sol") // -> Now, since the capacity of names2 is one, when it try to extend value it don't have more capacity
	// so it makes an copy of array and add the Sol in it, don't modifying names
	fmt.Printf("%p\n", &names2)
	printFunc(names)
	printFunc(names2)
	//***************************************
	array := make([]string, 3)
	numElement := copy(array, names) // this func make a copy of the array and return the number of copied elements
	printFunc(array)
	fmt.Println(numElement, "\n----------\n")

	//************************************************
	fmt.Printf("%p -> [%s]\n", &names, names)
	for _, v := range names {
		names = names[:1]
		fmt.Printf("%p -> [%s]\n", &v, v)
	}
	fmt.Println("*****************************************")
	for i := range names {
		names = names[:1]
		fmt.Printf("%p -> [%s]\n", &names, names[i])
	}

}
func printFunc(array []string) {
	for _, val := range array {
		fmt.Printf("[%s]\n", val)
	}
	fmt.Println("------------")
}
