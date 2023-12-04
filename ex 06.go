package main

import (
	"errors"
	"fmt"
	"strings"
)

func concatenarComVirgulas(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("Slice vazio")
	}
	
	concatenado := strings.Join(slice, ",")
	return concatenado, nil
}

func main() {
	palavras := []string{"Olá", "mundo", "Go"}
	resultado, err := concatenarComVirgulas(palavras)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Concatenação:", resultado)
	}
}
