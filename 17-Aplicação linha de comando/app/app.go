package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar aplicação de linha de comando par ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e nomes de servidor na internet"

	flags := []cli.Flag {
		cli.StringFlag {
			Name: "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command {
		{
			Name: "ip",
			Usage: "Busca IPS de endereços na internet",
			Flags: flags,
			Action: buscarIps,
		},
		{
			Name: "servidores",
			Usage: "Buscar o nome do servidores na internet",
			Flags: flags,
			Action: buscarServidores,
		},
	}

	return app
}

func buscarIps(c *cli.Context)  {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor)
	}
}