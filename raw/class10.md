aula 10. context

ctx

para que ele serve?
se voce for trabalhar com grandes aplicações esse pacote é essencial para que voce tenha performance e ter ocntrole de tudo que acontece no sistema

tudo é metrificado quando trabalhamos com grandes aplicações e tudo tem que ter controle

vamos supor que vamos fazer uma consulta no banco, e uma requisição dando gargalo, ele bloqueia o resto. o contexto, define nas linguagens 

ele é um carregador de informações 
se algo der errado você pode parar tudo que estava fazendo na hora

vamos entender como esse pacote funciona

vamos imaginar que temos essa relação de arquivos:
/hotel
	main.go
go.mod

somos tipo um booking da vida

precisamos dar limites para quando precisamos interromper coisas e dar uma resposta para o usuário

conextos trabalham como se fossem árvores, então um contexto pode chamar o outro, para que voce consiga ter controle da sua aplicação

nossa aplicação sempre vai ter um contexto em branco, o background

```go
package main

import "context"

func main() {
	ctx := context.Background() // contexto pai
	ctx, cancel := context.WithCancel(ctx) // WithCancel -> ele retorna uma função de cancelamento. e quando essa função for disparada, todos que tiverem com esse contexto receberão um sinal falando que esse contexto já era. por boas práticas voce sempre cancela no final da execução a não se que ele seja cancelado antes
	defer cancel() // defer espera tudo ser executado para executar o cancel
	// eu quero um relógio que espera algo acontecer
	// se o relogio disparar, ele para o bookHotel
	
	// se essa funcao go func der cancel, antes do `case <- time.After(time.Second * 5):` então ele entra no primeiro case ctx.Done()
	go func() { // roda no background essa go
		time.Sleep(time.Second * 4) // se fosse 10 o quarto seria reservado. Caso fosse 4 o quarto não seria reservado, ia entrar no ctx.Done
		cancel()
	}()
	
	bookHotel(ctx) // cancel cancela no final
}

// iniciar o processo de reserva de um quarto no hotel, mas caso demore muito, eu encerro o processo para não pendurar o cliente e ele ficar esperando muito tempo

// ctx sempre o primeiro
func bookHotel(ctx context.Context) {
	select { // fica aguardando alguma condição dar certo -> que são channels no final das contas
		case <- ctx.Done(): // se nosso contexto não for cancelado (se alguem não cancelar o processo)
			fmt.Println("Tempo excedido para bookar o quarto")
		case <- time.After(time.Second * 5): // até 5 segundos o quarto é reservado. durante esse tempo, você que estiver reservando, pode cancelar o processo
			fmt.Println("Quarto reservado com sucesso")
	}
}
```

nesse caso acima utilizamos go func com time.Sleep

mas poderia ser qualquer outra coisa

vamos imaginar que duas APIs estão tentando reservar o mesmo quarto, então o primeiro que reservar cancela o bookHotel do outro

```go
package main

import "context"

func main() {
	ctx := context.Background() 
	ctx, cancel := context.WithCancel(ctx) 
	defer cancel()
	
	go func() { 
		// aqui voce pode colocar qualquer tarefa para executar o cancel, não precisa ser o time.Sleep
		cancel()
	}()
	
	bookHotel(ctx) 
}

func bookHotel(ctx context.Context) {
	select { 
		case <- ctx.Done():
			fmt.Println("Tempo excedido para bookar o quarto")
		case <- time.After(time.Second * 5): 
			fmt.Println("Quarto reservado com sucesso")
	}
}
```

isso faz interromper processos desnecessários
imagina que a sua aplicação ta fazendo vários processos, para no final dizer pra o usuário que aquele quarto foi reservado por outra pessoa
não precisa, voce pode avisar isso antes

além do WithCancel, temos o WithTimeout (após um certo tempo, ele cancela automaticamente)

nesse caso abaixo não precisaria da go func, porque ele usa o WithTimeout que faz a mesma coisa nesse caso
```go
package main

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
		case <- ctx.Done():
			fmt.Println("Tempo excedido para bookar o quarto")
		case <- time.After(time.Second * 5):
			fmt.Println("Quarto reservado com sucesso")
	}
}
```

o mesmo acontece com o banco

o mesmo pode acontecer com uma requisição http