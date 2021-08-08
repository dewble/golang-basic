package main

import (
	"fmt"
)

func square(x int) int {
	return x * x
}

func main() {
	fmt.Printf("9 * 9 = %d\n", square(9))
}

// func TestSquare1(t *testing.T) {
// 	assert := assert.New(t)                               // ❶ 테스트 객체 생성
// 	assert.Equal(81, square(9), "square(9) should be 81") // ❷ 테스트 함수 호출
// }

// func TestSquare2(t *testing.T) {
// 	assert := assert.New(t)
// 	assert.Equal(9, square(3), "square(3) should be 9")
// }
