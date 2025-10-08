# Aula 7 - Structs, composição e json com golang

structs
recurso bastante utilizado em algumas linguagens
confundimos structs com classes (principalmente em poo)
go não é orientada a objeto, e sim, o go way

```go
package main

type ClienteNome string
type ClienteEmail string
type ClienteCPF int

func main() {
  
  var clienteNome ClienteNome = 'Caio'

}
// dados ficam muito soutos
```

// como podemos melhorar? com struct (tipo um esqueleto)
```go
package main

type Cliente struct {
  Nome string
  Email string
  CPF int
}

func main() {

  cliente := Cliente{
    Nome: "Caio"
    Email: "c@g.com"
    CPF: 1234567890
  }
  fmt.Println(cliente) // {Caio c@g.com 1234567890}
  
  cliente2 := Cliente{"Mari", "m@g.com", 9876543210}
  fmt.Printf("Nome: %s. Email: %s. CPF %d", cliente2.Nome, cliente2.Email, cliente2.CPF) // Nome: .... Email: .... CPF: ....
}
```

// codemos fazer algo parecido com herança e parecido como composição
```go
type Cliente struct {
  Nome string
  Email string
  CPF int
}

type ClienteInternacional struct {
  Nome string
  Email string
  CPF int
  Pais string
}

cliente := Cliente{
  Nome: "Davi",
  Email: "d@g.com",
  CPF: "1234567890"
}
fmt.Printf("Nome: %s. Email: %s. CPF: %d", cliente.Nome, cliente.Email, cliente.CPF)

cliente3 := ClienteInternacional{
  Nome: "Caio",
  Email: "c@g.com",
  CPF: 1234561230
  Pais: "EUA"
}
fmt.Printf("Nome: %s. Email: %s. Pais: %s\n", cliente3.Nome, cliente3.Email, cliente3.Pais)

// as duas structs são muito parecidas então o que podemos fazer:

type ClienteInternacional struct {
  Cliente // compondo cliente em um cliente internacional
  Pais string
}
// precisa alterar:
cliente3 := ClienteInternacional{
  CLiente: Cliente{
    Nome: "Caio",
    Email: "c@g.com",
    CPF: 12345674890
  },
  Pais: "EUA"
}
fmt.Printf("Nome: %s. Email: %s. Pais: %s\n", cliente3.Nome, cliente3.Email, cliente3.Pais) // voce chama pais dessa forma como se fosse herança
```

Métodos
podemos atachar uma função em uma struct
e ai falamos que structs tem métodos

```go
import "fmt"

type Cliente struct {
	Nome string
	Email string
	CPF int
}

// Atachamos uma função na struct e ai atachamos
func (c Cliente) Exibe() {
	fmt.Println("Exibindo cliente pelo método ", c.Nome)
}
```

```go
package main

import "fmt"

type Cliente struct {
  Nome  string
  Email string
  CPF   int
}

func (c Cliente) Exibe() {
	fmt.Println("Exibindo cliente pelo método ", c.Nome)
}

type ClienteInternacional struct {
  Cliente
  Pais string
}

func main() {
  cliente1 := Cliente{
    "Caio",
    "c@g.com",
    9876,
  }

  cliente2 := ClienteInternacional{
    Cliente: Cliente{
      Nome:  "Andréa",
      Email: "a@g.com",
      CPF:   4585,
    },

    Pais: "Brasil",
  }

  fmt.Printf("Nome: %s. Email: %s, CPF: %d\n", cliente1.Nome, cliente1.Email, cliente1.CPF)
  fmt.Printf("Nome: %s. Email: %s, Pais: %s\n", cliente2.Nome, cliente2.Email, cliente2.Pais)
  cliente1.Exibe() // Exibindo cliente pelo método...
  cliente2.Exibe() // Exibindo cliente pelo método... -> pela composição os métodos de Cliente vão para ClienteInternacional também, como se fosse uma herança, mas se chama composição
}
```


//Bonus - json
```go
 clienteJson, err := json.Marshal(cliente2) // quando voce pega algo que esta na struct e tranforma em json
  
  if err != nil {
	  log.Fatal(err.Error())
  }
  
  fmt.Println(clienteJson) // [123 34 78 111 ...] -> cada coisa é um slice de bit
  fmt.Println(string(clienteJson)) // json certinho a partir da struct
  {"Nome":"Caio","Email":"c@g.com",...}

//mas os atributos estao em uma formatacao que não é interessante (com a primeira letra maiuscula)
// no go temos algo chamado tags que pode resolver isso (mudar o nome dessas propriedades)

type ClienteInternacional struct {
	Cliente
	Pais string `json:"pais"` // com backticks, é como se fosse uma anoteção, quando rodar uma função de json, o go vai transformar o "Pais" em "pais"
	
fmt.Println(string(ClienteJson))
{"Nome":"Caio","Email":"c@g.com",..., "pais":"Brasil"}
}
```

// Preencher um json em uma struct sem ser de forma manual
// no vídeo ele fala de hidratar o json

```go
jsonCliente4 := {"Nome":"Caio","Email":"c@g.com","CPF":"321","pais":"Brasil"}
cliente4 := ClienteInternacional{} // em branco

json.Unmarshal([]byte(jsonCliente4), cliente4) // altera apenas nesse escopo. mas o cliente4 vai permanecer vazio, precisa colocar o valor por referencia

json.Unmarshal([]byte(jsonCliente4), &cliente4) // agora, manda o valor para o local exato de onde esta armazenando o valor de cliente4. Aqui voce esta passando o endereço, e a função vai receber o parametro em si (por referencia) e não uma copia dele (modifica o valor no endereço da memória)

fmt.Println(cliente4)
// {{Davi d@g.com 321} EUA} -> conseguiu hidratar perfeitamente e tranformar para struct
// quem usa muita api com go voce fica fazendo muito essas conversões
```






















