Eu estou estudando go lang, e faz tempo que eu não estudo sobre. Poderia me ajudar a me fazer relembrar dos assuntos reformulando e melhorando minhas anotações para eu ler tudo novamente, e dar continuidade?

Abaixo esta o meu rascunho das aulas:

# Aula 1

- Tipada

- Compilada

- Orientada a Go Lang (mistura de POO e PF)

- Multiplataforma

- Gera apenas um executavel (.exe no caso do windows)

# Aula 2

- Go Lang trabalha com pacotes

GOPATH

Local onde vão estar os arquivos do Go

- bin: arquivos binários

- Arquivos que o próprio Go utiliza

- Seus arquivos

- pkg: arquivos pré-compilados

- Ajuda o Go no processo de compilação

- src

- Todo pacote de biblioteca instalado vai para essa pasta

*Acredito que a localização da pasta src mudou*

Exemplo de como instalar um pacote:

```

go get github.com/google/uuid

```

Comandos básicos

```

go version // Versão

go env // Variáveis de ambiente

```

GOOS

Gera os executaveis da maneira que você precisar (windows, mac, linux, ...)

[Hello world](./codes/go/a1.go)

```go

package main

import "fmt"

// Para importar várias coisas

import (

"fmt"

)

func main() {

fmt.Println("Olá mundo!")

}

```

main: função principal do programa/ ponto de entrada

fmt: biblioteca

funções começão com func, e a sintaxe é bem parecida com outras

Para rodar o .go

```

go run main.go

```

Para buildar um projeto go

```

go build main.go

```

* Se não for specificado o arquivo, o próprio go vai buscar onde está a func main()

Para build um projeto para um SO específico

```

GOOS=windows go build main.go

```

Para o projeto ser instalado na pasta bin e eu poder utilizar em qualquer lugar do computador

````

go install

```

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

// Inicie já com o comentário com o nome da função

// SomaX é uma função que faz isso...

func SomaX(a int) int {

return a + 10

}

```

Quando trabalharmos com Go doc, esses comentários vão ser uteis

E isso é só pra funções, constantes, ... exportados

Esses comentários é documentação

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

# Aula 5 funções

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