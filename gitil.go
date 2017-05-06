package main

import (
	"bufio"
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"io/ioutil"
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
			Name:     "fetch-all",
			Category: "Init commands",
			Usage:    "Does a git fetch --all",
			Action: func(c *cli.Context) error {
				fetchAll()
				return nil
			},
		},
		{
			Name:     "to-master",
			Category: "Merge commands",
			Usage:    "Merges the selected branch into master",
			Action: func(c *cli.Context) error {
				fetchAll()
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
				fetchAll()
				var branch = c.Args().First()
				mergeBranch("master", branch)
				fmt.Println("Merged master into", branch)
				return nil
			},
		},
		{
			Name:     "master-all",
			Category: "Merge commands",
			Usage:    "Merges master in all branches",
			Action: func(c *cli.Context) error {
				scanner := bufio.NewScanner(os.Stdin)
				fmt.Print("DANGER! This merges master in all branches, do you want to continue [y,N]:")
				for scanner.Scan() {
					if scanner.Text() == "y" {
						fetchAll()
						var ownBranch = c.Args().First()
						for _, branch := range getBranches() {
							if branch != ownBranch && branch != "master" {
								mergeBranch("master", branch)
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

func cloneRepository(url string) {
	exec.Command("git", "clone", url).Output()
}

//downloads all remote branches
func fetchAll() {
	exec.Command("git", "fetch", "--all").Output()
}

//checkouts a branch
func checkoutBranch(branch string) {
	exec.Command("git", "checkout", branch).Output()
}

//merges source into destination and checkouts source again
func mergeBranch(source, destination string) {
	checkoutBranch(destination)
	exec.Command("git", "merge", source).Output()
	exec.Command("git", "push").Output()
	checkoutBranch(source)
}

//lists all local branches
func getBranches() []string {
	//get files from refs
	files, _ := ioutil.ReadDir("./.git/refs/heads")
	branches := make([]string, len(files))
	//cast files to string
	for key, file := range files {
		branches[key] = file.Name()
	}
	return branches
}
