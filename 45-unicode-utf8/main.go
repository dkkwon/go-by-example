package main

import "fmt"

// https://medium.com/@jeongdowon/unicode%EC%99%80-utf-8-%EA%B0%84%EB%8B%A8%ED%9E%88-%EC%9D%B4%ED%95%B4%ED%95%98%EA%B8%B0-b6aa3f7edf96

func main() {

	word := "abc가나다"
	for i, r := range word {
		fmt.Printf("%d %d 0x%x %c\n", i, r, r, r)
	}
	fmt.Println()
	for i := 0; i < len(word); i++ {
		fmt.Printf("%d(%c) ", word[i], word[i])
	}
	// output:
	// 0 97 0x61 a
	// 1 98 0x62 b
	// 2 99 0x63 c
	// 3 44032 0xac00 가
	// 6 45208 0xb098 나
	// 9 45796 0xb2e4 다
	// 97(a) 98(b) 99(c) 234(ê) 176(°) 128() 235(ë) 130() 152() 235(ë) 139() 164(¤)
}
