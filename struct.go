package main

import "fmt"

func displayType(p Talker) string {
	switch v := p.(type) {
	case Person:
		return fmt.Sprintf("%s is a Person.", v.Name)
	default:
		return fmt.Sprintf("The object is of type %T.", v)
	}
}

func displayType1(p Talker) string {
	c, ok := p.(Person)
	if !ok {
		return "Unknown type"
	} else {
		return fmt.Sprintf("The type is person with name %s and age %d", c.Name, c.Age) // Added closing parenthesis
	}
}

type Talker interface {
	Talk() string
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Talk() string {
	return fmt.Sprintf("%s is talking.", p.Name)
}

func saySomething(p Person) {
	fmt.Println(p.Talk())
}

func main() {
	person1 := Person{
		Name: "Abinash",
		Age:  22,
	}

	saySomething(person1)
	fmt.Println(displayType(person1))
}
