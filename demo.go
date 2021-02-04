package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func demo() {
	YamlFile := "config.yaml"
	app := cli.NewApp()

	flags := []cli.Flag{
		altsrc.NewStringFlag(&cli.StringFlag{Name: "host"}),
		&cli.StringFlag{Name: "load"},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Println("action")
		fmt.Println(c.String("host"))
		return nil
	}

	app.Before = altsrc.InitInputSourceWithContext(flags, func(context *cli.Context) (altsrc.InputSourceContext, error) {
		return altsrc.NewYamlSourceFromFile(YamlFile)
	})

	app.Flags = flags

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}