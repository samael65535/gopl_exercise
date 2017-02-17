package main
// 创建一个web服务器，查询一次GitHub，然后生成BUG报告、里程碑和对应的用户信息。
import (
	"html/template"
	"strings"
	 "../ex4.10/github"
	"fmt"
	// "os"
	"log"
	"net/http"
)

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

func main() {
	// result, err := github.SearchIssues(os.Args[1:])
	// if err != nil {
	//	log.Fatal(err)
	// }
	// if err := issueList.Execute(os.Stdout, result); err != nil {
	//	log.Fatal(err)
	// }
	 fmt.Println()


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
			return
		}
		var terms []string
		for k, v := range r.Form {
			if k == "q" {
				for _, s := range v {
					terms = append(terms, strings.Split(s, " ")...)
				}
				break;
			}
		}
		if len(terms) == 0 {
			return
		}
		result, err := github.SearchIssues(terms)
		if err != nil {
			log.Fatal(err)
		}
		if err := issueList.Execute(w, result); err != nil {
			log.Fatal(err)
		}
	})


	http.ListenAndServe("localhost:8001", nil)
}
