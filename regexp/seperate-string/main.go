package main

import (
	"fmt"
	"regexp"
)

func main() {

	// Split(string, 쪼갤 개수)
	// -1 전부, 0 nothing, 2 -> 2,3번째는 그냥둠 , 3-> 3번째는 그냥둠

	// 정규표현식 .영문. 에 해당하는 문자열을 기준으로 모든 문자열을 쪼갬
	re1, _ := regexp.Compile("\\.([a-z]+)\\.")
	s1 := re1.Split("mail.dewble.net www.dewble.com ftp.dewble.com", -1)
	fmt.Println(s1) // [mail net www com ftp com]

	// 정규표현식 .영문. 에 해당하는 문자열을 기준으로 모든 문자열을 쪼갬
	re2, _ := regexp.Compile("\\.([a-z]+)\\.")
	s2 := re2.Split("mail.dewble.net www.dewble.com ftp.dewble.com", 2)
	fmt.Println(s2) // [mail net www.dewble.com ftp.dewble.com]

	// 정규표현식 .영문. 에 해당하는 문자열을 기준으로 모든 문자열을 쪼갬
	re3, _ := regexp.Compile("\\.([a-z]+)\\.")
	s3 := re3.Split("mail.dewble.net www.dewble.com ftp.dewble.com", 3)
	fmt.Println(s3) // [mail net www com ftp.dewble.com]

	// 정규표현식 .영문. 에 해당하는 문자열을 기준으로 모든 문자열을 쪼갬
	re4, _ := regexp.Compile("\\.([a-z]+)\\.")
	s4 := re4.Split("mail.dewble.net www.dewble.com ftp.dewble.com", 0)
	fmt.Println(s4) // []
}
