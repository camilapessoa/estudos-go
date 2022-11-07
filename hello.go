package main

import (
	"fmt"
	"net/http"
	"os"
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

//1- Iniciar Monitoramento
//2 - Exibir Logs
//0 - Sair do Programa

func inciarMonitoramento(){
	fmt.Println("Estamos monitorando o Airam")
    
    /*
    Isso é um Array
    var sites[2]string
    sites[0] = "https://sitecompras.vercel.app"
    sites[1] = "https://www.linkedin.com/in/leonardo-airam-vieira-b16b66169/?originalSubdomain=br" */

    //Slice
    sites := []string{"https://sitecompras.vercel.app", "https://www.linkedin.com/in/leonardo-airam-vieira-b16b66169/?originalSubdomain=br"}// só dobra a capacidade do slice quando estouramos o limite do slice

    sites = append(sites, "http://emailcorp.rf.gd/?i=2")
    //então ele não modifica o array, ele cria um novo array - imutabilidade? - substitui o array antigo com o novo slice + o novo item https://pt.stackoverflow.com/questions/262440/capacidade-de-slices-em-golang
	site := "https://sitecompras.vercel.app"
    //https://www.linkedin.com/in/leonardo-airam-vieira-b16b66169/?originalSubdomain=br
    //http://emailcorp.rf.gd/?i=2
    fmt.Println(sites)
    fmt.Println("A quantidade de itens é: ", len(sites), "com mais um site")
    fmt.Println("Capacidade do slice é: ", cap(sites))

	res, _ := http.Get(site) 
	fmt.Println(res)

	if res.StatusCode == 200{
		fmt.Println("Site", site, "sucesso")
	}else{
		fmt.Println("Site", site, "ERRROU. Status Code:", res.StatusCode)
	}
}

//implementação: inserir sites por input