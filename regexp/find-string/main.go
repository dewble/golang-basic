package main

import (
	"fmt"
	"regexp"
)

func main() {

	// .영문 으로 끝나는 문자열 리턴
	re1, _ := regexp.Compile("\\.[a-zA-Z]+$")
	s1 := re1.FindString("hello dom.ain.com")
	fmt.Println(s1)
	// .com

	// 문자열을 검색한 뒤 찾은 문자열과
	// 괄호로 구분된 하위 항목 리턴
	re2, _ := regexp.Compile("([a-zA-Z]+\\.[a-zA-Z]+)$")
	s2 := re2.FindStringSubmatch("hello domain.com")
	fmt.Println(s2)

	re3, _ := regexp.Compile("([a-zA-Z]+)(\\.[a-zA-Z]+)$")
	s3 := re3.FindStringSubmatch("hello domain.com")
	fmt.Println(s3)

	// [domain.com domain.com]
	// [domain.com domain .com]

	// 문자열을 검색한 뒤 찾은 문자열의 위치와 괄호로 구분한 하위 항목의 위치를 리턴
	n3 := re3.FindStringSubmatchIndex("domain.com")
	fmt.Println(n3)
	// [0 10 0 6 6 10]
	// [시작 끝 시작 끝 시작 끝]

	// 정규표현식에 해당하는 모든 문자열 검색, 모두 리턴(-1)
	re4, _ := regexp.Compile("\\.[a-zA-Z.]+") // .영문.
	s4 := re4.FindAllString("hello.co.kr world hello.net example.com", -1)
	fmt.Println(s4)
	// [.co.kr .net .com]

	// 정규표현식에 해당하는 모든 문자열 검색, 2개 리턴(2)
	re5, _ := regexp.Compile("\\.[a-zA-Z.]+") // .영문.
	s5 := re5.FindAllString("hello.co.kr world hello.net example.com", 2)
	fmt.Println(s5)
	// [.co.kr .net]

	// 문자열을 검색한 뒤 찾은 문자열의 위치와 괄호로 구분한 하위 항목의 위치를 리턴

	re6, _ := regexp.Compile("\\.[a-zA-Z.]+") // .영문.
	s6 := re6.FindAllStringIndex("hello.co.kr world hello.net example.com", -1)
	fmt.Println(s6)
	// [[5 11] [23 27] [35 39]]
}
