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


func SymmetricNum() {
	    var a int64
    fmt.Scan(&a)
    
    if ((a % 10) == (a / 1000)) && (((a / 100) % 10) == (a / 10) % 10) {
        fmt.Print("YES")
    } else {
        fmt.Print("NO")
    }
}

func Homework() {
	var k, m int64
    fmt.Scan(&k, &m)

	fmt.Print(k / 2)
    
    fmt.Print(math.Ceil(float64(k) / float64(m)))
}

func InterestingNum(){
	var x int64
    fmt.Scan(&x)
    a := x / 100
    b := x / 10 % 10
    c := x % 10
    
    maxValue := max(a, b, c)
	minValue := min(a, b, c)
	midValue := a + b + c - minValue - maxValue
    if float64(midValue) == math.Abs(float64(maxValue) - float64(minValue)) {
        fmt.Print("Число интересное")
    } else {
        fmt.Print("Число неинтересное")
    }
}

func FootballCheck() {
	var age uint8
    var sex string
    fmt.Scan(&age, &sex)
    
    if sex == "m" && age >= 12 && age <= 18 {
        fmt.Print("YES")
    } else {
        fmt.Print("NO")
    }
}

func BicycleRent() {
	var rent, freeMinutes, duration, tax uint64
    fmt.Scan(&rent, &freeMinutes, &duration, &tax)
    
    var overLimit uint64 = 0 
    if duration > freeMinutes{
        overLimit = duration - freeMinutes
    }
    total := rent + tax * overLimit
    fmt.Print(total)
}