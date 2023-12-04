package main

import (
	"errors"
	"fmt"
)

func stringsComMaisDe5Caracteres(slice []string) ([]string, error) {
	if len(slice) == 0 {
		return nil, errors.New("Slice vazio")
	}

	var resultado []string
	for _, str := range slice {
		if len(str) > 5 {
			resultado = append(resultado, str)
		}
	}

	return resultado, nil
}

func main() {
	palavras := []string{"Go", "Linguagem", "Programação", "Incrível"}

	resultado, err := stringsComMaisDe5Caracteres(palavras)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Strings com Mais de 5 Caracteres:", resultado)
	}
}
