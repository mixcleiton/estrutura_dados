package selectionsort

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

func RunOrdenacaoSimples() {

	ordena(produtos, len(produtos))

	for _, produto := range produtos {
		fmt.Printf("Produto %v custa %f \n", produto.nome, produto.preco)
	}

}

func ordena(lista []Produto, quantidadeElementos int) {
	for atual := 0; atual < quantidadeElementos; atual++ {
		menor := buscaMenor(lista, atual, quantidadeElementos)
		produtoAtual := lista[atual]
		produtoMenor := lista[menor]
		lista[atual] = produtoMenor
		lista[menor] = produtoAtual
	}
}

func buscaMenor(produtos []Produto, inicio int, termino int) int {
	maisBarato := inicio
	for atual := inicio; atual < termino; atual++ {
		if produtos[atual].preco < produtos[maisBarato].preco {
			maisBarato = atual
		}
	}
	return maisBarato
}
