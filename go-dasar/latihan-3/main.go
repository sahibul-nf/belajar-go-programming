package main

import "fmt"

func main() {

	scores := [6]int{100, 10, 5, 4, 40, 2}

	// hitung nilai rata-rata
	var hitung int

	for _, score := range scores {
		// fmt.Println(score)
		hitung += score
	}

	total := float32(hitung) / float32(len(scores))
	fmt.Println(total)

	fmt.Println()

	// masukkan nilai <= 10 untuk dimasukkan ke slice
	var badScore []int

	for _, score := range scores {
		if score <= 10 {
			badScore = append(badScore, score)
		}
	}

	fmt.Println(badScore)
}
