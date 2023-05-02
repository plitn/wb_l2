package pattern

import "fmt"

/*
когда мы будем использовать этот паттерн, то
сервис сможет применять посетителя к любому набору объектов и не уточнять их тип
*/

type transporter interface {
	// + какие-то общие методы для структур транспорта

	accept(v visitorer)
}

type visitorer interface {
	rideBike(bike)
	rideTram(tram)
	rideBus(bus)
	rideTrain(train)
}

type personVisitor struct{}

func (p *personVisitor) rideBike(b bike) {
	// do smth
	fmt.Println("person visited bike")
}

func (p *personVisitor) rideTram(t tram) {
	// do smth
	fmt.Println("person visited tram")
}

func (p *personVisitor) rideBus(b bus) {
	// do smth
	fmt.Println("person visited bus")
}

func (p *personVisitor) rideTrain(t train) {
	// do smth
	fmt.Println("person visited train")
}

type bike struct{}

func (b *bike) accept(v visitorer) {
	v.rideBike(*b)
}

type tram struct{}

func (t *tram) accept(v visitorer) {
	v.rideTram(*t)
}

type bus struct{}

func (b *bus) accept(v visitorer) {
	v.rideBus(*b)
}

type train struct{}

func (t *train) accept(v visitorer) {
	v.rideTrain(*t)
}
