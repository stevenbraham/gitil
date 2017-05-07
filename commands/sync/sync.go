//all function relating to sending and retrieving information local -> remote
package sync

import "os/exec"

func CloneRepository(url string) {
	exec.Command("git", "clone", url).Output()
}

//downloads all remote branches
func FetchAll() {
	exec.Command("git", "fetch", "--all").Output()
}

//Adds a tag to the last commit and pushes the tag to the origin
func CreateTag(tagName string) {
	exec.Command("git", "tag", tagName).Output()
	exec.Command("git", "push", "--tags").Output()
}
