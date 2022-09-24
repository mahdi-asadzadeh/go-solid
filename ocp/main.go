package main

import "fmt"

type Calcule interface {
	Execute(int, int) int
}

type Add struct{}

func (ad *Add) Execute(a int, b int) int {
	return a + b
}

type Calculetor struct {
	c Calcule
}

func (ca *Calculetor) Execute(a int, b int) int {
	return ca.c.Execute(a, b)
}

func main() {

	ca := Calculetor{c: &Add{}}
	r := ca.c.Execute(2, 2)
	fmt.Println(r)

}
