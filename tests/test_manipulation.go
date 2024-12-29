package main

import (
	"fmt"
	sa "github.com/aquaticcalf/strings-attached"
)

func test_reverse() {
	fmt.Println("-------------")
	fmt.Println("testing Reverse...")
	fmt.Println("\"hello\" -> \"olleh\" :",sa.Reverse("hello") == "olleh")
	fmt.Println("\"\" -> \"\" :",sa.Reverse("") == "")
	fmt.Println("\"racecar\" -> \"racecar\" :",sa.Reverse("racecar") == "racecar")
}

func test_is_palindrome() {
	fmt.Println("-------------")
	fmt.Println("testing Is_palindrome...")
	fmt.Println("\"A man, a plan, a canal, Panama\" -> palindrome :",sa.Is_palindrome("A man, a plan, a canal, Panama") == true)
	fmt.Println("\"hello\" -> not a palindrome :",sa.Is_palindrome("hello") == false)
	fmt.Println("\"\" -> palindrome :",sa.Is_palindrome("") == true)
}

func test_truncate() {
	fmt.Println("-------------")
	fmt.Println("testing Truncate...")
	fmt.Println("\"hello world\" -> \"he...\" :",sa.Truncate("hello world", 5) == "he...")
	fmt.Println("\"short\" -> \"short\" :",sa.Truncate("short", 10) == "short")
}

func test_word_count() {
	fmt.Println("-------------")
	fmt.Println("testing Word_count...")
	fmt.Println("\"hello world\" -> 2 :",sa.Word_count("hello world") == 2)
	fmt.Println("\"  multiple   spaces  \" -> 2 :",sa.Word_count("  multiple   spaces  ") == 2)
	fmt.Println("\"\" -> 0 :",sa.Word_count("") == 0)
}

func main() {
	fmt.Println("running tests")
	test_reverse()
	test_is_palindrome()
	test_truncate()
	test_word_count()
}