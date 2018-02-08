package main

import "fmt"
import "os"

func main() {

	exibeIntroducao()
	exibeMenu()
	comando := leComando()

	switch comando {
	case 0:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	case 1:
		fmt.Println("Monitorando o sistema...")
	case 2:
		fmt.Println("Exibindo logs")
	default:
		fmt.Println("Opção inválida")
		os.Exit(-1)
	}
}

func exibeIntroducao() {
	nome := "Gopher"
	var versao float32 = 1.0
	fmt.Println("Ola, sr.", nome)
	fmt.Println("Versao: ", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi:", &comandoLido)
	return comandoLido
}
