package main

import "fmt"

func encontrarPosicao(slice []int, valor int) int {
	for i, v := range slice {
		if v == valor {
			return i
		}
	}
	return -1
}

func main() {
	numeros := []int{1, 2, 3, 4, 5}
	valorProcurado := 3
	posicao := encontrarPosicao(numeros, valorProcurado)

	fmt.Printf("Posição do valor %d: %d\n", valorProcurado, posicao)
}
