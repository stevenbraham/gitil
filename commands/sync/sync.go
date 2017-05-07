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
