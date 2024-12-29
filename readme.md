## strings attached

> trying to impose order upon the chaos of character squences we call "strings"


### why? ( the external question )

you might ask, "why another string manipulation library?"

and in that question lies the beautiful absurdity of our craft. 

just as sisyphus was destined to roll his boulder up the hill for eternity,

we developers are bound to keep creating new string utilities,

each time convinced that *this* time we'll get it right.

### features

the `stringsattached` package provides a set of useful string manipulation functions:

- Reverse: returns the reverse of a string (e.g., "abcdefgh" -> "hgfedcba")
- Is_palindrome: returns a boolean indicating whether a string is a palindrome, ignoring non-alphanumeric characters and case
- Truncate: cuts a string to a specified length and adds an ellipsis if necessary
- Word_count: counts the number of words in a string

### usage

to use the `stringsattached` package, simply import it in your go program:

```go
package main

import (
    "fmt"
    sa "github.com/yourusername/stringsattached"
)

func main() {
    fmt.Println(sa.Reverse("hello")) // output: "olleh"
    fmt.Println(sa.Is_palindrome("A man, a plan, a canal: Panama")) // output: true
    fmt.Println(sa.Truncate("This is a very long string", 10)) // output: "this is a..."
    fmt.Println(sa.Word_count("Hello world, this is a test")) // output: 6
}
```

### documentation

**Reverse**
```go
func Reverse(s string) string
```
returns the reverse of a string.

- parameters: s (string) - the input string
- returns: (string) - the reversed string

**Is_palindrome**
```go
func Is_palindrome(s string) bool
```
returns a boolean indicating whether a string is a palindrome, ignoring non-alphanumeric characters and case.

- parameters: s (string) - the input string
- returns: (bool) - true if the string is a palindrome, false otherwise

**Truncate**
```go
func Truncate(s string, max_length int) string
```
cuts a string to a specified length and adds an ellipsis if necessary.

- parameters:
  - s (string) - the input string
  - max_length (int) - the maximum length of the output string
- returns: (string) - the truncated string

**Word_count**
```go
func Word_count(s string) int
```
counts the number of words in a string.

- parameters: s (string) - the input string
- returns: (int) - the number of words in the string

### license

the `stringsattached` package is licensed under the mit license. See [`license.md`](license.md) for details.

### installation

to install the `stringsattached` package, run the following command:
```sh
go get github.com/aquaticcalf/stringsattached
```