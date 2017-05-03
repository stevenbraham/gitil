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
		},
		{
			Name:     "to-master",
			Category: "Merge commands",
			Usage:    "Merges the selected branch into master",
			Action: func(c *cli.Context) error {
				var branch = c.Args().First()
				mergeBranch(branch, "master")
				fmt.Println("Merged", branch, " into master")
				return nil
			},
		},
		{
			Name:     "from-master",
			Category: "Merge commands",
			Usage:    "Merges master into the selected branch",
			Action: func(c *cli.Context) error {
				var branch = c.Args().First()
				mergeBranch("master", branch)
				fmt.Println("Merged master into", branch)
				return nil
			},
		},
	}
	app.Run(os.Args)
}

func cloneRepository(url string) {
	exec.Command("git", "clone", url).Output()
}

//checkouts a branch
func checkoutBranch(branch string) {
	exec.Command("git", "checkout", branch).Output()
}

//merges source into destination and checkouts source again
func mergeBranch(source, destination string) {
	exec.Command("git", "fetch", "--all").Output()
	checkoutBranch(destination)
	exec.Command("git", "merge", source).Output()
	exec.Command("git", "push").Output()
	checkoutBranch(source)
}
