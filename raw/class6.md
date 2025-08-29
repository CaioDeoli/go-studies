# Aula 6 - Ponteiros

// Vai fazer uma grande diferença na sua carreira. Quanto mais fundamentos, base e conhecimento, mais voce vai poder criar e ser criativo

```go
package main

import "fmt"

func main() {

  a := 10
  // como se comporta no pc
  // essa variavel tem um apontamento para a memória
  /*
  memoria(10) <---- a <---- 10
  o valor 10 é atribuido para variavel a
  a variavel a aponta para um local da memora
  que vai guardar esse meu valor 10
  se eu mudar
  memoria(40) <---- a <---- 40

  essa memoria tem um endereço
  vamos chamar de memoria-endereco
  temos endereços na memoria onde nossos valores são chamados
  então toda vez que eu achamar a variavel a, vai ser retornado para mim o valor que esta no endereço
  */

  // Nós podemos saber qual é esse endereço de memória
  fmt.Println(a) // 10
  fmt.Println(&a) // 0xc0000160a0 --> local para onde a variavel esta apontando
  // memoria(10) (0xc0000160a0) <---- a <---- 10
}
```

Esse endereço na memoria fica catalogado nos ponteiros
```go
a := 10
var ponteiro *int = &a // o apontamento que esta mandando do a para memória esta guardado nessa variavel ponteiro (ele detem o endereço da memória)

```
// conseguimos fazer 'reference'. pega um endereço da memoria, consulta o endereço que esta la, e tras o resultado para nós
```go
a := 10
var ponteiro *int = &a
fmt.Println(ponteiro) // 0xc0000160a0 -> posicao da memoria que o ponteiro esta indicando
fmt.Println(*ponteiro) // 10
// ele capturou o endereço da memoria e retornou o valor guardado la

*ponteiro = 50
fmt.Println(*ponteiro) // 50 -> atribuiu como valor q estava na memoria
fmt.Println(ponteiro) // nao muda
// voce so muda o valor q estava armazenado na memoria
fmt.Println(a) // 50 -> o a estava sendo direcionado para o mesmo local da memoria que *ponteiro

b := &a
fmt.Println(*b) // 50 -> b é ponteiro
*b = 60
fmt.Println(*ponteiro) // 60
fmt.Println(a) // 60
// A, *ponteiro, *B estão apontando para o mesmo local da memória
// Eu tenho uma variavel, ela tem um endereço de memória, eu tenho um ponteiro que guarda o endereço da memória e assim eu consigo ter esses acessos
```

```go
func main() {
  variavel := 10

  abc(&variavel)

  fmt.Println(variavel) // 200
}

// funcao nao retorna nada mas altera o valor que voce passou por parametro porque é um endereçamento de memória
// por que reflete em todos os lugares
func abc(a *int) { // recebe um ponteiro
  *a = 200 // altera o valor que esta no endereço da memória
}
```

// palavras novas: structs, len

```go
type Carro struct {
  Name string
}

// c acessa o valor da struct
func (c Carro) andou() {
  fmt.Println(c.Name, "andou")
}

func main() {
  carro := Carro{
    Name: "Ka"
  }

  carro.andou() // Ka andou
}
```

Outra visao
```go
type Carro struct {
  Name string
}

// c acessa o valor da struct
func (c Carro) andou() {
  c.Name = "BMW"
  fmt.Println(c.Name, "andou")
}

func main() {
  carro := Carro{
    Name: "Ka"
  }

  carro.andou() // BMW andou
  fmt.Println(carro.Name) // Ka -> porque nao alterou o endereço da memória, porque voce só alterou o valor no escopo da funcao andou
}
```

o poder dos ponteiros
```go
type Carro struct {
  Name string
}

// vai alterar o valor real com o * porque aponta para o mesmo local
func (c *Carro) andou() {
  c.Name = "BMW"
  fmt.Println(c.Name, "andou")
}

func main() {
  carro := Carro{
    Name: "Ka"
  }

  carro.andou() // BMW andou
  fmt.Println(carro.Name) // Ka -> porque nao alterou o endereço da memória, porque voce só alterou o valor no escopo da funcao andou
}
```
// muito comum com Go pra mudar os valores da instancia
// as vezes voce nao quer alterar e ai voce nao usa o asteristico
// comportamentos diferentes
// importante entender para usar nas horas que voce precisar usar cada uma das propriedades