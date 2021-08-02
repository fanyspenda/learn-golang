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

//algoritmanya ternyata sama dengan encode karena tinggal dibalik saja

//Decode ...
//jika dengan bintang(pointer), maka merefer dari value instansiasinya.
//jika tanpa bintang, maka variable yang dijadikan parameter akan dicopy ke variable parameter.
//jadi, jika kita mengubah value dari parameter, value aslinya tidak akan berubah.
//misal, func ubahA(a int) {a=2}; ubahA(b) => nilai b tidak akan berubah
//misal, func ubahA(a *int) {a=2}; ubahA(b) => nilai b berubah
func (s *Student) Decode() (nameDecode string) {

	for _, v := range s.NameEncode {
		diff := 122 - int(v)
		nameDecode += string(int(v) - 122 + 97 + (diff * 2))
	}
	// your code here
	fmt.Println(nameDecode)
	return nameDecode
}
