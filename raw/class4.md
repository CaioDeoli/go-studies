# Aula 4 - Como o go trabalha com erros
Trabalhar e tratar erros

No go o erro sempre vai se tornar algo explicito quando voce for trabalhar com a maioria das funções

Uma função no go pode retornar mais de um valor. E um desses valores pode ser um erro. Então se tiver algum erro, paramos e tratamos aquele pedaço de código, o que vamos fazer pode mudar totalmente o fluxo do programa

```go
package main

import "net/http"

func main() {
  res, err := http.Get("http://google.com.br")
  // retorna um erro caso essa chamada get tenha alguma erro
  // O erro vai vim para err

  if err != nil { // se erro não for vazio // SE USA MUITO
    panic() // aborta toda a aplicação
    log.Fatal("Erro ao fazer comunicacao")
    log.Fatal(err.Error())
  }

  fmt.Println(res.Header)
}
```

Minha propria funcao que retorna um erro
```go
package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {

	res, err := soma(10, 11)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(res)

}

func soma(x int, y int) (int, error) {
  res := x + y

  if res > 10 {
    return 0, // nao importa o valor
    errors.New("Total maior que 10")
  }

  return res, nil // vou retornar vazio porque nao deu erro
}
```

blankidentify, joga o dado fora porque voce nao vai utilizar, mas voce nao pode deixar de receber
é o `_`
```go
package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {

	res, _ := soma(7, 2)

	fmt.Println(res)

}

func soma(x int, y int) (int, error) {
  res := x + y

  if res > 10 {
    return 0, // nao importa o valor
    errors.New("Total maior que 10")
  }

  return res, nil // vou retornar vazio porque nao deu erro
}
```