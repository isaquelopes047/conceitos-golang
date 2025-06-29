package app

import (
	"fmt"
	"log"
	"net"
	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicacao de linha de comando para executar
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e Nomes de servidor na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "devbook.cxom.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Busca IPS de endereço da internet",
			Flags: flags,
			Action: buscarIps,
		},
		{
			Name: "servidores",
			Usage: "Busca o nome do servidores da internet",
			Flags: flags,
			Action: buscaServidores,
		},
	}

	return app
}

func buscarIps(c * cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscaServidores(c *cli.Context){
	host := c.String("host")

	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidores := range servidores {
		fmt.Println(servidores.Host)
	}
}