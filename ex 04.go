package main

import "fmt"

func segundoMaiorValor(slice []int) int {
	if len(slice) < 2 {
		return 0
	}

	maior := slice[0]
	segundoMaior := slice[0]

	for _, valor := range slice {
		if valor > maior {
			segundoMaior = maior
			maior = valor
		} else if valor > segundoMaior && valor < maior {
			segundoMaior = valor
		}
	}

	return segundoMaior
}

func main() {
	numeros := []int{5, 2, 8, 1, 9, 3}
	segundoMaior := segundoMaiorValor(numeros)

	fmt.Printf("Segundo Maior Valor: %d\n", segundoMaior)
}
