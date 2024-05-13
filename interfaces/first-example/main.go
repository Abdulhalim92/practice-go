package main

import "fmt"

type Human struct {
	Firstname string
	Lastname  string
	Age       int
	Country   string
}

func (h Human) String() string {
	return fmt.Sprintf("Human named %s %s or age %d living in %s",
		h.Firstname,
		h.Lastname,
		h.Age,
		h.Country,
	)
}

type DomesticAnimal interface {
	ReceiveAffection(from Human)
	GiveAffection(to Human)
}

type Cat struct {
	Name string
}

func (c Cat) ReceiveAffection(from Human) {
	fmt.Printf("The car named %s has received affection from Human named %s\n", c.Name, from.Firstname)

}

func (c Cat) GiveAffection(to Human) {
	fmt.Printf("The cat named %s has given affection to Human named %s\n", c.Name, to.Firstname)
}

type Dog struct {
	Name string
}

func (d Dog) ReceiveAffection(from Human) {
	fmt.Printf("The dog named %s has received affection from Human named %s\n", d.Name, from.Firstname)
}

func (d Dog) GiveAffection(to Human) {
	fmt.Printf("The dog named %s has given affection to Human named %s\n", d.Name, to.Firstname)
}

type Snake struct {
	Name string
}

func (s Snake) ReceiveAffection(from Human) {
	//TODO implement me
	panic("implement me")
}

func (s Snake) GiveAffection(to Human) {
	//TODO implement me
	panic("implement me")
}

func Pet(animal DomesticAnimal, human Human) {
	animal.GiveAffection(human)
	animal.ReceiveAffection(human)
}

func main() {
	var john Human
	john.Firstname = "John"

	var c Cat
	c.Name = "Maru"

	var d Dog
	d.Name = "Medor"

	var s Snake
	s.Name = "lalala"

	Pet(c, john)
	Pet(d, john)
	Pet(s, john)
}
