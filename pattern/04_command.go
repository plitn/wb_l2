package pattern

import "fmt"

/*
комманд - запрос в виде объекта на выполнение
инвокер - объект-инициатор запроса
ресивер - объект-получатель запроса

интерфейс комманд
класс конкретной команды, который реализует команду.
класс инвокер записывает команду и провоцирует выполнение.
класс ресивер реализует получателя, имеет набор действий, который запрашивает команда.
*/

type Command interface {
	executeCommand()
}

type Reciever struct {
}

type worker struct {
	r *Reciever
}

func (w *worker) executeCommand() {
	w.r.doWork()
}

func (r *Reciever) doWork() {
	fmt.Println("reciever does work")
}

type Invoker struct {
	cmds []Command
}

func (i *Invoker) addComand(cmd Command) {
	i.cmds = append(i.cmds, cmd)
}

func (i *Invoker) executeCommand() {
	for _, cmd := range i.cmds {
		cmd.executeCommand()
	}
}
