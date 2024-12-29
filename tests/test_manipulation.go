package main

import (
	"fmt"
	"github.com/aquaticcalf/strings-attached"
)

func test_reverse() {
	fmt.Println("-------------")
	fmt.Println("testing Reverse...")
	fmt.Println("\"hello\" -> \"olleh\" :",stringsattached.Reverse("hello") == "olleh")
	fmt.Println("\"\" -> \"\" :",stringsattached.Reverse("") == "")
	fmt.Println("\"racecar\" -> \"racecar\" :",stringsattached.Reverse("racecar") == "racecar")
}

func test_is_palindrome() {
	fmt.Println("-------------")
	fmt.Println("testing Is_palindrome...")
	fmt.Println("\"A man, a plan, a canal, Panama\" -> palindrome :",stringsattached.Is_palindrome("A man, a plan, a canal, Panama") == true)
	fmt.Println("\"hello\" -> not a palindrome :",stringsattached.Is_palindrome("hello") == false)
	fmt.Println("\"\" -> palindrome :",stringsattached.Is_palindrome("") == true)

}

func main() {
	fmt.Println("running tests")
	test_reverse()
	test_is_palindrome()
}