package pattern

import "fmt"

/*
в зависимости от состояния воркера он будет отвечать разными фразами, когда мы будем вызывать
worker.Reply()
*/

type WorkerStater interface {
	Reply()
}
type Worker struct {
	state WorkerStater
}

func (w *Worker) Reply() {
	w.state.Reply()
}
func (w *Worker) SetState(state WorkerStater) {
	w.state = state
}
func NewWorker() *Worker {
	return &Worker{state: &SleepingWorker{}}
}

type SleepingWorker struct{}

func (sw *SleepingWorker) Reply() {
	fmt.Println("i'm sleeping")
}

type HardWorkingWorker struct{}

func (hw *HardWorkingWorker) Reply() {
	fmt.Println("i'm working really hard")
}
