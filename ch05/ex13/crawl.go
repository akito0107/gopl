package main

import (
	"bytes"
	"fmt"
	"io"
	"log"

	"golang.org/x/net/html"

	"os"

	"net/http"

	"net/url"

	"path"

	lib "github.com/akito0107/gopl/ch05/links"
	"github.com/k0kubun/pp"
)

const AGE = 3

type bodyHandler func(url string, r io.Reader) error

func main() {
	host := parseDomain(os.Args[1])
	handler := linkHandler(host, "data")

	lib.BreadthFirst(crawler(handler), os.Args[1:])
}

func parseDomain(u string) string {
	parsed, err := url.Parse(u)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(parsed.Host)
	return parsed.Host
}

func crawler(handler bodyHandler) func(url string) []string {
	return func(url string) []string {
		fmt.Println(url)
		list, err := extract(url, handler)
		if err != nil {
			log.Print(err)
		}
		return list
	}
}

func linkHandler(domain string, rootDir string) bodyHandler {
	return func(u string, r io.Reader) error {
		if parseDomain(u) != domain {
			return nil
		}
		raw, err := url.Parse(u)
		if err != nil {
			return err
		}

		pp.Print(raw)
		filepath := path.Dir(raw.Path)
		if filepath == "." {
			filepath = "/"
		}
		fileDir := rootDir + filepath

		filename := path.Base(raw.Path)
		if filename == "/" || filename == "." {
			filename = "index.html"
		}
		targetPath := fileDir + filename

		log.Printf("filepath: %s\n", targetPath)

		if err := os.MkdirAll(fileDir, 0777); err != nil {
			return err
		}

		file, err := os.Create(targetPath)
		if err != nil {
			return err
		}

		if _, err := io.Copy(file, r); err != nil {
			return err
		}

		return nil
	}
}

func extract(url string, handler bodyHandler) ([]string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	var buf bytes.Buffer
	tee := io.TeeReader(resp.Body, &buf)
	if err := handler(url, tee); err != nil {
		log.Fatal(err)
	}

	// doc, err := html.Parse(resp.Body)
	doc, err := html.Parse(&buf)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}

	lib.ForEachNode(doc, visitNode, nil)
	return links, nil
}
