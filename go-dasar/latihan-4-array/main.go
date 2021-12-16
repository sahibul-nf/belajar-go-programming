package main

import (
	"errors"
	"fmt"
)

func main() {

	scores := []int{100, 10, 5, 4, 40, 2}
	fmt.Println(sum(scores))

	fmt.Println()

	result, err := calculation(10, 2, "=")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result)

}

// hitung jumlah nilai array
func sum(scores []int) (sum int) {
	for _, score := range scores {
		// fmt.Println(score)
		sum += score
	}

	return
}

func calculation(value, value1 int, operator string) (result int, err error) {

	switch operator {
	case "+":
		result = value + value1
	case "*":
		result = value * value1
	case "-":
		result = value - value1
	case "/":
		result = value / value1
	default:
		err = errors.New("Error")
	}

	return
}
