package main

import (
	"errors"
	"fmt"
)

func aplicarFuncao(slice []int, funcao func(int) int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("Slice vazio")
	}

	resultados := make([]int, len(slice))
	for i, valor := range slice {
		resultados[i] = funcao(valor)
	}

	return resultados, nil
}

func main() {
	numeros := []int{1, 2, 3, 4, 5}
	funcaoDuplicar := func(x int) int {
		return x * 2
	}

	resultado, err := aplicarFuncao(numeros, funcaoDuplicar)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}
}
