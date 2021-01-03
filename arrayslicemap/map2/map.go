package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.0,
	}

	funcsESalarios["Rafael Silva"] = 1350.0
	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}

	// para ignorar a chave basta substituir pelo _
	for _, salario := range funcsESalarios {
		fmt.Println(salario)
	}

	// caso não seja informado um par chave e valor, será exibido apenas a chave do mapa, não o salario conforme esperado
	for salario := range funcsESalarios {
		fmt.Println(salario)
	}
}
