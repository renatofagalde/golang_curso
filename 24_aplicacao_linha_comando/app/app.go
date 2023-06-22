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

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}
	app.Commands = []cli.Command{{
		Name:   "ip",
		Usage:  "Usage ips de endereços",
		Flags:  flags,
		Action: buscarIps,
	}, {
		Name:   "servidores",
		Usage:  "busca servidores na internet",
		Flags:  flags,
		Action: buscarServidores,
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

func buscarServidores(c *cli.Context) {
	host := c.String("host")
	fmt.Println("Procurando servidores...")
	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	//exibindo struct
	for _, servidor := range servidores {
		fmt.Println(servidor)
	}

	//exibindo apenas o valor
	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}

}
