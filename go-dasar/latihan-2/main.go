package main

import "fmt"

func main() {

	sentence := "Golang is the best language"

	// cetak character jika index bil genap dari sebuah kalimat
	for i, letter := range sentence {
		if i%2 == 0 {
			fmt.Println(string(letter))
		}
	}

	fmt.Println()

	// cetak character vocal saja [a, i, u, e, o] dari sebuah kalimat
	for _, letter := range sentence {
		letterS := string(letter)

		if letterS == "o" || letterS == "a" || letterS == "i" || letterS == "e" || letterS == "u" {
			fmt.Println(string(letter))
		}
	}
}
