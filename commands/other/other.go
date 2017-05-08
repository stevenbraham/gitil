package other

import (
	"io"
	"net/http"
	"net/url"
	"os"
)

//Downloads a GitIgnore from gitignore.io to .gitignore
func CreateGitIgnore(params string) {
	Url, _ := url.Parse("https://www.gitignore.io/api/" + params)
	response, err := http.Get(Url.String())
	if err != nil {
		panic("Can't connect to gitignore api")
	} else {
		//write body to file
		defer response.Body.Close()
		out, _ := os.Create(".gitignore")
		io.Copy(out, response.Body)
		out.Close()
	}
}
