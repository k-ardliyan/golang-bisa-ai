package lib

import (
	"fmt"
	"strconv"
)

func GetInput() (float64, error) {
	var input string
	fmt.Scan(&input)
	value, err := strconv.ParseFloat(input, 64)
	return value, err
}
