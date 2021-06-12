package main

import (
	"fmt"
	"os"
)

func main() {
	// file-create.txt 파일 열기
	// 파일 없을 경우 err message

	file, err := os.Open("file-create.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // main 함수가 끝나기 직전에 파일을 닫음

	fi, err := file.Stat() // 파일 정보 가져오기
	if err != nil {
		fmt.Println(err)
		return
	}

	var data = make([]byte, fi.Size()) // 파일 크기만큼 바이트 슬라이스 생성

	n, err := file.Read(data) // 파일의 내용을 읽어서 바이트 슬라이스에 저장
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, "바이트 읽기 완료")
	fmt.Println(string(data)) // 문자열로 변환하여 data의 내용 출력
}
