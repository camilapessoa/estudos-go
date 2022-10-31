package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
    inicio()

    opcaoSelecionada := leitor()

    switch opcaoSelecionada {
    case 1:
			inciarMonitoramento()
    case 2:
        fmt.Println("Log: Nunca use reduce")
    case 0:
        fmt.Println("You died, saia do jogo")
				os.Exit(0)
    default:
        fmt.Println("Airam foi preso por não usar reduce")
				os.Exit(-1)
    }

}

func inicio() {
    nome := "Camila"
    versao := 1.1

    fmt.Println("Olá, tia", nome)
    fmt.Println("Aula GO em sua versão", versao)

    fmt.Println("1 - Iniciar Monitoramento")
    fmt.Println("2 - Exibir Logs")
    fmt.Println("0 - Sair do Programa")

}

func leitor() int {

    var instrucao int
    fmt.Scan(&instrucao)

    fmt.Println("O valor da variável instrucao é:", instrucao)
    return instrucao
}

//1- Iniciar Monitoramento
//2 - Exibir Logs
//0 - Sair do Programa

func inciarMonitoramento(){
	fmt.Println("Estamos monitorando o Airam")
	site := "https://sitecompras.vercel.app"
	res, _ := http.Get(site) 
	fmt.Println(res)

	if res.StatusCode == 200{
		fmt.Println("Site", site, "sucesso")
	}else{
		fmt.Println("Site", site, "ERRROU. Status Code:", res.StatusCode)
	}
}