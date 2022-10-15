package testamenorpreco

import "fmt"

var produtos []Produto = []Produto{
	{
		nome:  "Lamborghini",
		preco: 1000000,
	},
	{
		nome:  "Jipe",
		preco: 46000,
	},
	{
		nome:  "Brasília",
		preco: 16000,
	},
	{
		nome:  "Smart",
		preco: 46000,
	},
	{
		nome:  "Fusca",
		preco: 17000,
	},
}

func RunTestaMenorPreco() {

	maisBarato := buscaMenor(produtos)

	fmt.Println(produtos[maisBarato])
	fmt.Println("O carro mais barato custa: ", produtos[maisBarato].preco)
}

func buscaMenor(produtos []Produto) int {
	maisBarato := 0
	for index, produto := range produtos {
		if produto.preco < produtos[maisBarato].preco {
			maisBarato = index
		}
	}
	return maisBarato
}

func testaMenorPrecoSimples() {
	precos := []float64{1000000, 46000, 16000, 46000, 17000}

	maisBarato := 0
	atual := 0

	for atual = 0; atual < len(precos)-1; atual++ {
		if precos[atual] < precos[maisBarato] {
			maisBarato = atual
		}
	}

	fmt.Println(maisBarato)
	fmt.Println("O carro mais barato custa: ", precos[maisBarato])
}
