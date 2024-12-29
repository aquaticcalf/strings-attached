package main

import (
	"fmt"
	"github.com/aquaticcalf/strings-attached"
)

func test_reverse() {
	fmt.Println("testing Reverse...")
	fmt.Println("\"hello\" -> \"olleh\" :",stringsattached.Reverse("hello") == "olleh")
	fmt.Println("\"\" -> \"\" :",stringsattached.Reverse("") == "")
	fmt.Println("\"racecar\" -> \"racecar\" :",stringsattached.Reverse("racecar") == "racecar")
}

func main() {
	fmt.Println("running tests")
	fmt.Println("-------------")
	test_reverse()
}