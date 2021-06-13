package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"create-file.txt",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		// 파일 생성
		// 파일 읽기,쓰기
		// 파일 연 뒤 삭제

		os.FileMode(0644)) // 644 	권한
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // main 함수 끝나기 전에 파일 닫기

	n := 0
	s := "File Create and Read"
	n, err = file.Write([]byte(s)) // s 를 []byte 바이트 슬라이스로 변환, s를 파일에 저장
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n, "바이트 저장 완료")

	fi, err := file.Stat() // 파일 정보 가져오기
	if err != nil {
		fmt.Println(err)
		return
	}

	var data = make([]byte, fi.Size()) // 파일 크기만큼 바이트 슬라이스 생성

	file.Seek(0, os.SEEK_SET)
	// Write 함수로 인해 파일 읽기/쓰기 위치가 이동 했으므로
	// file.Seek 함수로 읽기/쓰기 위치를 파일의 맨 처음(0)으로 이동
	// 파일의 내용을 읽어서 바이트 슬라이스에 저장

	n, err = file.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, "바이트 읽기 완료")
	fmt.Println(string(data)) // 문자열로 변환하여 data의 내용 출력

}
