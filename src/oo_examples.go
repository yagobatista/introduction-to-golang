package main

import "fmt"

type Meta struct {
	ID int
}

type Person struct {
	Name      string
	BirthYear string
}

type Student struct {
	Meta
	Person

	Matricula string
}

func (this Student) SaveInstance() {
	fmt.Println("save student")
}

func (this Student) UpdateInstance() {
	fmt.Println("update student")

}

type Guardian struct {
	Person
	Students []Student
}

func (this Guardian) SaveInstance() {
	fmt.Println("save guardian")
}

func (this Guardian) UpdateInstance() {
	fmt.Println("update guardian")

}

type SaveOnDb interface {
	SaveInstance()
	UpdateInstance()
}

func callSave(saver SaveOnDb) {
	saver.SaveInstance()
}

func main() {
	var guardian Guardian
	var student Student
	callSave(guardian)
	callSave(student)
}
