package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2
	comprarSorvete := trab1 || trab2

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, false)
	fmt.Printf("Tv50: %t ,tv32: %t, sorvete : %t, saudavel: %t",
		tv50, tv32, sorvete, !sorvete)

}
