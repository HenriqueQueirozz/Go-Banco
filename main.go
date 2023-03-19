package main

import "fmt"

type ContaCorrente struct {
	titular 		string
	numeroAgencia	int
	numeroConta 	int
	saldo 			float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso", c.saldo
	}else{
		return "Saldo insuficiente ", c.saldo
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else{
		return "Valor do deposito menor que zero", c.saldo
	}
}

func main() {
	contaDaSilvia := ContaCorrente{}
	contaDaSilvia.titular = "Silvia'"
	contaDaSilvia.saldo = 500
	 
	fmt.Println(contaDaSilvia.saldo)
	fmt.Println(contaDaSilvia.Sacar(300))
	fmt.Println(contaDaSilvia.saldo)
	contaDaSilvia.Depositar(500)
	fmt.Println(contaDaSilvia.saldo)
}