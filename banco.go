package main

import (
	cli "banco/Clientes"
	c "banco/ContaCorrente"

	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func main() {
	for {
		MenuInicial()
		respostasMenuInicial()
	}
}

func MenuInicial() {
	fmt.Println("1- Verificar contas")
	fmt.Println("2- Criar Conta")
	fmt.Println("0- Sair do programa")
}

func respostasMenuInicial() {
	comando := comandoInt()
	switch comando {
	case 1:
		fmt.Println("")
		verificaContas()
		fmt.Println("")
	case 2:
		fmt.Println("")
		criaConta()
		fmt.Println("")
	case 0:
		fmt.Println("")
		fmt.Println("Saindo do programa!")
		fmt.Println("")
		os.Exit(0)
	default:
		fmt.Println("")
		fmt.Println("Este não é um valor válido, tente novamente")
		fmt.Println("")
	}
}

func menuCriaConta() {
	fmt.Println("1- Criar Nova Conta")
	fmt.Println("2- Voltar")
	fmt.Println("0- Sair do programa")
}

func respostasCriaConta() {
	comando := comandoInt()
	switch comando {
	case 1:
		fmt.Println("")
		criaConta()
		fmt.Println("")
	case 2:
		fmt.Println("")
		main()
		fmt.Println("")
	case 0:
		fmt.Println("")
		fmt.Println("Saindo do programa!")
		fmt.Println("")
		os.Exit(0)
	default:
		fmt.Println("")
		fmt.Println("Este não é um valor válido, tente novamente")
		fmt.Println("")
	}
}

func menuContas() {

}

func comandoInt() int64 {
	var comando int64 = -1
	fmt.Scan(&comando)

	return comando
}
func comandoFloat() float64 {
	var comando float64 = -1
	fmt.Scan(&comando)

	return comando
}
func comandoString() string {
	var comando string
	fmt.Scan(&comando)

	return comando
}

func verificaContas() {
	fmt.Println("")
	fmt.Println("Contas disponíveis para consulta:")
	fmt.Println("")
	arquivo, err := ioutil.ReadFile("contas.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(arquivo))
	fmt.Println("")
	menuContas()
}

func criaConta() {

	fmt.Println("Digite o nome do titular: ")
	titular := comandoString()
	fmt.Println("Digite o número do CPF: ")
	cpf := comandoString()
	fmt.Println("Digite o nome da profissão do titular: ")
	profissao := comandoString()
	fmt.Println("Digite o número da agência: ")
	agencia := comandoInt()
	fmt.Println("Digite o número da conta: ")
	conta1 := comandoInt()
	fmt.Println("Digite o saldo: ")
	saldo1 := comandoFloat()
	// id := uuid.New()

	conta := c.ContaCorrente{Titular: cli.Titular{
		Nome:      titular,
		CPF:       cpf,
		Profissao: profissao}, Agencia: agencia, Conta: conta1, Saldo: saldo1}

	arquivo, err := os.OpenFile("contas.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - Titular: " + conta.Titular.Nome + " - CPF: " + conta.Titular.CPF + " - Profissão: " + conta.Titular.Profissao + " -  Agência: " + strconv.FormatInt(conta.Agencia, 10) + " -  Conta Corrente: " + strconv.FormatInt(conta.Conta, 10) + " -  Saldo: " + strconv.FormatFloat(conta.Saldo, 'f', 2, 64) + "\n")
	arquivo.Close()
	fmt.Println("")
	fmt.Println("A conta foi criada com sucesso!")
	fmt.Println("")
	menuCriaConta()
	respostasCriaConta()
}
