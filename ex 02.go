package main

import "fmt"

func contarVogais(s string) int {
    quantidade := 0
    for _, char := range s {
        switch char {
        case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
            quantidade++
        }
    }
    return quantidade
}

func main() {
    texto := "Exemplo de String"
    resultado := contarVogais(texto)
    
    fmt.Printf("Quantidade de vogais: %d\n", resultado)
}
