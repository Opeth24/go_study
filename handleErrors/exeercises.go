package handleerrors

import "fmt"


func HandleErrors() {
	var a, b int
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Print("ошибка")
		return
	}
	_, err = fmt.Scan(&b)
	if err != nil {
		fmt.Print("ошибка")
		return
	}
	result, err := divide(a, b)
	if err != nil {
		fmt.Print("ошибка")
	}
	fmt.Print(result)
}

func divide(a, b int) (int, error) {
	result := a / b
	return result, nil
}