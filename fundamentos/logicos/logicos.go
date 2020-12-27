package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2 // ou exclusivo
	comprarSorvete := trab1 || trab2

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	comprarTv50, comprarTv32, comprarSorvete := compras(false, false)

	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t", comprarTv50, comprarTv32, comprarSorvete, !comprarSorvete)
}
