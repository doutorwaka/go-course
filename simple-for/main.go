package main

import "fmt"

func main() {
	var n int

	fmt.Print("Informe um numero inteiro positivo: ")
	fmt.Scanf("%d", &n)

	sumEvens := 0

	// 0, 1, 2, 3, 4
	// n = 4
	for i := 0; i <= n; i = i + 1 {
		if i%2 == 0 {
			sumEvens = sumEvens + i
		}
	}

	fmt.Printf("A somatoria dos numeros de 0, ..., %d é %d\n", n, sumEvens)
}
