package goroutines

import (
	"fmt"
	"sync"
	"time"
)

// ConcurrencyExample пример по работе с каналоми
func PipelineExample() {
	printValues(doubleFromChannel(generateInt()))
}

func WriterWithTwoReaderExample() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2) // добавляем 2 воркера

	go func() {
		for i := range 1_000 {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		defer wg.Done()
		for val := range ch {
			fmt.Printf("value = %d from worker 1\n", val)
		}
	}()

	go func() {
		defer wg.Done()
		for val := range ch {
			fmt.Printf("value = %d from worker 2\n", val)
		}
	}()

	wg.Wait()
	fmt.Println("ALL VALUES READ")
}

func TwoWriterReaderExample() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := range 2 {
			ch <- i
		}
	}()

	go func() {
		defer wg.Done()
		for i := range 2 {
			ch <- i * 2
		}
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	for val := range ch {
		fmt.Println(val)
	}
}

func generateInt() <-chan int {
	ch := make(chan int)

	go func() {
		for val := range 10 {
			ch <- val
		}
		close(ch)
	}()
	return ch
}

func doubleFromChannel(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for val := range in {
			time.Sleep(500 * time.Millisecond)
			out <- val * 2
		}
		close(out)
	}()

	return out
}

func printValues(in <-chan int) {
	for val := range in {
		fmt.Println(val)
	}
}
