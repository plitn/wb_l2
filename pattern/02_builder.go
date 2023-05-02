package pattern

import "fmt"

// Общий интерфейс строителя
type bottleBuilder interface {
	makeMaterial()
	makeCap()
	makePackage()
	addFluid()
}

// конкретный строитель, который реализует интерфейс
type waterBottle struct {
}

func (wb *waterBottle) makeMaterial() {
	fmt.Println("water bottle plastic made")
}

func (wb *waterBottle) makeCap() {
	fmt.Println("water bottle plastic cap made")
}

func (wb *waterBottle) makePackage() {
	fmt.Println("long water bottle package made")
}

func (wb *waterBottle) addFluid() {
	fmt.Println("added water to waterBottle")
}

// конкретный строитель, который реализует интерфейс
type cokeBottle struct {
}

func (cb *cokeBottle) makeMaterial() {
	fmt.Println("coke bottle glass made")
}

func (cb *cokeBottle) makeCap() {
	fmt.Println("coke bottle aluminum cap made")
}

func (cb *cokeBottle) makePackage() {
	fmt.Println("small coke bottle package made")
}

func (cb *cokeBottle) addFluid() {
	fmt.Println("added coke to cokeBottle")
}

// решает в каком порядке и как строить объект
type Director struct {
}

func (d *Director) makeSomeBottle(builder bottleBuilder) {
	builder.makeMaterial()
	builder.makePackage()
	builder.makeCap()
	builder.addFluid()
}
