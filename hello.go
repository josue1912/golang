package main

import "fmt"
import "os"
import "net/http"
import "time"

const monitoramentos = 3
const delay = 2

func main() {

	exibeIntroducao()
	for {
		exibeMenu()
		comando := leComando()
		switch comando {
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs")
		default:
			fmt.Println("Opção inválida")
			os.Exit(-1)
		}
	}
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando o sistema...")
	sites := []string{
		"https://random-status-code.herokuapp.com",
		"https://www.alura.com.br",
		"https://www.caelum.com.br"}

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		fmt.Println("-------------------------------------")
		time.Sleep(delay * time.Second)
	}

}

func testaSite(site string) {
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "Resposta:", resp.StatusCode)
	} else {
		fmt.Println("Site:", site, "esta com problema:", resp.StatusCode)
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
	fmt.Println("O comando escolhido foi:", comandoLido)
	return comandoLido
}
