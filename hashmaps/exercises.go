package hashmaps

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func HashmapTest() {
	s := []map[int]string{}
	fmt.Print(s)
}

func MemoSimple() {
	cache := map[int]int{}
	reader := bufio.NewReader(os.Stdin)
	data, err := reader.ReadString('\n')
	data = strings.Trim(data, "\r\n")
	if err != nil {
		panic("read string panic")
	}

	dataArr := strings.Split(data, " ")
	for _, elem := range dataArr {
		fmt.Println(elem)
		key, err := strconv.Atoi(string(elem))
		if err != nil {
			panic("convert panic")
		}

		if _, ok := cache[key]; !ok {
			cache[key] = work(key)
		}

		fmt.Print(cache[key], " ")
	}
}

func CityPopulation() {
	groupCity := map[int][]string{
		10:  {},
		100: {"Moscow", "Ekaterinburg", "Saint-Petersburg"},
	}
	population := map[string]int{
		"Ekaterinburg": 100,
		"Tashkent":     99,
		"Vladivostok":  58,
	}

	cityKeys := map[string]struct{}{}
	for _, city := range groupCity[100] {
		cityKeys[city] = struct{}{}
	}

	for key := range population {
		if _, ok := cityKeys[key]; !ok {
			delete(population, key)
		}
	}
	fmt.Println(population)

	// newPopulation := map[string]int{}
	// for _, city := range groupCity[100] {
	// 	newPopulation[city] = 100
	// }
	// population = newPopulation
	// fmt.Println(population)
}

func work(x int) int {
	return x * 10
}

func sanitazeString(s string) int64 {
	var res []rune
	for _, r := range s {
		if unicode.IsDigit(r) {
			res = append(res, r)
		}
	}
	num, err := strconv.ParseInt(string(res), 10, 64)
	if err != nil {
		return 0
	}
	return num
}

func AddingFirst(x, y string) int64 {
	xRunes := []rune(x)
	yRunes := []rune(y)

	temp := []rune{}
	for _, elem := range xRunes {
		if unicode.IsDigit(elem) {
			temp = append(temp, elem)
		}
	}
	xInt, err := strconv.Atoi(string(temp))
	if err != nil {
		panic("can't convert elem x")
	}

	temp = []rune{}
	for _, elem := range yRunes {
		if unicode.IsDigit(elem) {
			temp = append(temp, elem)
		}
	}
	yInt, err := strconv.Atoi(string(temp))
	if err != nil {
		panic("can't convert elem y")
	}
	return int64(xInt + yInt)
}

func Adding(x, y string) int64 {
	return sanitazeString(x) + sanitazeString(y)
}

func ParseCSV() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	data := scanner.Text()
	parts := strings.Split(data, ";")
	if len(parts) < 2 {
		return
	}
	var nums [2]float64
	for i := range 2 {
		parts[i] = strings.ReplaceAll(parts[i], " ", "")
		parts[i] = strings.ReplaceAll(parts[i], ",", ".")
		val, err := strconv.ParseFloat(parts[i], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "parse error: %v\n", err)
			return
		}
		nums[i] = val
	}
	
	if nums[1] != 0 {
		fmt.Printf("%.4f\n", nums[0]/nums[1])
	}
}
