package main

import "fmt"

func main() {

	var height float64
	var weight float64

	fmt.Print("Digite a sua altura: ")
	fmt.Scanf("%f", &height)

	fmt.Print("Digite o seu peso: ")
	fmt.Scanf("%f", &weight)

	var imc = weight / (height * height)

	fmt.Printf("O imc é %f\n", imc)

}
