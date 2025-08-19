# Aula 3 - Variáveis, Tipos e Pacotes em Go

## 📦 Escopo de Pacotes em Go
- Todo programa em Go é organizado em **pacotes** (`package`).
- O pacote **principal** é chamado `main` e precisa ter uma função `main()` para execução.
- Um pacote pode conter múltiplos arquivos `.go`, desde que todos comecem com `package `.

---

## 🔑 Declaração de Variáveis

### Forma Completa (com `var`)
```go
package main

var a string // declaração de variável

func main() {
    a = "Caio" // atribuição posterior
}
```
- Aqui o tipo foi declarado explicitamente, mas poderia ser omitido para inferência.

---

### Inferência de Tipos (com `:=`)
```go
package main

func main() {
    a := "Caio"
}
```
- O Go **infere o tipo** automaticamente a partir do valor.
- Uma vez atribuído, **não é possível mudar o tipo**:
```go
a := "Caio"
a = 2 // Erro de compilação
```

---

### Regras Importantes
- **Não é permitido redeclarar** variáveis no mesmo escopo:
```go
a := "Caio"
a := "Alice" // Erro
```
- **Go não compila** se houver variáveis declaradas mas **não utilizadas**.

---

## 🔢 Inferência e Tipos
Go possui variações de tipos primitivos (inteiros, floats, strings, booleanos, etc).

Exemplo com diferentes tipos:
```go
package main

import "fmt"

func main() {
    a := 10
    b := "World"
    c := 3.144
    d := false
    e := `Texto com
múltiplas linhas`

    fmt.Printf("%v \n", a) // mostra valor
    fmt.Printf("%T \n", b) // mostra tipo
    fmt.Printf("%v \n", c)
    fmt.Printf("%T \n", d)
    fmt.Printf("%v \n", e)
}
```

---

## 📚 Funções em Go
- Sintaxe de declaração:
```go
func Soma(a int, b int) int {
    return a + b
}
```
- Tipos sempre devem ser informados nos parâmetros.
- Se não houver `return`, não é preciso declarar tipo de retorno.

Exemplo de uso:
```go
package main

import "fmt"

func Soma(a int, b int) int {
    return a + b
}

func main() {
    resultado := Soma(1, 2)
    fmt.Println(resultado) // 3
}
```

---

## 📦 Pacotes Personalizados

Suponha a seguinte estrutura:
```
class3/
 ├── main.go
 └── math/
     ├── operations.go
     └── operationsX.go
```

### `main.go`
```go
package main

import (
    "fmt"
    "class3/math"
)

func main() {
    resultado := math.Soma(1, 2)
    fmt.Println(resultado)
}
```

### `math/operations.go`
```go
package math

func Soma(a int, b int) int {
    return a + b
}
```

### `math/operationsX.go`
```go
package math

var A string = "SHOWWWW"

func SomaX(a int) int {
    return a + 10
}
```

---

## ℹ️ Exportação em Go
- Em Go, **nomes que começam com letra maiúscula são exportados** (públicos).
- **Letra minúscula**: itens **internos ao pacote** (privados).
- Exemplo:
  ```
  pacote.Funcao() // acessível de fora
  pacote.funcao() // inacessível de fora
  ```

Isso vale para:
- Variáveis
- Funções
- Constantes
- Tipos e métodos

---

## 📘 Boas Práticas
- Sempre documente itens exportados:
```go
package math

// SomaX adiciona 10 ao valor passado como argumento.
func SomaX(a int) int {
    return a + 10
}
```
- Mais tarde, essa documentação aparece nos comandos `go doc` ou em ferramentas de documentação.

---

## 🚀 Módulos e Execução
- Se não estiver em `GOPATH`, é necessário inicializar um módulo:
```
go mod init class3
go run .
```

---

# ✅ Resumo da Aula
- `var` declara variáveis e pode ou não indicar tipo.
- `:=` permite declaração com inferência de tipo dentro de funções.
- Go **não permite variáveis não usadas ou mudarem de tipo**.
- Pacotes organizam o código; `main` é o ponto de entrada.
- Funções precisam definir tipos de parâmetros e, se necessário, de retorno.
- Itens exportados começam com **maiúscula**; minúscula restringe ao pacote.
- Documentar funções e variáveis exportadas é uma **boa prática**.
- `go mod init` inicializa módulos para organização e importação correta.

---