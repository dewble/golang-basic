package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// JSON 데이터 반환 웹서버 테스트 코드
func TestJsonHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/student", nil) // 1. "/student" 경로 테트

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	student := new(Student)
	err := json.NewDecoder(res.Body).Decode(student) // 2. 결과 변환
	assert.Nil(err)                                  // 3. 결과 확인
	assert.Equal("JM", student.Name)
	assert.Equal(29, student.Age)
	assert.Equal(88, student.Score)
}

/*
// 웹 서버 테스트 코드 만들기

func TestIndexHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil) // 1. "/" 경로 테트

	mux := MakeWebHandler() // 2.
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code) // 3. Code 확인
	data, _ := io.ReadAll(res.Body)       // 4. 데이터를 읽어서 확인
	assert.Equal("Hello World", string(data))
}

func TestBarHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil) // 5. "/bar" 경로 테트

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Eqaul(http.StatusOK, res.Code) // 3. Code 확인
	data, _ := io.ReadAll(res.Body)       // 4. 데이터를 읽어서 확인
	assert.Equal("Hello Bar", string(data))
}


*/
