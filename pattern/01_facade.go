package pattern

import "fmt"

/*
Паттерн фасад позволяет реализовать простой доступ к сложной подсистеме
*/

type CoffeeFacade struct {
	cg *coffeeGrinder
	wp *waterFaucet
	tm *tasteMaker
}

func NewCoffee() *CoffeeFacade {
	return &CoffeeFacade{
		cg: &coffeeGrinder{},
		wp: &waterFaucet{},
		tm: &tasteMaker{},
	}
}

func (cf *CoffeeFacade) MakeCoffee() {
	// тут производим какую-то работу с нужными классами
	cf.cg.grind()
	cf.wp.pour()
	cf.tm.makeGood()
}

type coffeeGrinder struct {
}

func (cg *coffeeGrinder) grind() {
	fmt.Println("grinding coffee")
}

type waterFaucet struct {
}

func (wf *waterFaucet) pour() {
	fmt.Println("pouring water")
}

type tasteMaker struct {
}

func (tm *tasteMaker) makeGood() {
	fmt.Println("making coffee good")
}
