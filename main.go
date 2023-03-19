package main

import (
	"fmt"
	"Go-Banco/contas"
	// "Go-Banco/clientes"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) (string, float64)
}

func main() {
	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	PagarBoleto(&contaDoDenis, 50)

	fmt.Println(contaDoDenis.ObterSaldo())
}