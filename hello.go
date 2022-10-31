package main

import (
	"fmt"
)

func main() {
    fmt.Println("Hello, World!")

		nome := "Camila"
		versao := 1.1

		fmt.Println("Olá, tia", nome)
		fmt.Println("Aula GO em sua versão", versao)

		fmt.Println("1 - Iniciar Monitoramento")
		fmt.Println("2 - Exibir Logs")
		fmt.Println("0 - Sair do Programa")

		var instrucao int 
		fmt.Scan(&instrucao)
		
		fmt.Println("O valor da variável instrucao é:", instrucao)

		//expressão
		// if instrucao == 1{
		// 	fmt.Println("Estamos monitorando o Marcos")
		// } else if instrucao == 2 {
		// 	fmt.Println("Exibindo Logs do ZAP do Marcos")
		// } else if instrucao == 3 {
		// 	fmt.Println("Desisti e vou sair")
		// } else {
		// 	fmt.Println("Brasil vai melhorar - YOU DIED!")
		// }

		switch instrucao {
		case 1:
			fmt.Println("Estamos monitorando o Airam")
		case 2:
			fmt.Println("Log: Nunca use reduce")
		case 0:
			fmt.Println("You died, Airam é do inglês")
		default:
			fmt.Println("Airam foi preso por não usar reduce")
			
		}


}

//1- Iniciar Monitoramento
//2 - Exibir Logs
//0 - Sair do Programa