package main

import (
	"errors"
	"fmt"
	"sort"
)

func ordenarCrescente(slice []int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("Slice vazio")
	}

	ordenado := make([]int, len(slice))
	copy(ordenado, slice)
	sort.Ints(ordenado)

	return ordenado, nil
}

func main() {
	numeros := []int{5, 2, 8, 1, 9, 3}

	resultado, err := ordenarCrescente(numeros)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Ordenado Crescente:", resultado)
	}
}
