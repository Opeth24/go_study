package condition

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func AgeChecker() {
	var age int64
	fmt.Scan(&age)

	if age >= 12 {
		fmt.Print("Доступ разрешен")
	} else {
		fmt.Print("Доступ запрещен")
	}

}

func PasswordChecker() {
	var password, accept string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	password = scanner.Text()

	scanner.Scan()
	accept = scanner.Text()

	if password == accept {
		fmt.Print("Пароль принят")
	} else {
		fmt.Print("Пароль не принят")
	}
}

func ThreeDigit() {
	var input int64

	fmt.Scan(&input)

	if input > 99 && input < 1000 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
		return
	}
}

func HappyTicket() {
	var input int64

	fmt.Scan(&input)
	first := input / 1000
	second := input % 1000

	first_part_acc := (first / 100) + ((first / 10) % 10) + (first % 10)
	second_part_acc := (second / 100) + ((second / 10) % 10) + (second % 10)

	if first_part_acc == second_part_acc {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}

func DoesBishopTake() {
	var i, j, i1, j1 int

	fmt.Scan(&i, &j, &i1, &j1)

	if math.Abs(float64(i-i1)) == math.Abs(float64(j-j1)) {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}

func Calculator() {
	var x, y int64
	var sign string

	fmt.Scan(&x, &y, &sign)

	if sign == "+" {
		fmt.Print(x + y)
	} else if sign == "-" {
		fmt.Print(x - y)
	} else if sign == "/" {
		if y == 0 {
			fmt.Print("На ноль делить нельзя!")
		} else {
			fmt.Print(float64(x) / float64(y))
		}
	} else if sign == "*" {
		fmt.Print(x * y)
	} else {
		fmt.Print("Неверная операция")
	}
}

func QuadraticEquation() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	d := math.Pow(b, 2) - 4.0 * a * c

	if d < 0 {
		return
	}

	if d > 0 {
		x1 := (-b + math.Sqrt(d)) / (2 * a)
		x2 := (-b - math.Sqrt(d)) / (2 * a)
		fmt.Println(min(x1, x2))
		fmt.Print(max(x1, x2))
	} else {
		x := -b / (2 * a)
		fmt.Print(x)
	}
}