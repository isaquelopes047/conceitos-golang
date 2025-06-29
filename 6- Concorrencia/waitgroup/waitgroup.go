package main

import (
	"fmt"
	"sync"
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
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		escrever("Ola mundo")
		waitGroup.Done()
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
