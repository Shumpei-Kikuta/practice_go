package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string // comma periodはいらない
	lastName  string
	contact   contactInfo
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"} // こういう定義の仕方をする
	// fmt.Println(alex)
	// var alex person
	// alex.firstName = "Alex" // 他のやり方
	// alex.lastName = "Andorson"
	// fmt.Println(alex) // zeroがassignされている。stringだとempty string ""
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		}, //最後にもcommaがいる
	}
	jim.print()

	jim.updateName("jimmy")
	jim.print()
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
