package main

import (
	"bufio"
	"fmt"
	"os"
)

func Scanner_exercise()  {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	delimeter := input.Text()

	input.Scan()
	first := input.Text()
	
	input.Scan()
	second:= input.Text()

	input.Scan()
	third := input.Text()

	fmt.Print(first,delimeter,second,delimeter,third)
	
}