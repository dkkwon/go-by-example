package main

import (
	"encoding/xml"
	"fmt"
)

// 구조체 필드 첫 글자 대문자(capital letters) -> if not, 생략됨
// xml elemet 를 변경하려면? 태그 지정 -> `xml: 키` --> `(백쿼트, GRAVE ACCENT) 사용

type Plant struct {
	XMLName xml.Name `xml:"plant"`   // XMLNAME : xml Plant root element
	Id      int      `xml:"id,attr"` // id : xml attribute
	Name    string   `xml:"name"`    // Name : xml element
	Origin  []string `xml:"origin"`  // Origin : xml element
}

// type Plant struct {
// 	XMLName xml.Name
// 	Id      int
// 	Name    string
// 	Origin  []string
// }

// func (p Plant) String() string {
// 	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
// 		p.Id, p.Name, p.Origin)
// }

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))

	fmt.Println(xml.Header + string(out))

	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
		// Plants  []*Plant `xml:"Parent>Child>lant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}
	// nesting.Plants = []*Plant{tomato, coffee}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}
