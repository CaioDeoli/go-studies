package main

import "fmt"

type stringer interface {
	String() string // método
}

type MyString string

// criamos um método para o MyString
func (m MyString) String() string {
	return string(m)
}

func concat[T stringer](vals []T) string {
	result := ""
	for _, val := range vals {
		result += val.String()
	}
	return result
}

func main() {
	// ainda vai continuar dando erro porque eu estou atéas do método de stringer
	fmt.Println(concat([]MyString{"a", "b", "c"}))
}
