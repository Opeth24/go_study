package main

import (
	"bufio"
	"fmt"
	"os"
)

func Scanner_learn() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	first := scanner.Text()

	scanner.Scan()
	second := scanner.Text()

	scanner.Scan()
	third := scanner.Text()

	fmt.Println(third)
	fmt.Println(second)
	fmt.Println(first)


	var s1 string
	var s2 string
	var s3 string
	countWords, _ := fmt.Scan(&s1, &s2, &s3)
	fmt.Println(countWords)
	fmt.Print(s1, s2, s3)

}