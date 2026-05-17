package timework

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func ParseDinner() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	d := scanner.Text()
	result, err := time.Parse(time.DateTime, d)
	if err != nil {
		fmt.Println(err)
	}

	if result.Hour() > 13 {
		result = result.Add(time.Hour * 24)
		fmt.Print(result.Format(time.DateTime))
	} else {
		fmt.Print(result.Format(time.DateTime))
	}
}
