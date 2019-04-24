package main

import (
	"fmt"
	"os"

	"github.com/hixi-hyi/localstack-go/localstack"
	"github.com/urfave/cli"
)

func checkRequired(ctx *cli.Context, flags ...string) {
	for _, flag := range flags {
		if value := ctx.String(flag); value == "" {
			fmt.Fprintf(os.Stderr, "--%s is required", flag)
			fmt.Fprintln(os.Stderr)
			ctx.App.Writer = os.Stderr
			cli.ShowAppHelpAndExit(ctx, -1)
		}
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "localstack-go"
	app.Usage = "Display the localstack endpoint"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "service, s",
			EnvVar: "LOCALSTACK_SERVICE",
		},
		cli.StringFlag{
			Name:   "domain, d",
			EnvVar: "LOCALSTACK_DOMAIN",
			Value:  "localhost",
		},
	}
	app.Action = func(ctx *cli.Context) error {
		checkRequired(ctx, "service")
		serviceName := ctx.String("service")
		l := localstack.New(&localstack.Config{Domain: ctx.String("domain")})
		service, err := localstack.Services.Get(serviceName)
		if err != nil {
			return err
		}
		fmt.Println(l.Endpoint(service))
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
