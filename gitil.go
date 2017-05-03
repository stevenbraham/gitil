package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"os"
	"os/exec"
)

func main() {
	//init app
	app := cli.NewApp()
	app.Name = "Gitil"
	app.Version = "0.1.0"
	app.Usage = "Gitil is a wrapper for git that has commands for common tasks"
	app.Action = func(c *cli.Context) error {
		fmt.Println("boom! I say!")
		return nil
	}
	app.Commands = []cli.Command{
		{
			Name:     "clone",
			Category: "Init commands",
			Usage:    "clones a repository",
			Action: func(c *cli.Context) error {
				cloneRepository(c.Args().First())
				fmt.Println("Cloned", c.Args().First())
				return nil
			},
		}}
	app.Run(os.Args)
}

func cloneRepository(url string) {
	exec.Command("git", "clone", url).Output()
}
