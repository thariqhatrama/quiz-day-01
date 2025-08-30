package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("findDivisors(6): ", findDivisors(6))
	fmt.Println("findDivisors(24): ", findDivisors(24))
	fmt.Println("findDivisors(7): ", findDivisors(7))
	fmt.Println("")
	fmt.Println("extractDigits(12234): ", extractDigit(12234))
	fmt.Println("extractDigits(5432): ", extractDigit(5432))
	fmt.Println("extractDigits(1278): ", extractDigit(1278))
	fmt.Println("extractDigits(0): ", extractDigit(0))
	fmt.Println("")
	// cetak pola bintang
	fmt.Println("Pattern 1:")
	printPattern1(5)

	fmt.Println("Pattern 2:")
	printPattern2(5)

	fmt.Println("Pattern 3:")
	printPattern3(5)
	fmt.Println("")

	// var rows int
	// fmt.Print("Masukkan jumlah baris piramid angka: ")
	// fmt.Scan(&rows)

	// fmt.Println("=== Piramid Angka ===")
	// printNumberPyramid(rows)

	fmt.Println(reverse("ABCD"))
	fmt.Println(reverse("tamaT"))
	fmt.Println(reverse("XYnb"))

	fmt.Println("")
	fmt.Println(checkBraces("(())"))
	fmt.Println(checkBraces("()()"))
	fmt.Println(checkBraces("((()"))
	fmt.Println(checkBraces("(()))((())"))

	fmt.Println("")
	fmt.Println(upperCaseExcept([]string{"code", "java", "cool"}, "java"))
	fmt.Println(upperCaseExcept([]string{"black", "pink", "venom"}, "venom"))

	fmt.Println("")
	fmt.Println(findMinMax([]int{2, 3, 4, 5, 6, 7, 8, 9, 1, 10}))
}

func findDivisors(n int) []int {
	var divisors []int
	for i := 1; i < n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}

func extractDigit(n int) []int {
	var digits []int
	for n > 0 {
		digit := n % 10
		digits = append(digits, digit)
		n /= 10
	}
	return digits
}

func printPattern1(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if j >= i {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

func printPattern2(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if j > n-i {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

func printPattern3(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if j <= i {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

func reverse(s string) string {
	rns := []rune(s) // ubah string jadi slice rune agar support UTF-8
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i] // tukar posisi karakter
	}
	return string(rns)
}

func checkBraces(s string) bool {
	braces := 0
	for _, ch := range s {
		if ch == '(' {
			braces++
		} else if ch == ')' {
			braces--
			if braces < 0 {
				return false
			}
		}
	}
	return braces == 0
}

// challenges day 2

func upperCaseExcept(words []string, except string) []string {
	result := []string{}
	for _, word := range words {
		if word == except {
			result = append(result, word)
		} else {
			result = append(result, strings.ToUpper(word))
		}
	}
	return result
}

func findMinMax(nums []int) [2]int {
	if len(nums) == 0 {
		return [2]int{0, 0} // default kalau kosong
	}

	min := nums[0]
	max := nums[0]

	for _, n := range nums {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	return [2]int{min, max}
}

// func printNumberPyramid(n int) {
// 	for i := n; i >= 1; i-- {
// 		// susun satu baris: turun i..1 lalu naik 2..i
// 		row := make([]int, 0, 2*i-1)
// 		for j := i; j >= 1; j-- {
// 			row = append(row, j)
// 		}
// 		for j := 2; j <= i; j++ {
// 			row = append(row, j)
// 		}

// 		// cetak baris dengan pemisah tab, tanpa tab terakhir
// 		for k, v := range row {
// 			if k > 0 {
// 				fmt.Print("\t")
// 			}
// 			fmt.Print(v)
// 		}
// 		fmt.Println()
// 	}
// }
