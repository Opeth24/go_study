package interfacepractice

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Давайте используем ваши знания структур, методов и интерфейсов на практике и реализуем объект, удовлетворяющий интерфейсу fmt.Stringer.
// Назовем его "Батарейка".

// Во-первых, вы должны объявить новый тип, удовлетворяющий интерфейсу fmt.Stringer.

// Ваш тип должен предусматривать, что на печати он будет выглядеть так: [      XXXX]: где пробелы - "опустошенная" емкость батареи, а X - "заряженная".

// Во-вторых, на стандартный ввод вы получаете строку, состоящую ровно из 10 цифр: 0 или 1 (порядок 0/1 случайный).
// Ваша задача считать эту строку любым возможным способом и создать на основе этой строки объект объявленного вами на первом этапе типа:
// надеюсь, вы понимаете, что строка символизирует емкость батарейки: 0 - это "опустошенная" часть, а 1 - "заряженная".

// В-третьих, созданный вами объект должен называться batteryForTest (использование этого имени обязательно).

// В вашем распоряжении фактически весь файл, НО завершающая фигурная скобка функции main() вам не видна, но она присутствует.
// Перед этой скобкой присутствует функция (которая вам тоже не видна), принимающая в качестве аргумента объект типа fmt.Stringer - batteryForTest,
// и направляющая его на стандартный вывод, поэтому вам не требуется выводить что-то на печать самостоятельно.

// Удачи!

type ToStringer interface {
	fmt.Stringer
}

type Battery struct {
	charges string
}

// func (b Battery) String() string {
// 	var result strings.Builder
// 	strToDraw := map[rune]string{
// 		'0': " ",
// 		'1': "X",
// 	}
// 	counter := map[rune]int{}
// 	result.WriteString("[")
// 	for _, r := range b.charges {
// 		counter[r] += 1
// 	}
// 	for _, key := range []rune{'0', '1'} {
// 		for range counter[key] {
// 			result.WriteString(strToDraw[key])
// 		}
// 	}
// 	result.WriteString("]")
// 	return result.String()

// }

func (b Battery) String() string {
	ones := strings.Count(b.charges, "1")
	zeros := len(b.charges) - ones

	var result strings.Builder
	result.Grow(len(b.charges) + 2)

	result.WriteByte('[')
	result.WriteString(strings.Repeat(" ", zeros))
	result.WriteString(strings.Repeat("X", ones))
	result.WriteByte(']')

	return result.String()
}

func Run() {
	reader := bufio.NewReader(os.Stdin)
	data, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error: %w", err)
	}
	data = strings.TrimSpace(data)
	batteryForTest := Battery{charges: data}
	fmt.Print(batteryForTest)
}

func ReaderExercise() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	result := 0

	for scanner.Scan() {
		text := scanner.Text()
		dataNum, err := strconv.Atoi(text)
		if err != nil {
			continue
		}

		result += dataNum
	}
	resultStr := strconv.Itoa(result)
	writer.WriteString(resultStr + "\n")
}

func isEmpty(row []string) bool {
	for _, cell := range row {
		cleanCell := strings.ReplaceAll(cell, "\x00", "")
		cleanCell = strings.TrimSpace(cleanCell)

		if cleanCell != "" {
			return false
		}
	}
	return true
}

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	r := csv.NewReader(file)

	line := 0
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		line += 1
		if isEmpty(row) {
			continue
		}
		if line == 5 {
			fmt.Println(row[2])
		}
	}
	return nil
}

func FindCSV() {
	const root = "./task"

	if err := filepath.Walk(root, walkFunc); err != nil {
		fmt.Print("error while walking")
	}
}

func BufferRead() {
	file, err := os.Open("./bufTask.txt")
	if err != nil {
		fmt.Print("error open: %w", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var delim byte = ';'
	counter := 0

	for {
		chunk, err := reader.ReadString(delim)
		chunk = strings.Trim(chunk, ";")
		counter++
		if len(chunk) == 1 && chunk == "0" {
			fmt.Printf("Ноль находится в позиции: %d", counter)
			break
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Ошибка чтения: %w", err)
			break
		}
	}
}

type user struct {
	Name     string
	Email    string
	Status   bool
	Language []byte
}

func UserToJson() {
	newUser := user{Name: "Anton", Email: "anton@mail.ru", Status: true, Language: []byte("ru")}

	data, err := json.Marshal(newUser)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", data)

	newUser.Language = []byte("en")
	err = json.Unmarshal(data, &newUser)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", newUser)
}

type student struct {
	LastName   string
	FirstName  string
	MiddleName string
	Birthday   string
	Address    string
	Phone      string
	Rating     []int
}

type studentGroup struct {
	ID       int
	Number   string
	Year     int
	Students []student
}

type responseAverage struct {
	Average float64
}

func MeanMarks() {
	// data, err := io.ReadAll(os.Stdin)
	// if err != nil {
	// 	panic(err)
	// }
	file, err := os.Open("K:/Projects/go_study/interfacePractice/text.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var jsonData studentGroup

	if err := json.Unmarshal(data, &jsonData); err != nil {
		fmt.Println(err)
		return
	}
	total := 0
	for _, stdt := range jsonData.Students {
		total += len(stdt.Rating)
	}
	mean := float64(total) / float64(len(jsonData.Students))

	rAvg := responseAverage{Average: mean}
	result, err := json.MarshalIndent(rAvg, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "%s", result)
}
