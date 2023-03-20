package main

import (
	"fmt"
)

type user struct {
	name string
	id   int
}

func main() {
	count := 10
	testingPointer1(&count) // pass a copy of the address of memory
	fmt.Println(count)
	user1 := createUser("Toinho", 0000) // factory func to create a user
	fmt.Println(user1)
}

func testingPointer1(num *int) { // store the copy of the address of memory in this pice of memory.
	*num++
}

func createUser(name string, id int) user {
	return struct {
		name string
		id   int
	}{
		name: name,
		id:   id,
	}
}
