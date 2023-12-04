package main

import (
	"errors"
	"fmt"
	"math"
)

func primosAteN(numero int) ([]int, error) {
	if numero < 0 {
		return nil, errors.New("Número negativo")
	}

	var primos []int
	for i := 2; i <= numero; i++ {
		if isPrimo(i) {
			primos = append(primos, i)
		}
	}

	return primos, nil
}

func isPrimo(numero int) bool {
	if numero < 2 {
		return false
	}

	limite := int(math.Sqrt(float64(numero)))

	for i := 2; i <= limite; i++ {
		if numero%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	numero := 20

	primos, err := primosAteN(numero)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("Números Primos até %d: %v\n", numero, primos)
	}
}
