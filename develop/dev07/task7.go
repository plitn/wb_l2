package dev07

import "sync"

func or(channels ...<-chan interface{}) <-chan interface{} {
	orChannel := make(chan interface{})
	var wg sync.WaitGroup

	// функция ждет закрытия канала, отправляем из него данные в мердж канал
	wait := func(channel <-chan interface{}) {
		defer wg.Done()
		for value := range channel {
			orChannel <- value
		}
	}
	wg.Add(len(channels))
	// ждем закрытия каналов
	for _, channel := range channels {
		go wait(channel)
	}

	// ждем закрытия всех каналов (вг), закрываем мердж канал
	go func() {
		wg.Wait()
		close(orChannel)
	}()
	return orChannel
}
