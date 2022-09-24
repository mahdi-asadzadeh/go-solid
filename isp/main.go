package main

import "fmt"

type Animal interface {
	Eat() bool
}

type FlaybleAnimal interface {
	Animal
	Flay() bool
}

type Dolfin struct{}

func (do *Dolfin) Eat() bool {
	return true
}

type Sparrow struct{}

func (sq *Sparrow) Flay() bool {
	return true
}
func (sq *Sparrow) Eat() bool {
	return true
}

type Animals struct {
	an Animal
}

type FlaybleAnimals struct {
	an FlaybleAnimal
}

func main() {
	an := Animals{an: &Dolfin{}}
	fmt.Println(an.an.Eat())

	anf := FlaybleAnimals{an: &Sparrow{}}
	fmt.Println(anf.an.Eat())
	fmt.Println(anf.an.Flay())
}
