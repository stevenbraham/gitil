package main

import (
	"bufio"
	"fmt"
	"github.com/stevenbraham/gitil/commands/branches"
	"github.com/stevenbraham/gitil/commands/other"
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
				repoUrl := c.Args().First()
				if repoUrl == "" {
					fmt.Println("Usage: gitil clone [url]")
					os.Exit(1)
				}
				sync.CloneRepository(repoUrl)
				fmt.Println("Cloned", repoUrl)
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
				if tag == "" {
					fmt.Println("Usage: gitil create-tag [tag]")
					os.Exit(1)
				}
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
				branches.CheckoutBranch(branch)
				return nil
			},
		},
		{
			Name:     "insta-commit",
			Aliases:  []string{"ic"},
			Category: "Sync commands",
			Usage:    "Adds all, commits and pushes",
			Action: func(c *cli.Context) error {
				message := c.Args().First()
				if message == "" {
					fmt.Println("Usage: gitil insta-commit [message]")
					os.Exit(1)
				}
				branches.AddAll()
				branches.Commit(message)
				sync.Push()
				fmt.Println("Commited", c.Args().First())
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
		{
			Name:     "reset-hard",
			Aliases:  []string{"rh"},
			Category: "Other commands",
			Usage:    "Deletes all changes",
			Action: func(c *cli.Context) error {
				scanner := bufio.NewScanner(os.Stdin)
				fmt.Print("DANGER! This deletes all your changes, do you want to continue [y,N]:")
				for scanner.Scan() {
					if scanner.Text() == "y" {
						branches.ResetHard()
						return nil
					}
					break
				}
				return nil
			},
		},
		{
			Name:     "create-gitignore",
			Aliases:  []string{"gi"},
			Category: "Other commands",
			Usage:    "Downloads a GitIgnore from gitignore.io to .gitignore",
			Action: func(c *cli.Context) error {
				params := c.Args().First()
				if params == "" {
					fmt.Println("Usage: gitil create-gitignore [platform,platform]")
					os.Exit(1)
				}
				other.CreateGitIgnore(params)
				fmt.Println("Created new gitignore")
				return nil
			},
		},
	}
	app.Run(os.Args)
}
