package gogun

import "fmt"

type GunI interface {
	Get()
	Map()
	Put()
	Once()
	Test() string
}

var peers []string

type Gun struct {
	name string
  	uuid string
}

func (g Gun) Get() {
	fmt.Println("Testing...!",g.name)
}

func (g Gun) Map() {
	fmt.Println("Testing...!",g.name)
}

func (g Gun) Put() {
	fmt.Println("Testing...!",g.name)
}

func (g Gun) Once() {
	fmt.Println("Testing...!",g.name)
}

func (g Gun) Test() string{
	fmt.Println("func test!")
  return g.name
}

func Test(a, b int) int {
	return a - b
}

func Add(a, b int) int {
	return a + b
}
