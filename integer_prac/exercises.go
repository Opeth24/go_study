package integerprac

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func StrToInt() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	val := input.Text()
	number, _ := strconv.Atoi(val)
	fmt.Println(number * number)
}

func MetresToKm() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	val := input.Text()
	number, _ := strconv.Atoi(val)
	float_number := float64(number) / 1000.0
	fmt.Print(float_number)
}

func ToSixPow() {
	var x int32

	fmt.Scan(&x)

	square := x * x
	fourth := square * square
	sixth := fourth * square
	fmt.Print(sixth)
}

func ThirdMultiply() {
	var x, y, z int64

	fmt.Scan(&x, &y, &z)
	fmt.Print(x * y * z)
}

func Apple() {
	var n, k int32

	fmt.Scan(&n, &k)
	fmt.Print(k / n)
}

func GenerateNextInt() {
	var n int

	fmt.Scan(&n)
	fmt.Println(n)
	fmt.Println(n + 1)
	fmt.Println(n + 2)
}

func MultipleSumBy3() {
	var i, j, k, z int64
	fmt.Scan(&i, &j, &k, &z)
	sum := (i + j + k + z) * 3
	fmt.Print(sum)
}

func PieForParty() {
	var rub, cent, mult uint64

	fmt.Scan(&rub, &cent, &mult)

	rub *= mult
	cent *= mult

	rub += (cent / 100)
	cent = cent % 100

	fmt.Printf("%d %d", rub, cent)

}

func ParseMinutes() {
	var minutes uint64
	fmt.Scan(&minutes)

	fmt.Printf("%d мин - это %d час %d минут.", minutes, minutes/60, minutes%60)
}

func NextAndPrev() {
	var n int64
	fmt.Scan(&n)

	fmt.Printf("Следующее за числом %d число: %d\n", n, n+1)
	fmt.Printf("Для числа %d предыдущее число: %d", n, n-1)
}

func GetDigitSum() {
	var x int64
	fmt.Scan(&x)
	var sum int64

	sum += x / 100
	sum += x % 10
	sum += (x / 10) % 10

	fmt.Print(sum)
}

func ReverseDigit() {
	var x int64
	fmt.Scan(&x)
	var sum int64

	sum += x / 100
	sum += ((x / 10) % 10) * 10
	sum += (x % 10) * 100

	fmt.Print(sum)
}
