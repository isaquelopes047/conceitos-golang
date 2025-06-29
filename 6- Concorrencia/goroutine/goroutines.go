package main

import (
	"fmt"
	"time"
)

/*
	CONCORRECIA
	Capacidade de lidar com múltiplas tarefas ao mesmo tempo,
	mas não necessariamente executadas simultaneamente. Usa alternância rápida entre tarefas.
	Ex: Um garçom atendendo várias mesas — ele faz uma tarefa de cada vez, mas gerencia várias "ao mesmo tempo".

	PARALELISMO
	Execução de múltiplas tarefas ao mesmo tempo, simultaneamente, usando múltiplos núcleos de CPU.
	Ex: Vários garçons atendendo mesas ao mesmo tempo.

	Concorrência = estrutura para multitarefa.
	Paralelismo = execução real simultânea.
*/

func main() {
	go escrever("Ola mundo")
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
