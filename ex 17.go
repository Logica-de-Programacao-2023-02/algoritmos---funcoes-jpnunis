package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

func ordenarStrings(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("Slice vazio")
	}

	sort.Strings(slice)
	resultado := strings.Join(slice, " ")

	return resultado, nil
}

func main() {
	palavras := []string{"Golang", "é", "uma", "linguagem", "incrível"}

	resultado, err := ordenarStrings(palavras)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Strings Ordenadas:", resultado)
	}
}
