package timework

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func FindTwoDiff() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	d := scanner.Text()
	datesSlice := strings.Split(d, ",")
	firstDate, err := time.Parse("02.01.2006 15:04:05", datesSlice[0])
	if err != nil {
		fmt.Print(err)
	}
	secondDate, err := time.Parse("02.01.2006 15:04:05", datesSlice[1])
	if err != nil {
		fmt.Print(err)
	}

	if secondDate.Before(firstDate) {
		firstDate, secondDate = secondDate, firstDate
	}
	fmt.Print(secondDate.Sub(firstDate))
}


func ConvertDuration() {
	const now = 1589570165
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	rawDuration := scanner.Text()
	minutes := strings.Split(rawDuration, "мин.")[0]
	minutes = strings.TrimSpace(minutes)
	var seconds string
	seconds = strings.Split(rawDuration, "сек.")[0]
	seconds = strings.Split(seconds, "мин.")[1]
	seconds = strings.TrimSpace(seconds)
	secondsInt, err := strconv.Atoi(seconds)
	if err != nil {
		fmt.Println(err)
	}
	minutesInt, err := strconv.Atoi(minutes)
	if err != nil {
		fmt.Println(err)
	}

	toTime := time.Unix(now, 0).UTC()
	fmt.Println(toTime)
	nextTime := toTime.Add(time.Second * time.Duration(secondsInt))
	nextTime = nextTime.Add(time.Minute * time.Duration(minutesInt))
	fmt.Println(nextTime.Format(time.UnixDate))
	
}

func EfficientDuration() {
	const now = 1589570165
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	rawDuration := scanner.Text()
	replacer := strings.NewReplacer(" мин.", "m", " сек.", "s", " ", "")
	durStr := replacer.Replace(rawDuration)
	duration, _ := time.ParseDuration(durStr)

	resultTime := time.Unix(now, 0).UTC().Add(duration)
	fmt.Println(resultTime.Format(time.UnixDate))
}