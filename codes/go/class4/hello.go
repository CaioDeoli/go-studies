// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	res, err := http.Get("http://go32131ogle.com.br")
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	fmt.Println(res.Header)
// }

package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {

	res, err := soma(17, 2)
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
