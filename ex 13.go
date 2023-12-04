package main

import (
	"errors"
	"fmt"
	"math"
)

func somaDigitos(numero int) (int, error) {
	if numero < 0 {
		return 0, errors.New("Número negativo")
	}

	var soma int
	for numero > 0 {
		soma += numero % 10
		numero /= 10
	}

	return soma, nil
}

func main() {
	numero := 12345

	soma, err := somaDigitos(numero)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("Soma dos Dígitos de %d: %d\n", numero, soma)
	}
}
