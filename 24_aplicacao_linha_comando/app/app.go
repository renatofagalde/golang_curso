package app

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
)

// gerar irá retornar aplicacao de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca ips e nomes de servidor na internet"

	app.Commands = []cli.Command{{
		Name:  "ip",
		Usage: "Usage ips de endereços",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "host",
				Value: "devbook.com.br",
			},
		},
		Action: buscarIps,
	}}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	//net
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
