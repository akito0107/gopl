package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	var buf bytes.Buffer

	buildHTML(db, &buf)

	b, err := ioutil.ReadAll(&buf)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(b)
}

func buildHTML(db database, w io.Writer) {
	html := template.Must(template.New("repoInfo").Parse(headerTemplate + trackTemplate + footerTemplate))
	if err := html.Execute(w, db); err != nil {
		log.Fatal(err)
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

const trackTemplate = `
<h2>Price List</h2>
<table>
  <tr style='text-align: left'>
    <th>Item</th>
    <th>Price</th>
  </tr>
  {{range $key, $value := .}}
  <tr>
    <td>{{$key}}</td>
    <td>{{$value}}</td>
  </tr>
  {{end}}
</table>
<br>
`

const footerTemplate = `
</body>
</html>
`
