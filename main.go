/*
 * Copyright (c) 2020.
 * Jim Filippou · jimfilippou8@gmail.com
 */

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

	// timeCounter is used to calculate time between operations
	var timeCounter time.Time

	// Generate the app instance
	var app = cli.NewApp()

	// Customize cli app
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
				err := utils.FeedTheDB(ctx)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:  "query",
			Usage: "Queries the database for given input file",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "queries",
					Usage:    "Load queries from `FILE`",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {

				searchResults, err := utils.Query(ctx, ctx.String("queries"))
				if err != nil {
					return err
				}

				// Experiment 20
				for index, item := range searchResults[0] {
					for j, res := range item.Hits.Hits {
						var x = index + 1
						var y = j + 1
						var pretty string = ""
						if x < 10 {
							pretty = "Q0"
						} else {
							pretty = "Q"
						}
						fmt.Fprintf(ctx.App.Writer, "%s%d 0 %s %d %f standard\n", pretty, x, res.Id, y, *res.Score)
					}
				}

				fmt.Fprintf(ctx.App.Writer, "Results are stored in %p:\n", searchResults)

				return nil

			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
