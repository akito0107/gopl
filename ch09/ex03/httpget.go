package memo5

import (
	"io/ioutil"
	"log"
	"net/http"
)

func httpGetBody(url string, done Done) (interface{}, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("newRequest Failed: %s, error: %v", url, err)
		return nil, err
	}

	cancelChan := make(chan struct{})
	req.Cancel = cancelChan

	go func() {
		select {
		case <-done:
			close(cancelChan)
		}
	}()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
