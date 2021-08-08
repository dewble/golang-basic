package main

import (
	"fmt"
	"regexp"
)

func main() {

	// 정규표현식에 해당하는 문자열을 지정된 문자열로 바꿈
	re1, _ := regexp.Compile("\\.[a-zA-Z]+$")
	s1 := re1.ReplaceAllString("dewble.com", ".co.kr")
	fmt.Println(s1) // dewble.co.kr

	// 두 분째로 찾은 문자열 world!만 남김
	re2, _ := regexp.Compile("([a-zA-Z,]+) ([a-zA-Z!]+)")
	s2 := re2.ReplaceAllString("hello, world!!", "${2}")
	fmt.Println(s2) // world!!

	re3, _ := regexp.Compile("([a-zA-Z,]+) ([a-zA-Z!]+)")
	s3 := re3.ReplaceAllString("hello, world!!", "${2} ${1}")
	fmt.Println(s3) // world!! hello,

	// 찾은 문자열에 alias 정하기
	re4, _ := regexp.Compile("(?P<first>[a-zA-Z,]+) (?P<second>[a-zA-Z!]+)")
	// 정규표현식에 해당하는 문자열을 지정된 문자열로 바꿈
	s4 := re4.ReplaceAllString("Hello, world!!", "${second} ${first}")
	fmt.Println(s4) // world!! hello,

	// ReplaceAllLiteralString 함수는 문자열을 바꿀 때 정규표현식 기호를 무시합니다. 즉 ${second} ${first}는 동작하지 않고 문자 그대로(Literal) 들어갑니다.
	re5, _ := regexp.Compile("(?P<first>[a-zA-Z,+) (?P<second>[a-zA-Z!]+)")
	s5 := re5.ReplaceAllLiteralString("Hello, world!", "${second} ${first}")
	fmt.Println(s5) // world!! hello,

}
