package main

import (
	"fmt"
	"regexp"
)

func main() {
	matched, _ := regexp.MatchString("^hello", "hello world")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("^hello", "no hello world")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("world$", "no hello world")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("\\.[a-z]+\\([0-9]+\\)$", "hello.world(12)")
	fmt.Println(matched) // true

	matched, _ = regexp.MatchString("\\.[a-z]+\\([0-9]+\\)$", "hello.world(12).com")
	fmt.Println(matched) // false
}
