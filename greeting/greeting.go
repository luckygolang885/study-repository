/* ТЕОРИЯ ПО МОДУЛЯМ И ПАКЕТАМ 
package greeting

import "fmt"

func SayHello() {
	fmt.Println("Hello, golang packages!")

	j := GiveMeInt()

	fmt.Println("from sayhello:", j)
}
*/


// 24.09.2026 ПРАКТИКА Модули и пакеты
package greeting

import "fmt"

func SayHello() {
	fmt.Println("Привет из пакета greeting!")
}