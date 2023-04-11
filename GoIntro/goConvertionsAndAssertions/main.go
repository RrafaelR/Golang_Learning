package main

import "fmt"

type Mover interface {
	Move()
}

type Locker interface {
	Lock()
	Unlock()
}

type MoveLocker interface {
	Mover
	Locker
}
type user struct {
	name  string
	email string
}

type bike struct{} //concrete type for the example
func (bike) Move() {
	fmt.Println("Moving the bike")
}

func (bike) Lock() {
	fmt.Println("Locking the bike")
}

func (bike) Unlock() {
	fmt.Println("Unlocking the bike")
}

// STRING IMPLEMENTS THE FMT.STRING INTERFACE
func (u *user) String() string {
	return fmt.Sprintf("My name is %q and my email is %q", u.name, u.email)
}

func main() {
	var ml MoveLocker // this two have zero value, because interface is valueless
	var m Mover

	m = bike{}
	// ml = m -> it won't works because m only see Move() func, this don't look for another methods inside bike

	// making an assertion
	b, ok := m.(bike) // if typeof m is bike, b receive a copy from m with bike value and this func returns true
	if ok {           // if m points to a bike value
		ml = b
		ml.Unlock()
	}
	u := user{
		name:  "Rafael",
		email: "rafael@email.com",
	}
	fmt.Println(u)
	fmt.Println(&u)
}
