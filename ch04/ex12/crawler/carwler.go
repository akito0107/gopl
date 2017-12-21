package crawler

import (
	"os"
	"strconv"

	"fmt"

	"net/http"

	"encoding/json"

	"io/ioutil"
	"log"
)

const baseURL = "https://xkcd.com/"
const workerSize = 10

func Crawl() {
	if len(os.Args) != 4 {
		log.Fatal("usage: ./xkcd crawl :start :stop  eg). xkcd crawl 100 200")
	}
	start, err := strconv.Atoi(os.Args[2])
	end, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatal(err)
	}
	jobs := make(chan int, 100)
	results := make(chan error, 100)

	for w := 0; w < workerSize; w++ {
		go work(w, jobs, results)
	}

	for j := start; j < end; j++ {
		jobs <- j
	}
	close(jobs)
	for r := 0; r < end-start; r++ {
		if err = <-results; err != nil {
			log.Fatal(err)
		}
	}
}

type Comic struct {
	Num        int    `json:"num"`
	Transcript string `json:"transcript"`
}

func work(id int, jobs <-chan int, results chan<- error) {
	for {
		comicId := <-jobs
		if comicId == 0 {
			return
		}
		url := fmt.Sprintf("%s%d/info.0.json", baseURL, comicId)
		log.Printf("worker %d crawl %s \n", id, url)
		resp, err := http.Get(url)
		if err != nil {
			results <- err
			continue
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			log.Printf("worker %d crawl %s failed with %s \n", id, url, resp.Status)
			results <- nil
			continue
		}
		var c Comic
		if err := json.NewDecoder(resp.Body).Decode(&c); err != nil {
			results <- err
			continue
		}
		data, err := json.Marshal(c)
		if err != nil {
			results <- err
			continue
		}
		err = ioutil.WriteFile(fmt.Sprintf("./data/%d.json", comicId), data, 0644)
		if err != nil {
			results <- err
			continue
		}
		results <- nil
	}
}
