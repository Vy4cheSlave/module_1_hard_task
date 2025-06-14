package fan_in

import (
	"sync"
)

// MergeChannels - принимает несколько каналов на вход и объединяет их в один
// Fan-in и merge channels синонимы
func MergeChannels(channels ...<-chan int) <-chan int {
	// TODO: ваш код
	outputChan := make(chan int)
	wg := &sync.WaitGroup{}

	for _, channel := range channels {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for value := range channel {
				outputChan <- value
			}
		}()
	}

	go func() {
		wg.Wait()
		close(outputChan)
	}()
	return outputChan

	panic("implement me")
}
