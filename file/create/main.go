package main

import (
	"fmt"
	"os"
)

func main() {
	// file-create.txt 파일 생성
	// file 존재할 경우 덮어쓴다
	file, err := os.Create("file-create.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // main 함수가 끝나기 직전에 파일을 닫음

	s := "Hello world~ Create file with go~"

	n, err := file.Write([]byte(s)) // s를 []byte 바이트 슬라이스로 변환, s를 파일에 저장
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, "바이트 저장 완료")
}
