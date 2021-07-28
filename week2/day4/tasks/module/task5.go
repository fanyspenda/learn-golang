package module

import (
	"fmt"
)

//Student ...
type Student struct {
	Name       string
	NameEncode string
	Score      int
}

//Cipher ...
type Cipher interface {
	Encode() string
	Decode() string
}

//Encode ...
func (s *Student) Encode() (nameEncode string) {

	fmt.Println(s.Name)

	for _, v := range s.Name {
		diff := 122 - int(v)
		nameEncode += string(int(v) - 122 + 97 + (diff * 2))
	}

	s.NameEncode = nameEncode
	return nameEncode
}

//Decode ...
//algoritmanya ternyata sama dengan encode karena tinggal dibalik saja
func (s *Student) Decode() (nameDecode string) {

	for _, v := range s.NameEncode {
		diff := 122 - int(v)
		nameDecode += string(int(v) - 122 + 97 + (diff * 2))
	}
	// your code here
	fmt.Println(nameDecode)
	return nameDecode
}
