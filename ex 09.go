package main

import (
	"errors"
	"fmt"
	"strings"
)

func extrairPalavras(s string) ([]string, error) {
	if s == "" {
		return nil, errors.New("String vazia")
	}

	palavras := strings.Fields(s)
	return palavras, nil
}

func main() {
	frase := "Olá mundo, Go é incrível!"

	resultado, err := extrairPalavras(frase)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Palavras:", resultado)
	}
}
