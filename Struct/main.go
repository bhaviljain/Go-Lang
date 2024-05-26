package main

import "fmt"

type Person struct {
	Fname string
	Lname string
	age   int
}
type Adress struct {
	House string
	Area  string
	State string
}
type Contact struct {
	Email string
	Phone int
	Fax   string
}
type Employee struct {
	Person_details Person
	Person_contact Contact
	Person_adress  Adress
}

func main() {
	//1st method
	var Bhavil Person
	Bhavil.Fname = "Bhavil"
	Bhavil.Lname = "jain"
	Bhavil.age = 23
	fmt.Println(Bhavil)
	//2nd method
	person1 := Person{
		Fname: "Bhavil",
		Lname: "Jain",
		age:   200,
	}
	fmt.Println(person1)

	var person2 = new(Person)
	person2.Fname = "kaka"
	person2.age = 10
	fmt.Println(person2.age)
	fmt.Println(person2.Fname)

	var emp Employee
	emp.Person_details = Person{
		Fname: "Bhavil",
		Lname: "jain",
		age:   20,
	}
	emp.Person_contact = Contact{
		Email: "bhav@gmail.com",
		Phone: 8989898989,
	}
	emp.Person_contact.Fax = "newFax@fax.com"
	emp.Person_adress.Area = "xyz"
	emp.Person_adress.House = "xyz"
	emp.Person_adress.State = "xyz"
	fmt.Println(emp.Person_contact)
}
