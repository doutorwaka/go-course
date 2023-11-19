package main

import "fmt"

func main() {
	var numberFromUser int32

	fmt.Print("Informe um número inteiro: ")
	fmt.Scanf("%d", &numberFromUser)

	if numberFromUser < 0 {
		fmt.Printf("O numero %d é negativo.\n", numberFromUser)
	} else if numberFromUser > 0 {
		fmt.Printf("O numero %d é postivo.\n", numberFromUser)
	} else {
		fmt.Printf("O numero %d é zero.\n", numberFromUser)
	}
}
