package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

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

func leSitesDoArquivo() []string {
	var sites []string
	// arquivo, err := os.Open("sites.txt")
	// arquivo, err := ioutil.ReadFile("sites.txt")
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("[Erro] Motivo:", err)
	} else {
		leitor := bufio.NewReader(arquivo)
		for {
			linha, err := leitor.ReadString('\n')
			linha = strings.TrimSpace(linha)
			sites = append(sites, linha)
			fmt.Println("Linha adicionada:", linha)
			if err == io.EOF {
				break
			}
		}
	}
	return sites
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando o sistema...")
	sites := leSitesDoArquivo()
	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		fmt.Println("-------------------------------------")
		time.Sleep(delay * time.Second)
	}
}

func testaSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("[ERRO] Motivo: ", err)
	} else {
		if resp.StatusCode == 200 {
			fmt.Println("Site:", site, "Resposta:", resp.StatusCode)
		} else {
			fmt.Println("Site:", site, "esta com problema:", resp.StatusCode)
		}
	}
}

func exibeIntroducao() {
	nome := "Gopher"
	var versao float32 = 1.0
	fmt.Println("Ola, sr.", nome)
	fmt.Println("Versão:", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	return comandoLido
}
