package main

import (
	"errors"
	"fmt"
)

func numeroEstaNoSlice(numero int, slice []int) (bool, error) {
	if len(slice) == 0 {
		return false, errors.New("Slice vazio")
	}

	for _, valor := range slice {
		if valor == numero {
			return true, nil
		}
	}

	return false, nil
}

func main() {
	numero := 3
	slice := []int{1, 2, 3, 4, 5}

	estaNoSlice, err := numeroEstaNoSlice(numero, slice)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("O número %d está no slice? %t\n", numero, estaNoSlice)
	}
}
