package main

import (
	"errors"
	"fmt"
	"strings"
)

func substituirVogaisPorAsterisco(s string) (string, error) {
	if s == "" {
		return "", errors.New("String vazia")
	}

	vogais := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true}

	resultado := strings.Builder{}
	for _, char := range s {
		if vogais[char] {
			resultado.WriteRune('*')
		} else {
			resultado.WriteRune(char)
		}
	}

	return resultado.String(), nil
}

func main() {
	texto := "Golang é incrível!"

	resultado, err := substituirVogaisPorAsterisco(texto)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Texto com Vogais Substituídas:", resultado)
	}
}
