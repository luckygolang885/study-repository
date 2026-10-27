/* ТЕОРИЯ 12.10.2026 Concurrency
package miner

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//                       передаем уголь но не берем
func Miner(ctx context.Context, wg *sync.WaitGroup, transferPoint chan<- int, n int, power int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Я шахтёр номер:", n, "Мой рабочий день закончен!")
			return

		// если не отменен буду использовать логику шахтера - default:
		default:
		fmt.Println("Я шахтёр номер:", n, "начал добывать уголь!")
		time.Sleep(1 * time.Second)
		fmt.Println("Я шахтёр номер:", n, "Добыл уголь!", power)

		transferPoint <- power
		fmt.Println("Я шахтёр номер:", n, "Передал уголь:", power)
		}
	}
}

//                                                 <- chan не дает записать в канал что либо из main 
func MinerPool(ctx context.Context, minerCount int) <-chan int{
	coalTransferPoint := make(chan int)

	wg := &sync.WaitGroup{}

	for i := 1; i <= minerCount; i++ {
		wg.Add(1)
		go Miner(ctx, wg, coalTransferPoint, i, i*10)
	}

	go func() {
		wg.Wait()
	    close(coalTransferPoint)
	}()

	return coalTransferPoint
}
*/