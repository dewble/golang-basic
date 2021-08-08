package main

import (
	"fmt"
	"net/http"
)

func MakeWebHandler() http.Handler { // 1. 핸들러 인스턴스를 생성하는 함수
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Bar")
	})
	return mux
}

func main() {
	http.ListenAndServe(":3000", MakeWebHandler()) // 3. mux 인스턴스 사용
}

/*
웹 서버 테스트 코드 만들기
*/

/*
특정 경로에 있는 파일 읽어오기
*/
// func main() {
// 	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static")))) // ❶
// 	http.ListenAndServe(":3000", nil)
// }

/*
파일 서버
*/

// func main() {
// 	http.Handle("/", http.FileServer(http.Dir("static"))) // ❶
// 	http.ListenAndServe(":3000", nil)
// }

/*
serveMux 이용하기
*/

// func main() {
// 	mux := http.NewServeMux() // 1. ServerMux instance 생성
// 	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, "Hello World") // 2. 인스턴스에 핸들러 등록
// 	})
// 	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, "Hello Bar")
// 	})

// 	http.ListenAndServe(":3000", mux) // 3. mux 인스턴스 사용
// }

/*
HTTP 쿼리 인수 사용하기
*/

// func barHandler(w http.ResponseWriter, r *http.Request) {
// 	values := r.URL.Query()    // 1. 쿼리 인수 가져오기
// 	name := values.Get("name") // 2. 특정 키값이 있는지 확인
// 	if name == "" {
// 		name = "World"
// 	}
// 	id, _ := strconv.Atoi(values.Get("id")) // 3. id값을 가져와서 int 타입 변환
// 	fmt.Fprintf(w, "Hello %s!, id:%d", name, id)
// }

// func main() {
// 	http.HandleFunc("/bar", barHandler) // 4. /bar 핸들러 등록
// 	http.ListenAndServe(":3000", nil)
// }
