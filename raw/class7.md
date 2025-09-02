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
fmt.Printf("Nome: %s. Email: $s. Pais: $s\n", cliente3.Nome, cliente3.Email, cliente3.Pais)

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

````