//All function related to merging branches
package branches

import (
	"io/ioutil"
	"os/exec"
	"strings"
	"github.com/stevenbraham/gitil/commands/sync"
)

//checkouts a branch
func CheckoutBranch(branch string) {
	exec.Command("git", "checkout", branch).Output()
}

//merges source into destination and checkouts source again
func MergeBranch(source, destination string) {
	CheckoutBranch(destination)
	exec.Command("git", "merge", source).Output()
	sync.Push()
	CheckoutBranch(source)
}

//executes git all all
func AddAll() {
	exec.Command("git", "add", ".").Output()
}

//Git commit
func Commit(message string) {
	exec.Command("git", "commit", "-m", message).Output()
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

//parses the HEAD file and returns current branch from it
func GetCurrentBranch() string {
	data, err := ioutil.ReadFile(".git/HEAD")
	if err != nil {
		panic("HEAD file missing")
	}
	head := strings.Trim(string(data), "\n")
	return strings.Replace(head, "ref: refs/heads/", "", 1)
}
