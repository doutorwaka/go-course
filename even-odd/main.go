package main

import "fmt"

func main() {
	var numberFromUser int32

	fmt.Print("Digite um número inteiro: ")
	fmt.Scanf("%d", &numberFromUser)

	if numberFromUser%2 == 0 {
		fmt.Printf("O numero %d é par.\n", numberFromUser)
	} else {
		fmt.Printf("O numero %d é ímpar.\n", numberFromUser)
	}
}
