package main

import "fmt"

type Person interface {
	sayInfo() string
}
type Student struct {
	name     string
	lastName string
}

type Professor struct {
	name     string
	lastName string
	class    string
}

func (s Student) sayInfo() string {
	return s.name + " " + s.lastName
}

func (p Professor) sayInfo() string {
	return p.name + " " + p.lastName + ", in class " + p.class
}

func main1() {
	student := Student{
		name:     "Joan",
		lastName: "Cabezas",
	}

	professor := Professor{
		name:     "Santiago",
		lastName: "Monroy",
		class:    "Android development",
	}
	fmt.Println(student.sayInfo())
	fmt.Println(professor.sayInfo())

	var schoolPersons []Person
	schoolPersons = append(schoolPersons, student)
	schoolPersons = append(schoolPersons, professor)
	fmt.Println(schoolPersons)
}
