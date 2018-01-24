package main

import (
	"bytes"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

func main() {
	http.HandleFunc("/", htmlHandler())
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type handler func(w http.ResponseWriter, r *http.Request)

func htmlHandler() handler {
	return func(w http.ResponseWriter, r *http.Request) {
		var buf bytes.Buffer

		sorter := &selectableSorter{Tracks: tracks}
		for k, _ := range r.URL.Query() {
			if k == "artist" {
				sorter.Select(byArtist)
			}
			if k == "title" {
				sorter.Select(byTitle)
			}
			if k == "album" {
				sorter.Select(byAlbum)
			}
			if k == "year" {
				sorter.Select(byYear)
			}
		}
		if len(sorter.order) == 0 {
			sorter.Select(byAlbum)
		}
		buildHTML(sorter, &buf)

		b, err := ioutil.ReadAll(&buf)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(b)
	}
}

func buildHTML(info *selectableSorter, w io.Writer) {
	sort.Sort(info)
	html := template.Must(template.New("repoInfo").Parse(headerTemplate + trackTemplate + footerTemplate))
	if err := html.Execute(w, info); err != nil {
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
<h2>Tracklist</h2>
<table>
  <tr style='text-align: left'>
    <th>Title</th>
    <th>Artist</th>
    <th>Album</th>
    <th>Year</th>
    <th>Length</th>
  </tr>
  {{range .Tracks}}
  <tr>
    <td>{{.Title}}</td>
    <td>{{.Artist}}</td>
	<td>{{.Album}}</td>
	<td>{{.Year}}</td>
	<td>{{.Length}}</td>
  </tr>
  {{end}}
</table>
<br>
`

const footerTemplate = `
</body>
</html>
`
