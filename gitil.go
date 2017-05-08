package main

import (
	"bufio"
	"fmt"
	"github.com/stevenbraham/gitil/commands/branches"
	"github.com/stevenbraham/gitil/commands/sync"
	"gopkg.in/urfave/cli.v1"
	"os"
)

func main() {
	//init app
	app := cli.NewApp()
	app.Name = "Gitil"
	app.Version = "0.1.0"
	app.Usage = "Gitil is a wrapper for git that has commands for common tasks"
	app.Action = func(c *cli.Context) error {
		app.Command("help").Run(c)
		return nil
	}
	app.Commands = []cli.Command{
		{
			Name:     "clone",
			Aliases:  []string{"c"},
			Category: "Sync commands",
			Usage:    "clones a repository",
			Action: func(c *cli.Context) error {
				sync.CloneRepository(c.Args().First())
				fmt.Println("Cloned", c.Args().First())
				return nil
			},
		},
		{
			Name:     "fetch-all",
			Aliases:  []string{"fa"},
			Category: "Sync commands",
			Usage:    "Does a git fetch --all",
			Action: func(c *cli.Context) error {
				sync.FetchAll()
				return nil
			},
		},
		{
			Name:     "create-tag",
			Aliases:  []string{"ct"},
			Category: "Sync commands",
			Usage:    "Adds a tag to the last commit and pushes the tag to the origin",
			Action: func(c *cli.Context) error {
				tag := c.Args().First()
				sync.CreateTag(tag)
				fmt.Println("Created tag", tag)
				return nil
			},
		},
		{
			Name:     "add-commit-push",
			Aliases:  []string{"ac"},
			Category: "Sync commands",
			Usage:    "Adds all files to commit and pushes the commit to the origin",
			Action: func(c *cli.Context) error {
				tag := c.Args().First()
				sync.CreateTag(tag)
				fmt.Println("Created tag", tag)
				return nil
			},
		},
		{
			Name:     "to-master",
			Aliases:  []string{"tm"},
			Category: "Merge commands",
			Usage:    "Merges the selected branch into master",
			Action: func(c *cli.Context) error {
				sync.FetchAll()
				var branch = branches.GetCurrentBranch()
				//if not empty use the branch provided by the command line
				if c.Args().First() != "" {
					branch = c.Args().First()
				}
				branches.MergeBranch(branch, "master")
				fmt.Println("Merged", branch, " into master")
				return nil
			},
		},
		{
			Name:     "from-master",
			Aliases:  []string{"fm"},
			Category: "Merge commands",
			Usage:    "Merges master into the selected branch",
			Action: func(c *cli.Context) error {
				sync.FetchAll()
				var branch = branches.GetCurrentBranch()
				//if not empty use the branch provided by the command line
				if c.Args().First() != "" {
					branch = c.Args().First()
				}
				branches.MergeBranch("master", branch)
				fmt.Println("Merged master into", branch)
				return nil
			},
		},
		{
			Name:     "insta-commit",
			Aliases:  []string{"ic"},
			Category: "Sync commands",
			Usage:    "Adds all, commits and pushes",
			Action: func(c *cli.Context) error {
				branches.AddAll()
				branches.Commit(c.Args().First())
				sync.Push()
				fmt.Println("Commited",c.Args().First())
				return nil
			},
		},
		{
			Name:     "master-all",
			Aliases:  []string{"ma"},
			Category: "Merge commands",
			Usage:    "Merges master in all branches",
			Action: func(c *cli.Context) error {
				scanner := bufio.NewScanner(os.Stdin)
				fmt.Print("DANGER! This merges master in all branches, do you want to continue [y,N]:")
				for scanner.Scan() {
					if scanner.Text() == "y" {
						sync.FetchAll()
						var ownBranch = branches.GetCurrentBranch()
						//if not empty use the branch provided by the command line
						if c.Args().First() != "" {
							ownBranch = c.Args().First()
						}
						for _, branch := range branches.GetBranches() {
							if branch != ownBranch && branch != "master" {
								branches.MergeBranch("master", branch)
								fmt.Println("Merged master into", branch)
							}
						}
						return nil
					}
					break
				}
				return nil
			},
		},
	}
	app.Run(os.Args)
}
