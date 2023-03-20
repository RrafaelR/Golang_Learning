package main

import "fmt"

// Whenever you pass an integer, float, string, or struct into a function, this function took a copy of the value
type contactInfo struct {
	phone int
	email string
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}
type person2 struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	cintia := person{
		firstName: "Cintia",
		lastName:  "Doncic",
		contact: contactInfo{
			email: "cintiaDoncic@go.com",
			phone: 22222222,
		},
	}
	alice := person2{
		firstName: "Alice",
		lastName:  "Doncic",
		contact: contactInfo{
			email: "aliceDoncic@go.com",
			phone: 22222222,
		},
	}

	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Ferguson"
	alex.contact.email = "alexferguson@go.com"
	alex.contact.phone = 11111111
	alex.updateFirstName("Alexander") // pass the address of this var to the func
	alex.print()                      // {firstName:Alex lastName:Ferguson contact:{phone:11111111 email:alexferguson@go.com}}
	cintia.print()                    // {Cintia Doncic {22222222 cintiaDoncic@go.com}}
	name := "Bill"
	pointerName := &name
	fmt.Println(pointerName)
	printName(pointerName)
	// alice = person2(cintia) // varible convertion for declared values
	// fmt.Println(alice)
	person3 := struct {
		firstName string
		lastName  string
		contact   contactInfo
	}{
		firstName: "Rafael",
		lastName:  "Brown",
		contact: contactInfo{
			email: "rafaelBrown@go.com",
			phone: 22222222,
		},
	}
	alice = person3 // you don't need to make the convertion because person3 is an implicity declaration
	fmt.Println(alice)

}

func (pointerToPerson *person) updateFirstName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
func printName(name *string) {
	fmt.Println(name)
}
