package hashmaps

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func HashmapTest() {
	var s = []map[int]string{}
	fmt.Print(s)
}


func MemoSimple() {
	var cache = map[int]int{}
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
		10: {},
		100: {"Moscow", "Ekaterinburg", "Saint-Petersburg"},
	}
	population := map[string]int{
		"Ekaterinburg": 100,
		"Tashkent": 99,
		"Vladivostok": 58,
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