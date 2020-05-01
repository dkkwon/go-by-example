package main

import (
	"fmt"
	"strings"
)

// We often need our programs to perform operations on collections of data,
// like selecting all items that satisfy a given predicate or mapping all items
// to a new collection with a custom function.

// In some languages it’s idiomatic to use generic data structures and
// algorithms. Go does not support generics; in Go it’s common to provide
// collection functions if and when they are specifically needed for your
// program and data types.

// Here are some example collection functions for slices of strings. You can
// use these examples to build your own functions. Note that in some cases it
// may be clearest to just inline the collection-manipulating code directly,
// instead of creating and calling a helper function.

// 일치하는 원소 위치
func Index(strs []string, s string) int {
	for i, v := range strs {
		if v == s {
			return i
		}
	}
	return -1
}

// 일치하는 원소 존재 여부
func Include(strs []string, s string) bool {
	return Index(strs, s) >= 0
}

// 조건 (f func())에 맞는 원소 존재 여부
func Any(strs []string, f func(string) bool) bool {
	for _, v := range strs {
		if f(v) {
			return true
		}
	}
	return false
}

// 조건 (f func())에 모두 만족 여부
func All(strs []string, f func(string) bool) bool {
	for _, v := range strs {
		if !f(v) {
			return false
		}
	}
	return true
}

//  조건 (f func())에 만족하는 원소 필터
func Filter(strs []string, f func(string) bool) []string {
	filtered := make([]string, 0)
	for _, v := range strs {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

// 입력 -> 함수 (f func()) -> 결과 저장
func Map(strs []string, f func(string) string) []string {
	mapped := make([]string, len(strs))
	for i, v := range strs {
		mapped[i] = f(v)
	}
	return mapped
}

func main() {
	strs := []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear"))
	fmt.Println(Include(strs, "grape"))

	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	fmt.Println(Map(strs, strings.ToUpper))
}
