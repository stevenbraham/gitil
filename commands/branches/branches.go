//All function related to merging branches
package branches

import (
	"io/ioutil"
	"os/exec"
)

//checkouts a branch
func CheckoutBranch(branch string) {
	exec.Command("git", "checkout", branch).Output()
}

//merges source into destination and checkouts source again
func MergeBranch(source, destination string) {
	CheckoutBranch(destination)
	exec.Command("git", "merge", source).Output()
	exec.Command("git", "push").Output()
	CheckoutBranch(source)
}

//lists all local branches
func GetBranches() []string {
	//get files from refs
	files, _ := ioutil.ReadDir("./.git/refs/heads")
	branches := make([]string, len(files))
	//cast files to string
	for key, file := range files {
		branches[key] = file.Name()
	}
	return branches
}
