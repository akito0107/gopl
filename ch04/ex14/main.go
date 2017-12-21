package main

import (
	"bytes"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	token := os.Getenv("TOKEN")
	if token == "" {
		log.Fatal("Please set github token")
	}
	if len(os.Args) != 2 {
		log.Fatal("usage: ./ex14 user/repo")
	}
	repos := strings.Split(os.Args[1], "/")
	if len(repos) != 2 {
		log.Fatal("usage: ./ex14 user/repo")
	}
	info, err := GetRepositoryInfo(token, repos[0], repos[1])
	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	buildHTML(info, &buf)

	http.HandleFunc("/", htmlHandler(&buf))
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type handler func(w http.ResponseWriter, r *http.Request)

func htmlHandler(buf io.Reader) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(buf)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(b)
	}
}

const headerTemplate = `
<html lang="ja">
<head>
  <meta charset="utf-8">
  <!--[if lt IE 9]>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/html5shiv/3.7.3/html5shiv.js"></script>
  <![endif]-->
</head>
<body>
`

const issueTemplate = `
<h2>Issues of {{.RepoName}}</h2>
<table>
  <tr style='text-align: left'>
    <th>#</th>
    <th>State</th>
    <th>Title</th>
    <th>Link</th>
  </tr>
  {{range .Issues}}
  <tr>
    <td>{{.Id}}</td>
    <td>{{.State}}</td>
	<td>{{.Title}}</td>
	<td><a href='{{.URL}}'>Here</a></td>
  </tr>
  {{end}}
</table>
<br>
`

const milestoneTemplate = `
<h2>Milestones of {{.RepoName}}</h2>
<table>
  <tr style='text-align: left'>
    <th>#</th>
    <th>State</th>
    <th>Title</th>
    <th>Link</th>
  </tr>
  {{range .Milestones}}
  <tr>
    <td>{{.Id}}</td>
    <td>{{.State}}</td>
	<td>{{.Title}}</td>
	<td><a href='{{.URL}}'>Here</a></td>
  </tr>
  {{end}}
</table>
<br>
`

const collaboratorTemplate = `
<h2>Collaborators of {{.RepoName}}</h2>
<table>
  <tr style='text-align: left'>
    <th>Id</th>
    <th>Login</th>
    <th>Link</th>
  </tr>
  {{range .Collaborators}}
  <tr>
    <td>{{.Id}}</td>
    <td>{{.Login}}</td>
	<td><a href='{{.URL}}'>Here</a></td>
  </tr>
  {{end}}
</table>
<br>
`

const footerTemplate = `
</body>
</html>
`

func buildHTML(info *RepositoryInfo, w io.Writer) {
	html := template.Must(template.New("repoInfo").Parse(headerTemplate + issueTemplate + milestoneTemplate + collaboratorTemplate + footerTemplate))
	if err := html.Execute(w, info); err != nil {
		log.Fatal(err)
	}
}
