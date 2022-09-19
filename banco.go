package main

import (
	"fmt"
	"github.com/ZacSv/projetosPessoais/contas"
)

func main() {

	clienteIsac := contas.Cliente{}
	clienteIsac.Cpf = "701.823.576-98"
	clienteIsac.Titular = "Isac"
	clienteIsac.Profissao = "Desenvolvedor"

	contaIsac := contas.ContaCorrente{}
	contaIsac.NomeTitular = clienteIsac
	contaIsac.SaldoConta = 500.00

	clienteKaryne := contas.Cliente{}
	clienteIsac.Cpf = "701.823.576-78"
	clienteIsac.Titular = "Karyne"
	clienteIsac.Profissao = "Desenvolvedor"

	contaKaryne := contas.ContaCorrente{}
	contaKaryne.NomeTitular = clienteKaryne
	contaKaryne.SaldoConta = 1200.00

	fmt.Println(contaKaryne.Transferir(200, &contaIsac)) //"&" identifica o endereço na memoria da conta que recebe o dinheiro
	fmt.Println(contaIsac.SaldoConta)
	fmt.Println(contaKaryne.SaldoConta)

}
