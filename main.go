/*
package main

// golang возьми с помощью import ящик с инструментами fmt ящик инструментов
import "fmt"

func main() { // func main - это значит обьявляем кусок кода в котом программа будет начинать свое выполнение
	// Println выводи на экран || и так же Println начинается текст с новой строки
	fmt.Println("Hellow, World!")
	fmt.Println("Hellow, World!")
	fmt.Println("Hellow, World!")

	// а Print выводит текст в одну строчку
	fmt.Print("Hellow, World!")
	fmt.Print("Hellow, World!")
	fmt.Print("Hellow, World!")
}
*/

/* Переменнные
package main

import "fmt"

func main() {
	// text := "Game Over" // у переменной text, тип строка
	// text = 15 || не может число только строку
	// score - счет
	// score зарезервивовал какуюто область в памяти
	// := создай и положи переменную с именнем которя нахоидтся слеова значение которое находится справо
	// двоеточие : пишется только тогда когда создаем переменную, если переменая уже создана то мы уже не пишем двоеточие :
    // 1.score := 111 // у переменной score ,тип целое число

	// score = 109358945
	// score = 50 + 10
	// score = score + 102


	text1 := "Get Ready"
	text2 := "Game Over"
	score := 0


	fmt.Println(text1)
	fmt.Println("Ваш счёт:", score)

	fmt.Println("Вы пролетели через трубу")
	score = score + 1
	fmt.Println("Ваш счёт:", score)

	fmt.Println("Вы пролетели через трубу")
	score = score + 1
	fmt.Println("Ваш счет:", score)

	fmt.Println("Вы пролетели через трубу")
	score = score + 1
	fmt.Println("Ваш счет:", score)

	fmt.Println("Врезались в трубу!")
	fmt.Println(text2)


	//


	text := "hello" // string
	score := 0 // int целые числа хранятся в разном виде
	drob := 0.5 // float дробные числа хранятся в разном виде

	// score = 0.7 // Не можем использовать(хранить) дробное число в целое число

	fmt.Println(text)
	fmt.Println(score)
	fmt.Println(drob)

	hidden := true

	fmt.Println("Профиль скрытый?", hidden)

	//

	score := 0

	score = score + 10

	score = score + 20

	score = score - 15

	score = score / 3

	score = score * 10

	fmt.Println("Счёт:", score)

	//

	score := 11.0

	celoe := score / 3

	fmt.Println("Счёт:", score, celoe)

}
*/

/*
package main

import "fmt"

func main() {
	//
	text := "hello"

	text = text + " world"

	fmt.Println(text)



	//
	var number int = 19 // Длинный тип переменной
	var text string = "hi" // При длинном можем не указывать значение при указывание

	fmt.Println("number:", number)
	fmt.Println("text:", text)

	number = 15
	text = "Hello, World!"

	fmt.Println("number", number)
	fmt.Println("text:", text)


	//
	var number int
	var text string
	var drob float64
	var boolean bool

	fmt.Println("number:", number)
	fmt.Println("text:", text)
	fmt.Println("drob:", drob)
	fmt.Println("boolean:", boolean)



	//
	text := "hello"
	text = text + " world"
	// number := 25.0

	// number = number + 5 // Длинная запись
	// number += 10 // Короткая запись

	// number -= 10 // Короткая запись
	//number = number - 10 // Длинная запись

	// number /= 10
	// number *= 10

	fmt.Println("text:", text)


	//
	number := 10

	number += 5

	fmt.Println("number:", number)


	//
	score := 5

	score++ // Для увилечения числа на +1 || Инкримент
	score-- // Для уменьшения числа на -1 || Динкримент

	fmt.Println("score:", score)
}
*/

/* ПРАКТИКА 13.09.2026
1.
package main

import "fmt"

func main() {
	name := "Арзу"

	fmt.Println(name)
}
*/

/*
2.
package main

import "fmt"

func main() {
	name := "Арзу"
	age := 17
	city := "Volgograg"

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(city)
}
*/

/*
3.
package main

import "fmt"

func main() {
	age := 4
	height := 4.5
	isStudent := true

	fmt.Println(age)
	fmt.Println(height)
	fmt.Println(isStudent)
}
*/

/*
4.
package main

import "fmt"

func main() {
	name := "Arzu"
	age := 17
	height := 190.5

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(height)
}
*/

/*
5.
package main

import "fmt"

func main() {
	score := 10
	score = 25

	fmt.Println(score)
}
*/

/*
6.
package main

import "fmt"

func main() {
	a := 15
	b := 7
	sum := a + b

	fmt.Println(sum)
}
*/

/*
7.
package main

import "fmt"

func main() {
	a := 20
	b := 5

	plus := a + b
	minus := a - b
	umnoshenie := a * b
	Division := a / b

	fmt.Println(plus)
	fmt.Println(minus)
	fmt.Println(umnoshenie)
	fmt.Println(Division)
}
*/

/*
8.
package main

import "fmt"

func main() {
	name := "Арзу"
	age := 17
	city := "Volgograd"
	hours := 2

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(city)
	fmt.Println(hours)
}
*/

/* ТЕОРИЯ Условные ветвления, 14.09.2026

package main

import "fmt"

func main() {
	score := 5

	if score > 10 {
		fmt.Println("Ты красавчик!")
	} else {
		fmt.Println("Тебе нужно ещё многому научиться")
	}

	fmt.Println("Счёт:", score)
	fmt.Println("Спасибо за использование нашей игры! Возвращайтесь ещё!")
}


// СТАРАЯ СХЕМА
package main

import "fmt"

func main() {
	score := 5

	if score > 10 {
		if score > 15 {
			fmt.Println("Ты мега-красавчик!")
		} else {
			fmt.Println("Ты красавчик!")
		}
	} else {
		fmt.Println("Тебе нужно ещё многому научиться")
	}
}



// НОВАЯ СХЕМА И ЛУЧШАЯ
package main

import "fmt"

func main() {
	score := 5


	score := 5

	if score > 15 {
	fmt.Println("Ты мега-красавчик")
	} else {
	 if score > 10 {
	 fmt.Println("Ты красавчик")
	 } else {
	  fmt.Println("Тебе нужно ещё многому научиться")
	  }
	 }

	//

	if score > 15 {
		fmt.Println("Ты мега-красавчик!")
	} else if score > 10 {
		fmt.Println("Ты красавчик!")
	} else {
		fmt.Println("Тебе нужно ещё многому научиться")
	}
}



//

package main

import "fmt"

func main() {
	score := 21

	if score == 12 { // Двойное == это сравнение значения
		fmt.Println("Дюжина!")
	} else if score == 21 {
		fmt.Println("Очко")
	} else if score == 50 {
		fmt.Println("Полтиник!")
	} else {
		fmt.Println("Ты не попал ни в какую категорию")
	}
}


//
package main

import "fmt"

func main() {
	subscribed := true

	if subscribed {
		fmt.Println("Ты подписан на мой ютуб канал! Спасибо!!!")
	} else {
		fmt.Println("Пожалуйста, подпишись на мой туб канал.")
	}
}


package main

import "fmt"

func main() {
	number := 15

	ravno5 := number == 5
	bolshe12 := number > 12

	if ravno5 {
		fmt.Println("number равно пяти!")
	}

	if bolshe12 {
		fmt.Println("Больше 12!")
	}

	fmt.Println(ravno5)
	fmt.Println(bolshe12)
}
*/

// ПРАКТИКА ПОЛОВИНА 14.09.2026
/*
1.
package main

import "fmt"

func main() {
	age := 17

	if age >= 18 {
		fmt.Println("Совершеннолетний")
	}
}



2.
package main

import "fmt"

func main() {
	temperature := 5

	if temperature > 0 {
		fmt.Println("Тепло")
	} else {
		fmt.Println("Холодно")
	}
}


3.
package main

import "fmt"

func main() {
	a := 15
	b := 10

	if a > b {
		fmt.Println("a больше b")
	} else {
		fmt.Println("b больше или равен")
	}
}



4.
package main

import "fmt"

func main() {
	number := 0

	if number > 0 {
		fmt.Println("Положительное")
	} else {
		fmt.Println("Ноль или отрицательное")
	}
}


5.
package main

import "fmt"

func main() {
	password := "12345"

	if password == "12345" {
		fmt.Println("Доступ разрешён")
	} else {
		fmt.Println("Неверный пароль")
	}
}



6.
package main

import "fmt"

func main() {
	age := 17

	if age >= 18 {
		fmt.Println("Можно")
	} else {
		fmt.Println("Нельзя")
	}
}



7.
package main

import "fmt"

func main() {
	number := 6

	if number%2 == 0 {
		fmt.Println("Чётное")
	} else {
		fmt.Println("Нечётное")
	}
}


8.
package main

import "fmt"

func main() {
	score := 55

	if score >= 90 {
		fmt.Println("Отлично")
	} else {
		fmt.Println("Можно лучше")
	}
}



9.
package main

import "fmt"

func main() {
	age := 17

	if age < 18 {
		fmt.Println("Скидка 20%")
	} else {
		fmt.Println("Скидки нет")
	}
}



10.
package main

import "fmt"

func main() {
	balance := 30000000
	price := 30000000

	if balance >= price {
		fmt.Println("Покупка совершена")
	} else {
		fmt.Println("Недостаточно средств")
	}
}



11.
package main

import "fmt"

func main() {
	level := 5

	if level >= 50 {
		fmt.Println("Легенда")
	} else if level >= 30 {
		fmt.Println("Очень сильный")
	} else if level >= 10 {
		fmt.Println("Неплохой уровень")
	} else {
		fmt.Println("Новичок")
	}
}



12.
package main

import "fmt"

func main() {
	temperature := 40

	if temperature >= 30 {
		fmt.Println("Жарко")
	} else if temperature >= 20 {
		fmt.Println("Тепло")
	} else if temperature >= 10 {
		fmt.Println("Прохладно")
	} else {
		fmt.Println("Холодно")
	}
}


13.
package main

import "fmt"

func main() {
	score := 90

	if score >= 90 {
		fmt.Println("5 - Отлично")
	} else if score >= 75 {
		fmt.Println("4 - Хорошо")
	} else if score >= 60 {
		fmt.Println("3 - Удовлетворительно")
	} else {
		fmt.Println("2 - Не сдал")
	}
}



// ПРАКТИКА 15.09.2026
// две вертикальные палочки || - или
package main

import "fmt"

func main() {
	// две вертикальные палочки || - или


	score := 5

	// score < 6 == false
	// score score > 16 == false
	// Хватит одного условий было истиной то все выражение будет истинным
	if score < 6 || score > 16 {
		fmt.Println("Ты попал в петушиную зону!")
	} else {
		fmt.Println("Ты не попал в зону")
	}
}



//
package main

import "fmt"

func main() {

	score := 16

	// Нужно чтобы каждый из выражения был истинной
	// score > 8 == true
	// score < 15 == false
	if score > 8 && score < 15 {
		fmt.Println("Ты попал в яблочко!")
	} else {
		fmt.Println("Ты НЕ попал в яблочко")
	}
}


//
package main

import "fmt"

func main() {

	sunny := false
	weekend := true

	// я поду гулять, только если будет солнечная погода и будет выходной
	// иначе я буду сопротивляться
	// нужно 2 true а не один true иначе будет false
	if sunny && weekend {
		fmt.Println("Я иду гулять")
	} else {
		fmt.Println("Я сопротивляюсь")
	}


	computerClub := true
	icecream := false

	if computerClub || icecream {
		fmt.Println("Я иду гулять")
	} else {
		fmt.Println("Я буду сопротивляться")
	}
}


//
package main

import "fmt"

func main() {
	score := 10

	if score >= 10 {
		fmt.Println("Ты красавчик")
	} else {
		fmt.Println("Тебе нужно ещё многому научиться")
	}
}


//
package main

import "fmt"

func main() {
	score := 15

	if score <= 10 {
		fmt.Println("Ты красавчик!")
	} else {
		fmt.Println("Тебе нужно ещё многому научиться")
	}
}



//
package main

import "fmt"

func main() {
	score := 7

	// если игрок попал ровно в 7 то он проиграл
	if score != 7 {
		fmt.Println("Ты выиграл!")
	} else {
		fmt.Println("Ты проиграл!")
	}
}


//
package main

import "fmt"

func main() {
	subscribed := false

	fmt.Println("до условия")

	// !
	if !subscribed {
		fmt.Println("Я вижу ты не подписан! Подпишись, пожалуйста!")
	}

	fmt.Println("после условия")
}
*/

/* ПРАКТИКА 15.09.2026
1.
package main

import "fmt"

func main() {
	age := 1

	if age <= 17 {
		fmt.Println("Ты несовершеннолетний")
	} else {
		fmt.Println("Ты совершеннолетний")
	}
}



2.
package main

import "fmt"

func main() {
	password := "1234"

	if password != "12345" {
		fmt.Println("Неверный пароль")
	} else {
		fmt.Println("Пароль правильный")
	}
}


3.
package main

import "fmt"

func main() {
	isBanned := false

	if !isBanned { // Если игрок не заблокировать вывести сделать !isBanned
		fmt.Println("Доступ разрешён")
	} else {
		fmt.Println("Доступ запрещён")
	}
}


4.
package main

import "fmt"

func main() {
	age := 19

	if age <= 18 {
		fmt.Println("Вход разрешён")
	} else {
		fmt.Println("Вход запрещён")
	}
}


5.
package main

import "fmt"

func main() {
	login := "admin"

	if login != "admin" { //  не равен
		fmt.Println("Неверный логин")
	} else {
		fmt.Println("Добро пожаловать!")
	}
}


6.
package main

import "fmt"

func main() {
	hasSubscription := false

	if !hasSubscription { // если подписки НЕТ // РАБОТАЕТ ТОЛЬКО С BOOLEAN
		fmt.Println("Оформите подписку")
	} else {
		fmt.Println("Доступ к контенту открыт")
	}
}



7.
package main

import "fmt"

func main() {
	temperature := -5

	if temperature <= 0 {
		fmt.Println("На улице холодно")
	} else {
		fmt.Println("На улице тепло")
	}
}


8.
package main

import "fmt"

func main() {
	code := "Go2027"

	if code != "Go2026" {
		fmt.Println("Доступ запрещён")
	} else {
		fmt.Println("Доступ разрешён")
	}
}



9.
package main

import "fmt"

func main() {
	isBlocked := true

	if !isBlocked {
		fmt.Println("Аккант активен")
	} else {
		fmt.Println("Аккаунт заблокирован")
	}
}
*/

/* ТЕОРИЯ ЦИКЛ 16.09.2026
// ЦИКЛ МНОГОКРАТНОЕ ПОВТОРЕНИЕ

package main

import "fmt"

func main() {

	// 1-я секцция -- объявление переменной
	// i := 1

	// 2-я секция -- условия, при котором цикл будет выполняться
	// i <= 5

	// 3-я секция -- изменение от одной итерации к другой
	// итерация это один прогон по всему циклу {}

	// i++ / i = i + 1
	number := 999
	for i := 1; i <= 3; i++ { // Тело цикла
		score := 5
		score = score + 3

		fmt.Println("Hello, World!", i) // Цикл закончится тогда когда 5 <= 5; а не 1 <= 5;
		fmt.Println("Score:", score)
		fmt.Println("number:", number)
		fmt.Println("")

		number++
	}
	// i == 1
	// i == 2
	// i == 3
	// i == 4
	// i == 5
	// i == 6 // 6 не подошла и сразу вышли из цикла

	fmt.Println("После")
	//


	fmt.Println("До")

	for i := 1; i <= 10; i *= 2 {
		fmt.Println("Hello, World", i)
	}

	fmt.Println("После")
}
*/

/* ПРАКТИКА 16.09.2026
1.
package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}



2.
package main

import "fmt"

func main() {

	for i := 2; i <= 20; i +=2 {
		fmt.Println(i)
	}
}



3.
package main

import "fmt"


func main() {

	for i := 10; i >= 1; i-- {
		fmt.Println(i)
	}
	fmt.Println("Старт")
}



4.
package main

import "fmt"

func main() {
	sum := 0
	for i:= 1; i <= 100; i++ {
		fmt.Println(i)
		sum = i + sum
	}
	fmt.Println(sum)
}



5.
// Таблица умножения
package main

import "fmt"

func main() {
	number := 5

	for i := 1; i <= 10; i++ {
		fmt.Println(number, i, number*i)
	}
}



6.
// Числа, которые делятся на 3
package main

import "fmt"

func main() {
	n := 30

	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
*/

/* ТЕОРИЯ 16.09.2026 ЦИКЛ

package main

import "fmt"

func main() {
	// 🔋
	// 🐤

	score := 0
	fmt.Println("Get Ready")
	fmt.Println("Счёт:", score)
	fmt.Println("")


	for i := 1; i <= 5; i++ {
	fmt.Println("Вы подлетаете к трубе!", i)
	fmt.Println("🐤 🔋🔋")

	if i % 2 == 0 { // Проверка на чётность | остаток дает 0 значит число чётное | если остаток деления не дает 0 значит число нечётное
		fmt.Println("🧨🧨")
	} else {
		fmt.Println("🔋🔋")
	}

	fmt.Println("Вы пролетаете через трубу!", i)
	fmt.Println("🔋🐤🔋")
	fmt.Println("")

	fmt.Println("Вы пролетели через трубу!", i)
	fmt.Println("🔋🔋 🐤")
	fmt.Println("")

	score++
	fmt.Println("Счёт:", score)
	fmt.Println("")
	}

	fmt.Println("Генерация уровня окончена!")
}



//
package main

import (
	"fmt"
	"math/rand"
)
// "math/rand" // ДОПОЛНИТЕЛЬНЫЕ ЯЩИК ИНСТРУМЕНТОВ
// "time" // ДОПОЛНИТЕЛЬНЫЕ ЯЩИК ИНСТРУМЕНТОВ

// 1. бесконечные циклы
// 2. ключевое слово break

func main() {
	fmt.Println("Hello, World!")
	//                             0 1 2 3 4
	fmt.Println("Рандомная цифра", rand.Intn(5))

	// 0 1
	// 50% того что выпадет 0 50% того что выпадет 1
	if rand.Intn(2) == 1 {
		fmt.Println("Я сгенерировал рандомно число один!")
	} else {
		fmt.Println("Я не попал в единицу")
	}
}


//
package main

import (
	"fmt"
	"time"
)

// "math/rand" // ДОПОЛНИТЕЛЬНЫЕ ЯЩИК ИНСТРУМЕНТОВ
// "time" // ДОПОЛНИТЕЛЬНЫЕ ЯЩИК ИНСТРУМЕНТОВ

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Новая итерация! i:", i)
		time.Sleep(500 * time.Millisecond)
	}
}


package main

import (
	"fmt"
	"time"
	"math/rand"
)

// 1. бесконечные циклы
// 2. ключевое слово break

func main() {
	// 🔋
	// 🐤
	// 🧨

	score := 0

	fmt.Println("Get Ready")
	fmt.Println("Счёт:", score)
	fmt.Println("")

	// без i := 1; i <= 6; i++, for будет просто бесконено воспроизводить
	for {
		fmt.Println("------------------------")
		fmt.Println("Подлетаю к трубе!")
		fmt.Println("🐤 🔋🔋")
		fmt.Println("")

		fmt.Println("Пролетаю через трубу!")
		fmt.Println("🔋🐤🔋")
		fmt.Println("")

		// 0 1 2 3 4 5 6 7
		// 1/8
		if rand.Intn(8) == 1 {
			fmt.Println("Я врезался в трубу :(")
			fmt.Println("🔋❌🔋")
			break // Прерывает цикл
		}

		fmt.Println("Я успешно пролетел через трубу!")
		fmt.Println("🔋🔋 🐤")
		fmt.Println("")

		score++
		fmt.Println("Счёт:", score)
		fmt.Println("------------------------")
		fmt.Println("")

		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("")
	fmt.Println("Game Over!")
	fmt.Println("Ваш итоговый счёт:", score)




	//
	for i := 1; i <= 5; i++ {
		fmt.Println("Итерация:", i)

		if i == 3 {
			break
		}
	}
}
*/

/* ПРОДОЛЖЕНИЕ ПРАКТИКИ 16.09.2026 циклы

1. Задание 1 — От 1 до 20
package main

import "fmt"

func main() {
	for i := 1; i <= 20; i++ {
		fmt.Println(i)
	}
}


//
2. Нечётные числа
package main

import "fmt"

func main() {
	for i := 1; i <= 30; i +=2 {
		fmt.Println(i)
	}
}




3. Обратный отсчёт
package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 20; i >= 1; i-- {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("Пуск")
}



4. Сумма
package main

import "fmt"

func main() {
	sum := 0

	for i := 1; i <= 50; i++ {
		fmt.Println(i)
		sum = i + sum
	}

	fmt.Println(sum)
}


5.
package main

import "fmt"

func main() {
	number := 0

	for i := 1; i >= 10; i++ {
		fmt.Println(number)
	}
}


// ТЕОРИЯ ФУНКЦИИ 18.09.2026

package main

import "fmt"

// 1 -- как работает функции
// 2 -- осмысленные примеры работы функций

func main() {

			score := 5
			fmt.Println("Счёт:", score)

			fmt.Println("До вызова функции")
			hello() // когда вызываем функцию hello() то можно сказать просто вставляем 2 кода
			hello()
			hello()
			fmt.Println("После вызова функции ")

			aboba()
		//

		fmt.Println("Вызываю sqaure(5)")
		square(5)
		fmt.Println("Закончил выполнение square(5)")
		fmt.Println("")

		fmt.Println("Вызываю sqaure(10)")
		square(10)
		fmt.Println("Закончил выполнение square(10)")
		fmt.Println("")

		fmt.Println("Вызываю sqaure(3)")
		square(3)
		fmt.Println("Закончил выполнение square(3)")
		fmt.Println("")
	//

	officiantWork("Иван")
	officiantWork("Олег")
	officiantWork("Василий")

	//

	number1 := 5
	number2 := 10
	text := "Привет, земляне!"

	sum(number1, number2, text)
	sum(10, 33, "Как дела друзья?")
}

// Допустим мы работаем офицциантом в ресторане
// и у нас задача:
// для каждого нового гостя, нужно:
// 1. накрыть ему стол
// 2. поприветствовать по имени
// 3. принять заказ
// 4. принести блюдо

func sum(a int, b int, greeting string) {
	s := a + b

	fmt.Println(greeting)

	fmt.Println("Сумма a+b =", s)
	fmt.Println()
}

func officiantWork(name string) {
	fmt.Println("Накрываю на стол!")
	fmt.Println("Привет,", name, "!")
	fmt.Println("Я принял заказ!")
	fmt.Println("Я принёс блюдо!")
	fmt.Println("")
}

//

// Заготовленные куска кода могут принимать различные аргументы
func square(x int) { // функции могут принимать аргументы
	fmt.Println("Мы приняли в функции переменную x:", x)
	fmt.Println("x в квадрате:", x*x)
}

func hello() {
	fmt.Println("Я функция")
	fmt.Println("Меня вызвали!")
	fmt.Println("Я завершаюсь")
	fmt.Println("")
}

func aboba() {
	a := 1
	b := 10
	c := 3

	fmt.Println("Сумма:", a+b+c)
}



//
package main

import "fmt"

func main() {
	number := 5
	text := "Привет"

	fmt.Println("number:", number)
	fmt.Println("text:", text)

	foo(number, text)

	fmt.Println("number:", number)
	fmt.Println("text:", text)
}

func foo(n int, t string) {
	fmt.Println("n:", n)
	fmt.Println("t:", t)

	n = 100
	t = "Меня поменяли :("

	fmt.Println("n:", n)
	fmt.Println("t:", t)
}



//

package main

import "fmt"

func main() {
	fmt.Println("До вызова функции")
	number := sum(1, 3)
	fmt.Println("После вызова функции")

	fmt.Println("number:", number)
}

func sum(a int, b int) int { // int возвращение числа || можно вернуть все что угодно
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	s := a + b

	return s
}

func sumString(s1 string, s2 string) string {
	s := s1 + s2

	return s
}




//
package main

import "fmt"

func main() {
	text1 := "Hello"
	text2 := "World"

	text3 := sumString(text1, text2)

	fmt.Println("text3:", text3)
}

func sum(a int, b int) int { // int возвращение числа || можно вернуть все что угодно
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	s := a + b

	return s
}

func sumString(s1 string, s2 string) string {
	s := s1 + s2

	return s
}



//
package main

import "fmt"

func main() {
	text1 := "Hello"
	text2 := "World"

	text3 := sumString(text1, text2)

	fmt.Println("text3:", text3)
}

func sumString(s1 string, s2 string) string {
	s := s1 + s2

	return s
}


//
package main

import "fmt"

func main() {
	text1 := "i'm "
	text2 := "Tbe best"

	text3 := gojo(text1, text2)

	fmt.Println("Слияние текста:", text3)
}

func gojo(m string, a string) string {
	resutl := m + a

	return resutl
}




// МОЖЕТ БЫТЬ НЕСКОЛЬКО RETURN
// Для того что написан несколько return нужно условное ветвления
package main

import "fmt"

func main() {
	s := str()

	fmt.Println(s)
}

func str() string {
	subscribed := true

	if subscribed {
		return "Ты подписан на канал!"
	} else {
		return "Ты НЕ подписана на канал!"
	}
}




//
package main

import "fmt"

// 1. return в функции, которая ничего не возвращает

// 2. Глобальные переменные
var number int = 5

var pi float64 = 3.14

func main() {
	kvadrat := square(5)

	fmt.Println("Квадрат:", kvadrat)

	greeting("Арзу")
}

func greeting(name string) {
	if name == "" {
		fmt.Println("Вы передали пустое имя :(")
		return
	}


	fmt.Println("Привет, уважаемый", name)
}

func square(n int) int {
	s := n * n

	return s
}
*/

/*ПРАКТИКА 20.09.2026
package main

import "fmt"

func main() {
	greet()
}


func greet() {
	fmt.Println("Привет, Go!")
}




//
package main

import "fmt"

func main() {
	greet("Арзу")
}

func greet(name string) {
	fmt.Println("Привет,", name)
}




//
package main

import "fmt"

func main() {
	result := add(10, 15)

	fmt.Println("Результат:", result)
}

func add(aizen int, sosuke int) int {
	s := aizen + sosuke

	return s
}




//
package main

import "fmt"

func main() {
	result := multiply(6, 6)

	fmt.Println("Результат:", result)
}

func multiply(a int, b int) int {
	s := a * b

	return s
}




//
package main

import "fmt"

func main() {
	checkAge(17)
}

func checkAge(age int) {
	if age >= 18 {
		fmt.Println("Доступ разрешён")
	} else {
		fmt.Println("Доступ запрещён")
	}
}




//
package main

import "fmt"

func main() {
	result := max(19, 18)

	fmt.Println("Результат:", result)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
*/

/* ТЕОРИЯ 21.09.2026 defer
// STACK - Структура данных это какой то формат хранений данных, данные можно хранить абсолютно по разному

package main

import "fmt"

// Я main и я начался
// Я hello и я начался
// Я функция hello 1
// Я функция hello 2
// Я функция hello 3
// Я hello и я закончился
// Я main и я закончился

func main() {
	fmt.Println("Я main и я начался")
	defer func() { // если мы запишим defer после него анонимную функцию - func() и т.д, и в самой функции напишем какую то логику, то содержимое этой анонимной функции будет выполнена в конце
		fmt.Println("Я main и я закончился") // выполнится самым последним
	} ()
}


func hello() {
	fmt.Println("Я hello и я начался")
	defer func() {
		fmt.Println("Я hello и я закончился")
	} ()

	fmt.Println("Я функция hello 1")
	fmt.Println("Я функция hello 2")
	fmt.Println("Я функция hello 3")
}

// Stack отложенных вызовов
// Stack кладем снизу вверх а достаем сверху вниз

func foo() {
	defer func() {
		fmt.Println("Defer1")
	}()

	defer func() {
		fmt.Println("Defer2")
	}()

	defer func() {
		fmt.Println("Defer3")
	}()

	fmt.Println("foo")
}




func database() {
	// Создать подключение

	// Использовать только этот вариант
	defer func() {
		// закрыть подключение в defer
	}()



	// не использовать этот вариант
	boolean := true

	if boolean {
		//
		// закрыть подключение
		return
	} else {

	}

	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			//
			// закрыть подключение
			return
		} else {
			//
			// закрыть подключение
			return
		}
	}

	// закрыть подключение
}
*/

/* ПРАКТИКА 21.09.2026 defer()

1.
package main

import "fmt"

func main() {
	test()
}

func test() {
	defer func() {
		fmt.Println("Второе сообщение")
	}()

	fmt.Println("Первое сообщение")
	fmt.Println("Третье сообщение")
}



2.
package main

import "fmt"

func main() {
	test()
}

func test() {
	defer func() {
		fmt.Println("Сообщение 3")
	}()

	fmt.Println("Сообщение 1")
	fmt.Println("Сообщение 2")
	fmt.Println("Сообщение 4")
}




3.
package main

import "fmt"

func main() {
	test()
}

func test() {
	defer func() {
		fmt.Println("Первое")
	}()

	defer func() {
		fmt.Println("Второе")
	}()

	defer func() {
		fmt.Println("Третье")
	}()

	fmt.Println("Основной код")
}



4.
package main

import "fmt"

func main() {
	test()
}

func test() {
	x := 10

	defer func() {
		fmt.Println("Значение:", x)
	}()

	x = 20
}




5.
package main

import "fmt"

func main() {
	x := test(100)

	fmt.Println(x)
}

func test(x int) int {
	y := x

	defer func() {
		fmt.Println("Функция завершена:")
	}()

	return y
}



6.
package main

import "fmt"

func main() {
	count(50)
}

func count(y int) {
	s := y

	defer func() {
		fmt.Println("Функция завершена")
	}()

	s++

	fmt.Println(s)
}
*/

/* ТЕОРИЯ 21.09.2026 УКАЗАТЕЛИ
// Указатели это тип переменых, есть тип указатели

package main

import "fmt"

// Такая запись говорит что мы берем адресс у какойто переменной тем самым берем на неё указатель, создаем укзаатель и берем переменную
func main() {
	number := 5

	// Указатель с английского
	// & - Амперсант, Что значит то что я взял указатель на перменную number
	// я сказал языку программироваю что я хочу создать какой то указатель и указывать он будет на переменную название которой ч напишу после амперсанта - &
	// & - амперсант означает что мы берем адресс у переменной
	// Амперсант number - &number это мы берем указатель на number
	pointer := &number

	// 0x3d9056e8e060 в терминале - это адресс в памяти, это адресс нашей перемены на которую мы взяли указатель
	// с помощью адресса - 0x3d9056e8e060, мы можем имея адресс имеем доступ

    foo(pointer)

	fmt.Println(number)
	fmt.Println("Разыменовывание pointer:", *pointer)

	name := "Игорь"

	ptr := &name

	fmt.Println("Имя до изменения:", name)

	changeName(ptr)

	fmt.Println("Имя после изменения:", name)
}

// функция foo будет ожидать переменную n(аргумент) который является указателем на тип данных - int
// * - звездачка говорит что тут будет указатель, звездачка помогает языку программирования отличить когда мы в функции ждем просто переменную int(без *) от случая когда мы в функции когда мы ждем указатель *
// *int это указатель на int
// Такая запись говорит что мы объявляем переменную или объявлем аргумент типа указатель на int - n
func foo(n *int) {
	fmt.Println(n)

	// Вот такая запись говорит о том что указатель разминоваем и тем самым хотим получить доступ к переменной
	fmt.Println(*n)

	*n = 10
}

func boo(n int) {
	n = 10
}

func changeName(s *string) {
	str := "hello"
	s = &str
}



//
package main

import "fmt"

// nil указатели
// что такое nil указатели - то нил указатель это указатель который у нас указывает в пустоту тоесть никуда, если указатель никуда не указывает то это nil
func main() {
	number := 15

	var ptr *int

	fmt.Println("number:", number)
	fmt.Println("ptr:", ptr)

	// чтобы не было ошибок
	if ptr != nil {
		fmt.Println("Разыменование:", *ptr)
	} else {
		fmt.Println("Получен nil-указатель")
	}

}
*/

/* ПРАКТИКА 21.09.2026 УКАЗАТЕЛИ
1.
package main

import "fmt"

func main() {
	x := 10

	// Создаем указатель p, связанный с переменной x (оператор & берет адрес)
	// p хранит адрес x.
	p := &x

	fmt.Println("Значение x:", x)

	// получает значение по этому адресу.
	fmt.Println("Значение p:", *p)
}



2.
package main

import "fmt"

func main() {
	x := 10

	// создать указатель, связанный с x
	y := &x

	// Используя указатель, меняем значение x на 50
	*y = 50

	fmt.Println(x)
}



3.
package main

import "fmt"

func main() {
	// создаю переменную x со значением 20.
	x := 20

	// передаю в функцию адрес переменной x.
	foo(&x)

	// теперь x равно 100.
	fmt.Println(x)
}

// функция получает указатель на int.
func foo(a *int) {
	// через указатель изменяешь значение переменной x.
	*a = 100
}




4.
package main

import "fmt"

func main() {
	x := 10
	y := 20

	fmt.Println("До обмена:", x, y)
	// Передаем адреса переменных с помощью оператора &
	swap(&x, &y)
	fmt.Println("После обмена:", x, y)
}

func swap(x *int, y *int) {
	*x, *y = *y, *x
}



5.
package main

import "fmt"

func main() {
	y := 25

	foo(&y)

	fmt.Println("Результат:", y)
}

func foo(x *int) {
	*x += 10
}



6.
package main

import "fmt"

func main() {
	num := -15

	fmt.Println("До вызова функции:", num)

	// передал адрес переменной.
	result(&num)

	fmt.Println("После вызова функции:", num)
}

// функция получила указатель.
func result(a *int) {
	// Если значение по указателю отрицательное, оно перезаписывается нулем.
	if *a < 0 { // получил значение по адресу.
		// изменяет исходную переменную.
		*a = 0
	}
}
*/

/* ТЕОРИЯ 22.09.2026 СТРУКТУРЫ
1.
package main

import "fmt"

type User struct {
	// Теперь я сказал user что у user есть имя и возраст и т.д
	// Описаное того что является user
	Name string // Значение по умолчание для это - ""
	Age int // Значение по умолчание для это - 0
	PhoneNumber string // Значение по умолчание это - ""
	isClose bool // Значение по умолчание это - false
	Rating float64 // Значение по умолчание это - 0.0 -- 10.0
}

// Структура это объединение каких то переменных, над которомы мы можем производить какие то действия
func main() {
	// Создадим нашего пользователя
	user := User{
		Name: "Сергей",
		Age: 64,
		PhoneNumber: "+7 (911) 911 11-22",
		isClose: true,
		Rating: 5.5,
	}

	user.Rating = 555

	fmt.Println("User:", user)

	// Я хочу выести только имя, нужно написать саму структуру и написать user и через точку выбрать что находится внутри
	// fmt.Println("Имя до:", user.Name)
	// user.Name = "Игорь"
	// fmt.Println("Имя после:", user.Name)
}



2.
package main

import "fmt"

// Если коллеги забыли что рейтинг должен быть от 0.0 до 10.0 - внизу продолжение
// На такие случаи существует механизм - методы
type User struct {
	Name string
	Rating float64
}

func (u User) Greeting() {
	fmt.Println("Всем привет!")
	fmt.Println("Меня зовут", u.Name)
	fmt.Println("Мой рейтинг:", u.Rating)
	fmt.Println("")
}

func (u User) GoodBye() {
	fmt.Println("Всем пока!")
	fmt.Println("Меня звали", u.Name)
	fmt.Println("Мой рейтинг:", u.Rating)
	fmt.Println("")
}

//
func (u User) RatinUP(rating float64) {
	if u.Rating + rating <= 10.0 {
		u.Rating += rating
		// Аналогичный код что и сверху u.Rating = u.Rating + rating
		fmt.Println("Я добавил рейтинг")
	} else {
		fmt.Println("Я не прошёл валидацию")
	}
}
//

// с * ресивер(user) не становится копией
func (u *User) RatinUP(rating float64) {
	if u.Rating + rating <= 10.0 {
		u.Rating += rating
		// Аналогичный код что и сверху u.Rating = u.Rating + rating
		fmt.Println("Я добавил рейтинг")
	} else {
		fmt.Println("Я не прошёл валидацию")
	}
}

func main() {
	user := User{
		Name: "Вася",
		Rating: 6.0,
	}

	// ptr := &user // & указатель

	fmt.Println("User до:", user)
	fmt.Println("")

	//user.Greeting()
	// user.GoodBye()
	// Старый способ GoodBye(user)
	user.RatinUP(1.5)

	fmt.Println("User после:", user)
}



//
package main

import "fmt"

type User struct {
	// Имя пользователя
	// Правило: иия должно быть не пустое
	Name string // ""

	// Возраст пользователя:
	// Правило: возраст должен быть больше 0 и меньше 150
	Age int // 0

	// Номер телефона
	// Правило: номер телефона не должен быть пустым
	PhoneNumber string // ""

	// Закрыт ли профиль пользователя?
	isClose bool // false

	// Рейтинг пользователя
	// Правило: рейтинг должен быть больше либо равен нулю, но меньше либо равен 10
	Rating float64 // 0.0 -- 10.0
}

func NewUser(
	name string,
	age int,
	phoneNumber string,
	isClose bool,
	rating float64,
) User {
	if name == "" {
		return User{}
	}

	if age <= 0 || age >= 150 {
		return User{}
	}

	if phoneNumber == "" {
		return User{}
	}

	if phoneNumber == "" {
		return User{}
	}

	if rating < 0.0 || rating > 10.0 {
		return User{}
	}

	return User{
		Name: name,
		Age: age,
		PhoneNumber: phoneNumber,
		isClose: isClose,
		Rating: rating,
	}
}

func (u *User) ChangeName(newName string) {
	if newName != "" {
		u.Name = newName
	}
}

func (u *User) ChangeAge(newAge int) {
	if newAge > 0 && newAge < 150 {
		u.Age = newAge
	}
}

func (u *User) ChangePhoneNumber(newPhoneNumber string) {
	if newPhoneNumber != "" {
		u.PhoneNumber = newPhoneNumber
	}
}

func (u *User) CloseAccount() {
	// Закрыть аккаунт
	u.isClose = true
}

func (u *User) OpenAccount() {
	// Открыть аккаунт
	u.isClose = false
}

func (u *User) RatingUP(rating float64) {
	if u.Rating + rating <= 10.0 {
		u.Rating += rating
	}
}

func (u *User) RatingDOWN(rating float64) {
	if u.Rating - rating >= 0.0 {
		u.Rating -= rating
	}
}

// func main() {
	// user := User{
		// Name: "Сергей",
		// Age: 64,
		// PhoneNumber: "+7 (911) 911 11-22",
		// isClose: true,
		// Rating: 5.5,
	// }

	// fmt.Println("Рейтинг до:", user.Rating)

	// user.RatingUP(2.4)

	// fmt.Println("Рейтинг после:", user.Rating)
// }


func main() {
	user := NewUser(
		"Данил",
		50,
		"79118707090",
		false,
		4.5,
	)

	fmt.Println("User:", user)
}
*/

/* ПРАКТИКА 24.09.2026 СТРУКТУРЫ

1.
package main

import "fmt"

type Person struct {
	Name string
	Age int
	City string
}

func main() {
	person := Person{
		Name: "Арзу",
		Age: 17,
		City: "Волгоград",
	}

	fmt.Println("Информация:", person)
}




2.
package main

import "fmt"

type Car struct {
	Brand string
	Model string
	Age int
}

func main() {
	car1 := Car{
		Brand: "Mercedes",
		Model: "CLS 63 AMG",
		Age: 2010,
	}

	car2 := Car{
		Brand: "Audi",
		Model: "RS7",
		Age: 2018,
	}

	fmt.Println("Информация о машине:", car1)
	fmt.Println("Информация о машине:", car2)
}




3.
package main

import "fmt"

type Student struct {
	Name string
	Age int
	Grade float64 // 0.0 -- 10.0
}

func main() {
	student := Student{
		Name: "Арзу",
		Age: 17,
		Grade: 5.5,
	}

	// Изменить оценку студенту
	student.Grade = 10.0

	fmt.Println("Информация о студенте:", student)
}




4.
package main

import "fmt"

type Book struct {
	Name string
	Author string
	NumberPages int
}

func main() {
	book := Book{
		Name: "Страник по звездам",
		Author: "Джек Лондон",
		NumberPages: 384,
	}

	fmt.Println("Информация о название книги:", book.Name)
	fmt.Println("Информация об авторе:", book.Author)
	fmt.Println("Информация о количестве страниц:", book.NumberPages)
}




5.
package main

import "fmt"

type Player struct {
	Name string
	Lvl int
	Points float64
}

func main() {
	player := Player{
		Name: "Арзу",
		Lvl: 17,
		Points:8.7,
	}

	player.Lvl += 1
	player.Points += 1.3

	fmt.Println("Информация об игроке:", player)
}



6.
package main

import "fmt"

type Game struct {
	GameTitle string
	Genre string
	NumberOfhours int
}

func main() {
	game := Game{
		GameTitle: "CS2",
		Genre: "Shooter",
		NumberOfhours: 235,
	}

	// Изменил количество часов
	game.NumberOfhours = 240

	fmt.Println("Название игры:", game.GameTitle)
	fmt.Println("Жанр:", game.Genre)
	fmt.Println("Количество часов:", game.NumberOfhours)
}
*/

/* ТЕОРИЯ 24.09.2026 МОДУЛИ И ПАКЕТЫ
// МОДУЛЬ ЭТО КАКОЙ ТО ПРОЕКТ В КОТОРОМ ЕСТЬ ГЛВ=АВНЫЙ ФАЙЛ ИЛИ ЕСТЬ КАКИЕ ТО ЕЩЕ ФАЙЛЫ ЭТО МОДУЛЬ
// 1 Модуль это программы которые мы можем запускать и 2 модуль это какие то отдельрные библиотеки в которых есть файл и подтерикторий
// Короче говоря Модуль это либо наша целая программа либо наша целая библиотека
// А пакет это просто какая то отдельная папочка с файлами и все файлы которые лежат в этой папочке они относятся  к пакету например: User/(Главный пакет), user1.go - user2.go - user3.go
// Но если в пакете другие папочки то это уже другие пакеты и для них будут другие правила

// go mod init study - Модуль который можно создать в Терминале
// go mod init study перевод с английского - Голанг возьми модуль проинииализируй его и назови его как я его напишу

// go get... - Это чтобы скачивать с какого то сайта например с GitHub

package main

// fmt - Это стандартная библиотека
import (
	"fmt"
	"study/greeting"
)

// Для того чтобы вызывать функцию с какого-то пакета нужно посмотреть название этого пакета в моем случае это greeting, написать название этго пакета в main.go

func main() {
	fmt.Println("Hello, world!")

	greeting.SayHello()

	greeting.SayBad()

	i := greeting.GiveMeInt()

	fmt.Println("i:", i)
}



//
package main

import (
	"fmt"
	"study/greeting"
)

func main() {
	fmt.Println("Hello, World!")
	greeting.Greeting()
}
*/

/* 24.09.2026 ПРАКТИКА Модули и пакеты
1.
package main

import (
	"fmt"
	"study/greeting"
)

func main() {
	fmt.Println("Hello, World!")
	greeting.SayHello()
}


2.
package main

import (
	"fmt"
	"study/calcult"
)

func main() {
	result := calcult.Result(15, 27)

	fmt.Println("Результат:", result)
}



3.
package main

import (
	"fmt"
	"study/calcult"
)

func main() {
	result1 := calcult.OneFunc(17, 18)

	result2 := calcult.TwoFunc(6, 7)

	fmt.Println("Результат1:", result1)
	fmt.Println("")
	fmt.Println("Результат2:", result2)
}



4.
package main

import (
	"fmt"
	plar "study/calcult/Plar"
)

func main() {
	finish := plar.Player{
		Name: "Арзу",
		Lvl: 17,
		Points: 10.0,
	}

	fmt.Println("Характеристики игрока:", finish)
}



5.
package main

import (
	"fmt"
	plar "study/calcult/Plar"
)

func main() {
	finish := plar.Player{
		Name: "Арзу",
		Lvl: 86,
		Points: 8.9,
	}

	plar.LvlUP(&finish)

	fmt.Println("Характеристики:", finish)
}
*/

/* ТЕОРИЯ 26.09.2026 Массивы
// Масив это - последовательность элементов в память, первый элемент, втрой элемент, третий элемент и т.д, они расположены последовательно и мы можем обращаться к ним с помощью индексации
1.
package main

import "fmt"

func main() {
	// arr - array - с английского Массив
	// Создали Массив целых чисел
	// 5 - это сколько в массиве
	//            0  1   2  3    4
	arr := [5]int{5, 66, 7, 100, 1}

	for i := 0; i < 5; i++ {
		        // %2 - это операция деления с остатком.
		if arr[i] % 2 == 0 { // == 0 - Остаток равен нулю только у чётных чисел (например: 4 / 2 = 2, остаток 0). Именно они умножаются на 2.
			arr[i] *= 2
		}
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i, "-", arr[i])
	}

	//fmt.Println(arr[1])

	// К 66 прибавил 5
	//arr[1] += 5

	//fmt.Println(arr[1])

	//fmt.Println(arr)
}




2.
package main

import "fmt"

type User struct {
	Name string
	Rating float64
	Premium bool
}

func main() {
	//
	user1 := User{
		Name: "Данил",
		Rating: 5.5,
		Premium: true,
	}

	user2 := User{
		Name: "Петя",
		Rating: 4.5,
		Premium: false,
	}

	user3 := User{
		Name: "Вася",
		Rating: 7.5,
		Premium: true,
	}
	//

	userArray := [3]User{
		// Теперь один источник правды, это замена которая выше
	User{
		Name: "Данил",
		Rating: 5.5,
		Premium: true,
	},

	User{
		Name: "Петя",
		Rating: 4.5,
		Premium: false,
	},

	User{
		Name: "Вася",
		Rating: 7.5,
		Premium: true,
	},
}

   fmt.Println("До")
   fmt.Println("-------")
   for i := 0; i < len(userArray); i++ {
	fmt.Println(userArray[i])
   }
   fmt.Println("")

   for i := 0; i < len(userArray); i++ {
	if userArray[i].Premium {
		userArray[i].Rating += 1.0
	}
   }

   fmt.Println("После")
   fmt.Println("-------")
   for i := 0; i < len(userArray); i++ {
	if userArray[i].Premium {
		fmt.Println(userArray[i])
	}
   }
   fmt.Println("")

//
	fmt.Println("До")
	fmt.Println("-------")
	fmt.Println(user1)
	fmt.Println(user2)
	fmt.Println(user3)
	fmt.Println("")

	if user1.Premium {
		user1.Rating += 1.0
	}

	if user2.Premium {
		user2.Rating += 1.0
	}

	if user3.Premium {
		user3.Rating += 1.0
	}

	fmt.Println("После")
	fmt.Println("-------")
	fmt.Println(user1)
	fmt.Println(user2)
	fmt.Println(user3)
//
}




2.
// Слайсы

package main

import "fmt"

type User struct {
	Name    string
	Rating  float64
	Premium bool
}

func main() {
	//
		user1 := User{
			Name: "Данил",
			Rating: 5.5,
			Premium: true,
		}

		user2 := User{
			Name: "Петя",
			Rating: 4.5,
			Premium: false,
		}

		user3 := User{
			Name: "Вася",
			Rating: 7.5,
			Premium: true,
		}
	//

	//               с тройкой это Массив без тройки слайс
	userArray := []User{
		// Теперь один источник правды, это замена которая выше
		User{
			Name:    "Данил",
			Rating:  5.5,
			Premium: true,
		},

		User{
			Name:    "Петя",
			Rating:  4.5,
			Premium: false,
		},

		User{
			Name:    "Вася",
			Rating:  7.5,
			Premium: true,
		},
	}

	fmt.Println("До")
	fmt.Println("-------")
	fmt.Println("len:", len(userArray))
	fmt.Println("cap:", cap(userArray))
	fmt.Println("")

	userArray = append(
		userArray,
		User{
		Name: "Виталик",
		Rating: 4.0,
		Premium: true,
	})

	fmt.Println("После")
	fmt.Println("-------")
	fmt.Println("len:", len(userArray))
	fmt.Println("cap:", cap(userArray))
	fmt.Println("")

	for i := 0; i < len(userArray); i++ {
		if userArray[i].Premium {
			userArray[i].Rating += 1.0
		}
	}

	for i := 0; i < len(userArray); i++ {
		fmt.Println(userArray[i])
	}

	// fmt.Println("До")
	// fmt.Println("-------")
	// for i := 0; i < len(userArray); i++ {
	//	fmt.Println(userArray[i])
	// }
	// fmt.Println("")

	// for i := 0; i < len(userArray); i++ {
	//	if userArray[i].Premium {
	//		userArray[i].Rating += 1.0
	//	}
	// }

	// fmt.Println("После")
	// fmt.Println("-------")
	// for i := 0; i < len(userArray); i++ {
	//	if userArray[i].Premium {
	//		fmt.Println(userArray[i])
	//	}
	// }
	// fmt.Println("")
}
//


package main

import "fmt"

func main() {
	//               создаться слайс целых чисел
	//                     если написать 5 то будет пять нулей
	// intSlice := make([]int, 0)
	intSlice := make([]int, 0, 5)

	intSlice = append(intSlice, 10, 15, 25, 150)

	// Добавляем с помощью слайса в intSlice цифры
	// intSlice = append(intSlice, 10)
	// intSlice = append(intSlice, 15)
	// intSlice = append(intSlice, 25)
	// intSlice = append(intSlice, 150)

	fmt.Println(intSlice)


	// СЛАЙС
	weather := []int{
		+3,
		+6,
		+9,
		-4,
		+1,
	}
}





// Мапы
package main

import "fmt"

func main() {
	weather := map[int]int{
	//  key -> value
		11: +3,
		12: +6,
		13: +9,
		14: -4,
		15: +1,
	}

	for k, _ := range weather {
		weather[k] += 1
	}

	fmt.Println(weather)

	weather[11] += 1
	weather[12] += 1
	weather[13] += 1
	weather[14] += 1
	weather[15] += 1
	//
	c, ok := weather[11]
	if ok {
		fmt.Println("OKAY")
	} else {
		fmt.Println("bad")
	}
	//

	fmt.Println(weather[11])
	weather[11] = -10
	fmt.Println(weather[11])
}



//
package main

import "fmt"

func main() {
	// нам нужно хранить данные о людях
	// нужно по имени человека сразу понимать
	// был ли он судим

	criminal := map[string]bool{
		"Вася": true,
		"Петя": false,
		"Антон": false,
		"Вова": false,
		"Алексей": true,
	}

	c, ok := criminal["Никита"]

	if !ok {
		fmt.Println("Человека нет в базе")
		return
	}

	fmt.Println("Человек найден в базе!")
	if c {
		fmt.Println("Человек судим!")
	} else {
		fmt.Println("Человек НЕ судим")
	}
}
*/

/* ДОПОЛНИТЕЛЬНО К Массивы/Слайсы/Мапы
package main

import "fmt"

func main() {
	// Массивы
	arr := [3]int{1, 2, 3}
	fmt.Println("Массив:", arr, len(arr), cap(arr))

	// Слайсы
	slice := make([]int, 0, 5)
	slice = append(slice, 4)
	slice = append(slice, 5)
	slice = append(slice, 6)
	fmt.Println("Слайс:", slice, len(slice), cap(slice))

	// Мапы       // месяц  дни
	mapa := make(map[string]int, 10)
	mapa["Январь"] = 31
	mapa["Февраль"] = 28
	mapa["Март"] = 31
	fmt.Println("Мапа:", mapa, len(mapa))
}
*/

/* ПРАКТИКА 28.09.2026 Массивы/Слайсы/Мапы
1. Массив чисел
package main

import "fmt"

func main() {
	//            0   1   2   3   4
	arr := [5]int{10, 20, 30, 40, 50}

	fmt.Println("Массив:", arr)
}




2. изменение элемента массива
package main

import "fmt"

func main() {
	//           0   1   2   3   4
	arr := [5]int{5, 10, 15, 20, 25}

	arr[2] = 100

	fmt.Println("Массивы:", arr)
}




3.
package main

import "fmt"

func main() {
	number := [5]int{10, 20, 30, 40, 50}

	result := 0

	for i := 0; i < len(number); i++ {
		result += number[i]
	}

	fmt.Println("Подсчёт всех элементов:", result)
}




3.1 деление с остатком + массив
package main

import "fmt"

func main() {
	number := [10]int{3, 8, 12, 15, 21, 24, 30, 33, 40, 47}

	for i := 0; i < len(number); i++ {
		if number[i] %2 == 0 {
			fmt.Println(number[i])
		}
	}
}




4. создание и добавление
package main

import "fmt"

func main() {
	users := []string{
		"Арзу",
		"Аизен",
		"Соске",
	}

	users = append(users, "Годжо", "Гриффит")

	fmt.Println("Слайсы:", users)
}




5. работа со слайсом
package main

import "fmt"

func main() {
	point := []int{10, 20, 30, 40, 50}

	point[2] = 100

	point = append(point, 60)

	fmt.Println("Вывод слайсов:", point)
}




6. сумма элементов слайса
package main

import "fmt"

func main() {
	score := []int{5, 10, 15, 20, 25}

	result := 0

	for i := 0; i < len(score); i++ {
		result += score[i]
	}

	fmt.Println("Результат сложения:", result)
}




7. создание Map
package main

import "fmt"

func main() {
	users := make(map[string]int)
//    Имена - ключ   Цифры - значения
	users["Арзу"] = 17
	users["Аизен"] = 25
	users["Годжо"] = 28

	fmt.Println(users)
}




8. получение и изменение значения
package main

import "fmt"

func main() {
	userName := make(map[string]int)
	userName["Арзу"] = 17
	userName["Аизен"] = 25
	userName["Годжо"] = 28

	fmt.Println("Возраст до:")
	fmt.Println(userName["Арзу"])
	fmt.Println("")

	userName["Арзу"] = 18

	fmt.Println("Возраст после:")
	fmt.Println(userName["Арзу"])
}



9.
package main

import "fmt"

func main() {
	userName := make(map[string]int)
	userName["Арзу"] = 17
	userName["Аизен"] = 25
	userName["Годжо"] = 28

	_, exists := userName["Соске"]

	if exists {
		fmt.Println("Человек найден")
	} else {
		fmt.Println("Человек не найден")
	}
}
*/

/* ТЕОРИЯ 29.09.2026 Пользовательский ввод
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Введите команду: добавить лук
// Вы хотите добавить лук

// Введите команду: удалить морковь
// вы кажется хотите удалить морковь

// неизвестная команда
// |
// help
// вот список команд!

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите команду:")

	if ok := scanner.Scan(); !ok {
		fmt.Println("Ошибка ввода")
		return
	}

	text := scanner.Text()

	fields := strings.Fields(text)

	if len(fields) == 0 {
		fmt.Println("Вы ничего не ввели!")
		return
	}

	fmt.Println("text:", text)

	fmt.Println("Слова:", fields)
	fmt.Println("Команда:", fields[0])

	cmd := fields[0]

	if cmd == "добавить" {
		str := ""
		for i := 1; i < len(fields); i++ {
			str += fields[i]

			if i != len(fields)-1 {
				str += " "
			}
		}
	}

	if cmd == "добавить" {
		fmt.Println("вы хотите добавить")
		fmt.Println("")
	} else if cmd == "удалить" {
		fmt.Println("вы кажется хотите удалить")
		fmt.Println("")
	} else if cmd == "help" {
		fmt.Println("-- эта команда выводит список доступных команд")
		fmt.Println("")
	} else {
		fmt.Println("Вы ввели неизвестную команду")
	}
}
*/

/* ПРАКТИКА 30.09.2026 Пользовательский ввод
1. Имя пользователя
package main

import "fmt"

func main() {
	var name string

	fmt.Print("Введите ваше имя: ")

	fmt.Scan(&name)

	fmt.Println("Привет:,", name + "!")
}



2. Возраст
package main

import "fmt"

func main() {
	var age int

	fmt.Print("Введите ваш возраст: ")

	fmt.Scan(&age)

	fmt.Println("Ваш возраст:", age)
}




3.Два числа
package main

import "fmt"

func main() {
	var first, second int

	fmt.Print("Введите первое число: ")
	_, err := fmt.Scan(&first)
	if err != nil {
		fmt.Println("Ошибка, введите целое число!")
		return
	}

	fmt.Print("Введите второе число: ")
	_, err = fmt.Scan(&second)
	if err != nil {
		fmt.Println("Ошибка, введите целое число!")
		return
	}

	sum := first + second
	fmt.Println("Сложение чисел:", sum)
}




4. Информация о пользователе
package main

import "fmt"

func main() {
	var name, city string
	var age int

	fmt.Print("Введите ваше имя: ")
	fmt.Scan(&name)

	fmt.Print("Введите ваш возраст: ")
	fmt.Scan(&age)

	fmt.Print("Введите ваш город: ")
	fmt.Scan(&city)

	fmt.Println("Имя:", name, "Возраст:", age, "Город:", city)
}




5. Калькулятор
package main

import "fmt"

func main() {
	var first, second int
	var action string

	fmt.Print("Введите первое число: ")
	fmt.Scan(&first)

	fmt.Print("Введите второе число: ")
	fmt.Scan(&second)

	fmt.Print("Введите действие (+, -, *, /): ")
	fmt.Scan(&action)

	switch action {
	case "+":
		fmt.Println("Результат:", first + second)

	case "-":
		fmt.Println("Результат:", first - second)

	case "*":
		fmt.Println("Результат:", first * second)

	case "/":
		if second == 0 {
			fmt.Println("Делить на ноль нельзя!")
			return
		}

		fmt.Println("Результат:", first/second)

	default:
		fmt.Println("Неизвестное действие!")
	}
}




6.
package main

import "fmt"

func main() {
	var first, second, third float64 // Если ты хочешь получать точное среднее, используй float64 для чисел:

	fmt.Print("Введите первое число: ")
	fmt.Scan(&first)

	fmt.Print("Введите второе число: ")
	fmt.Scan(&second)

	fmt.Print("Введите третье число: ")
	fmt.Scan(&third)

	average := (first + second + third) / 3

	fmt.Println("Среднее значение:", average)
}
*/

/* ТЕОРИЯ КУРС 2 01.10.2026 Интерфейсы - ПРОЙДЕНА НЕ ДО КОНЦА ТЕОРИЯ!!!
// Интерфейс для взаимодействия между пользователем и сайтом
1.
package main

import "fmt"

type Auto interface {
	StepOnGas()
	StepOnBrake()
}

type BMW struct {}

func (b BMW) StepOnGas() {
	fmt.Println("Я БМВ! 550 лошадинных сил! Жмём на газ!")
}

func (b BMW) StepOnBrake() {
	fmt.Println("Я БМВ! Тормоз в пол! Хорошее торможение!")
}

type Zhiga struct {}

func (z Zhiga) StepOnGas() {
	fmt.Println("Я жига! 2107! Пробую не развалиться!")
}

func (z Zhiga) StepOnBrake() {
	fmt.Println("Я жига! Вот блин тормоза чуть не отвалились(")
}

func rideBMW(bmw BMW) {
	fmt.Println("Я водитель!")
	fmt.Println("Я сажусь в свою машину!")
	fmt.Println("И нажимаю на газ...")
	bmw.StepOnGas()
}

func rideZhiga(zhiga Zhiga) {
	fmt.Println("Я водитель!")
	fmt.Println("Я сажусь в свою машину!")
	fmt.Println("И нажимаю на газ...")
	zhiga.StepOnGas()
}

func ride(auto Auto) {
	fmt.Println("Я водитель!")
	fmt.Println("Я сажусь в свою машину!")
	fmt.Println("И нажимаю на газ...")
	auto.StepOnGas()
	auto.StepOnBrake()
}

func main() {
	//
	bmw := BMW{}

	bmw.StepOnGas()

	zhiga := Zhiga{}

	zhiga.StepOnGas()
	//

	bmw := BMW{}
	//zhiga := Zhiga{}

	ride(bmw)
	//ride(zhiga)
}



package main

import (
	"fmt"
	//"github.com/k0kubun/pp"
	"strconv"
	"time"
)

// [ 10_000, 50_000, 12_000, 6_000, 100_000 ]
// 6_000
// 10_000
// 12_000
// 50_000
// 100_000

// slice vs map

// 0(n)
// us - микросекунд
// 1 -> 10us
// 10 -> 100us
// 100 -> 1000us
// 1_000_000 -> 10_000_000 -> 10us

// map:
// 0(1)
// 1 -> 10us
// 10 -> 10us
// 100 -> 10us
// 1_000_000 -> 10us

// slice[0] -> 10us
// slice[10_000_000] -> 10s

// 1. сохранять большое количество PaymentInfo
// 2. находить конкретный элемент PaymentInfo в множестве

type PaymentInfo struct {
	ID          int
	Description string
	Usd         int
	Cancelled   bool
}

type PaymentModuleWithSlice struct {
	s []PaymentInfo
}

func (m *PaymentModuleWithSlice) AddInfo(info PaymentInfo) {
	m.s = append(m.s, info)
}

func (m *PaymentModuleWithSlice) GetInfo(id int) PaymentInfo {
	for _, info := range m.s {
		if info.ID == id {
			return info
		}
	}

	return PaymentInfo{}
}

// ----------------------------------------------------

type PaymentModuleWithMap struct {
	m map[int]PaymentInfo
}

func (m *PaymentModuleWithMap) AddInfo(info PaymentInfo) {
	m.m[info.ID] = info
}

func (m *PaymentModuleWithMap) GetInfo(id int) PaymentInfo {
	info, ok := m.m[id]
	if !ok {
		// если не ok то возвращать PaymentInfo{}
		return PaymentInfo{}
	}

	// Если ok то возвращать info
	return info
}

func main() {
	pSlice := PaymentModuleWithSlice{}
	pMap := PaymentModuleWithMap{
		m: make(map[int]PaymentInfo),
	}

	interations := 10_000

	// ------------------------------------------------

	before := time.Now()

	for i := 0; i < interations; i++ {
		info := PaymentInfo{
			ID: i,
			Description: strconv.Itoa(i), // возвращет в строку: 1 -> "1", 123 -> "123"
		}

		pSlice.AddInfo(info)
	}

	fmt.Println("slice add:", time.Since(before)) // Вывожу сколько времени прошло с момента как я запомнил


	before = time.Now()
	for i := 0; i < interations; i++ {
		info := PaymentInfo {
			ID: i,
			Description: strconv.Itoa(i), // возвращет в строку: 1 -> "1", 123 -> "123"
		}

		pMap.AddInfo(info)
	}

	fmt.Println("map add:", time.Since(before)) // Вывожу сколько времени прошло с момента как я запомнил

	// ------------------------------------------------

	before = time.Now()

	for i := 0; i < interations; i++ {
		info := pSlice.GetInfo(i)
		_ = info // _ - это заглушка что типо ее использовал где то
	}

	fmt.Println("slice get:", time.Since(before)) // Вывожу сколько времени прошло с момента как я запомнил


	before = time.Now()

	for i := 0; i < interations; i++ {
		info := pMap.GetInfo(i)
		_ = info // _ - это заглушка что типо ее использовал где то
	}

	fmt.Println("map get:", time.Since(before)) // Вывожу сколько времени прошло с момента как я запомнил
	//
	info1 := PaymentInfo{
		ID:          10,
		Description: "BURGER",
		Usd:         5,
		Cancelled:   false,
	}

	pSlice.AddInfo(info1)
	pMap.AddInfo(info1)

	pp.Println("pSlice:", pSlice)
	pp.Println("pMap:", pMap)

	i1 := pSlice.GetInfo(10)
	i2 := pMap.GetInfo(10)

	fmt.Println("i1:", i1)
	fmt.Println("i2:", i2)
	//
}
*/

/* ПРАКТИКА 03.10.2026 Интерфейсы
1.
package main

import "fmt"

type Animal interface {
	MakeSound()
}

type Dog struct {}

func (d Dog) MakeSound() {
	fmt.Println("Я собака: Гав!")
}

type Cat struct {}

func (c Cat) MakeSound() {
	fmt.Println("Я кошка: Мяу!")
}

func main() {
	var dog Animal = Dog{}
	var cat Animal = Cat{}

	dog.MakeSound()
	cat.MakeSound()
}




2.
package main

import "fmt"

type Work interface {
	GetJobTitle() string
}

type Programmer struct {}

func (p Programmer) GetJobTitle() string {
	return "Программист"
}

type Designer struct {}

func (d Designer) GetJobTitle() string {
	return "Дизайнер"
}

func main() {
	var programmist Work = Programmer{}
	var designer Work = Designer{}

	fmt.Println("Профессия:", programmist.GetJobTitle())
	fmt.Println("Профессия:", designer.GetJobTitle())
}




3.
package main

import "fmt"

type Vehicle interface {
	CarsNameSpeed()
}

type Car struct {}

func (c Car) CarsNameSpeed() {
	fmt.Println("Название машины: AUDI RS7")
	fmt.Println("Максимальная скорость: 357км/ч")
}

type Bike struct {}

func (b Bike) CarsNameSpeed() {
	fmt.Println("Название велосипеда: Welt Edge 2.0")
	fmt.Println("Максимальная скорость: 37.5км/ч")
}

func main() {
	var car Vehicle = Car{}
	var bike Vehicle = Bike{}

	car.CarsNameSpeed()
	bike.CarsNameSpeed()
}




4.
package main

import "fmt"

type Player interface {
	GetName() string
	GetLevel() int
}

type Warrior struct {
	Name string
	Level int
}

func (w Warrior) GetName() string {
	return w.Name
}

func (w Warrior) GetLevel() int {
	return w.Level
}

type Mage struct {
	Name string
	Level int
}

func (m Mage) GetName() string {
	return m.Name
}

func (m Mage) GetLevel() int {
	return m.Level
}

func main() {
	var warrior Player = Warrior{
		Name: "Арзу",
		Level: 17,
	}

	var mage Player = Mage{
		Name: "Годжо",
		Level: 27,
	}

	fmt.Println("Имя:", warrior.GetName())
	fmt.Println("Уровень:", warrior.GetLevel())

	fmt.Println("")

	fmt.Println("Имя:", mage.GetName())
	fmt.Println("Уровень:", mage.GetLevel())
}



5.
package main

import "fmt"

type Calculator interface {
	Calculate(a int, b int) int
}

type Addition struct {}

func (a Addition) Calculate(x int, y int) int {
	return x + y
}

type Multiplication struct {}

func (m Multiplication) Calculate(x int, y int) int {
	return x * y
}

func main() {
	var addition Calculator = Addition{}
	var multiply Calculator = Multiplication{}

	result1 := addition.Calculate(10, 5)
	result2 := multiply.Calculate(10, 5)

	fmt.Println("Результат сложения:", result1)
	fmt.Println("Результат умножения:", result2)
}




6.
package main

import "fmt"

type Animal interface {
	NameSound()
}

type Dog struct {}

func (d Dog) NameSound() {
	fmt.Println("Собака: Гав!")
}

type Cat struct {}

func (c Cat) NameSound() {
	fmt.Println("Кошка: Мяу!")
}

type Bird struct {}

func (b Bird) NameSound() {
	fmt.Println("Птица: Чик-Чирик!")
}

func main() {
	var dog Animal = Dog{}
	var cat Animal = Cat{}
	var bird Animal = Bird{}

	dog.NameSound()
	cat.NameSound()
	bird.NameSound()
}
*/

/* ТЕОРИЯ 03.10.2026 Обработка ошибок

package main

import (
	"errors"
	"fmt"
	"math/rand"
	"github.com/k0kubun/pp"
)

type User struct {
	Name     string
	Ballance int
}

func Pay(user *User, usd int) error {
	if user.Ballance - usd < 0 {
		// return "Недостаточно средств!"
		return errors.New("Недостаточно средств!")
	}
		user.Ballance -= usd

		// сетевой запрос

	return nil

}

type Car struct {
	Armor int
}

func (c *Car) Gas() (int, error) {
	if c.Armor - 10 <= 0 {
		return 0, errors.New("Мы не стали газовать, чтобы не сломать машину!")
	}

	kmch := rand.Intn(150)

	c.Armor -= 10

	return kmch, nil
}

func main() {
	//
	user := User{
		Name:     "Олег",
		Ballance: 100,
	}

	pp.Println("User до:", user)
	err := Pay(&user, 50)
	pp.Println("User после:", user)

	if err != nil {
		fmt.Println("Оплаты не было! Причина:", err.Error())
	} else {
		fmt.Println("Была произведена оплата!")
	}
	//

	car := Car{
		Armor: 25,
	}

	pp.Println("Car до:", car)

	kmch, err := car.Gas()
	pp.Println("Car после:", car)
	if err != nil {
		fmt.Println("Ошиба нажатия на газ:", err.Error())
	}

	fmt.Println("Получившийся разгон:", kmch)

	pp.Println("Car после:", car)

	kmch, err = car.Gas()
	fmt.Println("Получившийся разгон:", kmch)

	pp.Println("Car после:", car)

	kmch, err = car.Gas()
	fmt.Println("Получившийся разгон:", kmch)

	pp.Println("Car после:", car)

	kmch, err = car.Gas()
	fmt.Println("Получившийся разгон:", kmch)

	pp.Println("Car после:", car)
}
*/

/* ПРАКТИКА 04.10.2026 Обработка ошибок
1.
package main

import (
	"fmt"
	"errors"
)

// Создай функцию, которая принимает два целых числа и выполняет деление.
func divide(a int, b int) (int, error) {
	// Если выполнить операцию нельзя, функция должна сообщить об ошибке.
	if b == 0 {
		return 0, errors.New("деление невозможно!")
	}
	// Если операция возможна, функция должна вернуть результат.
	return a / b, nil
}

func main() {
	// Вариант 1: деление 20 на 4
	res1, err1 := divide(20, 4)
	if err1 != nil {
		fmt.Println("Ошибка:", err1)
	} else {
		fmt.Println("Результат:", res1)
	}

	// Вариант 2: деление 20 на 0
	_, err2 := divide(20, 0)
	if err2 != nil {
		fmt.Println("Ошибка:", err2) // Выведется: Ошибка: деление невозможно
	}
}




2.
package main

import (
	"fmt"
	"errors"
)

func checkAge(age int) (string, error) {
	if age < 0 {
		return "", errors.New("возраст не может быть отрицательным")
	}
	return "Возраст корректный", nil
}

func main() {
	// Первая проверка возраста
	if msg, err := checkAge(17); err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println(msg)
	}

	// Вторая проверка возраста
	if msg, err := checkAge(-5); err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println(msg)
	}
}




3.
package main

import (
	"errors"
	"fmt"
)

func checkNumbers(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("число не может быть отрицательным")
	}
	return n, nil
}

func main() {
	numbers := []int{25, 10, 0}

	for _, num := range numbers {
		result, err := checkNumbers(num)
		if err != nil {
			fmt.Println("Ошибка:", err)
		} else {
			fmt.Println("Результат:", result)
		}
	}
}




4.
package main

import (
	"fmt"
	"errors"
)

func FindUser(username string) (string, error) {
	users := []string{"Арзу", "Аизен", "Годжо"}

	for _, user := range users {
		if user == username {
			return "пользователь найден", nil
		}
	}

	return "", errors.New("пользователь не найден")
}

func main() {
	result1, err1 := FindUser("Арзу")
	if err1 != nil {
		fmt.Println("Ошибка:", err1)
	} else {
		fmt.Println(result1)
	}

	result2, err2 := FindUser("Соске")
	if err2 != nil {
		fmt.Println("Ошибка:", err2)
	} else {
		fmt.Println(result2)
	}
}




5.
package main

import (
	"fmt"
	"errors"
)

func checkPassword(password string) (string, error) {
	if len(password) < 6 {
		return "", errors.New("пароль слишком короткий")
	}
	return "Пароль принят", nil
}

func main() {
	result1, err1 := checkPassword("123")
	if err1 != nil {
		fmt.Println("Ошибка:", err1)
	} else {
		fmt.Println(result1)
	}

	result2, err2 := checkPassword("qwerty123")
	if err2 != nil {
		fmt.Println("Ошибка:", err2)
	} else {
		fmt.Println(result2)
	}
}




6.
package main

import (
	"fmt"
	"errors"
)

func checkNumbers(n int) (string, error) {
	if n == 0 {
		return "", errors.New("число не может быть равно нулю")
	}

	return "число корректное", nil
}

func main() {
	result1, err1 := checkNumbers(10)
	if err1 != nil {
		fmt.Println("Ошибка:", err1)
	} else {
		fmt.Println(result1)
	}

	result2, err2 := checkNumbers(0)
	if err2 != nil {
		fmt.Println("Ошибка:", err2)
	} else {
		fmt.Println(result2)
	}
}
*/

/* ТЕОРИЯ 04.10.2026 ПАНИКА
package main

import "fmt"

func main() {
	defer func() {
		p := recover()
		if p != nil {
			fmt.Println("Была паника:", p)
		}
	}()

	slice := []int{1, 2, 3} // //
	fmt.Println(slice[4])
}
*/

/* ТЕОРИЯ 04.10.2026 Горутины и каналы
package main

import (
	"fmt"
	"time"
)

// Функция, описывающая поход рабочего в шахту
// Принимает:
// 1. Номер похода в шахту
// Возвращает:
// 1. Добытый уголь
func mine(transferPoint chan int, n int) {
	fmt.Println("Поход в шахту номер", n, "начался...")
	time.Sleep(1 * time.Second)
	fmt.Println("Поход в шахту номер", n, "закончился")

	transferPoint <- 10
	fmt.Println("Поход номер", n, "уголь передал!")
}

func main() {
	// coal - перевод с английского это `Уголь`
	coal := 0

	// Через функцию make создаются каналы можно создать слайс, мапы, канал
	// chan - с английского это `Канал`
	transferPoint := make(chan int, 3) // если убрать 3 то не будет буфера

	// Запомнил время которое было до похода в шахту
	initTime := time.Now()

	go mine(transferPoint, 1)
	go mine(transferPoint, 2)
	go mine(transferPoint, 3)

	//       получение значения из нашего канала
	time.Sleep(3 * time.Second)
	coal += <-transferPoint

	time.Sleep(3 * time.Second)
	coal += <-transferPoint

	time.Sleep(3 * time.Second)
	coal += <-transferPoint

	fmt.Println("Добыли", coal, "угля!")
	fmt.Println("Прошло времени:", time.Since(initTime)) // Потом воспользовался time.Since Для того чтобы узнать сколько прошло времени
}

// Отдельно:
// Когда не буфинизированный канал то рабочему нужно будет ждать пока заберут его 10 углей
// А с буфенизированный каналом не нужно ждать, рабочий может просто положить уголь в тележку и пойти обратно
*/

/* ПРАКТИКА 05.10.2026 Горутины и каналы
1.
package main

import (
	"fmt"
	"time"
)

func File(Files chan int, n int) {
	time.Sleep(1 * time.Second)
	fmt.Println("Получил файл под номером", n)

	switch n {
	case 1:
		Files <- 120
	case 2:
		Files <- 80
	case 3:
	    Files <- 150
	}
}

func main() {
	file := 0

	Files := make(chan int, 3)

	initTime := time.Now()

	go File(Files, 1)
	go File(Files, 2)
	go File(Files, 3)

	time.Sleep(1 * time.Second)
	file += <-Files

	time.Sleep(1 * time.Second)
	file += <-Files

	time.Sleep(1 * time.Second)
	file += <-Files

	fmt.Println("ОбЪём загруженных данных:", file, "мб")
	fmt.Println("Прошло времени:", time.Since(initTime))
}




2.
package main

import (
	"fmt"
	"time"
)

func internatShop(Shop chan int, n int) {
	time.Sleep(1 * time.Second)
	fmt.Println("Стоимость заказов:", n)

	switch n {
	case 1:
	    Shop <- 500
	case 2:
		Shop <- 1200
	case 3:
		Shop <- 750
	case 4:
		Shop <- 900
	}
}

func main() {
	zakaz := 0

	Shop := make(chan int, 4)

	initTime := time.Now()

	go internatShop(Shop, 1)
	go internatShop(Shop, 2)
	go internatShop(Shop, 3)
	go internatShop(Shop, 4)

	time.Sleep(1 * time.Second)
	zakaz += <- Shop

	time.Sleep(1 * time.Second)
	zakaz += <- Shop

	time.Sleep(1 * time.Second)
	zakaz += <- Shop

	time.Sleep(1 * time.Second)
	zakaz += <- Shop

	fmt.Println("Общая стоимость заказов:", zakaz)
	fmt.Println("Прошло времени:", time.Since(initTime))
}




3.
package main

import (
	"fmt"
	"time"
)

func Servercheck(Server chan int, n int) {
	time.Sleep(2 * time.Second)
	fmt.Println("Сервер:", n)

	switch n {
	case 1:
		Server <- 100
	case 2:
		Server <- 200
	case 3:
		Server <- 300
	}
}

func main() {
	threeserver := 0

	Server := make(chan int, 3)

	initTime := time.Now()

	go Servercheck(Server, 1)
	go Servercheck(Server, 2)
	go Servercheck(Server, 3)

	time.Sleep(1 * time.Second)
	threeserver += <- Server

	time.Sleep(1 * time.Second)
	threeserver += <- Server

	time.Sleep(1 * time.Second)
	threeserver += <- Server

	fmt.Println("Общий результат:", threeserver)
	fmt.Println("Прошло времени:", time.Since(initTime))
}




4.
package main

import (
	"fmt"
	"time"
)

func pointsCheck(PlayersPoints chan int, n int) {
	fmt.Println("Количество очков:", n)

	switch n {
	case 1:
	   PlayersPoints <- 15
	case 2:
		PlayersPoints <- 30
	case 3:
		PlayersPoints <- 25
	case 4:
		PlayersPoints <- 40
	case 5:
		PlayersPoints <- 10
	}
}

func main() {
	gamepoint := 0

	Pointes := make(chan int, 5)

	initTime := time.Now()

	go pointsCheck(Pointes, 1)
	go pointsCheck(Pointes, 2)
	go pointsCheck(Pointes, 3)
	go pointsCheck(Pointes, 4)
	go pointsCheck(Pointes, 5)

	time.Sleep(1 * time.Second)
	gamepoint += <- Pointes

	time.Sleep(1 * time.Second)
	gamepoint += <- Pointes

	time.Sleep(1 * time.Second)
	gamepoint += <- Pointes

	time.Sleep(1 * time.Second)
	gamepoint += <- Pointes

	time.Sleep(1 * time.Second)
	gamepoint += <- Pointes

	fmt.Println("Общее количество очков:", gamepoint)
	fmt.Println("Прошло времени:", time.Since(initTime))
}




5.
package main

import (
	"fmt"
	"time"
)

func Statistica(Static chan int, n int) {
	fmt.Println("Результат измерений:", n)

	switch n {
	case 1:
		Static <- 45
	case 2:
		Static <- 70
	case 3:
	    Static <- 35
	case 4:
		Static <- 50
	}
}

func main() {
	threestatic := 0

	firstStatic := make(chan int, 4)

	initTime := time.Now()

	go Statistica(firstStatic, 1)
	go Statistica(firstStatic, 2)
	go Statistica(firstStatic, 3)
	go Statistica(firstStatic, 4)

	time.Sleep(1 * time.Second)
	threestatic += <- firstStatic

	time.Sleep(1 * time.Second)
	threestatic += <- firstStatic

	time.Sleep(1 * time.Second)
	threestatic += <- firstStatic

	time.Sleep(1 * time.Second)
	threestatic += <- firstStatic

	average := threestatic / 4

	fmt.Println("Общая сумма:", threestatic)
	fmt.Println("Среднее значение:", average)
	fmt.Println("Прошло времени:", time.Since(initTime))
}



6.
package main

import (
	"fmt"
	"time"
)

func Competition(Fight chan int, n int) {
	fmt.Println("Время участников:", n)

	switch n {
	case 1:
	    Fight <- 12
	case 2:
		Fight <- 9
	case 3:
		Fight <- 15
	case 4:
		Fight <- 11
	case 5:
		Fight <- 10
	}
}

func main() {
	General := 0
	best := 0
	worst := 0

	fifth := make(chan int, 5)

	initTime := time.Now()

	go Competition(fifth, 1)
	go Competition(fifth, 2)
	go Competition(fifth, 3)
	go Competition(fifth, 4)
	go Competition(fifth, 5)

	time.Sleep(1 * time.Second)
	General += <- fifth

	time.Sleep(1 * time.Second)
	General += <- fifth

	time.Sleep(1 * time.Second)
	General += <- fifth

	time.Sleep(1 * time.Second)
	General += <- fifth

	time.Sleep(1 * time.Second)
	General += <- fifth


	for i := 0; i < 5; i++ {
		result := <-fifth

		General += result

		if i == 0 {
			best = result
			worst = result
		}

		if result < best {
			best = result
		}

		if result > worst {
			worst = result
		}
	}

	average := float64(General) / 5

	fmt.Println("Лучшее время:", best)
	fmt.Println("Худшее время:", worst)
	fmt.Println("Среднее время:", average)
	fmt.Println("Прошло времени:", time.Since(initTime))

}
*/

/* ТОЕРИЯ 06.10.2026 Анонимные функции в горутинах
package main

import (
	"fmt"
	"time"
)

func foo() {
	for {
		fmt.Println("Hello")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go foo()

	// ЛУЧШЕ ИСПОЛЬЗОВАТЬ ЭТО ЧЕМ В FUNC FOO()
	go func () {
	// for без ничего это бесконечный цикл
		for {
		fmt.Println("Anon")
		time.Sleep(100 * time.Millisecond)
	}
	}()

	time.Sleep(1 * time.Second)
}
*/

/* ТЕОРИЯ 06.10.2026 Конструкция select {}
1.
package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	intCh := make(chan int)
	strCh := make(chan string)

	go func() {
		i := 1
		for {
			intCh <- 1
			i++

			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		i := 1
		for {
			strCh <- "hi" + strconv.Itoa(i)
			i++

			time.Sleep(100 * time.Millisecond)
		}
	}()


	time.Sleep(50 * time.Millisecond)

	select {
	case number := <-intCh:
		fmt.Println("intCh:", number)
	case str := <-strCh:
			fmt.Println("strCh:", str)
	default:
		fmt.Println("Никакой канал не готов")
	}


	// for без ничего это бесконечный цикл
	for {
		// Лучший вариант
		select {
		case number := <-intCh:
			fmt.Println("intCh:", number)
		case str := <-strCh:
			fmt.Println("strCh:", str)
		default:
		    fmt.Println("Никакой канал не готов")
		}


		// не лучший вариант
		number := <-intCh
		fmt.Println("intCh:", number)

		str := <- strCh
		fmt.Println("strCh:", str)
		//
	}
	//
}



2.
package main

import (
	"fmt"
	"time"
)

type Message struct {
	Author string
	Text string
}

func main() {
	messageChan1 := make(chan Message)
	messageChan2 := make(chan Message)

	go func() {
		for {
			messageChan1 <- Message{
				Author: "Друг 1",
				Text: "Привет",
			}

			time.Sleep(10 * time.Second)
		}
	}()


	go func() {
		fmt.Println("Время отправки сообщения вторым другом:", time.Now())
		for {
			messageChan2 <- Message{
				Author: "Друг 2",
				Text: "Как дела?",
			}

			time.Sleep(100 * time.Millisecond)
		}
	}()

for {
	select {
	case msg1 := <-messageChan1:
		fmt.Println("Я получил сообщение от:", msg1.Author, "текст сообещния:", msg1.Text)
	case msg2 := <-messageChan2:
		fmt.Println("Я получил сообщение от:", msg2.Author, "текст сообщения:", msg2.Text)
	}
}
}
*/

/* ПРАКТИКА 06.10.2026 Конструкция select {}
1.
package main

import (
	"fmt"
)

type Message struct {
	Author string
}

func main() {
	message1 := make(chan Message)
	message2 := make(chan Message)

	go func() {
		for i := 1; i <= 2; i++{
			message1 <- Message{
				Author: "Привет",
			}
		}
	}()


	go func() {
		for i := 1; i <= 2; i++{
			message2 <- Message{
				Author: "GO",
			}
		}
	}()

	for i := 1; i <= 2; i++{
		select {
		case msg1 := <- message1:
			fmt.Println("Я получил сообщение от:", msg1.Author)
		case msg2 := <- message2:
			fmt.Println("Я получил сообщение от:", msg2.Author)
		}
	}
}
x



2.
package main

import (
	"fmt"
	"time"
)

type Fight struct {
	Number int
}

func main() {
	messageOne := make(chan Fight)
	messageFirst := make(chan Fight)

	go func() {
		time.Sleep(2 * time.Second)
		for i := 1; i <= 1; i++ {
			messageOne <- Fight{
				Number: 100,
			}

		}
	}()

	go func() {
		time.Sleep(1 * time.Second)
		for i := 1; i <= 1; i++{
			messageFirst <- Fight{
				Number: 200,
			}

		}
	}()


	for i := 1; i <= 1; i++{
		select {
		case fgt1 := <- messageOne:
			fmt.Println("Первый получившийся результат:", fgt1.Number)
		case fgt2 := <- messageFirst:
			fmt.Println("Получившийся результат:", fgt2.Number)
		}
	}
}




3.
package main

import "fmt"

func main() {
	// Создаем три канала
	chan1 := make(chan int)
	chan2 := make(chan int)
	chan3 := make(chan int)

	// Запускаем три горутины
	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- 20
	}()

	go func() {
		chan3 <- 30
	}()

	sum := 0

	for i := 1; i <= 3; i++ {
		select {
		case x := <- chan1:
			fmt.Println("Получено из первого канала:", x)
			sum += x
		case x := <- chan2:
			fmt.Println("Получено из второго канала:", x)
			sum += x
		case x := <- chan3:
			fmt.Println("Получено из третьего канала:", x)
			sum += x
		}
	}

	fmt.Println("Итоговая сумма:", sum)
}




4.
package main

import (
	"fmt"
	"time"
)

func main() {
	// Создаем канал для получения числа
	chan1 := make(chan int)

	// Запускаем горутину
	go func() {
		// Имитируем долгую работу в 3 секунды
		time.Sleep(3 * time.Second)
		chan1 <- 100
	}()

	// Ожидаем результат с таймаутом в 2 секунды
	select {
	case x := <- chan1:
		// Если данные пришли раньше таймаута
		fmt.Println("Результат:", x)
	case <- time.After(2 * time.Second):
		// Если прошло 2 секунды, а из ch ничего не пришло
		fmt.Println("Время ожидания истекло")
	}
}




5.
package main

import (
	"fmt"
	"time"
)

func main() {
	server1 := make(chan string)
	server2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		server1 <- "Сервер 1 ответил первым"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		server2 <- "Сервер 2 ответил первым"
	}()

	select {
	case x := <- server1:
		fmt.Println(x)
	case x := <- server2:
		fmt.Println(x)
	}
}




6.
package main

import (
	"fmt"
	"time"
)

func main() {
	task1 := make(chan string)
	task2 := make(chan string)
	task3 := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		task1 <- "Задача 1 завершится через 3 секунды"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		task2 <- "Задача 2 завершится через 1 секунду"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		task3 <- "Задача 3 завершится через 2 секунды"
	}()

for i := 1; i <= 3; i++ {
	select {
	case x := <- task1:
		fmt.Println(x)
	case x := <- task2:
		fmt.Println(x)
	case x := <- task3:
		fmt.Println(x)
	}
}

	fmt.Println("Все задачи завершены")
}
*/

/* ТЕОРИЯ 06.10.2026 Закрытие каналов. Аксиомы каналов.
// Аксиома - Положение, принимаемое без доказательств.
// Есть закрытый канал - closed channel и nil - nil channel
// РАЗБОР CLOSED CHANNEL
package main

import "fmt"

func main() {
	// var ch chan int // Это nil Канал, не создаем через make

	// создаём новый открыты канал
	ch := make(chan int)

	// закрываем канал
	// 1. канал можно закрыть только один раз
	// 2. если у нас закрыты канал и я попробую оттуда прочитать , я получу значение по умолчанию
	// 3. Писать в закрытый канал я не могу
	close(ch) // Таким способом можно закрыть каналы

	// я записываю значение в закрытый канал
	ch <- 10

	// читаю значение из закрытого канала
	// если вывести ok и там будет false то это значит что были значения по умолчанию
	v1, ok1 := <- ch
	v2, ok2 := <- ch
	v3, ok3 := <- ch

	fmt.Println("v:", v1, v2, v3)
	fmt.Println("ok:", ok1, ok2, ok3)
	//
}

2.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	transferPoint := make(chan int)

	// miner
	go func() {
	// 3 4 5 6
	iterations := 3 + rand.Intn(4) // Intn принимает в себя какой то число(аргумент) и возвращает  рандомное число от 0 вкулючительно
	fmt.Println("iterations:", iterations)

	for i := 1; i <= iterations; i++ {
		transferPoint <- 10
		time.Sleep(300 * time.Millisecond)
	}

	close(transferPoint)
	}()

	coal := 0
	//for {
	//	v, ok := <-transferPoint // передаем уголь из transferPoint в coal
	//	if !ok {
	//		fmt.Println("Шахтёр обработал всю шахту!")
	//		break
	//	}

	//	coal += v

	//	fmt.Println("coal:", coal)
	//}


	// КОРОТКАЯ ВЕРСИЯ ЧЕМ ВЫШЕ
	for v := range transferPoint {
		coal += v
		fmt.Println("coal:", coal)
	}
	fmt.Println("Суммарно добытый уголь:", coal)
}



// НИКОГДА НЕ ИСПОЛЬЗОВАТЬ nil КАНАЛЫ
// РАЗБОР NIL CHANNEL
package main

import "fmt"

func main() {
	// не nil канал
	// ch := make(chan string)

	// nil канал
	var ch chan string

	// ГОРУТИНА
	go func() {
		ch <- "hello"
	}()

	v := <-ch

	fmt.Println("v:", v)
}
*/

/* ПРАКТИКА 06.10.2026 Закрытие каналов. Аксиомы каналов.
1.
package main

import "fmt"

func main() {
	chisla := make(chan int)

	go func() {
		chisla <- 10
		chisla <- 20
		chisla <- 30
		chisla <- 40
		chisla <- 50

		close(chisla)
	}()

	for i := 1; i <= 5; i++ {
		x := <- chisla
		fmt.Println(x)
	}

	fmt.Println("Канал закрыт")
}




2.
package main

import "fmt"

func main() {
	CheckClose := make(chan int)

	go func() {
		CheckClose <- 100

		close(CheckClose)
	}()

	for i := 1; i <= 1; i++ {
		x := <- CheckClose

		fmt.Println("Получено:", x)
	}

	_, ok := <- CheckClose
	if !ok {
		fmt.Println("Канал закрыт")
	}
}





3.
package main

import "fmt"

func main() {
	ReadsClose := make(chan int)

	go func() {
		ReadsClose <- 5
		ReadsClose <- 10
		ReadsClose <- 15

		close(ReadsClose)
	}()

	for i := 1; i <= 3; i++ {
		x := <- ReadsClose

		fmt.Println("Получено:", x)
	}

	_, ok := <- ReadsClose
	if !ok {
		fmt.Println("Данные закончились")
	}
}





4.
package main

import "fmt"

func SendClose(numbers chan int) {
	numbers <- 1
	numbers <- 2
	numbers <- 3
	numbers <- 4
	numbers <- 5

	close(numbers)
}

func main() {
	ch := make(chan int)

	go SendClose(ch)

	for i := 1; i <= 5; i++ {
		x := <- ch
		fmt.Println("Получено:", x)
	}

	fmt.Println("Получение завершено")
}



5.
package main

import "fmt"

func foo(number chan int) {
	number <- 10
	number <- 20
	number <- 30

	close(number)
}

func main() {
	ch := make(chan int)

	go foo(ch)

	for i := 1; i <= 3; i++ {
		x := <- ch

		fmt.Println("Получено:", x)
	}

	_, ok := <- ch
	if !ok {
		fmt.Println("Канал полностью закрыт")
	}
}




6.
package main

import "fmt"

func Aksioma(number chan int) {
	number <- 5
	number <- 10
	number <- 15
	number <- 20

	close(number)
}

func main() {
	ch := make(chan int)

	go Aksioma(ch)

	for i := 1; i <= 4; i++ {
		x := <- ch

		fmt.Println("Получено:", x)
	}

	_, ok := <- ch
	if !ok {
		fmt.Println("Канал закрыт, данных больше нет")
	}
}
*/

/* ТЕОРИЯ 07.10.2026 Контекст
package main

import (
	"context"
	"fmt"
	"time"
)

// функция foo будет завязана на родительский контекст
func foo(ctx context.Context, n int) {
	for {
		select{
		case <-ctx.Done():
			fmt.Println("Foo завершилась :(", n)
			return
		default:
			fmt.Println("Foo продолжает своё выполнение!", n)
		}



		time.Sleep(100 * time.Millisecond)
	}
}

// функция boo будет завязана на дочерном контексте
func boo(ctx context.Context, n int) {
	for {
		select{
		case <-ctx.Done():
			fmt.Println("Boo завершилась :(", n)
			return
		default:
			fmt.Println("Boo продолжает своё выполнение!", n)
		}



		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// родительский конекст
	//             parentCancel дает возможность отменить родительский контекст
	parentContext, parentCancel := context.WithCancel(context.Background())
	childContext, childCancel := context.WithCancel(parentContext)


	go foo(parentContext, 1)
	go foo(parentContext, 2)
	go foo(parentContext, 3)

	go boo(childContext, 1)
	go boo(childContext, 2)
	go boo(childContext, 3)

	time.Sleep(1 * time.Second)
	// благодаря функции childCancel() я отменяю дочерный контекст
	childCancel()

	time.Sleep(1 * time.Second)
	// через ещё одну секунду я отменю родительский контекст
	parentCancel()

	time.Sleep(3 * time.Second)
	fmt.Println("main завершился :(")
}
*/

/* ПРАКТИКА 07.10.2026 Контекст
1.
package main

import (
	"fmt"
	"time"
	"context"
)

func worker(ctx context.Context) {
	for i := 1; i <= 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("Работа отменена")
			return
		default:
			fmt.Println("Выполняется шаг:", i)
			time.Sleep(500 * time.Millisecond)
		}
	}
	fmt.Println("Работа завершена")
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() ставит дополнительную отмену при завершении main

	go worker(ctx)

	time.Sleep(2200 * time.Millisecond)
	cancel()

	time.Sleep(500 * time.Millisecond)
}





2.
package main

import (
	"fmt"
	"time"
	"context"
)

func worker(ctx context.Context) {
	for i := 1; i <= 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("Время вышло")
			return
		default:
			fmt.Println("Выполняется шаг:", i)
			time.Sleep(500 * time.Millisecond)
		}
	}
	fmt.Println("Работа завершена")
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()

	go worker(ctx)

	time.Sleep(2 * time.Second)
	time.Sleep(500 * time.Millisecond)
}




3.
package main

import (
	"time"
	"fmt"
	"context"
)

func worker(ctx context.Context, numbers chan int) {
	defer close(numbers)

	for i := 1; i <= 10; i++ {
		select {
		case <-ctx.Done():
			return
		case numbers <- i:
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()

	numbers := make(chan int)

	go worker(ctx, numbers)

	for number := range numbers {
		fmt.Println("Получено:", number)
	}

	fmt.Println("Работа остановлена")
}




4.
package main

import (
	"fmt"
	"time"
	"context"
	"sync"
)

func worker(ctx context.Context, numbers chan string, id int, wg *sync.WaitGroup) {
	defer wg.Done() // сообщает WaitGroup, что этот worker завершил работу

	for i := 1; i <= 10; i++ {
		select {
		case <-ctx.Done():
			return // если контекст завершён — worker останавливается

		case numbers <- fmt.Sprintf("Worker %d: шаг %d", id, i):// Sprintf создаёт строку и подставляет в неё id и номер шага

			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second) // создаём контекст, который автоматически завершится через 2 секунды
	defer cancel() // освобождаем ресурсы контекста после завершения main

	numbers := make(chan string) // создаём канал для передачи строк от workers в main

	var wg sync.WaitGroup // создаём WaitGroup для отслеживания workers

	wg.Add(2) // говорим WaitGroup, что должны завершиться 2 worker

	go worker(ctx, numbers, 1, &wg) // запускаем первого worker

	go worker(ctx, numbers, 2, &wg) // запускаем второго worker

	go func() {
		wg.Wait() // ждём, пока оба worker завершат работу

		close(numbers) // после завершения обоих workers закрываем канал
	}()

	for message := range numbers {
		fmt.Println(message) // получаем сообщения из канала и выводим их
	}

	fmt.Println("Все работники остановлены") // выполняется после закрытия канала
}




5.
package main

import (
	"fmt"
	"time"
	"context"
)

func serverRequest(ctx  context.Context) {
	for i := 1; i <= 5; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("Запрос отменён")
			return

		default:
			fmt.Println("Обработка запроса:", i)
			time.Sleep(1 * time.Second)
		}
	}

	fmt.Println("Запрос выполнен")
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel() // освобождаем ресурсы контекста

	go serverRequest(ctx) // запускаем запрос отдельно

	time.Sleep(4 * time.Second) // даём worker время завершиться
}





6.
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, message chan string, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		select {
		case <- ctx.Done():
			return

		case message <- fmt.Sprintf("Worker %d: шаг %d", id, i):
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()

	messages := make(chan string)

	var wg sync.WaitGroup
	wg.Add(3)

	go worker(ctx, messages, 1, &wg)
	go worker(ctx, messages, 2, &wg)
	go worker(ctx, messages, 3, &wg)

	go func() {
		wg.Wait()
		close(messages)
	}()

	for messages := range messages {
		fmt.Println(messages)
	}

	fmt.Println("Все workers завершены")
}
*/

/* ТЕОРИЯ 09.10.2026 WaitGroup
package main

import (
	"fmt"
	"sync"
	"time"
)

// С WaitGroup мы рабоатет через указатели - *, &
func postman(wg *sync.WaitGroup, text string) {
	defer wg.Done() // с английского Done будет готово

	for i := 1; i <= 3; i ++ {
		fmt.Println("Я почтальон, я отнёс газету", text, "в", i, "раз")
		time.Sleep(250 * time.Millisecond)
	}
}

// не использую каналы потомучто я ничего не возвращаю, почтальон просто приносит и ему ничего не нужно возвращать
func main() {
	wg := &sync.WaitGroup{} // WaitGroup с английского это будет Ожидание Группы горутин

	wg.Add(1) // Говорим что щас будет запущена гарутина
	go postman(wg, "Новости")

	wg.Add(1)
	go postman(wg, "Игровой журнал")

	wg.Add(1)
	go postman(wg, "Автомобильная хроника")

	wg.Wait() // Заблокируемся в main пока не будет 0

	fmt.Println("main завершился!")
}
*/

/* ПРАКТИКА 10.10.2026 WaitGroup
1.
package main

import (
	"fmt"
	"time"
	"sync"
)

func worker(wg *sync.WaitGroup, id int) {
	defer wg.Done()

	fmt.Printf("Работник № %d начал работу\n", id)

	time.Sleep(2 * time.Second)

	fmt.Printf("Работник № %d закончил работу\n", id)
}

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go worker(wg, 1)

	wg.Add(1)
	go worker(wg, 2)

	wg.Wait()

	fmt.Println("Все работники завершили работу")
}




2.
package main

import (
	"fmt"
	"time"
	"sync"
)

func worker(wg *sync.WaitGroup, id int) {
	defer wg.Done()

	fmt.Printf("Обработка числа: %d\n", id)

	time.Sleep(1 * time.Second)

	fmt.Printf("Обработка завершена: %d\n", id)
}

func main() {
	wg := &sync.WaitGroup{}

	numbers := []int{10, 20, 30}

	for _, id := range numbers {
		wg.Add(1)
		go worker(wg, id)
	}

	wg.Wait()

	fmt.Println("Все вычисления завершены")
}




3.
package main

import (
	"fmt"
	"time"
	"sync"
)

func worker(wg *sync.WaitGroup, id int) {
	defer wg.Done()

	fmt.Printf("Загрузка № %d началась\n", id)

	time.Sleep(time.Duration(id) * time.Second)

	fmt.Printf("Загрузка № %d завершена\n", id)
}

func main() {
	wg := &sync.WaitGroup{}

	numbers := []int{1, 2, 3, 4}

	for _, id := range numbers {
		wg.Add(1)
		go worker(wg, id)
	}

	wg.Wait()

	fmt.Println("Все загрузки завершены")
}




4.
package main

import (
	"fmt"
	"time"
	"sync"
)

func worker(wg *sync.WaitGroup, id int) {
	defer wg.Done()

	fmt.Printf("Обработано число: %d\n", id)

	time.Sleep(1 * time.Second)
}

func main() {
	wg := &sync.WaitGroup{}
	defer wg.Wait()

	numbers := []int{2, 4, 6, 8, 10}

	for _, id := range numbers {
		wg.Add(1)
		go worker(wg, id)
	}

	wg.Wait()

	fmt.Println("Все числа обработаны")
}




5.
package main

import (
	"fmt"
	"time"
	"sync"
)

func worker(wg *sync.WaitGroup, id int) {
	defer wg.Done()

	fmt.Printf("Работник № %d приступил\n", id)

	time.Sleep(time.Duration(4-id) * time.Second)

	fmt.Printf("Работник № %d завершил\n", id)
}

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go worker(wg, 1)

	wg.Add(1)
	go worker(wg, 2)

	wg.Add(1)
	go worker(wg, 3)

	wg.Wait()

	fmt.Println("Работа завершена")
}




6.
package main

import (
	"fmt"
	"time"
	"sync"
)

func worker(wg *sync.WaitGroup, id int) {
	defer wg.Done()

	fmt.Printf("Задача № %d выполняется\n", id)

	time.Sleep(time.Duration(id) * 500 * time.Millisecond)

	fmt.Printf("Задача № %d выполнена\n", id)
}

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go worker(wg, 1)

	wg.Add(1)
	go worker(wg, 2)

	wg.Add(1)
	go worker(wg, 3)

	wg.Add(1)
	go worker(wg, 4)

	wg.Wait()

	fmt.Println("Все задачи выполнены")
}




// ТОЕРИЯ 10.10.2026 Состояние гонки. Атомики. Мьютексы.
package main

import (
	"fmt"
	"sync"
	//"sync/atomic"
)

// 10 000
// var number int = 0

// atomic изначально имеет значение 0
//var number atomic.Int64

var slice []int

var mtx sync.Mutex

func increase(wg *sync.WaitGroup) {
	defer wg.Done() // делает зевершение

	for i := 1; i <= 1000; i++ {
		mtx.Lock() // Берет число
		slice = append(slice, i) // Закидывает число в Lock
		mtx.Unlock() // Освобождает Lock что-бы можно было сново положить число в Lock и так до бесконечности
	}
}

// Если у нас где то состяоние гонки то значит я написал код неверно
func main() {
	wg := &sync.WaitGroup{}
	wg.Add(10)

	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)

	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)

	wg.Wait()

	mtx.Lock()
	fmt.Println("slice len:", len(slice))
	mtx.Unlock()
}
*/

/* ПРАКТИКА 10.10.2026 Состояние гонки. Атомики. Мьютексы.
1.
package main

import (
	"fmt"
	"sync"
)

// переменная
var counter int

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 100000; i++ {
		counter++
	}
}

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(5)

	go worker(wg)
	go worker(wg)
	go worker(wg)
	go worker(wg)
	go worker(wg)

	wg.Wait()

	fmt.Println(counter)
}





2.
package main

import (
	"fmt"
	"sync"
)

// переменная
var balance int = 100

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 1000; i++ {
		balance++
	}
}

func workerfirst(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1000; i > 0; i-- {
		balance--
	}
}

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go worker(wg)

	wg.Add(1)
	go workerfirst(wg)

	wg.Wait()

	fmt.Println("balance:", balance)
}





3.
package main

import (
	"fmt"
	"sync"
)

var counter int

func work(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 1000; i ++ {
		counter++
	}
}

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go work(wg)

	wg.Add(1)
	go work(wg)

	wg.Add(1)
	go work(wg)

	wg.Add(1)
	go work(wg)

	wg.Add(1)
	go work(wg)

	wg.Add(1)
	go work(wg)

	wg.Add(1)
	go work(wg)

	wg.Add(1)
	go work(wg)

	wg.Add(1)
	go work(wg)

	wg.Add(1)
	go work(wg)

	wg.Wait()

	fmt.Println(counter)
}





4.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter atomic.Int64

func work(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 1000; i++ {
		counter.Add(1)
	}
}

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go work(wg)

	wg.Add(1)
	go work(wg)

	wg.Add(1)
	go work(wg)

	wg.Add(1)
	go work(wg)

	wg.Add(1)
	go work(wg)

	wg.Add(1)
	go work(wg)

	wg.Add(1)
	go work(wg)

	wg.Add(1)
	go work(wg)

	wg.Add(1)
	go work(wg)

	wg.Add(1)
	go work(wg)

	wg.Wait()

	fmt.Println("counter:", counter.Load())
}





5.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var check atomic.Int64

func work(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 5000; i++ {
		check.Add(1)
	}
}

func main() {
	wg := &sync.WaitGroup{}

	for i := 1; i <= 8; i++ {
		wg.Add(1)
		go work(wg)
	}

	wg.Wait()

	fmt.Println("check:", check.Load())
}





6.
package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var country atomic.Bool

func work() {
	time.Sleep(2 * time.Second)

	country.Store(true)
}

func main() {
	go work()

	for !country.Load() {
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("Данные готовы!")
}





7.
package main

import (
	"fmt"
	"sync"
)

var counter int
var score sync.Mutex

func point(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 1000; i++ {
		score.Lock()
		counter++
		score.Unlock()
	}
}

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go point(wg)

	wg.Add(1)
	go point(wg)

	wg.Add(1)
	go point(wg)

	wg.Add(1)
	go point(wg)

	wg.Add(1)
	go point(wg)

	wg.Add(1)
	go point(wg)

	wg.Add(1)
	go point(wg)

	wg.Add(1)
	go point(wg)

	wg.Add(1)
	go point(wg)

	wg.Add(1)
	go point(wg)

	wg.Wait()

	score.Lock()
	fmt.Println("counter:", counter)
	score.Unlock()
}





8.
package main

import (
	"fmt"
	"sync"
)

var bank int = 1000
var score sync.Mutex

func work(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		score.Lock()
		bank += 100
		score.Unlock()
	}
}

func job(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 10; i > 0; i-- {
		score.Lock()
		bank -= 100
		score.Unlock()
	}
}

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go work(wg)

	wg.Add(1)
	go work(wg)

	wg.Add(1)
	go job(wg)

	wg.Add(1)
	go job(wg)

	wg.Wait()

	score.Lock()
	fmt.Println("bank:", bank)
	score.Unlock()
}





9.
package main

import (
	"fmt"
	"sync"
)

var users = make(map[string]int)
var score sync.Mutex

func addUser(wg *sync.WaitGroup, group int) {
	defer wg.Done()

	for i := 1; i <= 100; i++ {
		name := fmt.Sprintf("user_%d_%d", group, i)
		age := 18 + i%10

		score.Lock()
		users[name] = age
		score.Unlock()
	}
}

func main() {
	wg := &sync.WaitGroup{}

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go addUser(wg, i)
	}

	wg.Wait()

	fmt.Println("Всего пользователей:", len(users))
}




10.
package main

import (
	"fmt"
	"sync"
)

type Stats struct {
	totalVisits int
	pages map[string]int
	mutex sync.Mutex
}

func (s *Stats) Visit(page string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.totalVisits++
	s.pages[page]++
}

func worker(wg *sync.WaitGroup, stats *Stats, id int) {
	defer wg.Done()

	pages := []string{"home", "about", "contacts"}

	for i := 0; i < 1000; i++ {
		page := pages[(i+id)%len(pages)]
		stats.Visit(page)
	}
}

func main() {
	stats := Stats{
		pages: make(map[string]int),
	}

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(&wg, &stats, i)
	}

	wg.Wait()

	stats.mutex.Lock()
	fmt.Println("Всего посещений:", stats.totalVisits)
	fmt.Println("Посещения страниц:")

	for page, visits := range stats.pages {
		fmt.Printf("%s: %d\n", page, visits)
	}
	stats.mutex.Unlock()
}






// ТЕОРИЯ 10.10.2026 RWMutex
package main

import (
	"fmt"
	"sync"
	"time"
)

var likes int = 0
var mtx sync.RWMutex

func setLike(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 100_000; i++ {
		mtx.Lock()
		likes++
		mtx.Unlock()
	}
}

func getLike(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 100_000; i++ {
		mtx.RLock()
		_ = likes
		mtx.RUnlock()
	}
}

func main() {
	wg := &sync.WaitGroup{}

	initTime := time.Now()

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go setLike(wg)
	}

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go getLike(wg)
	}

	wg.Wait() // жду когда итоговое значение станет равно нулю

	fmt.Println("Время выполнение:", time.Since(initTime))
}
*/

/* ПРАКТИКА 11.10.2026 RWMutex
// sync.RWMutex позволяет разделить блокировки на два типа:
// - RLock() / RUnlock() — для чтения. Несколько горутин могут читать данные одновременно.
// - Lock() / Unlock() — для изменения данных. В этот момент другие читатели и писатели не могут получить блокировку.
package main

import (
	"fmt"
	"sync"
)

Общая переменная score изначально равна 100.
var score int = 100
var more sync.RWMutex

func point(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 1000; i++ {
		more.RLock()
		value := score
		more.RUnlock()

		fmt.Println("value:", value)
	}
}

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go point(wg)

	wg.Add(1)
	go point(wg)

	wg.Add(1)
	go point(wg)

	wg.Add(1)
	go point(wg)

	wg.Add(1)
	go point(wg)

	wg.Wait()

	fmt.Println("score:", score)
}





2.
package main

import (
	"fmt"
	"sync"
)

var score int
var strong sync.RWMutex

func work(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 1000; i++ {
		strong.Lock()
		score++
		strong.Unlock()
	}
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 1000; i++ {
		strong.RLock()
		value := score
		strong.RUnlock()

		fmt.Println("value:", value)
	}
}

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(7)

	for i := 0; i < 2; i++ {
		go work(wg)
	}

	for i := 0; i < 5; i++ {
		go increment(wg)
	}

	wg.Wait()

	strong.RLock()
	fmt.Println("score:", score)
	strong.RUnlock()
}




3.
package main

import (
	"fmt"
	"sync"
)

var prices = map[string]int {
	"phone": 1000,
	"laptop": 2000,
	"tablet": 1500,
}

var more sync.RWMutex

func increment(wg *sync.WaitGroup, product string) {
	defer wg.Done()

	for i := 1; i <= 1000; i++ {
		more.RLock()
		value := prices[product]
		more.RUnlock()

		_ = value
	}
}

func price(wg *sync.WaitGroup, product string) {
	defer wg.Done()

	for i := 1; i <= 1000; i++ {
		more.Lock()
		prices[product] += 1
		more.Unlock()
	}
}

func main() {
	wg := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go increment(wg, "phone")
	}

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go price(wg, "phone")
	}

	wg.Wait()

	more.RLock()
	fmt.Println("prices:", prices)
	more.RUnlock()
}





4.
package main

import (
	"fmt"
	"sync"
)

var carta = map[string]int {
	"Iqoo": 3,
	"MSI": 13,
	"Iphone": 27,
}

var copins sync.RWMutex

func work(wg *sync.WaitGroup, product string) {
	defer wg.Done()

	for i := 1; i <= 1000; i++ {
		copins.RLock()
		value := carta[product]
		copins.RUnlock()

		_ = value
	}
}

func worker(wg *sync.WaitGroup, product string) {
	defer wg.Done()
	for i := 1; i<= 1000; i++ {
		copins.Lock()
		carta[product] += 1
		copins.Unlock()
	}
}

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(7)
	for i := 0; i < 5; i++ {
		go work(wg, "Iqoo")
	}

	for i := 0; i < 2; i++ {
		go worker(wg, "MSI")
	}

	wg.Wait()

	copins.Lock()
	fmt.Println("carta:", carta)
	copins.Unlock()
}





5.
package main

import (
	"fmt"
	"sync"
)

var carta = map[string]string{
	"volume": "Громкость",
	"brightness": "Яркость",
	"language": "Язык",
}

var mtx sync.RWMutex

func work(wg *sync.WaitGroup, product string) {
	defer wg.Done()

	for i := 1; i <= 1000; i++ {
		mtx.RLock()
		value := carta[product]
		mtx.RUnlock()

		_ = value
	}
}

func worker(wg *sync.WaitGroup, product string) {
	defer wg.Done()

	for i := 1; i <= 1000; i++ {
		mtx.Lock()
		carta[product] += "Яркость "
		mtx.Unlock()
	}
}

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(7)
	for i := 0; i < 5; i++ {
		go work(wg, "language")
	}

	for i := 0; i < 2; i++ {
		go worker(wg, "brightness")
	}

	wg.Wait()

	mtx.RLock()
	fmt.Println("carta", carta)
	mtx.RUnlock()
}




6.
package main

import (
	"fmt"
	"sync"
)

var kash = map[string]int {
	"Арзу": 17,
	"Годжо": 28,
	"Айзен": 100,
}

var mtx sync.RWMutex

func work(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 1000; i++ {
		mtx.RLock()
		carta := kash["Арзу"]
		mtx.RUnlock()

		_ = carta
	}
}

func ages(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 1000; i++ {
		mtx.Lock()
		kash["Арзу"] += 1
		mtx.Unlock()
	}
}

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(7)

	for i := 0; i < 5; i++ {
		go work(wg)
	}

	for i := 0; i < 2; i++ {
		go ages(wg)
	}

	wg.Wait()

	mtx.Lock()
	fmt.Println("kash:", kash)
	mtx.Unlock()
}






7.
package main

import (
	"fmt"
	"sync"
)

type WebsiteStats struct {
	totalVisits int
	pageVisits  map[string]int
	mtx         sync.RWMutex
}

func (s *WebsiteStats) ReadStats(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		s.mtx.RLock()
		total := s.totalVisits
		pages := len(s.pageVisits)
		s.mtx.RUnlock()

		_, _ = total, pages
	}
}

func (s *WebsiteStats) RegisterVisit(wg *sync.WaitGroup, page string) {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		s.mtx.Lock()
		s.totalVisits++
		s.pageVisits[page]++
		s.mtx.Unlock()
	}
}

func main() {
	stats := &WebsiteStats{
		pageVisits: map[string]int{
			"/home":    0,
			"/about":   0,
			"/contact": 0,
		},
	}

	var wg sync.WaitGroup

	// 5 горутин читают статистику.
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go stats.ReadStats(&wg)
	}

	// 2 горутины регистрируют посещения.
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go stats.RegisterVisit(&wg, "/home")
	}

	wg.Wait()

	stats.mtx.RLock()
	fmt.Println("Total visits:", stats.totalVisits)
	fmt.Println("Visits by page:", stats.pageVisits)
	stats.mtx.RUnlock()
}
*/

/* ТЕОРИЯ 12.10.2026 Concurrency большой пример
package main

import (
	"context"
	"fmt"
	"study/miner"
	"study/postman"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var coal atomic.Int64

	mtx := sync.Mutex{}
	var mails []string

	minerContext, minerCancel := context.WithCancel(context.Background())
	postmanContext, postmanCancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("--->Рабочий день шахтёров окончен!")
		minerCancel()
	}()

	go func() {
		time.Sleep(6 * time.Second)
		fmt.Println("--->Рабочий день почтальона окончен!")
		postmanCancel()
	}()

	coalTransferPoint := miner.MinerPool(minerContext, 300)
	mailTransferPoint := postman.PostmanPool(postmanContext, 300)

	initTime := time.Now()

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		for v := range coalTransferPoint {
			coal.Add(int64(v))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for v := range mailTransferPoint {
			mtx.Lock()
			mails = append(mails, v)
			mtx.Unlock()
		}
	}()

	wg.Wait()

	fmt.Println("Суммарно добытый уголь:", coal.Load())

	mtx.Lock()
	fmt.Println("Суммарное количество полученных писем:", len(mails))
	mtx.Unlock()

	fmt.Println("Затраченное время:", time.Since(initTime))
}
*/

/* ПРАКТИКА 12.10.2026 Concurrency
1.
package main

import (
	"fmt"
	"sync"
)

func work(wg *sync.WaitGroup)  {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		fmt.Println("Привет из горутины 5 раз.")
	}
}

func main() {
	wg := &sync.WaitGroup{}

	fmt.Println("Программа запущена")

	wg.Add(1)
	go work(wg)

	fmt.Println("Главная функция работает")

	wg.Wait()
}





2.
package main

import (
	"fmt"
	"sync"
)

func number(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}

func alfvit(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 'A'; i <= 'E'; i++ {
		// %c чтобы выводить буквы
		fmt.Printf("%c\n", i)
	}
}

func main() {
	wg := &sync.WaitGroup{}

	fmt.Println("Обе задачи запущены")

	wg.Add(1)
	go number(wg)

	wg.Add(1)
	go alfvit(wg)

	wg.Wait()
}





3.
package main

import (
	"fmt"
	"sync"
)

func work(wg *sync.WaitGroup, far chan<- int) {
	defer wg.Done()

	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	far <- sum
}

func fifthinc(wg *sync.WaitGroup, far chan<- int) {
	defer wg.Done()

	sum := 0
	for i := 1; i <= 50; i++ {
		sum += i
	}
	far <-sum
}

func main() {
	wg := &sync.WaitGroup{}
	far := make(chan int, 2)

	wg.Add(1)
	go work(wg, far)

	wg.Add(1)
	go fifthinc(wg, far)

	wg.Wait()
	close(far)

	for sum := range far {
		fmt.Println("Сумма:", sum)
	}
}





4.
package main

import (
	"fmt"
	"sync"
)

func work(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		ch <- i // отправляет каждое число от 1 до 10.
	}

	close(ch) // close(ch) сообщает, что новых чисел больше не будет.
}

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(1)
	go work(wg, ch) // запускает горутину.

	// получает и выводит все числа.
	for num := range ch {
		fmt.Println(num)
	}

	wg.Wait()
}





5.
package main

import (
	"fmt"
	"sync"
)

func work(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		ch <- i
	}
}

func worker(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	for i := 6; i <= 10; i++ {
		fmt.Println(i)
		ch <- i
	}
}

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(1)
	go work(wg, ch)

	wg.Add(1)
	go worker(wg, ch)

	go func() {
		wg.Wait()
		close(ch)
	}()

	for num := range ch {
		fmt.Println("num:", num)
	}
}





6.
package main

import (
	"fmt"
)

func calculateFactorial(n int, ch chan int) {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	ch <- result // отправляет результат в канал.
}

func main() {
	number := 5
	ch := make(chan int)

	go calculateFactorial(number, ch) // запускает вычисление в отдельной горутине.

	factorial := <- ch // получает результат в главной функции.

	fmt.Println(number, factorial) // выводит число и его факториал.
}





7.
package main

import (
	"fmt"
	"sync"
)

func work(wg *sync.WaitGroup, ch  chan int) {
	defer wg.Done()

	for i := 10; i > 0; i-- {
		result := i * i
		fmt.Println(i)
		ch <- result
	}
}

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(1)
	go work(wg, ch)

	for i := 1; i <= 10; i++ {
		number := <- ch
		fmt.Println("number:", number)
	}
	wg.Wait()
}





8.
package main

import (
	"fmt"
	"sync"
)

func work(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	ch <- sum
}

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(1)
	go work(wg, ch)

	num := <- ch
	fmt.Println("Результат:", num)

	wg.Wait()
}





9.
package main

import (
	"fmt"
	"sync"
)

func work(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	sum := 0
	for i := 1; i <= 50; i++ {
		sum += i
	}
	ch <- sum
}

func worker(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	score := 0
	for i := 51; i <= 100; i++ {
		score += i
	}
	ch <- score
}

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(1)
	go work(wg, ch)

	wg.Add(1)
	go worker(wg, ch)

	num := <- ch
	fmt.Println("num:", num)

	point := <- ch
	fmt.Println("point:", point)

	wg.Wait()
}





10.
package main

import (
	"fmt"
	"sync"
)

func work(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(1)
	go work(wg, ch)

	for num := range ch {
		fmt.Println("num:", num)
	}

	wg.Wait()
}
*/

/* ТЕОРИЯ 14.10.2026 Первый HTTP сервер
package main

import (
	"fmt"
	"net/http"
)

func payHandler(w http.ResponseWriter, r *http.Request) {
	str := "Новый платёж обработан!!"
	b := []byte(str)

	// с помощью _ я сказал что не хочу использовать int и error у write
	_, err := w.Write(b)

	if err != nil {
		fmt.Println("ВО ремя записи HTTP ответа произошла ошибка:", err.Error())
	} else {
		fmt.Println("я корректно совершил оплату!")
	}
}

func cancelHandler(w http.ResponseWriter, r *http.Request) {
	str := "Оплата отменена!"
	b := []byte(str)

	// с помощью _ я сказал что не хочу использовать int и error у write
	_, err := w.Write(b)

	if err != nil {
		fmt.Println("ВО ремя записи HTTP ответа произошла ошибка:", err.Error())
	} else {
		fmt.Println("Я корректно отменил оплату!")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	str := "Hello, world!"
	b := []byte(str)

	// с помощью _ я сказал что не хочу использовать int и error у write
	_, err := w.Write(b)

	if err != nil {
		fmt.Println("ВО ремя записи HTTP ответа произошла ошибка:", err.Error())
	} else {
		fmt.Println("я корректно обработал HTTP запрос!")
	}
}

func main() {
	// HandleFunc - функция из пакета http
	http.HandleFunc("/default", handler)
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/cancel", cancelHandler)

	fmt.Println("Запускаю HTTP сервер!")
	// запускает HTTP сервер
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Произошла ошибка:", err.Error())
	}

	fmt.Println("Программа закончисла своё выполнение!")
}
*/

/* ПРАКТИКА 14.10.2026 Первый HTTP сервер
1.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func payInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Привет этой мой первый HTTP-сервер на GO.")
}

func main() {
	http.HandleFunc("/", payInfo)

	fmt.Println("Сервер запущен: http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}





2.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func look(w http.ResponseWriter, r *http.Request) {
	// Здесь w передаётся первым аргументом, чтобы текст отправился в HTTP-ответ и появился в браузере.
	fmt.Fprintln(w, "Главная страница.")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Обо мне.")
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Контакты.")
}

func main() {
	http.HandleFunc("/", look)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)

	fmt.Println("Сервер запщуен: http://localhost:8080/")
	fmt.Println("Сервер запущен: http://localhost:8080/about")
	fmt.Println("Сервер запущен: http://localhost:8080/contact")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}





3.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func port(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<!DOCTYPE html>")
	fmt.Fprintln(w, "<html lang=\"ru\">")
	fmt.Fprintln(w, "<head>")
	fmt.Fprintln(w, "<meta charset=\"UTF-8\">")
	fmt.Fprintln(w, "<title>Мой первый сервер</title>")
	fmt.Fprintln(w, "</head>")
	fmt.Fprintln(w, "<body>")
	fmt.Fprintln(w, "<h1>Добро пожаловать!</h1>")
	fmt.Fprintln(w, "<h2>Мой первый сервер на Go</h2>")
	fmt.Fprintln(w, "<p>Я учусь создавать веб-сайты.</p>")
	fmt.Fprintln(w, "</body>")
	fmt.Fprintln(w, "</html>")
}

func main() {
	http.HandleFunc("/", port)

	fmt.Println("Сервер запущен: http://localhost:8080/")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}





4.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func addres(w http.ResponseWriter, r *http.Request) {
	// путь, который пользователь открыл в браузере.
	if r.URL.Path == "/" {
		fmt.Fprintln(w, "Главная страница")
		// return прекращает выполнение функции после ответа главной страницы.
		return
	}

	// отправляет ответ со статусом 404 Not Found.
	http.NotFound(w, r)
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Обо мне")
}

func main() {
	http.HandleFunc("/", addres)
	http.HandleFunc("/about", about)

	fmt.Println("Сервер запущен: http://localhost:8080/")
	fmt.Println("Сервер запущен: http://localhost:8080/about")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}





5.
package main

import (
	"fmt"
	"log"
	"net/http"
)
func adres(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet{
		// http.Error отправляет текст ошибки и устанавливает нужный HTTP-статус.
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintln(w, "GET запрос принят!")
}

func main() {
	http.HandleFunc("/", adres)

	fmt.Println("Сервер запущен: http://localhost:8080/")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}





6.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func adres(w http.ResponseWriter, r *http.Request) {
	// Проверяем, что адрес строго совпадает с /hello
	if r.URL.Path != "/hello" {
		http.NotFound(w, r)
		return
	}

	// Получаем значение параметра "name" из URL (например, /hello?name=Arzu)
	name := r.URL.Query().Get("name")

	// Формируем ответ в зависимости от наличия имени
	if name != "" {
		fmt.Fprintf(w, "Привет, %s!\n", name)
	} else {
		fmt.Fprintln(w, "Привет, гость!")
	}
}

func main() {
	// Регистрируем обработчик для пути /hello
	// В Go стандартный ServeMux трактует "/hello" как префикс,
	// поэтому внутри хендлера выше добавлена строгая проверка пути r.URL.Path != "/hello"
	http.HandleFunc("/hello", adres)

	fmt.Println("Сервер запущен: http://localhost:8080/hello")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}





7.
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var (
	visits int
	mtx sync.Mutex
)

func home(w http.ResponseWriter, r *http.Request) {
mtx.Lock() и mtx.Unlock() защищают счётчик от одновременного изменения несколькими запросами.
	mtx.Lock()
	хранит количество посещений
	visits++
	сохраняет значение для ответа.
	currentVisits := visits
	mtx.Lock() и mtx.Unlock() защищают счётчик от одновременного изменения несколькими запросами.
	mtx.Unlock()

	fmt.Fprintf(w, "Количество посещений:  %d\n", currentVisits)
}

func main() {
	http.HandleFunc("/", home)

	fmt.Println("Сервер запущен: http://localhost:8080/")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}





8.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		fmt.Fprintln(w, `<!DOCTYPE html>
<html lang="ru">
<head>
	<meta charset="UTF-8">
	<title>Обратная связь</title>
</head>
<body>
	<h1>Форма обратной связи</h1>
	<form action="/" method="POST">
		<label>Имя:</label>
		<input type="text" name="name">
		<br><br>
		<label>Сообщение:</label>
		<textarea name="message"></textarea>
		<br><br>
		<button type="submit">Отправить</button>
	</form>
</body>
</html>`)
		return
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Ошибка обработки формы", http.StatusBadRequest)
			return
		}

		name := r.FormValue("name")
		message := r.FormValue("message")

		if name == "" || message == "" {
			fmt.Fprintln(w, "Заполните все поля")
			return
		}

		fmt.Fprintf(w, "Имя: %s\nСообщение: %s\n", name, message)
		return
	}

	http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
}

func main() {
	http.HandleFunc("/", home)

	fmt.Println("Сервер запущен: http://localhost:8080/")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}





9.
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"sync"
)

var (
	notes []string
	mtx   sync.Mutex
)

var page = template.Must(template.New("page").Parse(`
<!DOCTYPE html>
<html lang="ru">
<head>
	<meta charset="UTF-8">
	<title>Мои заметки</title>
</head>
<body>
	<h1>Мои заметки</h1>

	<form action="/" method="POST">
		<label for="note">Новая заметка:</label><br>
		<textarea id="note" name="note" required></textarea><br><br>
		<button type="submit">Добавить</button>
	</form>

	<h2>Все заметки</h2>
	{{if .}}
		<ul>
			{{range .}}
				<li>{{.}}</li>
			{{end}}
		</ul>
	{{else}}
		<p>Пока заметок нет.</p>
	{{end}}
</body>
</html>
`))

func home(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		mtx.Lock()
		currentNotes := append([]string(nil), notes...)
		mtx.Unlock()

		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		if err := page.Execute(w, currentNotes); err != nil {
			log.Printf("Ошибка отображения страницы: %v", err)
		}

	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Ошибка обработки формы", http.StatusBadRequest)
			return
		}

		note := strings.TrimSpace(r.FormValue("note"))

		if note == "" {
			http.Error(w, "Заметка не может быть пустой", http.StatusBadRequest)
			return
		}

		mtx.Lock()
		notes = append(notes, note)
		mtx.Unlock()

		http.Redirect(w, r, "/", http.StatusSeeOther)

	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", home)

	fmt.Println("Сервер запущен: http://localhost:8080/")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
*/

/* ТОЕРИЯ 15.10.2026 Тело HTTP запроса
1.
package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var mtx = sync.Mutex{}
var money = 1000 // usd
var bank = 0  // usd

//                                  http.Request содержит всю информацию об входящем HTTP запроса и в том числе его тела
func payHandler(w http.ResponseWriter, r *http.Request) {
	// []byte
	// byte == 10101010
	// bit = либо 0 либо 1
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("fail to read HTTP body:", err)
		return
	}

	httpRequestBodyString := string(httpRequestBody)

	paymentAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		fmt.Println("fail to convert HTTP body to integer:", err)
		return
	}

	mtx.Lock()
	if money-paymentAmount >= 0 {
		time.Sleep(3 * time.Second)
		money-=paymentAmount
		fmt.Println("Оплата прошла успешно:", money)
	}
	mtx.Unlock()
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("fail to read HTTP request body:", err)
		return
	}

	httpRequestBodyString := string(httpRequestBody)

	saveAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		fmt.Println("fail to convert HTTP body to int:", err)
		return
	}

	mtx.Lock()
	if money >= saveAmount {
		// взять из кошелька
		money -= saveAmount

		// положить в копилку
		bank += saveAmount

		fmt.Println("Новое значение переменной money:", money)
		fmt.Println("Новое значение переменной bank:", bank)
	}
	mtx.Unlock()
}

func main() {
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/save", saveHandler)

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("HTTP server error", err)
	}
}



2.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	mtx := sync.Mutex{}
	a := 0

	mtx.Lock()
	a++
	mtx.Unlock()

	// Лучше выбирать Atomic чем Mutex, потомучто Atomic быстрее
	b := atomic.Int64{}
	b.Add(1)
}
*/

/* ПРАКТИКА 16.10.2026 Тело HTTP запроса
// ПРАКТИКА ПО claude ai
1.
package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
)

// mtx защищает переменную visits от одновременного доступа из разных горутин.
// Go запускает каждый HTTP-запрос в отдельной горутине, поэтому без мьютекса
// два запроса могли бы одновременно изменить счётчик и дать неверный результат.
var mtx = sync.Mutex{}

// visits - общий счётчик посетителей, его видят все запросы.
var visits = 0

// visitHandler обрабатывает запрос POST /visit.
// w - это то, куда мы пишем ответ клиенту, r - это сам запрос от клиента.
func visitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Только POST", http.StatusMethodNotAllowed)
		return
	}

	// Читаем тело запроса целиком. Вернётся срез байт и ошибка.
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		// Если прочитать не удалось, сообщаем клиенту, а не молчим.
		http.Error(w, "не удалось прочитать тело", http.StatusBadRequest)
		return
	}

	// Превращаем байты в строку и убираем пробелы и переводы строк по краям.
	// Так тело из одних пробелов тоже будет считаться пустым.
	name := strings.TrimSpace(string(httpRequestBody))

	// Пустое имя - это некорректный запрос, отвечаем 400.
	if name == "" {
		http.Error(w, "имя не может быть пустым", http.StatusBadRequest)
		return
	}

	// Захватываем мьютекс: с этого момента только наша горутина
	// может работать со счётчиком, остальные ждут на Lock.
	mtx.Lock()

	// Увеличиваем общий счётчик на 1
	visits++

	// Копируем значение в локальную переменную ВНУТРИ блокировки.
	// Если читать visits после Unlock, другой запрос успеет его изменить,
	// и два человека получили бы один и тот же номер.
	number := visits

	// Освобождаем мьютекс как можно раньше, чтобы не задерживать другие запросы.
	// Дальнейшая работа (запись ответа) общих данных не трогает.
	mtx.Unlock()

	// Пишем ответ клиенту. Статус 200 ставится автоматически при первой записи.
	// Используем локальную копию number, а не глобальную visits.
	fmt.Fprintf(w, "Привет, %s! Ты посетитель номер %d\n", name, number)
}

func main() {
	http.HandleFunc("/visit", visitHandler)

	fmt.Println("Запуск сервера HTTP: http://localhost:9091/")

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Ошибка во время работы HTTP сервера", err)
	}
}





2.
package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var mtx = sync.Mutex{}
var stock = 100

func buy(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Только POST", http.StatusMethodNotAllowed)
		return
	}

	httpRequestBost, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "не удалось прочитать тело", http.StatusBadRequest)
		return
	}

	text := strings.TrimSpace(string(httpRequestBost))
	if text == "" {
		http.Error(w, "число не может быть пустым", http.StatusBadRequest)
		return
	}

	// Превращаем строку в целое число.
	// Atoi возвращает два значения: само число и ошибку.
	amount, err := strconv.Atoi(text)
	if err != nil {
		// Отвечаем клиенту статусом 400 (Bad Request) и понятным сообщением.
		// http.Error сразу ставит статус и пишет текст в ответ.
		http.Error(w, "нужно целое число", http.StatusBadRequest)

		// return обязателен: без него код пошёл бы дальше
		// и работал бы с неверным значением amount.
		return
	}

	if amount <= 0 {
		http.Error(w, "число должно быть больше нуля", http.StatusBadRequest)
		return
	}

	mtx.Lock()
	defer mtx.Unlock()

	if amount > stock {
		http.Error(w, "недостаточно товара", http.StatusConflict)
		return
	}

	stock -= amount
	fmt.Println("Оплата прошла успешно:", stock)

	fmt.Fprintln(w, "куплено товаров", amount, "осталось", stock)
}

func restock(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Только POST", http.StatusMethodNotAllowed)
		return
	}

	RequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "не удалось прочитать тело", http.StatusBadRequest)
		return
	}

	texttwo := strings.TrimSpace(string(RequestBody))
	if texttwo == "" {
		http.Error(w, "число не может быть пустым", http.StatusBadRequest)
		return
	}

	amonc, err := strconv.Atoi(texttwo)
	if err != nil {
		http.Error(w, "нужно целое число", http.StatusBadRequest)
		return
	}

	if amonc <= 0 {
		http.Error(w, "число должно быть больше нуля", http.StatusBadRequest)
		return
	}

	mtx.Lock()
	defer mtx.Unlock()

	stock += amonc
	fmt.Println("Пополнение склада прошло успешно", stock)

	fmt.Fprintln(w, "пополнен склад на", amonc, "счёт склада", stock)
}

func main() {
	http.HandleFunc("/buy", buy)
	http.HandleFunc("/restock", restock)

	fmt.Println("Сервер запущен: http://localhost:9091/buy")
	fmt.Println("Сервер запущен: http://localhost:9091/restock")

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Ошибка во ремя работы HTTP сервера", err)
	}
}





3.
package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var accounts = make(map[string]int)
var mtx = sync.Mutex{}

func deposit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Только POST", http.StatusMethodNotAllowed)
		return
	}

	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "не удалось прочитать тело", http.StatusBadRequest)
		return
	}

	text := strings.TrimSpace(string(httpRequestBody))

	about := strings.Fields(text)

	// about - срез слов, который вернула strings.Fields.
	// Для тела "anna 100" там лежит ["anna", "100"], len(about) равен 2.
	// Проверяем длину ДО обращения about[1]: если слов меньше двух,
	// индекса 1 не существует, и программа упала бы с паникой.
	// Условие != 2 заодно отсекает лишнее: "anna 100 200" тоже ошибка.
	if len(about) != 2 {
		http.Error(w, "нужно два слова: имя и сумма", http.StatusBadRequest)

		// return обязателен: без него код пошёл бы дальше к about[1]
		// и всё равно упал бы с паникой.
		return
	}

	saveAmount, err := strconv.Atoi(about[1])
	if err != nil {
		http.Error(w, "сумма должна быть числом", http.StatusBadRequest)
		return
	}

	name := about[0]

	if saveAmount <= 0 {
		http.Error(w, "сумма должна быть больше нуля", http.StatusBadRequest)
		return
	}

	mtx.Lock()
	defer mtx.Unlock()

	accounts[name] += saveAmount

	// accounts[name] - баланс только этого имени, а не вся мапа.
	// Читаем его под той же блокировкой (defer Unlock ещё не сработал),
	// поэтому другой запрос не успеет изменить значение между
	// пополнением и ответом.
	// Раньше здесь стояло просто accounts, и клиент видел чужие счета.
	fmt.Fprintln(w, "Баланс пополнен на имя", name, "и баланс", accounts[name])
}

func spend(w http.ResponseWriter, r *http.Request) {
	// Списывать можно только через POST.
	// Если метод другой, отвечаем 405 и выходим
	if r.Method != http.MethodPost {
		http.Error(w, "Только POST", http.StatusMethodNotAllowed)
		return
	}

	// Читаем тело запроса целиком. Вернётся срез байт и ошибка.
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "не удалось прочитать тело", http.StatusBadRequest)
		return
	}

	// Превращаем байты в строку и убираем пробелы и переводы строк по краям.
	// Без этого "anna 30\n" давал бы лишний мусор в последнем слове.
	text := strings.TrimSpace(string(httpRequestBody))

	// Делим строку на слова. Для "anna 30" получится ["anna", "30"].
	// Fields игнорирует лишние пробелы между словами.
	about := strings.Fields(text)

	// Слов должно быть ровно два: имя и сумма.
	// Проверка стоит ДО обращения about[1], иначе при теле "anna"
	// индекса 1 нет, и программа упала бы с паникой
	if len(about) != 2 {
		http.Error(w, "нужно два слова: имя и сумма", http.StatusBadRequest)
		return
	}

	// Второе слово превращаем в число. Atoi возвращает число и ошибку.
	// Ошибка значит, что в теле было не целое число ("abc", "3.5").
	amount, err := strconv.Atoi(about[1])
	if err != nil {
		http.Error(w, "сумма должна быть числом", http.StatusBadRequest)
		return
	}

	// Ноль и отрицательные суммы запрещены.
	// Иначе "списание" минус 50 увеличило бы баланс.
	if amount <= 0 {
		http.Error(w, "сумма должна быть больше нуля", http.StatusBadRequest)
		return
	}

	// положил имя в отдельную переменную из about[0]
	// Первое слово - имя. Кладём в отдельную переменную:
	// оно будет ключом в мапе accounts.
	name := about[0]

	// Все проверки данных клиента позади, дальше работаем с общей мапой.
	// Захватываем мьютекс, чтобы другие запросы не меняли accounts одновременно.
	mtx.Lock()
	// defer снимет блокировку при любом выходе из функции:
	// и при return с ошибкой, и в конце. Забыть Unlock невозможно.
	defer mtx.Unlock()

	// Читаем баланс с проверкой существования.
	// ok равен false, если ключа name нет в мапе.
	// Обычное чтение accounts[name] вернуло бы 0, и мы не отличили бы
	// "счёта нет" от "счёт с нулевым балансом".
	balance, ok := accounts[name]
	if !ok {
		// Счёта нет: 404. Мапа остаётся без изменений.
		http.Error(w, "такого счёта нет", http.StatusNotFound)
		return
	}

	// Проверяем, хватает ли денег. Условие "меньше", а не "меньше или равно":
	// потратить ровно весь баланс можно, тогда останется 0.
	// Это 409, а не 400: запрос корректный, но не подходит под текущее
	// состояние счёта. Мапа остаётся без изменений.
	if balance < amount {
		http.Error(w, "на балансе недостаточно средств", http.StatusConflict)
		return
	}

	// Если денег хватает — списываем
	// Все проверки пройдены, только теперь списываем.
	// Проверка и списание идут под одной блокировкой, поэтому между ними
	// другой запрос не успеет изменить баланс.
	accounts[name] -= amount

	// Отвечаем клиенту: имя и новый баланс только этого имени.
	// Выводим accounts[name], а не всю мапу, чтобы не показывать чужие счета.
	// Статус 200 ставится автоматически при первой записи в w
	fmt.Fprintln(w, "у", name, "осталось", accounts[name])
}

func main() {
	http.HandleFunc("/deposit", deposit)
	http.HandleFunc("/spend", spend)

	fmt.Println("Сервер запущен: http://localhost:9091/deposit")
	fmt.Println("Сервер запущен: http://localhost:9091/spend")

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Ошибка во время HTTP сервера:", err)
	}
}
*/

/* 18.10.2026 ПРАКТИКА Тело HTTP запроса
1.
package main

import (
	"errors"  // errors.As: проверить, что ошибка именно "тело слишком большое"
	"fmt"     // Fprintf: писать форматированный текст в ответ
	"io"      // io.ReadAll: прочитать тело запроса целиком
	"net/http" // всё для HTTP: сервер, запрос, ответ, коды статусов
	"strings" // Fields, Count, HasSuffix: работа со строками
)

// stats - обработчик для адреса /stats.
// w - через него мы пишем ответ клиенту.
// r - входящий запрос (метод, заголовки, тело)
func stats(w http.ResponseWriter, r *http.Request) {
	// 1. Проверяем метод. Если это не POST, отвечаем 405 и выходим.
	if r.Method != http.MethodPost {
		http.Error(w, "только POST", http.StatusMethodNotAllowed)
		// без return код пошёл бы дальше
		return
	}

	// 2. Ставим лимит на размер тела: 1<<20 = 1 048 576 байт = 1 МБ.
	// MaxBytesReader оборачивает r.Body и вернёт ошибку,
	// как только клиент пришлёт больше лимита.
	// Благодаря этому огромный запрос не попадёт в память целиком.
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20) // 1 МБ

	// 3. Читаем тело. body - это []byte (срез байт).
	body, err := io.ReadAll(r.Body)
	if err != nil {
		// Различаем два вида ошибок чтения.
		var maxErr *http.MaxBytesError
		if errors.As(err, &maxErr) {
			// Тело больше 1 МБ: 413.
			http.Error(w, "body too large", http.StatusRequestEntityTooLarge)
			return
		}
		// Любая другая ошибка чтения: 400.
		http.Error(w, "не удалось прочитать тело", http.StatusBadRequest)
		return
	}

	// 4. Пустое тело: 400. Проверяем уже после чтения,
	// по длине полученного среза байт.
	if len(body) == 0 {
		http.Error(w, "empty body", http.StatusBadRequest)
		return
	}

	// 5. Превращаем байты в строку, чтобы работать с функциями пакета strings
	text := string(body)

	// 6. Считаем статистику.

	// Байты: длина среза. Русская буква даст 2 байта, а не 1.
	bytesCount := len(body)

	// Слова: Fields делит строку по любым пробельным символам
	// (пробел, таб, \n) и не создаёт пустых слов при нескольких пробелах подряд.
	wordsCount := len(strings.Fields(text))

	// Строки: сначала считаем символы перевода строки.
	linesCount := strings.Count(text, "\n")

	// Если текст не заканчивается на \n, последняя строка осталась
	// не посчитанной, поэтому добавляем её.
	if !strings.HasSuffix(text, "\n") {
		linesCount++
	}

	// 7. Отвечаем. Заголовок нужно ставить ДО первой записи в w,
	// иначе он уже не применится.
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// Fprintf пишет в w так же, как Printf пишет в консоль.
	// %d - подставить целое число, \n - перевод строки.
	// Статус 200 ставится автоматически при первой записи.
	fmt.Fprintf(w, "bytes: %d\nlines: %d\nwords: %d\n", bytesCount, linesCount, wordsCount)
}

func main() {
	http.HandleFunc("/stats", stats)

	fmt.Println("Запуск сервера: http://localhost:8080/stats")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка HTTP сервера", err)
	}
}





2.
package main

import (
	"fmt"      // Fprintf, Println: вывод в ответ и в консоль
	"net/http" // сервер, запрос, ответ, коды статусов
	"strconv"  // Atoi: перевод строки в число
	"strings"  // Contains, TrimSpace, Join: работа со строками
)

// register - обработчик для адреса /register.
func register(w http.ResponseWriter, r *http.Request) {
	// 1. Проверяем метод: только POST.
	if r.Method != http.MethodPost {
		http.Error(w, "только POST", http.StatusMethodNotAllowed)
		return
	}

	// 2. Разбираем тело формы (name=Ivan&age=25&email=a@b.c).
	// После этого значения лежат в r.PostForm.
	// Если тело сломано, вернётся ошибка.
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "не удалось разобрать форму", http.StatusBadRequest)
		return
	}

	// 3. Достаём значения полей.
	// PostFormValue берёт данные только из тела запроса (не из адреса).
	// Если поля нет, вернётся пустая строка, а не ошибка.
	// TrimSpace убирает пробелы по краям, чтобы имя "   " считалось пустым.
	name := strings.TrimSpace(r.PostFormValue("name"))
	ageStr := strings.TrimSpace(r.PostFormValue("age"))
	email := strings.TrimSpace(r.PostFormValue("email"))

	// 4. Сюда будем складывать ВСЕ найденные ошибки,
	// а не останавливаться на первой.
	var problems []string

	// Проверка name: не пустое.
	if name == "" {
		problems = append(problems, "name: не должно быть пустым")
	}

	// Проверка age: целое число от 1 до 120.
	// Если Atoi вернул ошибку, return не делаем:
	// нужно проверить остальные поля тоже.
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		problems = append(problems, "age: должно быть целым числом")
	} else if age < 1 || age > 120 {
		problems = append(problems, "age: должно быть от 1 до 120")
	}

	// Проверка email: достаточно наличия символа @.
	if !strings.Contains(email, "@") {
		problems = append(problems, "email: должен содержать @")
	}

	// 5. Если ошибки есть, отвечаем 400 и выводим их все списком.
	if len(problems) > 0 {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		// WriteHeader нужно вызвать ПОСЛЕ установки заголовков
		// и ДО записи тела, иначе статус останется 200.
		w.WriteHeader(http.StatusBadRequest)
		// Join склеивает срез строк в одну, разделяя через \n.
		fmt.Fprintln(w, strings.Join(problems, "\n"))
		return
	}

	// 6. Всё верно: 200 и приветствие.
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "Привет, %s!\n", name)
}

func main() {
	http.HandleFunc("/register", register)

	fmt.Println("Запуск сервера: http://localhost:8080/register")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка HTTP сервера:", err)
	}
}





3.
package main

import (
	"bytes"    // bytes.Count: посчитать символы \n в куске данных
	"errors"   // errors.As, errors.Is: распознавать виды ошибок
	"fmt"      // Fprintf, Println: вывод в ответ и в консоль
	"io"       // io.EOF: признак конца файла
	"net/http" // сервер, запрос, ответ, коды статусов
)

// upload - обработчик для адреса /upload.
func upload(w http.ResponseWriter, r *http.Request) {
	// 1. Проверяем метод: только POST.
	if r.Method != http.MethodPost {
		http.Error(w, "только POST", http.StatusMethodNotAllowed)
		return
	}

	// 2. Лимит на весь запрос: 10<<20 = 10 МБ.
	// Всё, что больше, вызовет ошибку при чтении.
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)

	// 3. Разбираем multipart-запрос.
	// 1<<20 - сколько МБ данных держать в памяти (1 МБ).
	// Части, которые не поместились, автоматически уходят во
	// временные файлы на диске, поэтому память не раздувается.
	// Временные файлы сервер удалит сам после завершения обработчика.
	err := r.ParseMultipartForm(1 << 20)
	if err != nil {
		var maxErr *http.MaxBytesError
		if errors.As(err, &maxErr) {
			// Запрос больше 10 МБ: 413.
			http.Error(w, "запрос слишком большой", http.StatusRequestEntityTooLarge)
			return
		}
		http.Error(w, "не удалось разобрать multipart-запрос", http.StatusBadRequest)
		return
	}

	// 4. Достаём файл из поля "file".
	// FormFile возвращает три значения:
	// file - сам файл (его можно читать как поток),
	// header - информация о файле (имя, размер),
	// err - ошибка.
	file, header, err := r.FormFile("file")
	if err != nil {
		// Различаем "поля file нет" и другие ошибки.
		if errors.Is(err, http.ErrMissingFile) {
			http.Error(w, "в запросе нет поля file", http.StatusBadRequest)
			return
		}
		http.Error(w, "не удалось получить файл", http.StatusBadRequest)
		return
	}
	// Закрываем файл при выходе из функции.
	// defer ставим ПОСЛЕ проверки ошибки: если err != nil, file пустой,
	// и закрывать там нечего.
	defer file.Close()

	// 5. Читаем файл потоком, кусками по 32 КБ.
	// Целиком файл в память не загружается: в буфере всегда
	// только текущий кусок.
	buf := make([]byte, 32*1024)
	var size int64 // сколько байт прочитали всего
	lines := 0     // сколько символов \n встретили
	var last byte  // последний прочитанный байт (нужен для подсчёта строк)

	for {
		// Read кладёт данные в buf и возвращает, сколько байт прочитано (n).
		// Важно: сначала обрабатываем n байт, и только потом смотрим на ошибку,
		// потому что вместе с последним куском может прийти io.EOF.
		n, err := file.Read(buf)
		if n > 0 {
			size += int64(n)
			lines += bytes.Count(buf[:n], []byte{'\n'}) // \n только в прочитанной части
			last = buf[n-1]
		}
		if err == io.EOF {
			break // файл закончился, это нормальный выход
		}
		if err != nil {
			http.Error(w, "ошибка чтения файла", http.StatusInternalServerError)
			return
		}
	}

	// Если файл не пустой и не заканчивается на \n,
	// последняя строка осталась не посчитанной.
	if size > 0 && last != '\n' {
		lines++
	}

	// 6. Обычное поле формы (не файл). Если его нет, вернётся пустая строка.
	note := r.PostFormValue("note")

	// 7. Отвечаем. Заголовок нужно ставить до первой записи в w.
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "filename: %s\nsize: %d\nlines: %d\nnote: %s\n",
		header.Filename, size, lines, note)
}

func main() {
	http.HandleFunc("/upload", upload)

	fmt.Println("Запуск сервера: http://localhost:8080/upload")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка HTTP сервера:", err)
	}
}
*/

/* ПРАКТИКА 18.10.2026 Тело HTTP запроса
1.
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var balance = 1000
var mtx = sync.Mutex{}

func add(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "только POST", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "не удалось прочитать тело", http.StatusBadRequest)
		return
	}

	amount, err := strconv.Atoi(string(body))
	if err != nil {
		http.Error(w, "нужно целое число", http.StatusBadRequest)
		return
	}

	mtx.Lock()
	defer mtx.Unlock()
	balance += amount

	fmt.Fprintln(w, "Ваш баланс пополнен", balance)
}

func main() {
	http.HandleFunc("/add", add)

	fmt.Println("Запус сервера: http://localhost:8080/add")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}





2.
package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

// Общее состояние сервера (глобальные переменные).
// К ним обращаются сразу несколько запросов, поэтому нужен мьютекс.
var alice = 500
var bob = 300
var fees = 0
var mtx = sync.Mutex{}

// balance: GET /balance. Показывает текущие балансы.
func balance(w http.ResponseWriter, r *http.Request) {
	// Разрешён только GET. Иначе 405 Method Not Allowed.
	if r.Method != http.MethodGet {
		http.Error(w, "только GET", http.StatusMethodNotAllowed)
		return
	}

	// Закрываем замок на время чтения: пока мы читаем три переменные,
	// другой запрос не должен их менять, иначе увидим «рваные» данные.
	mtx.Lock()
	// defer выполнится при выходе из функции, поэтому Unlock не забудется.
	defer mtx.Unlock()

	// Fprintln пишет в w (ответ клиенту), ставит пробел между аргументами
	// и перевод строки в конце.
	fmt.Fprintln(w, "alice:", alice)
	fmt.Fprintln(w, "bob:", bob)
	fmt.Fprintln(w, "fees:", fees)
}

// deposit: POST /deposit, тело "alice 100". Кладёт деньги на счёт.
func deposit(w http.ResponseWriter, r *http.Request) {
	// 1. Метод: только POST.
	if r.Method != http.MethodPost {
		http.Error(w, "только POST", http.StatusMethodNotAllowed)
		return
	}

	// 2. Читаем тело запроса целиком. body это []byte.
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "не удалось прочитать тело", http.StatusBadRequest)
		return
	}

	// 3. Превращаем байты в строку и убираем пробелы и \n по краям.
	text := strings.TrimSpace(string(body))

	// 4. Делим строку на слова: "alice 100" → ["alice", "100"]
	words := strings.Fields(text)

	// сравнивать с двойкой длину среза: встроенная функция
	// 5. Проверяем количество слов ДО обращения по индексу.
	// Иначе words[1] при одном слове вызвал бы панику.
	if len(words) != 2 {
		http.Error(w, "ошибка, нужно ввести имя и сумму", http.StatusBadRequest)
		return
	}

	// 6. Вторая часть это сумма. Atoi переводит "100" в число 100.
	// Если там не число, вернётся ошибка.
	summa, err := strconv.Atoi(words[1])
	if err != nil {
		http.Error(w, "нужно целое число", http.StatusBadRequest)
		return
	}

	// 7. Сумма должна быть положительной. Проверяем ДО Lock:
	// эта проверка не касается общих данных, держать из-за неё замок незачем.
	if summa <= 0 {
		http.Error(w, "сумма должна быть больше нуля", http.StatusBadRequest)
		return
	}

	// 8. Имя счёта. TrimSpace здесь лишний: Fields уже убрал пробелы.
	stroka := strings.TrimSpace(words[0])

	// 9. Эта проверка дублирует ветку default у switch ниже.
	// Одну из двух можно убрать.
	if stroka != "alice" && stroka != "bob" {
		http.Error(w, "ошибка, имя счёта не alice и не bob", http.StatusBadRequest)
		return
	}

	// 10. Закрываем замок: дальше меняем общие данные.
	mtx.Lock()
	defer mtx.Unlock()

	// 11. Выбираем счёт по имени и прибавляем сумму.
	// В Go break в конце case не нужен.
	switch stroka {
	case "alice":
		alice += summa
	case "bob":
		bob += summa
	default:
		http.Error(w, "ошибка вклада в депозит", http.StatusBadRequest)
		return
	}

	// 12. Успех. Статус 200 ставится автоматически при первой записи.
	fmt.Fprintln(w, "ok")
}

// withdraw: POST /withdraw, тело "bob 50". Снимает деньги со счёта.
func withdraw(w http.ResponseWriter, r *http.Request) {
	// Метод, чтение тела, разбор: то же, что в deposit.
	if r.Method != http.MethodPost {
		http.Error(w, "только POST", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "не удалось прочесть тело", http.StatusBadRequest)
		return
	}

	text := strings.TrimSpace(string(body))

	words := strings.Fields(text)

	// Ровно два слова: имя и сумма.
	if len(words) != 2 {
		http.Error(w, "ошибка, нужно имя и сумма", http.StatusBadRequest)
		return
	}

	summa, err := strconv.Atoi(words[1])
	if err != nil {
		http.Error(w, "нужно целое число", http.StatusBadRequest)
		return
	}

	// проверка суммы до Lock
	if summa <= 0 {
		http.Error(w, "сумма должна быть больше нуля", http.StatusBadRequest)
		return
	}

	mtx.Lock()
	defer mtx.Unlock()

	// Переключаемся по имени счёта (words[0]).
	switch words[0] {
	case "alice":
		// Проверка денег ВНУТРИ замка. Если её сделать до Lock,
		// два одновременных запроса оба увидят «хватает» и вместе
		// снимут больше, чем есть.
		if alice < summa {
			// 409 Conflict: запрос корректный, но конфликтует
			// с текущим состоянием счёта.
			http.Error(w, "недостаточно средств", http.StatusConflict)
			return
		}
		alice -= summa
	case "bob":
		if bob < summa {
			http.Error(w, "недостаточно средств", http.StatusConflict)
			return
		}
		bob -= summa
	default:
		// Неизвестное имя счёта. Без default запрос "carol 10"
		// вернул бы ok, ничего не сделав.
		http.Error(w, "ошибка снятия из депозита", http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "ok")
}

// transfer: POST /transfer, тело "alice bob 200".
// Перевод от alice к bob с комиссией 1 доллар.
func transfer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "только POST", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "не удалось прочитать тело", http.StatusBadRequest)
		return
	}

	// убирает пробелы и переводы строки по краям строки. " a b\n" → "a b"
	text := strings.TrimSpace(string(body))

	// "alice bob 200" → ["alice", "bob", "200"]
	words := strings.Fields(text)

	// Ровно три значения: отправитель, получатель, сумма
	if len(words) != 3 {
		http.Error(w, "нужно три значения - отправитель, получатель и сумма.", http.StatusBadRequest)
		return
	}

	// Сумма это ТРЕТЬЕ слово, индекс 2 (счёт с нуля).
	summa, err := strconv.Atoi(words[2])
	if err != nil {
		http.Error(w, "нужно целое число", http.StatusBadRequest)
		return
	}

	if summa <= 0 {
		http.Error(w, "ошибка, сумма должна быть больше нуля", http.StatusBadRequest)
		return
	}

	// Отправитель и получатель не могут совпадать.
	if words[0] == words[1] {
		http.Error(w, "нельзя переводить сумму самому себе", http.StatusBadRequest)
		return
	}

	// Оба имени должны быть alice или bob.
	if (words[0] != "alice" && words[0] != "bob") || (words[1] != "alice" && words[1] != "bob") {
		http.Error(w, "неизвестный счёт", http.StatusBadRequest)
		return
	}

	// Все проверки формата позади. Теперь работаем с общими данными.
	// Один Lock на всю операцию: три изменения (отправитель, получатель, fees)
	// должны произойти как единое целое, иначе другой запрос
	// увидит «деньги исчезли», пока перевод в процессе.
	mtx.Lock()
	defer mtx.Unlock()

	// from и to это указатели: они хранят АДРЕСА переменных alice или bob.
	// Благодаря этому не нужен отдельный switch для каждого из двух участников.
	var from, to *int

	if words[0] == "alice" {
		from = &alice // &alice: адрес переменной alice
	} else {
		from = &bob
	}

	if words[1] == "alice" {
		to = &alice
	} else {
		to = &bob
	}

	// *from: значение по адресу, то есть баланс отправителя.
	// Нужно покрыть сумму и комиссию (summa+1).
	// Проверяем ДО изменений, чтобы не пришлось ничего откатывать.
	if *from < summa +1 {
		http.Error(w, "недостаточно средств", http.StatusConflict)
		return
	}

	*from -= summa + 1
	*to += summa
	fees++

	fmt.Fprintln(w, "ok")
}

func main() {
	// Привязываем адреса к функциям-обработчикам.
	http.HandleFunc("/balance", balance)
	http.HandleFunc("/deposit", deposit)
	http.HandleFunc("/withdraw", withdraw)
	http.HandleFunc("/transfer", transfer)

	fmt.Println("Запуск сервера: http://localhost:9093/balance")
	fmt.Println("Запуск сервера: http://localhost:9093/deposit")
	fmt.Println("Запуск сервера: http://localhost:9093/withdraw")
	fmt.Println("Запуск сервера: http://localhost:9093/transfer")

	// Запускаем сервер на порту 9093. Функция блокирует выполнение,
	// пока сервер работает, и возвращает ошибку, если он упал
	// (например, порт занят).
	err := http.ListenAndServe(":9093", nil)
	if err != nil {
		fmt.Println(err)
	}
}




2.
package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

var stock = 100   // свободно на складе
var reserved = 0  // в резерве
var sold = 0      // продано
var paying = 0    // штуки, за которые прямо сейчас идёт оплата
var mtx = sync.Mutex{}

// readAmount: общий разбор для всех POST-обработчиков.
// Проверяет метод, читает тело, возвращает положительное число.
// Если что-то не так, сама отвечает клиенту и возвращает ok=false,
// поэтому в обработчике достаточно сделать return.
func readAmount(w http.ResponseWriter, r *http.Request) (int, bool) {
	if r.Method != http.MethodPost {
		http.Error(w, "только POST", http.StatusMethodNotAllowed)
		return 0, false
	}

	// Ограничиваем тело 1 КБ: для одного числа больше не нужно.
	body, err := io.ReadAll(http.MaxBytesReader(w, r.Body, 1<<10))
	if err != nil {
		http.Error(w, "не удалось прочитать тело", http.StatusBadRequest)
		return 0, false
	}

	n, err := strconv.Atoi(strings.TrimSpace(string(body)))
	if err != nil {
		http.Error(w, "нужно целое число", http.StatusBadRequest)
		return 0, false
	}
	if n <= 0 {
		http.Error(w, "число должно быть больше нуля", http.StatusBadRequest)
		return 0, false
	}
	return n, true
}

// status: GET /status
func status(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "только GET", http.StatusMethodNotAllowed)
		return
	}

	mtx.Lock()
	defer mtx.Unlock()

	fmt.Fprintln(w, "stock:", stock)
	fmt.Fprintln(w, "reserved:", reserved)
	fmt.Fprintln(w, "sold:", sold)
	// paying нужен, чтобы инвариант stock+reserved+paying+sold не нарушался,
	// пока идёт оплата. Если не хочешь его показывать, убери эту строку.
	fmt.Fprintln(w, "paying:", paying)
}

// reserve: POST /reserve, из stock в reserved
func reserve(w http.ResponseWriter, r *http.Request) {
	n, ok := readAmount(w, r)
	if !ok {
		return
	}

	mtx.Lock()
	defer mtx.Unlock()

	// Проверка ВНУТРИ замка, иначе два запроса увидят одно и то же количество.
	if stock < n {
		http.Error(w, "недостаточно товара на складе", http.StatusConflict)
		return
	}
	stock -= n
	reserved += n

	fmt.Fprintln(w, "ok")
}

// cancel: POST /cancel, из reserved обратно в stock
func cancel(w http.ResponseWriter, r *http.Request) {
	n, ok := readAmount(w, r)
	if !ok {
		return
	}

	mtx.Lock()
	defer mtx.Unlock()

	if reserved < n {
		http.Error(w, "в резерве меньше, чем нужно отменить", http.StatusConflict)
		return
	}
	reserved -= n
	stock += n

	fmt.Fprintln(w, "ok")
}

// restock: POST /restock, добавляет товар на склад
func restock(w http.ResponseWriter, r *http.Request) {
	n, ok := readAmount(w, r)
	if !ok {
		return
	}

	mtx.Lock()
	stock += n
	mtx.Unlock()

	fmt.Fprintln(w, "ok")
}

// buy: POST /buy, из reserved в sold, оплата занимает 3 секунды.
// Главная идея: замок НЕ держим во время сна.
func buy(w http.ResponseWriter, r *http.Request) {
	n, ok := readAmount(w, r)
	if !ok {
		return
	}

	// Шаг 1 (под замком): проверяем и «забираем» штуки из резерва.
	// Теперь второй параллельный /buy уже не сможет их купить повторно.
	mtx.Lock()
	if reserved < n {
		mtx.Unlock() // здесь нет defer, поэтому открываем замок вручную
		http.Error(w, "в резерве меньше, чем нужно купить", http.StatusConflict)
		return
	}
	reserved -= n
	paying += n // штуки «в пути»: не потеряны, просто оплачиваются
	mtx.Unlock()

	// Шаг 2 (БЕЗ замка): имитация платежа.
	// Пока мы спим, /status, /reserve и другие запросы работают.
	time.Sleep(3 * time.Second)

	// Шаг 3 (под замком): платёж прошёл, фиксируем продажу.
	mtx.Lock()
	paying -= n
	sold += n
	mtx.Unlock()

	fmt.Fprintln(w, "ok")
}

func main() {
	http.HandleFunc("/status", status)
	http.HandleFunc("/reserve", reserve)
	http.HandleFunc("/cancel", cancel)
	http.HandleFunc("/restock", restock)
	http.HandleFunc("/buy", buy)

	fmt.Println("Запуск сервера: http://localhost:9092")

	err := http.ListenAndServe(":9092", nil)
	if err != nil {
		fmt.Println(err)
	}
}





3.
package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
	"strings"
)

// bodyErr: общая реакция на ошибку чтения тела.
// Слишком большое тело даёт 413, любая другая ошибка даёт 400.
func bodyErr(w http.ResponseWriter, err error) {
	var maxErr *http.MaxBytesError
	if errors.As(err, &maxErr) {
		http.Error(w, "тело слишком большое", http.StatusRequestEntityTooLarge)
		return
	}
	http.Error(w, "не удалось прочитать тело", http.StatusBadRequest)
}

// ---------- Задание 1: POST /sum ----------

func sum(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "только POST", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20) // лимит 1 МБ

	body, err := io.ReadAll(r.Body)
	if err != nil {
		bodyErr(w, err)
		return
	}

	// Fields для пустой строки вернёт пустой срез,
	// поэтому одна проверка покрывает и «пусто», и «одни пробелы».
	tokens := strings.Fields(string(body))
	if len(tokens) == 0 {
		http.Error(w, "в теле нет чисел", http.StatusBadRequest)
		return
	}

	total := 0
	for _, t := range tokens {
		n, err := strconv.Atoi(t)
		if err != nil {
			// Говорим, КАКОЙ именно токен плохой.
			http.Error(w, "bad number: "+t, http.StatusBadRequest)
			return
		}
		total += n
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "sum: %d\n", total)
}

// ---------- Задание 2: POST /checksum ----------

func checksum(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "только POST", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 100<<20) // лимит 100 МБ

	// sha256.New() возвращает хеш, в который можно писать как в поток.
	h := sha256.New()

	// io.Copy копирует кусками из тела в хеш, память не растёт.
	// Возвращает количество скопированных байт, это и есть size.
	size, err := io.Copy(h, r.Body)
	if err != nil {
		bodyErr(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	// %x на []byte печатает шестнадцатеричную строку.
	fmt.Fprintf(w, "size: %d\nsha256: %x\n", size, h.Sum(nil))
}

// ---------- Задание 3: POST /orders ----------

func orders(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "только POST", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	// csv.Reader читает прямо из тела и сам разбирает кавычки:
	// "Хлеб, ржаной",45.50,2 даст три поля, а не четыре.
	cr := csv.NewReader(r.Body)
	// -1 отключает встроенную проверку числа полей: проверим сами,
	// чтобы вернуть свой текст с номером записи.
	cr.FieldsPerRecord = -1

	// Первая запись это заголовок, пропускаем.
	_, err := cr.Read()
	if err == io.EOF {
		http.Error(w, "пустое тело", http.StatusBadRequest)
		return
	}
	if err != nil {
		bodyErr(w, err)
		return
	}

	count := 0
	var totalCents int64 // деньги считаем в копейках целыми, не во float
	recNum := 1          // номер записи (заголовок это 1)

	for {
		rec, err := cr.Read()
		if err == io.EOF {
			break // данные закончились
		}
		if err != nil {
			bodyErr(w, err)
			return
		}
		recNum++

		if len(rec) != 3 {
			http.Error(w, fmt.Sprintf("запись %d: нужно 3 поля", recNum), http.StatusBadRequest)
			return
		}

		price, err := strconv.ParseFloat(strings.TrimSpace(rec[1]), 64)
		// ParseFloat принимает "NaN" и "Inf", их тоже отсекаем.
		if err != nil || math.IsNaN(price) || math.IsInf(price, 0) || price < 0 {
			http.Error(w, fmt.Sprintf("запись %d: неверное поле price", recNum), http.StatusBadRequest)
			return
		}

		qty, err := strconv.Atoi(strings.TrimSpace(rec[2]))
		if err != nil || qty <= 0 {
			http.Error(w, fmt.Sprintf("запись %d: неверное поле qty", recNum), http.StatusBadRequest)
			return
		}

		// Округляем до копеек и умножаем на количество.
		totalCents += int64(math.Round(price*100)) * int64(qty)
		count++
	}

	if count == 0 {
		http.Error(w, "нет ни одной строки после заголовка", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	// %02d добивает копейки нулём: 5 даст 05.
	fmt.Fprintf(w, "items: %d\ntotal: %d.%02d\n", count, totalCents/100, totalCents%100)
}

// ---------- Задание 4: middleware ----------

// logBody принимает любой обработчик и возвращает обёрнутый.
// Перед вызовом реального обработчика печатает размер и первые 50 байт тела,
// при этом обработчик по-прежнему видит тело целиком.
func logBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		const limit = 1 << 20

		// ContentLength берём из заголовка запроса, тело при этом не читаем.
		// Если длина неизвестна, там будет -1.
		if r.ContentLength > limit {
			fmt.Println("тело больше 1 МБ, не логируем")
			next.ServeHTTP(w, r)
			return
		}

		// ReadFull читает ровно 50 байт или меньше, если тело короче.
		// Ошибки EOF и ErrUnexpectedEOF здесь нормальны: тело просто короткое.
		head := make([]byte, 50)
		n, _ := io.ReadFull(r.Body, head)
		head = head[:n]

		fmt.Printf("%s %s: размер=%d, начало=%q\n", r.Method, r.URL.Path, r.ContentLength, head)

		// Возвращаем прочитанное обратно: сначала наши 50 байт, потом остаток потока.
		// MultiReader склеивает два источника в один.
		// Close нужно оставить от оригинального тела, поэтому собираем ReadCloser.
		r.Body = struct {
			io.Reader
			io.Closer
		}{
			Reader: io.MultiReader(bytes.NewReader(head), r.Body),
			Closer: r.Body,
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// /sum обёрнут в middleware, остальные нет (можно обернуть любой).
	http.Handle("/sum", logBody(http.HandlerFunc(sum)))
	http.HandleFunc("/checksum", checksum)
	http.HandleFunc("/orders", orders)

	fmt.Println("Запуск сервера: http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
*/

/* ТЕОРИЯ 19.10.2026 Тело HTTP ответа и статус коды
package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

var mtx = sync.Mutex{}
var money = 1000 // usd
var bank = 0  // usd

//                                  http.Request содержит всю информацию об входящем HTTP запроса и в том числе его тела
func payHandler(w http.ResponseWriter, r *http.Request) {
	// []byte
	// byte == 10101010
	// bit = либо 0 либо 1
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		// "http status code"
		w.WriteHeader(http.StatusInternalServerError)

		msg := "fail to read HTTP body:" + err.Error()
		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response:", err)
		}
		return
	}

	httpRequestBodyString := string(httpRequestBody)

	paymentAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		msg := "fail to convert HTTP body to integer:" + err.Error()
		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response:", err)
		}
		return
	}

	mtx.Lock()
	if money-paymentAmount >= 0 {
		money-=paymentAmount
		//                            превращаю money в строку с помощью strconv
		msg := "Оплата прошла успешно:" + strconv.Itoa(money)
		fmt.Println()
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response:", err)
		}
	}
	mtx.Unlock()
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("fail to read HTTP request body:", err)
		return
	}

	httpRequestBodyString := string(httpRequestBody)

	saveAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		fmt.Println("fail to convert HTTP body to int:", err)
		return
	}

	mtx.Lock()
	if money >= saveAmount {
		// взять из кошелька
		money -= saveAmount

		// положить в копилку
		bank += saveAmount

		fmt.Println("Новое значение переменной money:", money)
		fmt.Println("Новое значение переменной bank:", bank)
	}
	mtx.Unlock()
}

func main() {
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/save", saveHandler)

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("HTTP server error", err)
	}
}
*/

/* ПРАКТИКА 19.10.2026 Тело HTTP ответа и статус коды
1.
package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "только GET", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Сервер работает" )
}

func main() {
	http.HandleFunc("/hello", hello)

	fmt.Println("Запус сервера: http://localhost:8080/hello")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("HTTP server error", err)
	}
}





2.
package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func about(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "только GET", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Сервер работает")
}

func unknow(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Страница не найдена", http.StatusNotFound)
}

func hello(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "только GET", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Привет, клиент!")
}

func age(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "только POST", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "не удалось прочитать тело", http.StatusBadRequest)
		return
	}

	// преобразует данные в целое число.
	score, err := strconv.Atoi(string(body))
	if err != nil {
		http.Error(w, "нужно целое число", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "ваш возраст", score)
}

func name(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "только POST", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "не удалось прочитать тело", http.StatusBadRequest)
		return
	}

	// Проверяем, пустое ли оно или содержит только пробелы
	if len(strings.TrimSpace(string(body))) == 0 {
		http.Error(w, "Тело запроса не может быть пустым", http.StatusBadRequest)
		return
	}

	name := strings.TrimSpace(string(body))

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Привет,", name)
}

func main() {
	http.HandleFunc("/about", about)
	http.HandleFunc("/", unknow)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/age", age)
	http.HandleFunc("/name", name)

	fmt.Println("Запуск сервера: http://localhost:8080/about")
	fmt.Println("Запуск сервера: http://localhost:8080/unknow")
	fmt.Println("Запуск сервера: http://localhost:8080/hello")
	fmt.Println("Запуск сервера: http://localhost:8080/age")
	fmt.Println("Запуск сервера: http://localhost:8080/name")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("HTTP server error", err)
	}
}
*/

/* ТЕОРИЯ 19.10.2026 HTTP хеддеры
package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

var mtx = sync.Mutex{}
var money = 1000 // usd
var bank = 0  // usd

//                                  http.Request содержит всю информацию об входящем HTTP запроса и в том числе его тела
func payHandler(w http.ResponseWriter, r *http.Request) {
	// []byte
	// byte == 10101010
	// bit = либо 0 либо 1
	for k, v := range r.Header {
		fmt.Println("k:", k, "-- v:", v)
	}


	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		// "http status code"
		w.WriteHeader(http.StatusInternalServerError)

		msg := "fail to read HTTP body:" + err.Error()
		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response:", err)
		}
		return
	}

	httpRequestBodyString := string(httpRequestBody)

	paymentAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		msg := "fail to convert HTTP body to integer:" + err.Error()
		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response:", err)
		}
		return
	}

	mtx.Lock()
	if money-paymentAmount >= 0 {
		money-=paymentAmount
		//                            превращаю money в строку с помощью strconv
		msg := "Оплата прошла успешно:" + strconv.Itoa(money)
		fmt.Println()
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response:", err)
		}
	}
	mtx.Unlock()
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("fail to read HTTP request body:", err)
		return
	}

	httpRequestBodyString := string(httpRequestBody)

	saveAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		fmt.Println("fail to convert HTTP body to int:", err)
		return
	}

	mtx.Lock()
	if money >= saveAmount {
		// взять из кошелька
		money -= saveAmount

		// положить в копилку
		bank += saveAmount

		fmt.Println("Новое значение переменной money:", money)
		fmt.Println("Новое значение переменной bank:", bank)
	}
	mtx.Unlock()
}

func main() {
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/save", saveHandler)

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("HTTP server error", err)
	}
}
*/

/* ТЕОРИЯ 19.10.2026 HTTP методы
package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

var mtx = sync.Mutex{}
var money = 1000 // usd
var bank = 0  // usd

//                                  http.Request содержит всю информацию об входящем HTTP запроса и в том числе его тела
func payHandler(w http.ResponseWriter, r *http.Request) {
	// []byte
	// byte == 10101010
	// bit = либо 0 либо 1


	// for k, v := range r.Header {
	//	fmt.Println("k:", k, "-- v:", v)
	//}

	fmt.Println("HTTP method:", r.Method)

	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		// "http status code"
		w.WriteHeader(http.StatusInternalServerError)

		msg := "fail to read HTTP body:" + err.Error()
		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response:", err)
		}
		return
	}

	httpRequestBodyString := string(httpRequestBody)

	paymentAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		msg := "fail to convert HTTP body to integer:" + err.Error()
		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response:", err)
		}
		return
	}

	mtx.Lock()
	if money-paymentAmount >= 0 {
		money-=paymentAmount
		//                            превращаю money в строку с помощью strconv
		msg := "Оплата прошла успешно:" + strconv.Itoa(money)
		fmt.Println()
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response:", err)
		}
	}
	mtx.Unlock()
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("fail to read HTTP request body:", err)
		return
	}

	httpRequestBodyString := string(httpRequestBody)

	saveAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		fmt.Println("fail to convert HTTP body to int:", err)
		return
	}

	mtx.Lock()
	if money >= saveAmount {
		// взять из кошелька
		money -= saveAmount

		// положить в копилку
		bank += saveAmount

		fmt.Println("Новое значение переменной money:", money)
		fmt.Println("Новое значение переменной bank:", bank)
	}
	mtx.Unlock()
}

func main() {
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/save", saveHandler)

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("HTTP server error", err)
	}
}
*/

/* ТЕОРИЯ 20.10.2026 JSON
1.
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Payment struct {
	// Описание покупки
	Description string `json:"description"`

	// Сумма покупки
	USD int `json:"usd"`

	// ФИО человека, совершающего покупку
	FullName string `json:"fullName"`

	// Место прописки человека совершающего покупку
	Address string `json:"address"`

	// Время проведения оплаты
	Time time.Time
}

func (p Payment) Println() {
	fmt.Println("Description:", p.Description)
	fmt.Println("USD:", p.USD)
	fmt.Println("FullName:", p.FullName)
	fmt.Println("Address:", p.Address)
}

type HttpResponse struct {
	Money int
	PaymentHistory []Payment
}

func (p Payment) Validate() bool {
	if p.USD == 0 {
		return false
	}
}

var mtx = sync.Mutex{}
var money = 1000
var paymentHistory = make([]Payment, 0)

func payHandler(w http.ResponseWriter, r *http.Request) {
	var payment Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	payment.Time = time.Now()

	payment.Println()

	mtx.Lock()
	if money-payment.USD >= 0 {
		money -= payment.USD
	}

	paymentHistory = append(paymentHistory, payment)

	httpResponse := HttpResponse{
		Money: money,
		PaymentHistory: paymentHistory,
	}

	b, err := json.Marshal(httpResponse)
	if err != nil {
		fmt.Println("err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("err:",err)
		return
	}

	fmt.Println("money:", money)
	fmt.Println("payment histrory:", paymentHistory)
	mtx.Unlock()
}

func main() {
	http.HandleFunc("/pay", payHandler)


	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("Ошибка во время работы HTTP сервера:", err)
	}
}
*/

/* ПРАКТИКА 20.1O.2026 JSON
1.
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `json:"email"`
}

func userHeader(w http.ResponseWriter, r *http.Request) {
	user := User{
		Name: "Arzu",
		Age: 17,
		Email: "Arzu@gmail.com",
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}

func main() {
	http.HandleFunc("/userHeader", userHeader)

	fmt.Println("Запсу сервера: http://localhost:8080/userHeader")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("HTTP fail to server", err)
	}
}





2.
package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `json:"email"`
}

func userHeader(w http.ResponseWriter, r *http.Request) {
	var user User
	//             вывести данные в консоль.
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println("Значение JSON:", user)
}

func main() {
	http.HandleFunc("/userHeader", userHeader)

	fmt.Println("Запуск сервера: http://localhost:8080/userHeader")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("HTTP server error", err)
	}
}





3.
package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `json:"email"`
}

func userHeader(w http.ResponseWriter, r *http.Request) {
	var user User
	// JSON из Postman → User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// User → JSON в Postman
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func main() {
	http.HandleFunc("/userHeader", userHeader)

	fmt.Println("Запуск сервера: http://localhost:8080/userHeader")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("HTTP server error", err)
	}
}





4.
package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

// структура
type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `json:"email"`
}

func userHeader(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user.Age = 18

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&user)
}

func main() {
	http.HandleFunc("/userHeader", userHeader)

	fmt.Println("Запуск сервера: http://localhost:8080/userHeader")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("HTTP server error", err)
	}
}





5.
package main

import (
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `json:"email"`
}

func userHeader(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "только DELETE", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintln(w, "User deleted")
}

func main() {
	http.HandleFunc("/userHeader", userHeader)

	fmt.Println("Запуск сервера: http://localhost:8080/userHeader")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("HTTP server error", err)
	}
}
*/

/* ТЕОРИЯ 21.10.2026 HTTP Query параметры
1.
package main

import (
	"fmt"
	"net/http"
)

// localhost:9091/default?foo=x&boo=y

func handler(w http.ResponseWriter, r *http.Request) {
	fooParam := r.URL.Query().Get("foo")
	booParam := r.URL.Query().Get("boo")

	fmt.Println("foo param:", fooParam)
	fmt.Println("boo param:", booParam)
}

func main() {
	http.HandleFunc("/default:", handler)

	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("failed to run http server:", err)
	}
}
*/

/* ТЕОРИЯ 22.10.2026 REST API
ПРАВИЛА:

1. Определённый стил именования паттернов (эндпоинтов):
- Паттерны именуются существительными
- существительное - название ресурса, над которыми производится работа
1. Создавать новую задачу: Неверно - /createTask || Верно - /tasks
2. Получить информацию по уже созданной задаче: Неверно - /getTask || Верно - /tasks
3. Получить информацию по всем существующим задачам: Неверно - /getAllTask || Верно - /tasks
4. Получить список всех невыполненных задач: Неверно - /getUnfulfilledTasks || Верно - /tasks
5. Отметить Определённую задачу как выполненную: Неверно - /doneTask || Верно - /tasks
6. Удалить существующую задачу: Неверно - /deleteTask || Верно - /tasks

//

2. Сторого определённые HTTP-методы под каждый вид манипуляций над ресурсами
- Создание нового ресурса: POST
- Получение уже существующего ресурса: GET
- Частичное изменение существующего ресурса: PATCH
- Полное изменение существующего ресурса: PUT
- Удаление существующего ресурса: DELETE
1. Создать новую задачу: Неверно - random: GET, POST, PATCH, PUT, DELETE || Верно - POST
2. Получить информацию по уже созданной задаче: Неверно - random: GET, POST, PATCH, PUT, DELETE || Верно - GET
3. Получить информацию по всем существующим задачам: Неверно - random: GET, POST, PATCH, PUT, DELETE || Верно - GET
4. Получить список всех невыполненных задач: Неверно - random: GET, POST, PATCH, PUT, DELETE || Верно - GET
5. Отметить определённую задачу как выполненную: Неверно - random: GET, POST, PATCH, PUT, DELETE || Верно - PATCH
6. Удалить существующую задачу: Неверно - random: GET, POST, PATCH, PUT, DELETE || Верно - DELETE


//

3. Строго определённые места для передачи дополнительной информации в запросе
- Создание нового ресурса: JSON в теле запроса
- Получение одного уже существующего ресурса: Индентификатор ресурса зашит в паттерне
- Получение всех существующих ресурсов: Дополнительная информация не нужна
- Получение всех существующих ресурсов, учитывая фильтры: Query параметры
- Частичное изменение существующего ресурса: Индентификатор ресурса зашит в паттерне - JSON в теле запроса
- Полное изменение существующего ресурса: Индентификатор ресурса зашит в паттерне + JSON В теле запроса
- Удаление существующего ресурса; Индентификатор ресурса зашит в паттерне
1. Создать новую задачу: Неверно - || Верно - JSON в теле запроса: {"title":"title1", "description": "description1"}
   random:
   Тело запроса: "title"="title1","description"="description1"
   Query параметры: /tasks?title=title1&description=description1
   Защита в pattern: /tasks/title1/description1
2. Получить информацию по уже созданной задаче: Неверно - || Верно - Зашита в pattern: /tasks/title1
   random:
   Тело запроса: "title"="title1"
   Query параметры: /tasks?title=title1
   Зашита в pattern: /tasks/title1
3. Получить информацию по всем существующим задачам: Неверно - || Верно - дополнительная информация не нужна
   random:
   Тело запроса: "title"="ALL"
   Query параметры: /tasks?title=ALL
   Зашита в pattern: /tasks/ALL
4. Получить список всех невыполненных задач: Неверно - || Верно - Query параметры: /tasks?completed=true
   random:
   Тело запроса: "completed"=true
   Query параметры: /tasks?completed=true
   Зашита в pattern: /tasks/completed
5. Отметить определённую задачу как выполненную: Неверно - || Верно - Зашита pattern: /tasks/title1 + JSON в теле запроса {"completed":true}
   random:
   Тело запроса: "title"="title1", "completed"=true
   Query параметры: /tasks?title=title1&completed=true
   Зашита в pattern: /tasks/title1
6. Удалить существующую задачу: Неверно - || Верно - Зашита в pattern: /tasks/title1
   random:
   Тело запроса: "title"="title1"
   Query параметры: /tasks?title=title1
   Зашита в pattern: /tasks/title1

//

4. Строго определённая информация в HTTP теле ответа
   Успех:
- Создание нового ресурса: ---------------------------- JSON, представляющий созданный ресурс, в теле ответа
- Получение одного уже существующего ресурса: ----------- JSON, представляющий найденный ресурс, в теле ответа
- Получение всех существующих ресурсов: ---------------- JSON, представляющий все существующие ресурсы, в теле ответа
- Получние всех существующих ресурсов, учитовая фильтры: -- JSON, представляющий все найденные ресурсы, в теле ответа
- Частичное изменение существующего ресурса: ------------ JSON, представляющий новый изменённый ресурс, в теле ответа
- Полное изменение существующего ресурса: --------------- JSON, представляющий новый изменённый ресурс, в теле ответа
- Удаление существующего ресурса: --------------------- Пустое тело ответа

   Ошибка:
- Всегда: JSON, содержащий поле "message", которое описывает конкретный сбой + Опциональные поля по желанию и по договорённости




5. Строго определёная информация в HTTP статус-кода ответа
   Успех:
- Создание нового ресурса: --------------------------- 201 Created
- Получение одного уже существующего ресурса: ----------- 200 Ok
- Получение всех существующих ресурсов: --------------- 200 Ok
- Получение всех существующих ресурсов, учитывая фильтры: -- 200 Ok
- Частичное изменение существующего ресурса: -------------- 200 Ok
- Полное изменение существующего ресурса: --------------- 200 Ok
- Удаление существующего ресурса: ----------------------- 204 No Content

   Ошибка:
- Созданеие нового ресурса: --------------------------- 400 Bad Request ---------------- 409 Conflict -- 500 Internal Server Error
- Получение одного уже существующего ресурса: ----------- 400 Bad Request -- 404 Not Found ---------------- 500 Internal Server Error
- Получение всех существуюзих ресурсов: --------------- 400 Bad Request --------------------------- 500 Internal Server Error
- Получение всех существуюзих ресурсов, учитывая фильтры: -- 400 Bad Request --------------------------- 500 Internal Server Error
- Частичное изменение существующего ресурса: --------------- 400 Bad Request -- 404 Not Found-- 409 Conflict -- 500 Internal Server Error
- Полное изменение существующего ресурса: --------------- 400 Bad Request -- 404 Not Found-- 409 Conflict -- 500 Internal Server Error
- Удалепие существующего ресурса: ----------- 400 Bad Request -- 404 Not Found ---------------- 500 Internal Server Error


package main

import (
	"fmt"
	"restapi/http"
	"restapi/todo"
)

func main() {
	todoList := todo.NewList()
	httpHandlers := http.NewHTTPHandlers(todoList)
	httpServer := http.NewHTTPServer(httpHandlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed to start http server:", err)
	}
}

//
1. Как назвать паттерн?
2. Какой HTTP метод использовать?
3. Использовать тело запроса или query параметры?
4. Что возвращать в теле ответа?
N. ...

На все вопросы и ответы дает парадигма - REST API
//
*/

/* ПРАКТИКА 25.10.2026 REST API
1. GET - получить данные
package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)

// User - структура пользователя.
// Теги json говорят, под каким именем поля будут
// передаваться в json
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Email string `json:"email"`
}

// users - временное хранилище пользователей.
// Пока вместо базы данных используем обычный срез.
var users = []User{
	{Name: "Arzu", Age: 17, Email: "Arzu@gmail.com"},
	{Name: "Satoru", Age: 25, Email: "Satoru@gmail.com"},
	{Name: "Sukuna", Age: 30, Email: "Sukuna@gmail.com"},
}

// userHandler обрабатывает запрос к /users.
func userHandler(w http.ResponseWriter, r *http.Request) {

	// Проверяем HTTP метод.
	// Этот endpoint должен работать только с GET
	if r.Method != http.MethodGet {
		http.Error(w, "только GET", http.StatusMethodNotAllowed)
		return
	}

	// Говорим клиенту например Postman
	// Что сервер отправляет JSON
	w.Header().Set("Content-Type", "application/json")

	// Превращаем срез users в JSON
	// и отправляем его в ответ.
	json.NewEncoder(w).Encode(users)
}

func main() {
	http.HandleFunc("/userHandler", userHandler)

	fmt.Println("Запуск сервера: http://localhost:8080/userHandler")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("HTTP server error", err)
	}
}





2. POST - создать/добавить данные
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type User struct {
	Name string   `json:"name"`
	Age  int      `json:"age"`
	Email string  `json:"email"`
}

var users = []User{
	{Name: "Arzu", Age: 17, Email: "Arzu@gmail.com"},
	{Name: "Sukuna", Age: 30, Email: "Sukua@gmail.com"},
	{Name: "Gojo", Age: 29, Email: "Gojo@gmail.com"},
}

var mtx sync.Mutex

func userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "только POST", http.StatusMethodNotAllowed)
		return
	}

	var newUser User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "не удалось прочитать тело", http.StatusBadRequest)
		return
	}

	mtx.Lock()
	defer mtx.Unlock()

	users = append(users, newUser)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&newUser)
}

func main() {
	http.HandleFunc("/user", userHandler)

	fmt.Println("Запуск сервера: http://localhost:8080/user")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("HTTP server error", err)
	}
}





3. GET - получить данные
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type User struct {
	Name string   `json:"name"`
	Age  int      `json:"age"`
	Email string  `json:"email"`
	Id int        `json:"id"`
}

var users = []User{
	{Id: 1, Name: "Arzu", Age: 17, Email: "Arzu@gmail.com"},
	{Id: 2, Name: "Sukuna", Age: 30, Email: "Sukua@gmail.com"},
	{Id: 3, Name: "Gojo", Age: 29, Email: "Gojo@gmail.com"},
}

var mtx sync.Mutex

func userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "только GET", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "не удалось преоброзовать id из строки в int", http.StatusBadRequest)
		return
	}

	// Проверить, что такой индекс существует в срезе users
	// Внимание: если id — это индекс массива (0, 1, 2...), проверяем границы среза.
	for _, user := range users {
		if user.Id == id {
			w.Header().Set("Content-type", "application/json")
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	// Потому что 404 означает: пользователь не найден.
	http.Error(w, "Пользователь не найден", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/user", userHandler)

	fmt.Println("Запуск сервера: http://localhost:8080/user")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("HTTP server error", err)
	}
}





4. PUT - используется для изменения существующего ресурса.
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	Name string   `json:"name"`
	Age  int      `json:"age"`
	Email string  `json:"email"`
	Id int        `json:"id"`
}

var users = []User{
	{Name: "Arzu", Age: 17, Email: "Arzu@gmail.com", Id: 1},
	{Name: "Gojo", Age: 29, Email: "Gojo@gmail.com", Id: 2},
	{Name: "Sukuna", Age: 30, Email: "Sukuna@gmail.com", Id: 3},
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "только PUT", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "не удалось преоброзовать id из строки в int", http.StatusBadRequest)
		return
	}

	var newUser User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "не удалось прочитать тело", http.StatusBadRequest)
		return
	}

	for i := range users {
		if users[i].Id == id {
			newUser.Id = id
			users[i] = newUser

			w.Header().Set("Content-type", "application/json")
			json.NewEncoder(w).Encode(users[i])
			return
		}
	}

	http.Error(w, "Пользователь не найден", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/userHandler", userHandler)

	fmt.Println("Запуск сервера: http://localhost:8080/userHandler")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("HTTP server error", err)
	}
}





5. DELETE /user?id=... — удалить пользователя.
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	Name string   `json:"name"`
	Age  int      `json:"age"`
	Email string  `json:"email"`
	Id   int      `json:"id"`
}

var users = []User{
	{Name: "Arzu", Age: 17, Email: "Arzu@gmail.com", Id: 1},
	{Name: "Gojo", Age: 29, Email: "Gojo@gmail.com", Id: 2},
	{Name: "Satoru", Age: 17, Email: "Satoru@gmail.com", Id: 3},
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "только DELETE", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "не удалось преоброзовать id из строки в int", http.StatusBadRequest)
		return
	}

	for i, user := range users {
		if user.Id == id {
			fmt.Println("Пользователь найден", users[i])
			users = append(users[:i], users[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Пользователь удалён",
			})
			return
		}
	}

	http.Error(w, "Пользователь не найден", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/userHandler", userHandler)

	fmt.Println("Запуск сервера: http://localhost:8080/userHandler")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("HTTP server error", err)
	}
}





1. GET - Получить всех пользователей
package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)

type User struct {
	Id   int      `json:"id"`
	Name string   `json:"name"`
	Age  int      `json:"age"`
	Email string  `json:"email"`
}

var users = []User{
	{Id: 1, Name: "Arzu", Age: 17, Email: "Arzu@gmail.com"},
	{Id: 2, Name: "Gojo", Age: 29, Email: "Gojo@gmail.com"},
	{Id: 3, Name: "Satoru", Age: 30, Email: "Satoru@gmail.com"},
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "только GET", http.StatusMethodNotAllowed)
		return
	}

	// указать, что ответ — JSON:
	w.Header().Set("Content-Type", "application/json")
	// используй json.NewEncoder(w).Encode(...), чтобы отправить users клиенту.
	json.NewEncoder(w).Encode(&users)
}

func main() {
	http.HandleFunc("/userHandler", userHandler)

	fmt.Println("Запуск сервера: http://localhost:8080/userHandler")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("HTTP server error", err)
	}
}





2.POST — добавить нового пользователя.
package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type User struct {
	Id   int      `json:"id"`
	Name string   `json:"name"`
	Age  int      `json:"age"`
	Email string  `json:"email"`
}

var users = []User{
	{Id: 1, Name: "Arzu", Age: 17, Email: "Arzu@gmail.com"},
	{Id: 2, Name: "Gojo", Age: 29, Email: "Gojo@gmail.com"},
	{Id: 3, Name: "Satoru", Age: 18, Email: "Satoru@gmail.com"},
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "только POST", http.StatusMethodNotAllowed)
		return
	}

	var newUser User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "не удалось прочитать тело", http.StatusBadRequest)
		return
	}

	users = append(users, newUser)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

func main() {
	http.HandleFunc("/userHandler", userHandler)

	fmt.Println("Запуск сервера: http://localhost:8080/userHandler")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("HTTP server error", err)
	}
}





3. GET /user?id=... — получить одного пользователя.
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	Id   int      `json:"id"`
	Name string   `json:"name"`
	Age  int      `json:"age"`
	Email string  `json:"email"`
}

var users = []User{
	{Id: 1, Name: "Arzu", Age: 17, Email: "Arzu@gmail.com"},
	{Id: 2, Name: "Gojo", Age: 29, Email: "Gojo@gmail.com"},
	{Id: 3, Name: "Satoru", Age: 18, Email: "Satoru@gmail.com"},
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "только GET", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "не удалось преоброзовать id из строки в число", http.StatusBadRequest)
		return
	}

	// Проверить, что такой индекс существует в срезе users
	// Внимание: если id — это индекс массива (0, 1, 2...), проверяем границы среза.
	for _, user := range users {
		if user.Id == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.Error(w, "такого пользователя нет", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/userHandler", userHandler)

	fmt.Println("Запуск сервера: http://localhost:8080/userHandler")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("HTTP server error", err)
	}
}





4. PUT /user?id=... — изменить пользователя.
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	Id   int      `json:"id"`
	Name string   `json:"name"`
	Age  int      `json:"age"`
	Email string  `json:"email"`
}

var users = []User{
	{Id: 1, Name: "Arzu", Age: 17, Email: "Arzu@gmail.com"},
	{Id: 2, Name: "Gojo", Age: 29, Email: "Gojo@gmail.com"},
	{Id: 3, Name: "Satoru", Age: 18, Email: "Satoru@gmail.com"},
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "только PUT", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")

	// преоброзование id из строки в число через strconv.Atoi
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "не удалось преоброзовать id из строки в число", http.StatusBadRequest)
		return
	}

	var newUser User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "не удалось прочитать тело", http.StatusBadRequest)
		return
	}

	for i := range users {
		if users[i].Id == id {
			newUser.Id = id
			users[i] = newUser

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(&newUser)
			return
		}
	}
	http.Error(w, "не удалось найти пользователя", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/userHandler", userHandler)

	fmt.Println("Запуск сервера: http://localhost:8080/userHandler")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("HTTP server error", err)
	}
}





5. DELETE /user?id=... — удалить пользователя.
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	Id   int      `json:"id"`
	Name string   `json:"name"`
	Age  int      `json:"age"`
	Email string  `json:"email"`
}

var user = []User{
	{Id: 1, Name: "Arzu", Age: 17, Email: "Arzu@gmail.com"},
	{Id: 2, Name: "Gojo", Age: 29, Email: "Gojo@gmail.com"},
	{Id: 3, Name: "Satoru", Age: 18, Email: "Satoru@gmail.com"},
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "только DELETE", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")

	id , err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "не удалось преоброзовать строку в число", http.StatusBadRequest)
		return
	}

	for i := range user {
		if user[i].Id == id {
			user = append(user[:i], user[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Пользователь удален",
			})
			return
		}
	}

	http.Error(w, "Не удалось найти пользователя", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/userHandler", userHandler)

	fmt.Println("Запуск сервера: http://localhost:8080/userHandler")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("HTTP server error", err)
	}
}
*/

// 26.10.2026 3 КУРС - Теория Git
// 1. С помощью GitHub я могу свой проект сохранить в облачное хранилище
// GitHub и шанс на то что я потеряю свой проект будет гораздно меньше чем я осталю код на жетском диске на пк

// 2. Проблема версионирования кода.
// Ctrl + Z может сохранять некоторое количество кода

// 26.10.2026 Практика Git
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, Git!")
}