package cycle

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func HelloFor() {
	for i := 0; i < 3; i++ {
		fmt.Println("BB")
	}
	fmt.Println("CCC")

	for i := 0; i < 4; i++ {
		fmt.Println("DDDD")
	}
	fmt.Println("E")
	for i := 0; i < 5; i++ {
		fmt.Println("FFFFF")
	}
	fmt.Println("GG")
}


func RepeatString() {
    var text string
    var count int
    
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    text = scanner.Text()
    
    scanner.Scan()
    count, _ = strconv.Atoi(scanner.Text())
    
    for range count {
        fmt.Println(text)
    }
}

func CountMax() {
	// var x, best, curr uint64
    
    // fmt.Scan(&x)
    // curr = x
    // best = 1
    // for x != 0 {
    //     fmt.Scan(&x)
    //     if x > curr {
    //         curr = x
    //         best = 1
    //     } else if x == curr {
    //         best += 1
    //     }
    // }

	var a, max, count int  
   
    for fmt.Scan(&a); a != 0 ; fmt.Scan(&a){ 
        switch {
            case a > max:
            max, count = a, 1
            case a == max:
            count++
        }
    }
    fmt.Println(count)
}
