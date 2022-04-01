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

const monitoring = 3
const secondDelay = 3
const statusSucess = 200

func main() {

	exibeIntro()

	for {
		exibeMenu()

		option := comandoLido()

		switch option {
		case 1:
			startMonitoring()
		case 2:
			exibeLog()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Opção inválida")
			os.Exit(-1)
		}
	}
}

func exibeIntro() {
	name := "Davi Cruz"
	version := 1.1
	fmt.Println("Olá Sr.", name)
	fmt.Println("Version:", version)
}

func exibeMenu() {
	fmt.Println("1 - Inicar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func comandoLido() int {
	var option int
	fmt.Scan(&option)
	return option
}

func startMonitoring() {
	fmt.Println("Monitorando...")
	sites := readFile()

	for i := 0; i < monitoring; i++ {
		fmt.Println("Testando sites ----")
		for _, site := range sites {
			requestGet(site)
		}
		time.Sleep(time.Second * secondDelay)
		fmt.Println("")
	}
}

func requestGet(site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if response.StatusCode == statusSucess {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		writeFile(site, true)
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:", response.StatusCode)
		writeFile(site, false)
	}
}

func readFile() []string {
	var sites []string
	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
	}

	scanner := bufio.NewReader(file)
	for {
		linha, err := scanner.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	file.Close()
	fmt.Println(sites)

	return sites
}

func writeFile(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + fmt.Sprint(status) + "\n")

	file.Close()
}

func exibeLog() {
	file, err := os.Open("log.txt")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
	}

	scanner := bufio.NewReader(file)
	for {
		linha, err := scanner.ReadString('\n')
		linha = strings.TrimSpace(linha)
		fmt.Println(linha)

		if err == io.EOF {
			break
		}
	}

	file.Close()
}
