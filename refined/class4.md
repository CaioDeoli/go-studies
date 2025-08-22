# Aula 4 - Trabalhando com Erros em Go

---

## 🔑 Conceito Fundamental

No Go, **tratar erros é sempre explícito**. Diferente de outras linguagens que utilizam exceções, em Go os erros são valores retornados por funções. Isso força o desenvolvedor a lidar diretamente com a possibilidade de falha.

* Uma função pode retornar **múltiplos valores**, sendo um deles o erro (`error`).
* Caso um erro seja retornado, o fluxo do programa pode ser alterado imediatamente.

---

## 📦 Exemplo prático com `http.Get`

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://google.com.br")

	if err != nil { // Verificação explícita do erro
		// log.Fatal imprime a mensagem e encerra o programa
		log.Fatal("Erro ao fazer comunicação:", err)
	}

	// Se não houver erro, o programa continua normalmente
	fmt.Println("Headers da resposta:", res.Header)
}
```

### ℹ️ Observações

* `panic()` aborta toda a execução, mas **não é recomendado** para erros comuns.
* O uso de `log.Fatal()` é mais adequado, pois loga a mensagem e encerra de forma limpa.

---

## 🚀 Criando Funções que Retornam Erros

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
		log.Fatal(err.Error()) // Tratamento explícito do erro
	}

	fmt.Println("Resultado:", res)
}

// Função que retorna resultado e um erro
func soma(x int, y int) (int, error) {
	res := x + y

	if res > 10 {
		// Retorna um erro personalizado
		return 0, errors.New("Total maior que 10")
	}

	return res, nil // nil indica ausência de erro
}
```

### 📚 Explicação

* A função `soma` retorna **dois valores**: o resultado da soma e um possível erro.
* Caso o resultado seja inválido, retorna-se `0` junto com um erro.
* Caso contrário, retorna o valor válido e `nil` (sem erro).

---

## 🔄 Ignorando o Erro com Blank Identifier `_`

Em alguns casos, você pode **optar por ignorar** o erro. Para isso, utiliza-se o `_` (blank identifier):

```go
package main

import (
	"errors"
	"fmt"
)

func main() {
	// Ignorando o erro com _
	res, _ := soma(7, 2)
	fmt.Println("Resultado:", res)
}

func soma(x int, y int) (int, error) {
	res := x + y

	if res > 10 {
		return 0, errors.New("Total maior que 10")
	}

	return res, nil
}
```

### ⚠️ Observação Importante

* Embora possível, **não é uma boa prática ignorar erros** em Go.
* O `_` deve ser usado apenas quando temos certeza de que o erro não impacta o fluxo.

---

# ✅ Resumo da Aula

* Em Go, **erros são valores retornados explicitamente**.
* Uma função pode retornar múltiplos valores, sendo um deles `error`.
* Sempre verifique `if err != nil` para tratar erros.
* `log.Fatal()` é mais recomendado do que `panic()` para abortar em falhas.
* O blank identifier `_` permite ignorar valores (incluindo erros), mas deve ser usado com cautela.
* **Boa prática:** sempre trate erros, mesmo que seja apenas para logar e seguir.
