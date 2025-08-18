# Aula 2: Pacotes, Ambiente e Comandos Básicos do Go

## Estrutura de Pacotes e GOPATH

O Go organiza código em **pacotes** para modularidade e reutilização. Cada pacote é um diretório contendo arquivos `.go` relacionados.

### GOPATH (local de trabalho)

Historicamente, o **GOPATH** era a pasta onde ficavam os arquivos do Go:

* `bin/`: arquivos binários compilados, incluindo os do próprio Go e os seus.
* `pkg/`: arquivos pré-compilados que ajudam no processo de compilação.
* `src/`: todo pacote de biblioteca instalado (código-fonte).

> Observação: Com as versões modernas do Go (1.11+), o **GOPATH tradicional não é mais obrigatório**, pois o Go agora usa módulos (`go.mod`) para gerenciamento de dependências. Assim, a pasta `src` e a estrutura GOPATH se tornaram menos relevantes em projetos novos.

### Exemplo de instalação de pacote via módulos

```bash
go get github.com/google/uuid
```

Este comando baixa e adiciona o pacote `uuid` ao projeto, registrando no arquivo `go.mod`.

## Comandos Básicos do Go

* **Versão do Go:**

```bash
go version
```

* **Variáveis de ambiente do Go:**

```bash
go env
```

* **Compilação cruzada:**

```bash
GOOS=windows go build main.go
```

O `GOOS` permite gerar executáveis para sistemas operacionais específicos.

## Hello World em Go

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Olá mundo!")
}
```

* `package main`: define que este é um pacote executável.
* `import "fmt"`: importa a biblioteca de formatação e impressão.
* `func main()`: função principal, ponto de entrada do programa.

## Execução e Build

* **Executar diretamente:**

```bash
go run main.go
```

* **Compilar executável:**

```bash
go build main.go
```

> Se não especificado, o Go busca automaticamente o arquivo contendo `func main()`.

* **Instalar globalmente na pasta bin:**

```bash
go install
```

Permite executar o programa de qualquer lugar do computador.

## Resumo dos Principais Pontos da Aula

* Go organiza código em **pacotes** para modularidade.
* Estrutura antiga do **GOPATH**: `bin/`, `pkg/`, `src/`, mas hoje é recomendado usar **módulos** (`go.mod`).
* Comando `go get` instala pacotes de terceiros.
* Comandos básicos: `go version`, `go env`, `go run`, `go build`, `go install`.
* `GOOS` permite **compilação cruzada** para diferentes sistemas operacionais.
* Todo programa Go tem `package main` e função `main()` como ponto de entrada.
* Importação de pacotes é feita com `import`.
* Para executar ou gerar executáveis, utiliza-se `go run` ou `go build`, e `go install` para instalar globalmente.
