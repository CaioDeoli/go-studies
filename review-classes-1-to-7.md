# Go — Resumo e Anotações Reformuladas

> Versão organizada das suas anotações das Aulas 1–7, com correções, explicações, exemplos e próximos passos.

---

## Índice

1. Visão geral (Aula 1)
    
2. Pacotes, GOPATH e módulos (Aula 2)
    
3. Variáveis, tipos e escopos (Aula 3)
    
4. Erros e tratamento (Aula 4)
    
5. Funções e métodos (Aula 5)
    
6. Ponteiros (Aula 6)
    
7. Structs, composição e JSON (Aula 7)
    
8. Boas práticas & ferramentas
    
9. Exercícios práticos e próximos passos
    

---

## 1. Visão geral (Aula 1)

- **Go** é uma linguagem **tipada**, **compilada** e **multiplataforma**.
    
- Mistura conceitos de POO e programação funcional/imperativa — conhecida como _the Go way_.
    
- Compila para um único executável (ex.: `.exe` no Windows).
    
- Aplicações típicas: serviços web, CLI, ferramentas de rede, sistemas concorrentes.
    

---

## 2. Pacotes, GOPATH e módulos (Aula 2)

### Estrutura histórica (GOPATH)

Tradicionalmente (`GOPATH`) havia três pastas importantes:

- `bin` — binários instalados
    
- `pkg` — pacotes pré-compilados
    
- `src` — código fonte dos pacotes
    

> **Nota:** Desde Go 1.11+ o sistema de **módulos** (`go modules`) é o padrão (com `go.mod`). Você ainda verá `GOPATH` em projetos legados, mas para projetos novos use `go mod init`.

### Comandos básicos

```bash
go version   # mostra a versão do Go
go env       # mostra variáveis de ambiente (inclui GOPATH, GOMOD, GOOS, GOARCH)
```

### Instalar/obter pacotes

- `go get github.com/google/uuid` — adiciona/modifica dependência do módulo (e em versões antigas instalava binários no GOPATH/bin).
    
- `go mod init <module/name>` — inicia um módulo dentro da pasta atual.
    
- `go install ./...` — compila e instala binários do módulo (atualmente `go install pkg@version` pode instalar binários sem modificar `go.mod`).
    

### Cross-compilation

Você pode compilar para outro SO definindo `GOOS` e `GOARCH`:

```bash
GOOS=windows GOARCH=amd64 go build -o app.exe
```

### Execução e build

```bash
go run main.go         # compila em memória e executa
go build main.go       # gera um binário no diretório atual
go install             # instala o binário no GOPATH/bin ou em $(go env GOBIN)
```

---

## 3. Variáveis, tipos e escopos (Aula 3)

### Declaração e inferência

```go
var a string     // declaração explícita
a = "Caio"      // atribuição

b := "Alice"    // declaração + inferência (curto)
```

- `:=` só funciona dentro de funções.
    
- Tipos são estáticos: não é possível reatribuir um tipo diferente à mesma variável.
    

### Zero values

Cada tipo tem um valor zero (int -> 0, string -> "", bool -> false, ponteiro/interface -> nil).

### Impressão formatada

```go
fmt.Printf("%v\n", a) // valor
fmt.Printf("%T\n", a) // tipo
```

### Escopo de pacote e exportação

- Identificadores que começam com **maiúscula** são exportados (visíveis fora do pacote).
    
- Identificadores com **minúscula** são não-exportados (visíveis apenas dentro do pacote).
    

Exemplo de organização com módulos:

```text
myproj/
  go.mod
  main.go
  math/
    operations.go  // package math
```

`main.go` importa como `import "myproj/math"` (ou o path do módulo). Lembre-se de rodar `go mod init` quando não estiver em GOPATH.

---

## 4. Como Go trata erros (Aula 4)

- Erros são valores do tipo `error` e são retornados explicitamente.
    
- Padrão idiomático:
    

```go
res, err := http.Get("https://example.com")
if err != nil {
    // tratar ou retornar
}
```

- Evite `panic` a menos que seja um erro fatal irrecuperável (ex.: estado inconsistente). Use `log.Fatal` para encerrar com log.
    
- Criando erros:
    

```go
if cond { return 0, errors.New("mensagem") }
```

- Desde Go 1.13 há suporte para wrapping com `fmt.Errorf("...: %w", err)` e `errors.Is` / `errors.As`.
    
- Se não precisar de um retorno (ignorar), use o identificador em branco `_`.
    

---

## 5. Funções (Aula 5)

### Assinaturas e retornos

```go
func soma(a int, b int) int {
  return a + b
}
```

### Retornos múltiplos

```go
func soma(a, b int) (int, string) {
  return a + b, "somou"
}
```

### Retornos nomeados

```go
func soma(a, b int) (result int) {
  result = a + b
  return
}
```

### Variádicas

```go
func somaTudo(xs ...int) int {
  s := 0
  for _, v := range xs {
    s += v
  }
  return s
}
```

### Funções anônimas e closures

```go
f := func(x int) int { return x * x }
fmt.Println(f(3))

// função que retorna função
func makeSum(xs ...int) func() int {
  s := 0
  for _, v := range xs { s += v }
  return func() int { return s }
}

fmt.Println(makeSum(1,2,3)())
```

---

## 6. Métodos e ponteiros (Aula 6)

### Structs e métodos

```go
type Carro struct { Nome string }

func (c Carro) Andar() { fmt.Println(c.Nome, "andou") }
```

- **Receiver por valor** (`c Carro`) copia a instância — alterações no método **não** afetam o valor original.
    
- **Receiver por ponteiro** (`c *Carro`) permite modificar o valor original:
    

```go
func (c *Carro) SetNome(n string) { c.Nome = n }
```

### Ponteiros básicos

- `&x` pega o endereço de `x`.
    
- `*p` desreferencia o ponteiro `p`.
    

Exemplo:

```go
func incrementar(a *int) { *a = *a + 1 }

x := 10
incrementar(&x) // x agora é 11
```

Quando usar ponteiro em métodos: quando precisar modificar o estado do receiver ou evitar cópias grandes (ex.: structs grandes).

---

## 7. Structs, composição e JSON (Aula 7)

### Struct básico

```go
type Cliente struct {
  Nome  string
  Email string
  CPF   int
}
```

### Composição (embedding)

```go
type ClienteInternacional struct {
  Cliente
  Pais string
}

ci := ClienteInternacional{
  Cliente: Cliente{Nome: "Caio", Email: "c@g.com", CPF: 123456789},
  Pais: "EUA",
}
fmt.Println(ci.Nome) // acesso direto
```

### Tags JSON e encoding/json

```go
import "encoding/json"

type Cliente struct {
  Nome  string `json:"nome"`
  Email string `json:"email"`
  CPF   int    `json:"cpf"`
}

c := Cliente{"Caio", "c@g.com", 123}
js, err := json.Marshal(c)
if err != nil { log.Fatal(err) }
fmt.Println(string(js))

// Unmarshal
var c2 Cliente
if err := json.Unmarshal(js, &c2); err != nil {
  log.Fatal(err)
}
```

---

## 8. Boas práticas & ferramentas

- **Formato:** `gofmt` / `go fmt` — formate sempre o código.
    
- **Lint:** `go vet` e linters (ex.: `golangci-lint`) ajudam a detectar problemas.
    
- **Documentação:** comente entidades exportadas. Ex.:
    
    ```go
    // Soma retorna a soma de a e b.
    func Soma(a, b int) int { return a + b }
    ```
    
- **Módulos:** use `go.mod` para gerenciar dependências.
    
- **Teste:** escreva testes com `go test` e use `table-driven tests`.
    
- **Controle de erros:** trate erros imediatamente; prefira retornar erros ao caller.
    
- **Evite `panic`** em código de biblioteca — use em `main` para falhas irreversíveis.
    

---

## 9. Exercícios práticos e próximos passos

### Exercícios (curto prazo)

1. Converter suas pastas de `class3` para usar `go mod init` e rodar `go run ./...`.
    
2. Escrever uma função variádica que calcule média e testar com diferentes entradas.
    
3. Criar um pequeno CLI que lê JSON (um slice de `Cliente`), filtra por país e imprime resultados.
    
4. Escrever teste para a função `somaTudo`.
    

### Tópicos para continuar (médio prazo)

- Concorrência: goroutines, canais, `select`, `sync`, `context`.
    
- Interfaces e polimorfismo — como escrever código orientado a interfaces.
    
- Testes avançados e benchmarks (`testing.B`).
    
- Profiling e performance (`pprof`).
    
- Construir e distribuir binários (`go install`, release cross-compile).
    

### Mini-projeto sugerido

Crie uma pequena API REST em Go que:

- Expõe endpoints para criar/listar clientes (JSON),
    
- Armazena em memória (map),
    
- Suporta filtros (por país) e paginação simples,
    
- Inclui testes unitários para os handlers e funções de filtragem.
    

---

## Observações finais

- Revise conceitos de módulos (`go.mod`) e `GOPATH` — hoje em dia módulos são o padrão.
    
- Se quiser, posso transformar essas notas em: flashcards, exercícios com soluções, cheatsheet PDF, ou adicionar exemplos resolvidos (por aula).
    

_FIM_