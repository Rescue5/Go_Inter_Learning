package main

import (
	"fmt"
	"sync"
)

const CHUNCSIZE = 315

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	chaptSlice := len(slice)/CHUNCSIZE + 1
	wg := sync.WaitGroup{}
	sumCh := make(chan int)

	for i := 0; i < chaptSlice; i++ {
		pastI := i * CHUNCSIZE
		if pastI > len(slice) {
			break
		}

		lastI := pastI + CHUNCSIZE
		if lastI > len(slice) {
			lastI = len(slice)
		}

		sliceI := slice[pastI:lastI]

		// Запуск горутины с wg
		wg.Add(1)
		go func(s []int) {
			defer wg.Done()
			getSumElem(s, sumCh)
		}(sliceI)

	}

	// Закрытие канала после завершения всех wg
	go func() {
		wg.Wait()
		close(sumCh)
	}()

	var res int
	// По сути программа застревает на этом цикле, пока канал sumCh не будет закрыт
	// Следовательно, пока все wg (горутины) не будут завершены мы будем сидеть и сумировать поступающие элементы в канал
	for i := range sumCh {
		res += i
	}

	fmt.Println(res)

}

func getSumElem(slice []int, sumCh chan int) {
	var res int
	for _, i := range slice {
		res += i
	}
	sumCh <- res
}
