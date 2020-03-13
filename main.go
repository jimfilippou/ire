package main

import (
	"fmt"
	"github.com/jimfilippou/ire/utils"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	var timeCounter time.Time
	var app = cli.NewApp()
	app.Name = "ire"
	app.Usage = "Information Retrieval Project 2020"
	app.Version = "0.0.1"
	cli.ErrWriter = ioutil.Discard
	app.Authors = []cli.Author{
		{
			Name:  "Jim Filippou",
			Email: "jimfilippou8@gmail.com",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "generate",
			Aliases: []string{"g", "gen"},
			Usage:   "Generates files based on other data",
			Subcommands: []cli.Command{
				{
					Name:  "json",
					Usage: "Converts documents.txt to documents.json in a matter of milliseconds",
					Action: func(context *cli.Context) error {
						err := utils.CreateFile()
						if err != nil {
							return err
						}
						return nil
					},
					Before: func(c *cli.Context) error {
						timeCounter = time.Now()
						return nil
					},
					After: func(c *cli.Context) error {
						elapsed := time.Since(timeCounter)
						_, err := fmt.Fprintf(c.App.Writer, "Finished in %s", elapsed)
						if err != nil {
							return err
						}
						return nil
					},
				},
			},
		},
		{
			Name:  "feed",
			Usage: "Feeds the database with documents from documents.json",
			Action: func(ctx *cli.Context) error {
				err := utils.FeedTheDB()
				if err != nil {
					return err
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
