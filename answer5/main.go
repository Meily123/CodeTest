package main

import (
	"fmt"
)

/*
Question :
Buatkan sebuah function untuk mengecek sebuah kata palindrom atau bukan dengan
golang
*/

func isPalindrome(kata string) bool {
	// Case Sensitive Function
	// uncomment line below to make this function case insensitive
	// kata = strings.ToLower(kata)
	panjangKata := len(kata)
	for indeks := 0; indeks < panjangKata/2; indeks++ {
		if kata[indeks] != kata[panjangKata-1-indeks] {
			return false
		}
	}
	return true
}

func main() {
	// test case
	fmt.Println(isPalindrome("kodok"))
	fmt.Println(isPalindrome("KodOk"))
	fmt.Println(isPalindrome("meja"))
}
