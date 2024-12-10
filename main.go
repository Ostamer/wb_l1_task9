package main

import (
	"fmt"
)

func main() {
	// Исходный массив чисел
	square_numbers := []int{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}

	// Канал для исходных чисел
	inputChannel := make(chan int)

	// Канал для результатов операции x * 2
	outputChannel := make(chan int)

	// Горутина для записи чисел в inputChannel
	go func() {
		for _, num := range square_numbers {
			inputChannel <- num
		}
		close(inputChannel)
	}()

	// Горутина для выполнения операции и записи результата в outputChannel
	go func() {
		for num := range inputChannel {
			outputChannel <- num * 2
		}
		close(outputChannel)
	}()

	// Цикл для вывода чисел из outputChannel
	for result := range outputChannel {
		fmt.Println(result)
	}
}
