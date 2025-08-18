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

*Acredito que a  localização da pasta src mudou*

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