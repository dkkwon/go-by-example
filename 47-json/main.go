package main

import (
	"encoding/json"
	"fmt"
)

/*

// 구조체 필드 첫 글자 대문자(capital letters) -> if not, 생략됨
// 키를 소문자하려면? 태그 지정 -> `json: 키` --> `(백쿼트, GRAVE ACCENT) 사용

type 구조체명 struct {
	필드명 자료형 `json:"키"`
}
*/

type Response struct {
	Page   int
	Fruits []string
}

type ResponseWithTags struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

// JSON (JavaScript Object Notation) is a lightweight data-interchange format.
// Marshal 의미 : n. 군대 원수, 경찰 서장, v. 통제하다, 모의다. --> 정렬시키기, 순위결정

// func Marshal(v interface{}) ([]byte, error): Go 자료형 -> JSON byte steam (type byte = uint8)
// func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error): Go 자료형 -> JSON 텍스트 with Indent
// func Unmarshal(data []byte, v interface{}) error: JSON byte steam -> Go 자료형

func main() {
	booleanAsJSON, _ := json.Marshal(true)
	fmt.Println("boolean", string(booleanAsJSON))

	integerAsJSON, _ := json.Marshal(1)
	fmt.Println("integer", string(integerAsJSON))

	floatAsJSON, _ := json.Marshal(42.42)
	fmt.Println("float", string(floatAsJSON))

	stringAsJSON, _ := json.Marshal("gopher")
	fmt.Println("string", string(stringAsJSON))

	sliceOfStringsAsJSON, _ := json.Marshal([]string{"apple", "peach", "pear"})
	fmt.Println("slice:", string(sliceOfStringsAsJSON))

	mapFromStringsToIntegers, _ := json.Marshal(map[string]int{"apple": 5, "lettuce": 7})
	fmt.Println("map:", string(mapFromStringsToIntegers))
	// fmt.Println("map:", mapFromStringsToIntegers)

	res1 := &Response{Page: 1, Fruits: []string{"apple", "peach", "pear"}}
	res1AsJSON, _ := json.Marshal(res1)
	fmt.Println("response:", string(res1AsJSON))

	res2 := &ResponseWithTags{Page: 2, Fruits: []string{"apple", "peach", "pear"}}
	res2AsJSON, _ := json.Marshal(res2)
	fmt.Println("response with annotations:", string(res2AsJSON))
	res2AsJSON, _ = json.MarshalIndent(res2, "", " ")
	fmt.Println("response with annotations:", string(res2AsJSON))

	// bytesInInput := []byte(`{"num":6.13,"strs":["a","b"]}`)
	doc := `
	{
		"num":6.13,
		"strs":["a","b"]
	}
	`
	bytesInInput := []byte(doc)

	// We need to provide a variable where the JSON package can put the decoded
	// data. This map[string]interface{} will hold a map of strings to arbitrary
	// data types.

	var data map[string]interface{}

	if err := json.Unmarshal(bytesInInput, &data); err != nil {
		panic(err)
	}
	fmt.Println(data)

	// Accessing nested data requires a series of casts.
	num := data["num"].(float64)
	fmt.Println(num)

	// We can also decode JSON into custom data types. This has the advantages of
	// adding additional type-safety to our programs and eliminating the need for
	// type assertions when accessing the decoded data.

	stringInInput := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := &ResponseWithTags{}
	json.Unmarshal([]byte(stringInInput), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits)
}
