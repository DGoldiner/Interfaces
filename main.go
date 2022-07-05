package main

import "fmt"

func main() {

	var allsFood int
	animals := []animal{
		dog{weight: 7, norm: 2, number: 2},
		cat{weight: 3, norm: 7, number: 3},
		cow{weight: 20, norm: 25, number: 3},
	}

	for _, a := range animals {
		sum := a.getWeight() * a.getNorm() * a.getNumber() //потреба корму на один вид тварини
		allsFood += sum
		fmt.Println("name:", a.getAnimalName(), "weight:", a.getWeight(), "norm:", a.getNorm(), 
		"nomber of animal",a.getNumber(), "amount of food:", sum)
		
	}
	fmt.Println ("total amount of food:", allsFood)
}

type animal interface {
	summWeight
	animalNameGetter
	norma
	numbers
}

type summWeight interface {
	getWeight() int
}

type animalNameGetter interface {
	getAnimalName() string
}

type norma interface {
	getNorm() int
}

type numbers interface {
	getNumber() int
}

type dog struct {
	weight int
	norm   int
	number int
}

func (d dog) getWeight() int {
	return d.weight
}

func (d dog) getAnimalName() string {
	return "dog"
}

func (d dog) getNorm() int {
	return d.norm
}

func (d dog) getNumber() int {
	return d.number
}

type cat struct {
	weight int
	norm   int
	number int
}

func (c cat) getWeight() int {
	return c.weight
}

func (c cat) getAnimalName() string {
	return "cat"
}

func (c cat) getNorm() int {
	return c.norm
}

func (c cat) getNumber() int {
	return c.number
}

type cow struct {
	weight int
	norm   int
	number int
}

func (w cow) getWeight() int {
	return w.weight
}
func (w cow) getAnimalName() string {
	return "cow"
}

func (w cow) getNorm() int {
	return w.norm
}

func (w cow) getNumber() int {
	return w.number
}
