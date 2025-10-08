### Teoria
#### Go routines e channels

 conceito de concorrencia e paralelismo
não é a mesma coisa

concorrencia: lidar com varias coisas ao mesmo tempo. ex: no trabalho voce fala no telefone, ai em seguida voce escreve um codigo, em seguida voce volta ao telefone...
paralelismo: voce faz essas coisas ao mesmo tempo. falando no telefone e ao mesmo tempo esta digitando

concorrencia:

+----------------------                                   +----------------------
				  +----------------------

paralelismo

+----------------------
+---------------------- 

toda vez que iniciamos um programa, isso é iniciar um processo, que é gerenciado pelo SO de forma geral

ele é independente de outro processo

processo 1 não interfere no espaço de memória do processo 2

pro 1
a = 10

pro 2
a = 20

cada "a" não é igual ao outro os processos não interferem nas variaveis dos outros processo de forma geral

um processo nao compartilha a memoria com outro

IPC -> excessao a regra q consegue compartilhar

quando inicia um processo, o sistema operacional aloca um espaço/endereço de memória

--

thread = fio de execução
ex.: executa isso, depois executa isso, depois executa isso
+------------ +-------------- +----------

--

Vamos supor que a thread ira executar um calculo, eu tenho varias threads que realizam cada um um fio de execução (que executa um cálculo final)

--

pelo que eu entendi, um thread é um processo que pode compartilhar do mesmo espaço alocado

--

Vamos imaginar que temos duas threads e queremos executar elas ao mesmo tempo

Thread 1 e thread 2, ambas realizam um calculo especifico
que podemos executa-las de maneira concorrente ou paralela, e de uma forma ou de outra, trabalhamos com o programa mais rapidamente

quando estamos trabalhando com threads esse espaço de memória é compartilhado

thread 1 e thread 2
utilizam
variavel a = 10

thread 1 utiliza o a = 10 pois ele ta guardado na memoria que tem um endereço
imagina que thread 2 altera a variavel a
thread 2 -> variavel agora é 20
thread 1 vai dar um tilte porque a thread 2 alterou
pontex: thread 1 fala  "o thread 2, eu to mexendo com a, enquanto eu to fazendo isso, voce nao pode alterar o valor dessa variavel, quando eu termino, ai sim eu libero essa thread 2 pode mexer" -> isso da muito trabalho e é muito confuso e complicado

Go consegue trabalhar com threads e sincronização dessa forma, mas também pode trabalhar com goroutines e channels

toda vez que o SO começa a trabalhar e cria uma nova thread ele aloca 1mb. Para cada thread é 1mb que ele aloca. Que é bastante

Cada vez que cria uma nova thread o programa chama o SO com o SysCall para que ele crie uma nova thread, que isso é custoso

Quando trabalhamos com varias threads temos o scheduler que determina quando cada thread vai ser executada
- scheduler preeptivo
	- Thread A voce tem 10s pra executar o calculo, depois disso eu vou fazer a Thread B fazer o calculo dela, depois de 10s volta pra A, depois de 10s... Ele coloca um tempo especifico para a thread ser interrompida para a outra ser executada. Isso tem um custo: cada vez que eu mudo de A para B temos a mudança de CONTEXTO (conceito). Ex.: dois pizzaiolos vao fazer cada um uma pizza diferente, o pizzaiolo A tem 1 minuto para fazer as coisas no balcao disponivel, quando passar 1 minuto tudo que esta no balcao vai ser removido, e ai o pizzaiolo B vai entrar com suas coisas no balcão... Essa mudança de contexto tem um preço
- scheduler cooperativa
	- trabalha pizzaiolo A que o B ta esperando para iniciar a tarefa dele. Existe algumas condições, eu to esperando a massa da pizza crescer (então eu estou esperando algo acontecer) nesse momento que o pizzaiolo A ta esperando voce retira o pizzaiolo A para ele aguardar e coloca o pizzaiolo B para trabalhar. Tendo a vantagem de nao ficar interrompendo uma e outra. Mas uma thread pode monopolizar a mesa fazendo as outras esperarem.

Go trabalha como se fosse com o scheduler cooperativo (que o nome vira RUNTIME)
Mas de forma diferente
As threads abaixo sao criadas pelo SO
Essas threads ex.:
Thread A - CALCULO
Thread B - CALCULO - dividem a = 20

O go tem o runtime (todas as bibliotecas dele)


[ MEU CODIGO ]
+
[ RUNTIME ] 
=binario

ao inves de chamar o SO, o proprio GO cria o scheduler e suas threads, e ele próprio gerencia isso
quem gerencia essas threads é o runtime

chamamos a threads de USER LAND || GREEN THREADS
vantagem Não é realizado uma chamada para o SO
vantagem uma thread no go ocupa 2kb de memória
e ai podemos compartilhar a memória entre threads
dificilmente voce vai querer utilizar multitex ou coisa do tipo, pois o go trabalha com chanels

como faz pra nao fazer uma "thread nao monopolizar o processador" voce faz o go pular coisas que monopolizam
ex.:
- se tiver de acessar disco ou fazer qualquer operação bloqueante, vai pra o proximo
- se tiver que fazer uma chamada externa/api que bloquei vai pra proxima
- timeout ou sleepy -> vai pra proxima

ele lida de forma coperativa

mas o go também pode trabalhar de forma pre-eptiva

---


---
### Prática

```go
package main

func contador(tipo string) {
	for i := 0; i < 5; i++ {
		fmt.Println(tipo, i)
	}
}

func main() {
	contador(tipo: "sem go routine")
	go contador(tipo: "com go routine")
	fmt.Println("Hello 1")
	fmt.Println("Hello 2")
}

go run main.go
// não apareece o com go routine
// isso cria uma thread
// não deu tempo de essa thread acontecer por isso não apareceu
```

!=

```go
package main

func contador(tipo string) {
	for i := 0; i < 5; i++ {
		fmt.Println(tipo, i)
	}
}

func main() {
	contador(tipo: "sem go routine")
	go contador(tipo: "com go routine")
	fmt.Println("Hello 1")
	fmt.Println("Hello 2")
	time.Sleep("fim...")
}

go run main.go

/* 
sem go routine 0
sem go routine 1
sem go routine 2
sem go routine 3
sem go routine 4
Hello 1
Hello 2
com go routine 0
com go routine 1
com go routine 2
com go routine 3
com go routine 4
fim...

// Tudo é executado e o go contador() esta executando de forma concorrente

*/
```


```go
package main

import "fmt"

func contador(tipo string) {
	for i := 0; i < 5; i++ {
		fmt.Println(tipo, i)
		time.Sleep(time.Second)
	}
}

func main() {
	contador(tipo: "sem")
}

// go run
sem 0
sem 1
sem 2
sem 3
sem 4
```

```go
package main

import "fmt"

func contador(tipo string) {
	for i := 0; i < 5; i++ {
		fmt.Println(tipo, i)
		time.Sleep(time.Second)
	}
}

func main() {
	contador(tipo: "a")
	contador(tipo: "b")
}

// go run
a 0
a 1
a 2
a 3
a 4
b 0
b 1
b 2
b 3
b 4
// esperou acabar um para ir para o outro
// isso é serial
```

Agora se fosse dessa

```go
package main

import "fmt"

func contador(tipo string) {
	for i := 0; i < 5; i++ {
		fmt.Println(tipo, i)
		time.Sleep(time.Second)
	}
}

func main() {
	go contador(tipo: "a")
	go contador(tipo: "b")
	time.Sleep(time.Second * 10)
}

go run
//a0
//b0
//b1
//a1
//a2
//b2
//b3
//a3
//b4
//a5
// parece que estão rodando ao mesmo tempo, mas não estão. um executa e enquanto esta esperando a resposta dessa, já começa a executar a outra e assim vai
// É mais ou menos algo assincrono do NodeJS. Mas no caso do node é um loop
```


preeptiva e coperativa
```go
package main

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println("Começou")
	
	go func() {
		for {
		} // loop infinito
	}() 
	
	time.Sleep(time.Second)
	fmt.Println("Terminou")
}
```
como nosso computador é multicore, ele consegue rodar esses cores em paralelo
o go agora só ta rodando em um core nessa funcao

o go trabalha de forma cooperativa, ele espera terminar tudo, para depois começar o que vem depois dessa func()

func main no go é uma thread
esse go func() {} roda um loop infinito

na versão 1.13 do go
ele vai travar o meu problema ate que o for termine para rodar o resto
-> o loop dor for vai ficar infinito e nunca vai terminar
-> porque o programa nao faz nenhum chamada em outra função, para o scheduler do Go (runtime), retornar e continuar o codigo, esperando receber algo dessa outra função que ta como se fosse (async)
-> ja que essa mudança nao ta acontecendo, trava tudo

Na versão 1.17
Ele imprime terminou
Não trava porque o Go adicionou um recurso no scheduler, para que o Go também trabalhe de forma preepitiva quando for necessário. Nesse caso ele viu um loop infinito (uma função que ta enrolando), e ai da continuidade ao codigo (o processo continua)

---

### Channels

um channel serve para fazer um thread se comunicar com a outra

no go, no lugar de ter conflito que acontece nas threads tradicionais. Com os channels as threads se comunicam compartilhando memória, e não apenas compartilhar um espaço da memória. Ex.: a thread A fala com a thread B para ambos compartilharem uma variavel especifica



```go
package main

func main() { // func main é uma thread e go func() é outra thread

	hello := make(chan string)
	
	go func() {
	  hello <- "Hello world" // diz que hello vai receber hello world
	}() 
	
	result := <-hello // toda vez que voce tiver um valor nesse canal manda para essa variavel result
	// deixa a variavel result pegar esse valor
	fmt.Println(result) // vai aparecer hello world por que voce criou um canal, com a linha "hello <- 'Hello world'" pela thread "go func()", e pela linha "hello := make(chan string)" voce consegue fazer com que a thread main fale com a thread go func T1 <-> T2

}
```


```go
func main() {
	forecver := make(chan string)
	
	go func() {
		x := true
		for {
			if x == true {
				continue
			}
		}
	}()
	
	fmt.Println("aguardando para sempre")
	<- forever // ele ta aguardando alguem colocar algo no forever
	// e fica aguardando para sempre independente do resultado
	
}
```

==

```go
func main() {
	forecver := make(chan string)
	
	go func() {
		for {
		}
	}()
	
	fmt.Println("aguardando para sempre")
	<- forever // ele tambem fica aguardando para sempre
	
}
```

ele  fica mesmo que o for tenha parado de rodar ja que o Go funciona de forma preeptiva nesses casos, mas ele só parou o for, ele nao parou o aguardo do forever receber um valor. O nosso chanel fica parado



---
```go
package main

func main() {
	hello := make(chan string)
	
	go func() {
	    hello <- ""hello world"
	}()
	
	select {
	case x := <-hello // caso receba algo do hello
		fmt.Println(x)
	} // caso venha algum valor de hello e vem atribuido de x, imprima x
	// imprime hello world porque chegou algo prar o X
	// isso fica esperando a condicao ser executada
}
```


```go
package main

func main() {
	hello := make(chan string)
	
	go func() {
		time.Sleep(time.Second * 2) // demora 2 segundos para ir pra proxima linha
		hello <- ""hello world"
	}()
	
	select {
	case x := <-hello
		fmt.Println(x)
	default:
		fmt.Println("default")
	}
	
	// imprime default porque recebeu default antes que algo tivesse sido colocado em hello pela go func
}
```

```go
package main

func main() {
	hello := make(chan string)
	
	go func() {
		time.Sleep(time.Second * 2)
		hello <- "hello world"
	}()
	
	select {
	case x := <-hello
		fmt.Println(x)
	default:
		fmt.Println("default")
	}
	
	time.Sleep(time.Second * 5) // mesmo com isso, o select ja sai default e nao executa mais o case
}
```

---
```go
func main() {
	queue := make(chan int)
	
	go func() {
		i := 0
		for {
			time.Sleep(time.Second)
			queue <- i // se o queuoe for esvaziado é que ele volta nesse loop. se ninguem ler esse queue em outra thread, essa linha queue vai parar a execucao e aguardando a sua leitura
			i++
		}
	}()
	
	for x := range queue { // captura valor de queue, atribuir a x, e vai imprimir
		fmt.Println(x)
	}
}
```
resultado:
1
2
3
4
5
6
7
por que o range fica lendo o X, enquanto ele nao termina de ler, o codigo fica parado no for

essa variavel queue voce esta compartilhando entre main e go func

a cada vez que o range itera por x. o for de fo func é acionado pra adicionar mais 1

---
webservice GO

```go
package main

func worker(workerId int, msg chan) {
	for res := range msg {
		fmt.Println("Worker", workerId, "Msg:", msg)
		time.Sleep(time.Second) // toda vez que roda uma mensagem voce faz isso aqui dormir
	}
}

func main() {
	msg := make(chan int)
	go worker(1, msg) // vai deixar worker funcionando no background
	
	for i:= 0; i < 10; i++ {
		msg <- i
	}
	
	// a cada loop desse for, o worker vai pegar essa mensagem e vai imprimir
	
	Woker: 1 Msg: 0
	Woker: 1 Msg: 1
	Woker: 1 Msg: 2
	Woker: 1 Msg: 3
	Woker: 1 Msg: 4
	...
}

/*
O for de cima fica no background executando por consta do go worke
Enquanto ele nao terminar, ele nao pode dar o proximo passo
*/
```

Ex.: imagina o webservice assim: toda vez que acessarmos uma pagina, precisamos esperar o servidor processar o processo, para que outro visitante possa acessar a pagina. Tem que estar um para executar o outro

Como fazer pra 2 pessoas acessarem a mesma tela? cria 2 workers, enquanto um ta esperando, voce passa pra o outro com outra tarefa

NAO ENTENDI ESSA PARTE DE BAIXO

```go
package main

func worker(workerId int, msg chan) {
	for res := range msg {
		fmt.Println("Worker", workerId, "Msg:", msg)
		time.Sleep(time.Second) // toda vez que roda uma mensagem voce faz isso aqui dormir
	}
}

func main() {
	msg := make(chan int)
	go worker(1, msg) // vai deixar worker funcionando no background
	
	for i:= 0; i < 10; i++ {
		msg <- i
	}
	
	// a cada loop desse for, o worker vai pegar essa mensagem e vai imprimir
	
	Woker: 1 Msg: 0
	Woker: 1 Msg: 1
	Woker: 1 Msg: 2
	Woker: 1 Msg: 3
	Woker: 1 Msg: 4
	...
}
```

exemplo com 2 workers, onde enquanto um esta dormindo, o ouitro executa
```go
package main

func worker(workerId int, msg chan) {
	for res := range msg {
		fmt.Println("Worker", workerId, "Msg:", msg)
		time.Sleep(time.Second)
	}
}

func main() {
	msg := make(chan int)
	go worker(1, msg)
	go worker(2, msg)
	
	for i:= 0; i < 10; i++ {
		msg <- i
	}
	...
}
```
euqnato o worker 1 tava rodando, o 2 pegou uma outra mensagem

nesse caso abaixo encerra o programa por que tem baste worker
```go
package main

func worker(workerId int, msg chan) {
	for res := range msg {
		fmt.Println("Worker", workerId, "Msg:", msg)
		time.Sleep(time.Second)
	}
}

func main() {
	msg := make(chan int)
	go worker(1, msg)
	go worker(2, msg)
	go worker(3, msg)
	go worker(4, msg)
	go worker(5, msg)
	
	for i:= 0; i < 10; i++ {
		msg <- i
	}
	...
}
```
no webservice podemos colocar muitos desses workers e fazer eles trabalharem ao mesmo tempo
imprime os 5

imagina um webservice processando uma pagina, e eu ja posso deixar resolver

com channels vou pode pegar variaveis que estao em um processo e lhe dando nesse outro processo

voce pode fazer um worker mandar mensagem pra outro e pra o outro...

*Tem outras coisas. Mas isso é uma base um pouco aprofundada

go tem garbarege collector que é uma go route

no SO só tem um processo do go rodando, mas no go voce tem as go threads e o runtime (que é como se fosse o scheduler)

