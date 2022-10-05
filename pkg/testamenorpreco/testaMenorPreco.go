package testamenorpreco

import "fmt"

func RunTestaMenorPreco() {
	precos := []float64{1000000, 46000, 16000, 46000, 17000}

	maisBarato := 0
	atual := 0

	for atual = 0; atual < len(precos); atual++ {
		if precos[atual] < precos[maisBarato] {
			maisBarato = atual
		}
	}

	fmt.Println(maisBarato)
	fmt.Println("O carro mais barato custa: ", precos[maisBarato])
}
