package main

import (
	"bufio"
	"fmt"
	"os"
)

func HelloGo() {
	fmt.Print("Input your name: ")
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	name := scanner.Text()

	fmt.Printf("Hello, %s!\n", name)
}
