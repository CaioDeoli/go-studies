Aula 5 funções
Usar uma função como um método

dica: eu não aprendia um certo nivel de profundidade que as linguagens tinham. preste atenção sobre a linguagem e si, e fique a frente e torne-se mais produtivo

```go
package main

func main() {}
```

variações das funções go

```go
func soma(a int, b int) {
  return a + b // erro, porque nossa função nao especifica o que retorna
}
```

```go
func soma(a int, b int) int {
  return a + b
}

func main() {
  retorn := soma(10, 30) // retorno normal
}
```

```go
// go pode retornar mais de um valor
func soma(a int, b int) (int, string) {
  return a + b, "somou"
}

func main() {
  resultado, str := soma(10, 20)

  fmt.Println(resultado, str) // 30 somou // retorno multiplo
}
```

aprofundando
```go
func soma(a int, b int) (result int) { // retorno nomeado
  result = a + b
  return
}

func main() {
  resultado := soma(10, 20)
  fmt.Println(resultado)
}
```

ta na moda -> funções variadicas
posso passar quantos valores inteiros eu quiser
sempre lembre de tratar essas entradas
```go
func somaTudo(x ...int) int { // posso passar um array de inteiros
  resultado := 0

  for _, v := range x { // faz o loop em todos os valores de x, sem precisar de key `_`
    resultado += v
  }
  return resultado
}

func main() {
  resultado := somaTudo(3, 5, 10, 4, 6, 35)
  fmt.Println(resultado) // resultado da soma toda
}
```

pra quem gosta de programação funcional ou coisa do tipo
podemos trabalhar com funções anonimas e funções dentro de funções
no go quando trabalhamos com multtread a gente faz
```go
func main() {
  go func() { // roda como um thread

  }
}
```

```go
func main() {
  resultado := func(x ...int) func() int {
    res := 0

    for _, v := range x {
      res += v
    }

    return func() int {
      return res * res
    }
  }

  fmt.Println(resultado) // retorna um ponteiro para a função
  // um endereçamento de memoria
  // porque eu estou retornando uma função e não o inteiro em si
}
```
o correto então é:
```go
fmt.Println(resultado(1, 2, 3, 4, 5)()) // 225
```

como transformar uma função em um método
voce consegue trabalhar com métodos em golang
o go possui uma especie de orientação a objeto, mas é o go way (é uma forma diferente de programar)
os metodos em go são baseados em estruturas (type)

```go
package main

import "fmt"

type Carro struct {
  Nome string
}

// para dizer que a função é um método do caro use o "c Carro"
// c Carro é o relacionamento da função com o struct, um bind
func(c Carro) andar() {
  fmt.Println(c.Nome, "andou")
}

func main() {
  carro := Carro{
    Nome: "Gol"
  }

  carro.andar() // Gol andou
}
```