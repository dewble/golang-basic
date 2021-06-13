package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	// Create file

	s := "create and read file with iotuil "
	err := ioutil.WriteFile("ioutil.txt", []byte(s), os.FileMode(0644))
	// s 를 []byte 바이트 슬라이스로 변환, s를 hello.txt 파일에 저장

	if err != nil {
		fmt.Println(err)
		return
	}

	// Read file
	data, err := ioutil.ReadFile("ioutil.txt")
	// hello.txt 의 내용을 읽어서 바이트 슬라이스 리턴
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}
