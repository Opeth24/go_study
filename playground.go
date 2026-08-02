package main

import (
	// array "example.com/go_study/array"
	// s "example.com/go_study/stringsPractice"
	// cycle "example.com/go_study/cycle"
	// structures "example.com/go_study/structures"
	// "example.com/go_study/condition"
	// h "example.com/go_study/handleerrors"

	// i "example.com/go_study/interfacePractice"
	"fmt"
	"time"
	// t "example.com/go_study/TimeWork"
	// g "example.com/go_study/goroutines"
	// l "example.com/go_study/linkedlist"
	"example.com/go_study/network_practise"
)

func main() {
	// hashmap.ParseCSV()

	// x := lambda.DeleteOddDigit(uint(72))
	// fmt.Print(x)

	// i.DataParse()
	// t.ParseDinner()
	// t.ConvertDuration()
	// t.EfficientDuration()
	// topKFrequent([]int{1, 1, 1, 2, 2,2, 3}, 2)

	// g.TwoWriterReaderExample()
	
	go networkpractise.Server()
	time.Sleep(100 * time.Millisecond)
	networkpractise.StartClient()

	
}

func topKFrequent(nums []int, k int) []int {
	counter := make(map[int]int, len(nums))
	for _, elem := range nums {
		counter[elem] += 1
	}
	frequence := make([][]int, len(nums))
	for elem, freq := range counter {
		fmt.Println(elem, freq)
		frequence[freq-1] = append(frequence[freq-1], elem)
	}
	result := make([]int, k)
	fmt.Printf("frequence = %v", frequence)

	idx := 0
	for i := len(frequence) - 1; i >= 0; i-- {
		current := frequence[i]
		fmt.Println(current, len(current))
		if len(current) == 0 {
			continue
		}
		for j, elem := range current {
			result[idx] = elem
			fmt.Println(j, elem)
			idx++
			if idx == k {
				return result
			}
		}
	}
	return result
}
