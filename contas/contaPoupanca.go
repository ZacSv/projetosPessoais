package contas

type Cliente struct {
	Titular   string
	Cpf       string
	Profissao string
}

type contaCorrente struct {
	nomeTitular   Cliente
	numeroAgencia int
	numeroConta   int
	saldoConta    float64
}

func (c *contaCorrente) Transferir(valorDaTransferencia float64, contaDestino *contaCorrente) string {
	podeTransferir := valorDaTransferencia <= c.saldoConta && valorDaTransferencia >= 0

	if podeTransferir {
		c.saldoConta -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return "Transferencia realizada com sucesso"
	} else {
		return "Saldo insuficiente para a transferencia, ou menor do que 0"
	}

}

func (c *contaCorrente) Sacar(valorDoSaque float64) (string, float64) { //Definindo uma função para sacar

	podeSacar := valorDoSaque <= c.saldoConta && valorDoSaque >= 0 //Definindo como condiional para sacar o valor do saque ser menor ou igual ao valor da conta

	if podeSacar { //Aplicando a condicional caso for verdadeira
		c.saldoConta -= valorDoSaque                                         //Subtraia o valor
		return "Saque realizado com sucesso, saldo atual de: ", c.saldoConta //Imprima o sucesso
	} else {
		return "Saldo insuficiente para o saque, ou saque menor do que 0", c.saldoConta //Caso a condicional seja falsa, imprima!
	}

}

func (c *contaCorrente) Depositar(valorDeposito float64) (string, float64) { //Definindo uma função para depositar

	podeDepositar := valorDeposito > 0 //Definindo como condicional para depositar o valor do depósito ser maior do que 0

	if podeDepositar { //Aplicando a condicional caso for verdadeira
		c.saldoConta += valorDeposito                                           //Some o valor
		return "Deposito realizado com sucesso, saldo atual de: ", c.saldoConta //Imprima o sucesso
	} else {
		return "O valor que está sendo depositado é imcompatível", c.saldoConta //Caso a condicional seja falsa, imprima!
	}

}
