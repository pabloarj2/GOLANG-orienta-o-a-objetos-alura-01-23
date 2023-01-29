package main

import (
	"banco/contas"
	"fmt"
)

func Pagarboleto(conta verificaConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificaConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	Pagarboleto(&contaDoDenis, 60)

	fmt.Println(contaDoDenis.ObterSaldo())

	contaDaLuisa := contas.ContaCorrente{}
	contaDaLuisa.Depositar(500)
	Pagarboleto(&contaDaLuisa, 400)

	fmt.Println(contaDaLuisa.ObterSaldo())

}
