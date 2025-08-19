# Aula 3 - Variáveis, tipos e pacotes

Escopo dos pacotes go
Declaração de variaveis

Declaração de variaveis

---
No go não é obrigatório declarar o tipo da variavel para em outra linha você atribuir o valor
```go
package main

var a string // declaração de variavel

func main() {
	a = "Caio" // atribuição de valor a variavel
}

```
---
Nesse caso se declarou a variavel e atribuiu um valor a ela
O go tenta inferir/adivinhar o tipo do valor que voce colocou
```go
package main

func main() {
	a := "Caio"
}
```

Voce nao pode mudar o tipo da variavel
```go
package main

func main() {
	a := "Caio"
	a = 2 // Erro
}
```

// Go não compila de tiver uma variavel que nao ta sendo utilizada
// Go não permite fazer besteira

Como a variável já foi declarada, você não pode redeclarar
```go
package main

func main() {
	a: "Caio" // Erro
	a: "Alice" // Erro
}
```

Inferencias
// Existem várias variações de tipos (variações de inteiros, floats, ...)
```go
package main

import "fmt"

func main() {
	a := 10
	b := "World"
	c := 3.144
	d := false
	e := `uouuuuuuu
	
	legal
	`

	fmt.Printf("%v \n", a) // o valor de a passa para %v nessa formatação
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
}
```

```go
package main

import "fmt"

func main() {
	a := 10
	b := "World"
	c := 3.144
	d := false
	e := `uouuuuuuu
	
	legal
	` // Back ticks consegue passar valores multi line sem problema

	fmt.Printf("%T \n", a) // %T mostra o tipo da variavel
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)
}
```

Como, quando e onde usar essas variaveis
```go
func main() {}

func Soma(a int, b int) int {} // Assim que se declara uma função e depois de () vem o tipo do retorno
// Se voce não passa nada, é porque nao tem retorno a função
```

muito comum em go usar pacotes, métodos, funções

```go
package main

import "fmt"

func main() {
	resultado := Soma(1, 1) // Soma faz parte desse meu pacote main, por isso quando chamo soma ela busca dentro do meu pacote
	fmt.Printf("%T", resultado)
	fmt.Printf("%n", resultado)
}

func Soma(a int, b int) {
	return a + b
}
```

Vamos supor que minha função soma esta em outro pacote
cod 1 class3/main.go
```go
package main

import (
	"fmt"
	"class3/math" // Importa
)

func main() {
	resultado := math.Soma(1, 1)
	fmt.Printf("%T", resultado)
	fmt.Printf("%n", resultado)
}
```

cod 2 class3/math/operations.go
```go
package math

func Soma(a int, b int) {
	return a + b
}
```

// Se voce nao estiver dentro de GOPATH, voce precisará dar um cmd go mod init class3 para criar um módulo
cd codes\go\class3
go run .

Se voce criar mais um arquivo por exemplo
cod 3 class3/math/operationsX.go
```go
package math

func SomaX(a int) int {
	return a + 10
}
```

e em

cod 1 class3/main.go
```go
package main

import (
	"fmt"
	"class3/math" // Importa
)

func main() {
	resultado := math.SomaX(1, 1)
	fmt.Printf("%T", resultado)
	fmt.Printf("%n", resultado)
}
```

---

cod 3 class3/math/operationsX.go
```go
package math

var A string = "SHOWWWW"

func SomaX(a int) int {
	return a + 10
}
```

Sempre que você tiver importado um pacote em go, não importa quantos arquivos esse pacote tenha, eu vou conseguir acessar todas as funções dele

```go
package main

import (
	"class3/math"
	"fmt"
)

func main() {
	fmt.Printf("%v", math.A)
}
```
A imprime porque a variavel esta dentro de escopo de math -> muito importante

```go
func main() {
	resulta := math.Soma(1, 2)
	fmt.Printf("%v", resultado)
}
```
diferente de
```go
func main() {
	resulta := math.soma(1, 2)
	fmt.Printf("%v", resultado)
}
```
soma() da erro
o go trabalha com metodos, variaveis, funcoes exportadas e não exportadas
significa que toda vez que voce acessa um pacote de fora main acessa math
voce so consegue acessar essas coisas se elas forem exportadas
para voce dizer que algo é exportado a letra tem que ser maiuscula
primeira letra minuscula significa que eu não posso acessar algo de fora (função, variavel ou método) do pacote

é parecido com private e public em POO

boas práticas:
toda vez que voce exportar algo para fora
coloque uma documentação do que isso faz

```go
package math

// SomaX é uma função que faz iss...
func Soma

```
//////////