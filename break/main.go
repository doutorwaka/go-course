// Criar um programa que calcula a somatória de números inteiros positivos até que
// o usuário digite um número negativo. A partir daí, o programa deve imprimir
// a somatória na tela.
package main

import "fmt"

func main() {
	sum := 0
	numberFromUser := 0

	for numberFromUser >= 0 {
		fmt.Print("Informe um número inteiro positivo para a somatória ou negativo para sair: ")
		fmt.Scanf("%d", &numberFromUser)

		// if numberFromUser < 0 {
		// 	break
		// }

		sum = sum + numberFromUser
	}

	fmt.Printf("A somatória dos números informados é: %d\n", sum)
}
