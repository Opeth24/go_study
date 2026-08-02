package goroutines

import (
	"sync"
)

// Необходимо написать функцию func merge2Channels(fn func(int) int, in1 <-chan int, in2 <- chan int, out chan<- int, n int).

// Описание ее работы:

// n раз сделать следующее

//     прочитать по одному числу из каждого из двух каналов in1 и in2, назовем их x1 и x2.
//     вычислить f(x1) + f(x2)
//     записать полученное значение в out

// Функция merge2Channels должна быть неблокирующей, сразу возвращая управление.

// Функция fn может работать долгое время, ожидая чего-либо или производя вычисления.

// Формат ввода:

//     количество итераций передается через аргумент n.
//     целые числа подаются через аргументы-каналы in1 и in2.
//     функция для обработки чисел перед сложением передается через аргумент fn.

// Формат вывода:

//     канал для вывода результатов передается через аргумент out.

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	// порядок может быть нарушен и в i индекс массива запишется разные значения
	res1 := make([]int, n, n)
	res2 := make([]int, n, n)
	var wg sync.WaitGroup

	for i := range n {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			val := <- in1
			res1[idx] = fn(val)
		}(i)
		
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			val := <- in2
			res2[idx] = fn(val)
		}(i)
	}

	go func() {
		wg.Wait()
		for i := range n {
			out <- res1[i] + res2[i]
		}
	}()
}


func merge2Channels2(
	fn func(int) int,
	in1 <-chan int,
	in2 <-chan int,
	out chan<- int,
	n int,
) {
	go func() {
		res := make([]int, n)

		var wg sync.WaitGroup

		for i := 0; i < n; i++ {
			x1 := <-in1
			x2 := <-in2

			wg.Add(1)
			go func(idx, a, b int) {
				defer wg.Done()

				var r1, r2 int
				var calc sync.WaitGroup
				calc.Add(2)

				go func() {
					defer calc.Done()
					r1 = fn(a)
				}()

				go func() {
					defer calc.Done()
					r2 = fn(b)
				}()

				calc.Wait()
				res[idx] = r1 + r2
			}(i, x1, x2)
		}

		wg.Wait()

		for _, v := range res {
			out <- v
		}
	}()
}