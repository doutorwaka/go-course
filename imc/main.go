package main

import "fmt"

func main() {
	var peso = 76.0
	var altura = 1.76

	// IMC = peso / altura²
	var imc = peso / (altura * altura)

	fmt.Printf("O IMC de uma pessoa com peso = %f e altura = %f é %f\n", peso, altura, imc)
}
