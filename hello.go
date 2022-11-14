package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
    for{ 
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
func inciarMonitoramento(){
	fmt.Println("Estamos monitorando o Airam")
     //Slice
    sites := []string{"https://sitecompras.vercel.app", "https://www.linkedin.com/", "https://shouldideploy.today/", "http://www.pudim.com.br/"}

    sites = append(sites, "http://emailcorp.rf.gd/?i=2")
   
for i:= 0; i < 5; i++{
    for i, site := range sites{
        fmt.Println("Site", i, " -> ", site)
        testarSite(site)
    }
    time.Sleep(4 * time.Second)
    fmt.Println(" ")
}
    
}

func testarSite(pingSite string){
    res, _ := http.Get(pingSite) 
	fmt.Println(res)

	if res.StatusCode == 200{
		fmt.Println("Site", pingSite, "sucesso")
	}else{
		fmt.Println("Site", pingSite, "ERRROU. Status Code:", res.StatusCode)
	}
    fmt.Println("")
}