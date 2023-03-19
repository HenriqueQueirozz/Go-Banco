package main

import (
	"fmt"
	"Go-Banco/contas"
)



func main() {
	contaDaSilvia := contas.ContaCorrente{Titular:"Silvia", Saldo: 300}
	contaDoGustavo := contas.ContaCorrente{Titular:"Gustavo", Saldo: 100}

	status := contaDoGustavo.Transferir(200, &contaDaSilvia)

	fmt.Println(status)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)

	status_2 := contaDaSilvia.Transferir(-200, &contaDoGustavo)

	fmt.Println(status_2)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)


}