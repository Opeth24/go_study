package lambdafunctions

import "strconv"

func DeleteOddDigit(digit uint) uint {
	// fn := func(x uint) uint {
	// 	xStr := strconv.FormatUint(uint64(x), 10)
	// 	var result []rune
	// 	for _, num := range xStr {
	// 		if d, _ := strconv.Atoi(string(num)); d % 2 == 0 && d != 0 {
	// 			result = append(result, num)
	// 		}
	// 	}

	// 	y := string(result)
	// 	if y == "" {
	// 		y = "100"
	// 	}
	// 	r, err := strconv.ParseUint(y, 10, 32)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	if r == 0 {
	// 		r = 100
	// 	}
	// 	return uint(r)
	// }

	optimizeFn := func(x uint) uint {
		xStr := strconv.FormatUint(uint64(x), 10)
		var result []rune
		for _, r := range xStr {
			digit := uint64(r - '0')
			if digit != 0 && digit%2 == 0 {
				result = append(result, r)
			}
		}
		if len(result) == 0 {
			return 100
		}
		res, _ := strconv.ParseUint(string(result), 10, 64)
		return uint(res)
	}

	return optimizeFn(digit)
}
