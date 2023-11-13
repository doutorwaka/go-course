package main

import "fmt"

func main() {
	var height float64
	var weight float64

	fmt.Print("Informe a sua altura: ")
	fmt.Scanf("%f", &height)

	fmt.Print("Informe o seu peso: ")
	fmt.Scanf("%f", &weight)

	var imc = weight / (height * height)

	fmt.Printf("O seu imc é %f\n", imc)

	if imc > 25 {
		fmt.Println("Obesidade")
	}

	if imc < 25 {
		fmt.Println("Normal")
	}

	if imc < 18.5 {
		fmt.Println("Abaixo do peso")
	}
}
