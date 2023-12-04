package main

import (
	"errors"
	"fmt"
)

func aplicarEsomar(slice []int, funcao func(int) int) (int, error) {
	if len(slice) == 0 {
		return 0, errors.New("Slice vazio")
	}

	soma := 0
	for _, valor := range slice {
		soma += funcao(valor)
	}

	return soma, nil
}

func main() {
	numeros := []int{1, 2, 3, 4, 5}
	funcaoDuplicar := func(x int) int {
		return x * 2
	}

	resultado, err := aplicarEsomar(numeros, funcaoDuplicar)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Soma dos Resultados:", resultado)
	}
}
