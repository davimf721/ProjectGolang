package main

import "fmt"

func leComando() int {
	var comandoLido int

	fmt.Scan(&comandoLido)

	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}
