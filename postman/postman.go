/* ТЕОРИЯ 12. 10. 2026 Concurrency
package postman

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Postamn(ctx context.Context, wg *sync.WaitGroup ,transferPoint chan<- string, n int, mail string) {
	defer wg.Done()

	for {
		select {
		case <- ctx.Done():
			fmt.Println("Я почтальон номер:", n, "Мой рабочий день закончен!")
			return
		default:
			fmt.Println(" Я почтальон номер:", n, "Взял письмо")
		    time.Sleep(1 * time.Second)
		    fmt.Println("Я почтальон номер:", n, "Донёс письмо до почты:", mail)

		    transferPoint <- mail

		    fmt.Println("Я почтальон номер:", n, "Передал письмо:", mail)
		}
		
	}
}

func PostmanPool(ctx context.Context, postamnCount int) <-chan string {
	mailTransferPoint := make(chan string)

	wg := &sync.WaitGroup{}

	for i := 1; i <= postamnCount; i++ {
		wg.Add(1)
		go Postamn(ctx, wg, mailTransferPoint, i, postamanToMail(i))
	}

	go func() {
		wg.Wait()
		close(mailTransferPoint)
	}()

	return mailTransferPoint
}

func postamanToMail(postamnNumber int) string {
	ptm := map[int]string {
		1: "Семейный привет",
		2: "Приглашение друга",
		3: "Информация из автосервиса",
	}

	mail, ok := ptm[postamnNumber]
	if !ok {
		return "Лотерея"
	}

	return mail
}
*/