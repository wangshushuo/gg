package main

import (
	"log"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
)

func myCli() {
	prompt := promptui.Select{
		Label: "Select Day",
		Items: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday",
			"Saturday", "Sunday"},
	}

	_, result, err := prompt.Run()
	app := &cli.App{
		Flags: []cli.Flag {
			&cli.StringFlag{
				Name:    "lang",
				Aliases: []string{"l"},
				Value:   result,
				Usage:   "language for the greeting",
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}