package array

import (
	"fmt"
)

func ThirdElem() {
	var n int
	fmt.Scan(&n)
	fmt.Print(n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println(arr[3])

}


func DeleteDigit() {

	var n, j, k int
    fmt.Scan(&n)
	fmt.Scan(&k)
    
    var digits []int
    for n > 0 {
		if n % 10 == k {
			n /= 10
			continue
		}
        digits = append(digits, n % 10)
        n /= 10
    }
    fmt.Println(digits)
	for i := len(digits) - 1; i >= 0; i-- {
		j = len(digits) - 1 - i
		if j >= i {
			break
		}
		digits[i], digits[j] = digits[j], digits[i]
	}
    fmt.Println(digits)
}