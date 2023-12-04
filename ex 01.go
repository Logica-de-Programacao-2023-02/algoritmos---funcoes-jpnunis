package main

import "fmt"

func calcularMedia(slice []int) float64 {
    if len(slice) == 0 {
        return 0
    }

    soma := 0
    for _, valor := range slice {
        soma += valor
    }

    return float64(soma) / float64(len(slice))
}

func main() {
    numeros := []int{1, 2, 3, 4, 5}
    media := calcularMedia(numeros)
    
    fmt.Printf("Média: %.2f\n", media)
}
