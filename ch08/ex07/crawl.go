package main

import (
	"fmt"
	"io"
	"log"
	"strings"

	"golang.org/x/net/html"

	"os"

	"net/http"

	"net/url"
)

type bodyHandler func(url string, r io.Reader) error

type message struct {
	resp    *http.Response
	url     string
	path    string
	host    string
	hostURL string
}

func main() {
	if err := crawl(os.Args[1]); err != nil {
		log.Fatalf("%+v\n", err)
	}
}

func parseDomain(u string) string {
	parsed, err := url.Parse(u)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(parsed.Host)
	return parsed.Host
}

func crawl(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	host := resp.Request.URL.Host
	path := resp.Request.URL.Path
	schema := resp.Request.URL.Scheme
	hostURL := schema + "://" + host
	log.Printf("host: %s path: %s \n", host, path)

	if err := os.Mkdir(host, os.ModePerm); err != nil {
		if os.IsExist(err) {
			return nil
		}
		return err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	mes := message{
		resp:    resp,
		url:     url,
		path:    host + "/root.html",
		host:    host,
		hostURL: hostURL,
	}
	c := make(chan message, 10)
	go func(mes message, c chan message) {
		c <- mes
	}(mes, c)

	return extractByType(c)
}

func extractByType(resChan chan message) error {
	cnt := 1
	for ; cnt > 0; cnt-- {
		message := <-resChan
		resp := message.resp
		url := message.url
		path := message.path
		host := message.host
		hostURL := message.hostURL

		contentType := extractContentType(resp.Header)

		// img / css
		if contentType[0] != "text/html" {
			func() {
				f, err := os.Create(path)
				if err != nil {
					log.Fatalf("create file failed: %s error: %v", path, err)
				}
				defer f.Close()
				defer resp.Body.Close()
				if _, err := io.Copy(f, resp.Body); err != nil {
					log.Fatalf("copy failed: %s error: %v", hostURL, err)
				}
			}()
			continue
		}

		doc, err := html.Parse(resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Fatalf("parsing %s as HTML: %+v", url, err)
		}

		visitNode := func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "a" {
				for i, a := range n.Attr {
					if a.Key != "href" {
						continue
					}
					link, err := resp.Request.URL.Parse(a.Val)
					if err != nil {
						continue
					}
					if strings.HasPrefix(link.String(), hostURL) {
						n.Attr[i].Val = "file:" + a.Val
						if a.Val != "/" && a.Val != "#" && link.String() != hostURL {
							cnt++
							if a.Val[0] == '/' {
								extractAsFile(link.String(), host+a.Val, host, hostURL, resChan)
							} else {
								extractAsFile(link.String(), host+"/"+a.Val, host, hostURL, resChan)
							}

						}
					}
				}
			} else if n.Type == html.ElementNode && n.Data == "img" {
				for _, a := range n.Attr {
					if a.Key != "src" {
						continue
					}
					link, err := resp.Request.URL.Parse(a.Val)
					if err != nil {
						continue
					}
					cnt++
					extractAsFile(link.String(), host+"/"+a.Val, host, hostURL, resChan)
				}
			}
		}
		forEachNode(doc, visitNode, nil)

		func() {
			f, err := os.Create(path)
			if err != nil {
				log.Fatalf("create file failed %s error: %v", hostURL, err)
			}
			defer f.Close()
			html.Render(f, doc)
		}()
	}

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) *html.Node {
	var stack []*html.Node
	stack = append(stack, n)

	for len(stack) > 0 {
		n = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if pre != nil {
			pre(n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			stack = append(stack, c)
		}
		if post != nil {
			post(n)
		}
	}
	return nil
}

func extractContentType(header http.Header) []string {
	contentType, ok := header["Content-Type"]
	if !ok {
		return nil
	}
	return strings.Split(contentType[0], ";")
}

func extractAsFile(url, path, host, hostURL string, resChan chan<- message) error {
	go func(url, path, host, hostURL string) {
		resp, err := http.Get(url)

		if err != nil {
			log.Printf("http get to %s, got error: %+v\n", hostURL, err)
			return
		}

		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			log.Printf("hostURL: %s, error response code: %+v\n", hostURL, resp.StatusCode)
			return
		}
		resChan <- message{
			resp,
			url,
			path,
			host,
			hostURL,
		}
	}(url, path, host, hostURL)

	return nil
}
