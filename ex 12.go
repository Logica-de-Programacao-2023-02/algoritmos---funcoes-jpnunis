package main

import (
	"errors"
	"fmt"
	"unicode"
)

func stringsComLetraMaiuscula(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("Slice vazio")
	}

	var resultado string
	for _, str := range slice {
		if len(str) > 0 && unicode.IsUpper([]rune(str)[0]) {
			resultado += str + " "
		}
	}

	return resultado[:len(resultado)-1], nil
}

func main() {
	palavras := []string{"Go", "é", "Incrível", "golang"}

	resultado, err := stringsComLetraMaiuscula(palavras)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Palavras com Letra Maiúscula:", resultado)
	}
}
