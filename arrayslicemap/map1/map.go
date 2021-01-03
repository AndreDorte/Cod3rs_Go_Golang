package main

import "fmt"

func main() {
	//var aprovados map[int]string
	// mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678900] = "Maria"
	aprovados[12345678901] = "Pedro"
	aprovados[12345678902] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12345678900])
	delete(aprovados, 12345678900)
	fmt.Println(aprovados[12345678900])
}
