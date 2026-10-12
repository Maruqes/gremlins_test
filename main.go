package main

const (
	limiteDescontoVIP  = 10000 // 100 €
	limitePortesGratis = 5000  // 50 €
	custoPortes        = 500   // 5 €
)

// Os valores são em cêntimos. Clientes VIP têm 10% de desconto
// a partir de 100 €. Encomendas abaixo de 50 € pagam 5 € de portes.
func totalEncomenda(subtotal int, clienteVIP bool) int {
	total := subtotal

	if clienteVIP && subtotal >= limiteDescontoVIP {
		total = total - subtotal/10
	}

	if subtotal < limitePortesGratis {
		total = total + custoPortes
	}

	return total
}
