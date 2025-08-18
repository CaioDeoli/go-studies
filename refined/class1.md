# Aula 1: Introdução ao Go

## Características da Linguagem Go

* **Tipada**: Go possui tipagem estática, ou seja, o tipo de cada variável é conhecido em tempo de compilação. Isso ajuda a evitar erros comuns de tipo em tempo de execução.

* **Compilada**: Go é compilada diretamente para código nativo, o que resulta em programas rápidos e eficientes. O compilador verifica erros antes de gerar o executável.

* **Paradigma**: Go adota uma abordagem híbrida entre **Programação Orientada a Objetos (POO)** e **Programação Funcional (PF)**. Embora não tenha classes tradicionais, possui structs, métodos e interfaces, permitindo abstração e modularidade.

* **Multiplataforma**: Go permite compilar o mesmo código para diferentes sistemas operacionais (Windows, Linux, macOS) e arquiteturas (x86, ARM).

* **Executável único**: Ao compilar, Go gera apenas um executável independente (por exemplo, `.exe` no Windows) sem a necessidade de instalar runtime externo.

## Exemplos Práticos

### Hello World

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

Este código demonstra a estrutura básica de um programa Go: pacote principal, importação de biblioteca e função `main`.

### Declaração de Variáveis Tipadas

```go
var idade int = 25
nome := "Caio" // inferência automática de tipo
```

O Go permite tanto declaração explícita de tipo quanto inferência de tipo com `:=`.

### Estrutura e Método (POO em Go)

```go
type Pessoa struct {
    Nome string
    Idade int
}

func (p Pessoa) Saudacao() string {
    return "Olá, meu nome é " + p.Nome
}
```

Aqui, `Pessoa` é uma struct com um método associado, ilustrando conceitos de POO.

## Resumo dos Principais Pontos da Aula

* Go é **tipada** e **compilada**, garantindo segurança de tipo e performance.
* Suporta **POO** e **programação funcional**, apesar de não ter classes tradicionais.
* É **multiplataforma**, permitindo compilação cruzada para diferentes sistemas operacionais.
* Gera **executáveis únicos**, facilitando distribuição de aplicativos.
* Estrutura básica de um programa Go envolve `package main`, `import` e `func main()`.
* Suporta **declaração de variáveis** explícita ou por inferência.
* **Structs e métodos** permitem abstração e modularidade no código.
