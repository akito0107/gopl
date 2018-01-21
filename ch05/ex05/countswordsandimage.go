package main

import (
	"bytes"
	"log"
	"net/http"

	"fmt"

	"os"

	"bufio"

	"strings"

	"golang.org/x/net/html"
)

func main() {
	buf := new(bytes.Buffer)
	buf.ReadFrom(os.Stdin)
	s := buf.String()
	words, images, err := CountsWordsAndImages(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("words: %d, images: %d \n", words, images)
}

func CountsWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	var stack []*html.Node
	stack = append(stack, n)

	for len(stack) > 0 {
		n = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if n.Type == html.ElementNode && n.Data == "img" {
			images++
		}
		if n.Type == html.TextNode {
			words += wordfreq(n.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			stack = append(stack, c)
		}
	}
	return
}

func wordfreq(src string) (count int) {
	scanner := bufio.NewScanner(strings.NewReader(src))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		count++
	}
	return
}
