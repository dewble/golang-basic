package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibo(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, fibonacci1(-1), "fibonacci1(-1) should be 0")
	assert.Equal(0, fibonacci1(0), "fibonacci1(0) should be 0")
	assert.Equal(1, fibonacci1(1), "fibonacci1(1) should be 1")
	assert.Equal(2, fibonacci1(3), "fibonacci1(2) should be 2")
	assert.Equal(233, fibonacci1(13), "fibonacci1(13) should be 233")
}

func TestFibo2(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, fibonacci2(-1), "fibonacci2(-1) should be 0")
	assert.Equal(0, fibonacci2(0), "fibonacci2(0) should be 0")
	assert.Equal(1, fibonacci2(1), "fibonacci2(1) should be 1")
	assert.Equal(2, fibonacci2(3), "fibonacci2(2) should be 2")
	assert.Equal(233, fibonacci2(13), "fibonacci2(13) should be 233")
}

func BenchmarkFibo1(b *testing.B) {
	// TODO: Initialize
	for i := 0; i < b.N; i++ { // 1. b.N 만큼 반복
		// TODO: Your Code Here
		fibonacci1(20)
	}
}

func BenchmarkFibo2(b *testing.B) {
	// TODO: Initialize
	for i := 0; i < b.N; i++ {
		// TODO: Your Code Here
		fibonacci2(20)
	}
}
