package main

import (
	"github.com/podtato-head/podtato-head-app/pkg/podtatoserver"
	"github.com/pterm/pterm"
	urcli "github.com/urfave/cli/v3"
	"log"
	"os"
)

var p podtatoserver.PodTatoServer

func main() {
	app := &urcli.App{
		Name: "podtato-head",
		Flags: []urcli.Flag{
			&urcli.StringFlag{
				Name:        "component",
				Value:       "all",
				Usage:       "all, leftArm, rightArm, leftLeg, rightLeg",
				EnvVars:     []string{"PODTATO_COMPONENT"},
				Destination: &p.Component,
			},
			&urcli.StringFlag{
				Name:        "port",
				Value:       "8080",
				EnvVars:     []string{"PODTATO_PORT"},
				Destination: &p.Port,
			},
			&urcli.StringFlag{
				Name:        "delay",
				Value:       "0s",
				EnvVars:     []string{"PODTATO_STARTUP_DELAY"},
				Destination: &p.StartUpDelay,
			},
			&urcli.StringFlag{
				Name:        "secretMessage",
				Value:       "",
				EnvVars:     []string{"PODTATO_SECRET_MESSAGE"},
				Destination: &p.SecretMessage,
			},
		},
		Action: func(*urcli.Context) error {
			execute()
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func execute() {
	s, _ := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString("podtato-head")).Srender()
	pterm.DefaultCenter.Println(s)

	err := p.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
