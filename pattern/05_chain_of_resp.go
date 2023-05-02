package pattern

import "fmt"

type Handler interface {
	passNext(int)
}

type HandlerFirst struct {
	next Handler
}

func (hf *HandlerFirst) passNext(value int) {
	fmt.Printf("first handler checked value of %d, passing further", value)
	if hf.next != nil {
		hf.next.passNext(value)
	}
	return
}

type HandlerSecond struct {
	next Handler
}

func (hs *HandlerSecond) passNext(value int) {
	fmt.Printf("second handler checked value of %d, passing further", value)
	if hs.next != nil {
		hs.next.passNext(value)
	}
	return
}

type HandlerLast struct {
	next Handler
}

func (hl *HandlerLast) passNext(value int) {
	fmt.Printf("last handler checked value of %d, passing further", value)
	if hl.next != nil {
		hl.next.passNext(value)
	}
	return
}
