### Aula 9 - Generics

```go
package main

import "fmt"

func SomarInteiros(m map[string]int) int {
  soma := 0

  for _, v := range m {
    soma += v
  }

  return soma
}

func main() {
  res := SomarInteiros(map[string]int{"a": 1, "b": 2, "c": 3})
  fmt.Println(res)
}
```


Como seria se eu precisasse de uma função para receber não inteiros, mas floats? nas versões antigas não existiam generics então precisava fazer uma gambiarra com interface ou duplicar a função
```go
package main

import "fmt"

func SomarInteiros(m map[string]int) int {
  soma := 0

  for _, v := range m {
    soma += v
  }

  return soma
}

func SomarFloat(m map[string]float64) float64 {
  soma := 0

  for _, v := range m {
    soma += v
  }

  return soma
}

func main() {
  fmt.Println(SomarInteiros(map[string]int{"a": 1, "b": 2, "c": 3}))
  fmt.Println(SomatFloat(map[string]float64{"a": 1.1, "b": 22.2, "c": 3.3}))
}
```

Como faria para criar uma função que serve para ambos?
```go
package main

import "fmt"
// Coloque uma letra qualquer. No caso usamos T
func SomaGenerica[T int64 | float64] (m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return noma
}

func main() {
	fmt.Println(SomaGenerica(map[string]int{"a": 1, "b": 2, "c": 3}))
	fmt.Println(SomaGenerica(map[string]float64{"a": 1.1, "b": 22.2, "c": 3.3}))
}
// mesmo resultado mas com uma unica funcao
```


// agora imagina que eu tenho vários tipos para T
// podemos criar um tipo para isso

```go
type Number interface {
	int64 | float64
}

// Constraints -> amarras dos tipos de dados que entram na função generica
func SomaGenerica[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return noma
}

```

vamos imaginar que criamos outro tipo
```go
type Number interface {
	int | int64 | float64
}

type MyNumber int

// Constraints -> amarras dos tipos de dados que entram na função generica
func SomaGenerica[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return noma
}

func main() {
	var x, y, z MyNumber
	x = 1
	y = 2
	z = 3
	fmt.Println(SomaGenerica(map[string]MyNumber{"a": x, "b": y, "c": z})) // vai dar erros neles locais
	// nao pode porque só podemos trabalhar com o type Number
	// como fazer o go entender isso, fazendo uma aproximação ja que esses inteiros fazem parte do number
	// fazemos então o uso de ~
	
	// substituindo int por ~int
	type Number interface {
		~int | int64 | float64
	}
	
	// e nao da erro; ~ é um apoderador de aproximação
	// assim o type MyNumber implementa a interface do type Number
	fmt.Println(SomaGenerica(map[string]MyNumber{"a": x, "b": y, "c": z}))
	
}



```

```go
func Soma[T Number](number1 T, number2 T) T {
	return number1 + number2
}
```

```go
// envetualmente podemos ter um problema
// esse exemplo abaixo funciona de boa
func Soma[T Number](number1 T, number2 T) T {
	if number1 == number2 {
		return number1
	}
	return number2
}

// mas imagina se
// any é como se fosse "interface {}" = pode ser qualquer tipo de dado
func Soma[T any](number1 T, number2 T) T {
	// essa comparacao da errado. ele não deixa fazer essa comparação porque não são comparaveis
	// pode ser number1 uma string e number2 um numéro
	if number1 == number2 {
		return number1
	}
	return number2
}

// outro Constraints/ comparable é tipo uma interface. e que permite comparar os dois porque eu sei
// que os dois são do mesmo tipo (eu especifico abaixo que sao do mesmo tipo)
func Soma[T comparable](number1 T, number2 T) T {
	if number1 == number2 {
		return number1
	}
	return number2
}

//
func Max[T any](number1 T, number2 T) T {
	// temos o mesmo problema
	if number1 > number2 {
		return number1
	}
	return number2
}

//
func Max[T comparable](number1 T, number2 T) T {
	// continuamos com o mesmo problema
	// porque com comparable apenas comparar igualdade, não podemos ver
	// se um é maior que o outro
	if number1 > number2 {
		return number1
	}
	return number2
}

// pacotes criados de Constraints
// nome do pacote: constraints
// Complex
// Ordered -> serve para ver se algo é maior que o outro
// $ go get golang.org/x/exp/constrains
// Essas constraints ajuda no dia a dia
func Max[T constraints.Ordered](number1 T, number2 T) T {
	// continuamos com o mesmo problema
	// porque com comparable apenas comparar igualdade, não podemos ver
	// se um é maior que o outro
	if number1 > number2 {
		return number1
	}
	return number2
}

```

quando trabalhamos com generics funciona tanto com funções como métodos
```go
// função de concatenação
// lista desse valor tipo T vai ser recebido
func concat[T any](vals []T)string {
	result := ""
	for _, val := range vals {
		result += val.String() // da erro porque em nenhum lugar estamos falando que o tipo any tem o método String()
	}
	return result
}
```

```go
type stringer interface {
	String() string
}

func concat[T stringer](vals []T)string {
	result := ""
	for _, val := range vals {
		result += val.String() // nao da problema porque o stringer tem esse método
	}
	return result
}
```


```go
type stringer interface {
	String() string
}

func concat[T stringer](vals []T)string {
	result := ""
	for _, val := range vals {
		result += val.String()
	}
	return result
}

func main() {
	// concat vai ta dando erro porque não tem o método string
	fmt.Println(concat([]string{"a", "b", "c"})) 
}
```


```go
type stringer interface {
	String() string
}

type MyString string

func concat[T stringer](vals []T)string {
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
```


```go
type stringer interface {
	String() string // método
}

type MyString string

// criamos um método para o MyString
func(m MyString) String() string {
	return string(m)
}

func concat[T stringer](vals []T)string {
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
```

Com generics você pode economizar muito código
refs:
https://go.dev/doc/tutorial/generics
https://bitfieldconsulting.com/posts/generics
https://www.boot.dev/courses/learn-golang