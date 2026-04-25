package cycle

import "fmt"


func Cycle1() {
	var x int
    
    for fmt.Scan(&x); x <= 100; fmt.Scan(&x) {
        fmt.Println("x:", x)
        if x < 10 {
            continue
        }
        fmt.Println(x)
    }
}

func Cycle2() {
	var rub, p, y, years uint64
    fmt.Scan(&rub, &p, &y)
    
    years = 0
    for rub < y {
        rub = rub + rub * p / 100
        years++
    }
    fmt.Print(years)
}