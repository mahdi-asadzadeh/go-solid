package main

import "fmt"

type DBConnection interface {
	SomeQuery() []string
}

type MySql struct{}

func (db *MySql) SomeQuery() []string {
	return []string{"a", "b", "c"}
}

type NewRepository struct {
	DB DBConnection
}

func (nr *NewRepository) GetData() []string {
	return nr.DB.SomeQuery()
}

func main() {
	nr := NewRepository{DB: &MySql{}}
	result := nr.GetData()
	fmt.Println(result)
}
