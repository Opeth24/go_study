package stringPractice

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func IsPalindrome() {
	data, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	dataRunes := []rune(data)

	left, right := 0, len(dataRunes)-1

	for left < right {
		if dataRunes[left] != dataRunes[right] {
			fmt.Print("Нет")
			return
		}
		left++
		right--
	}
	fmt.Print("Палиндром")

}

func FirstIndex() {
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')
	subData, _ := reader.ReadString('\n')
	data = strings.Trim(data, "\r\n")
	subData = strings.Trim(subData, "\r\n")

	dataRunes := []rune(data)
	subDataRunes := []rune(subData)
	if len(dataRunes) == 0 || len(subDataRunes) == 0 || len(dataRunes) < len(subDataRunes) {
		fmt.Print(-1)
		return
	}
	resultIdx := -1

	for i := 0; i <= len(dataRunes)-len(subDataRunes); i++ {
		match := true

		for j := 0; j < len(subDataRunes); j++ {
			if dataRunes[i+j] != subDataRunes[j] {
				match = false
				break
			}
		}
		if match {
			resultIdx = i
			break
		}
	}
	fmt.Println(resultIdx)
}

func OddIndex() {
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')
	fmt.Print(data)
	var result []rune

	for idx, elem := range data {
		fmt.Println(idx, elem)
		if idx%2 != 0 {
			result = append(result, elem)
		}
	}
	fmt.Print(string(result))
}

func RemoveRepeatedSybmols() {
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')
	data = strings.Trim(data, "\r\n")
	result := []rune{}

	counts := make([]int, 26)

	for _, elem := range data {
		elem = unicode.ToLower(elem)
		if elem >= 'a' && elem <= 'z' {
			idx := elem - 'a'
			counts[idx]++
		}
	}

	for _, elem := range data {
		elem = unicode.ToLower(elem)
		if elem >= 'a' && elem <= 'z' {
			idx := elem - 'a'
			if counts[idx] > 1 {
				continue
			}
			result = append(result, elem)
		}
	}
	fmt.Print(string(result))

}

func CheckPassword() {
	var password string
	fmt.Scan(&password)

	length := 0
	for _, elem := range password {
		fmt.Printf("%t\n", unicode.IsDigit(elem))
		fmt.Printf("%t\n", unicode.IsLetter(elem))
		fmt.Printf("%t\n", unicode.Is(unicode.Latin, elem))
		if !unicode.IsDigit(elem) && (!unicode.IsLetter(elem) || !unicode.Is(unicode.Latin, elem)) {
			fmt.Print("Wrong password")
			return
		}
		length++
	}

	if length < 5 {
		fmt.Print("Wrong password")
	} else {
		fmt.Print("Ok")
	}
}

func StringJoin() {
	var data string
	fmt.Scan(&data)
	var result strings.Builder
	runes := []rune(data)

	for idx, elem := range runes {
		result.WriteRune(elem)
		if idx == len(runes) - 1 {
			break
		} 
		result.WriteRune('*')
	}
	fmt.Print(result.String())

}

func SquarePrint() {
	var data string
	_, err := fmt.Scan(&data)
	if err != nil {
		panic("panic")
	}
	for _, elem := range data {
		currentInt, err := strconv.Atoi(string(elem))
		if err != nil {
			panic("panic")
		}
		fmt.Printf("%v", currentInt * currentInt)
	}

}

const(
	k = 10.0
	p = 20.0
	v = 30.0
)
func M() float64 {
	return 	p * v
}

func W() float64 {
	return math.Sqrt(k / M())
} 

func T() float64 {
	return 6 / W()
}
