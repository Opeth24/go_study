package floatprac

import (
	"fmt"
	"math"
)

func CircleSquare() {
	var r float64
	const pi = 3.14

	fmt.Scan(&r)
	fmt.Print(pi * r * r)
}

func Mean() {
	var a, b float64

	fmt.Scan(&a, &b)
	fmt.Print((a + b) / 2)
}

func FindExp() {
	var a float64

	fmt.Scan(&a)

	fmt.Print((a - float64(int(a))))
}

func BitToKb() {
	var a float64
	fmt.Scan(&a)

	fmt.Print(a / math.Pow(2.0, 13.0))

}

func Pifagor() {
	var a, b float64
	fmt.Scan(&a, &b)

	c := math.Sqrt(math.Pow(a, 2.0) + math.Pow(b, 2.0))
	fmt.Print(a + b + c)
}

func MaxDigit() {
	var a int
	var res int

	fmt.Scan(&a)
	a = int(math.Abs(float64(a)))

	res = max(a / 100, a % 10, a / 10 % 10)
	fmt.Print(res)
}

func CalculateDiscount() {
	var price, total float64
	var discount uint16

	fmt.Scan(&price, &discount)
	total = price - (price / 100 * float64(discount))

	fmt.Print(total)
}


func IfCondition() {
    x := 10 // Внешняя переменная 'x'
    fmt.Println("Внешняя переменная x:", x)

    if true {
        x := 5 // Затеняющая переменная 'x'
        fmt.Println("Внутренняя переменная x:", x)
    }

    fmt.Println("Внешняя переменная x после внутренней области видимости:", x)
}
