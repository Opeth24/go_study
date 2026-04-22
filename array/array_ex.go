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
