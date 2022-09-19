package contas

type contaPoupanca struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Operacao      int
	saldo         float64
}

func (c *contaPoupanca) Sacar(valorDoSaque float64) (string, float64) { //Definindo uma função para sacar

	podeSacar := valorDoSaque <= c.saldo && valorDoSaque >= 0 //Definindo como condiional para sacar o valor do saque ser menor ou igual ao valor da conta

	if podeSacar { //Aplicando a condicional caso for verdadeira
		c.saldo -= valorDoSaque                                         //Subtraia o valor
		return "Saque realizado com sucesso, saldo atual de: ", c.saldo //Imprima o sucesso
	} else {
		return "Saldo insuficiente para o saque, ou saque menor do que 0", c.saldo //Caso a condicional seja falsa, imprima!
	}

}

func (c *contaPoupanca) Depositar(valorDeposito float64) (string, float64) { //Definindo uma função para depositar

	podeDepositar := valorDeposito > 0 //Definindo como condicional para depositar o valor do depósito ser maior do que 0

	if podeDepositar { //Aplicando a condicional caso for verdadeira
		c.saldo += valorDeposito                                           //Some o valor
		return "Deposito realizado com sucesso, saldo atual de: ", c.saldo //Imprima o sucesso
	} else {
		return "O valor que está sendo depositado é imcompatível", c.saldo //Caso a condicional seja falsa, imprima!
	}

}

func (c *contaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
