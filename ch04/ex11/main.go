package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"

	"net/http"

	"encoding/json"

	"strings"

	"bytes"

	"github.com/akito0107/gopl/ch04/github"
)

const base = "https://api.github.com/repos/"

var token = ""

func main() {
	token = os.Getenv("TOKEN")
	if token == "" {
		log.Fatal("please set github token")
	}
	switch os.Args[1] {
	case "list":
		if len(os.Args) != 3 {
			log.Fatal("usage: list user/repo")
		}
		repos := strings.Split(os.Args[2], "/")
		if len(repos) != 2 {
			log.Fatal("usage: list user/repo")
		}
		res, err := listIssue(repos[0], repos[1])
		if err != nil {
			log.Fatal(err)
		}
		for _, i := range res {
			fmt.Printf("%+v\n", i)
		}
		break
	case "open":
		if len(os.Args) != 4 {
			log.Fatal("usage: open user/repo title")
		}
		repos := strings.Split(os.Args[2], "/")
		if len(repos) != 2 {
			log.Fatal("usage: open user/repo title")
		}
		var issue IssueRequest
		issue.Title = os.Args[3]
		issue.Body = input("put your body")
		fmt.Printf("%+v\n", issue)
		id, err := openIssue(repos[0], repos[1], &issue)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id)
		break
	case "edit":
		if len(os.Args) != 3 {
			log.Fatal("usage: edit user/repo/:id")
		}
		repos := strings.Split(os.Args[2], "/")
		if len(repos) != 3 {
			log.Fatal("usage: edit user/repo/:id")
		}
		id, err := strconv.Atoi(repos[2])
		if err != nil {
			log.Fatal(err)
		}
		issue, err := getIssue(repos[0], repos[1], id)
		req := &IssueRequest{
			Title: issue.Title,
			Body:  issue.Body,
			State: issue.State,
		}
		req.Body = input(issue.Body)
		err = updateIssue(repos[0], repos[1], id, req)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id)
		break
	case "close":
		if len(os.Args) != 3 {
			log.Fatal("usage: close user/repo/:id")
		}
		repos := strings.Split(os.Args[2], "/")
		if len(repos) != 3 {
			log.Fatal("usage: close user/repo/:id")
		}
		id, err := strconv.Atoi(repos[2])
		req := &IssueRequest{
			State: "close",
		}
		err = updateIssue(repos[0], repos[1], id, req)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id)
		break
	default:
		log.Fatalf("unsupported action %s\n", os.Args[1])
	}
}

func input(initial string) string {
	fpath := os.TempDir() + "/message.txt"
	f, err := os.Create(fpath)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile(fpath, []byte(initial), os.ModePerm)
	f.Close()
	editor := "vim"
	if edt := os.Getenv("EDITOR"); edt != "" {
		editor = edt
	}
	cmd := exec.Command(editor, fpath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
	mes, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Fatal(err)
	}
	return string(mes)
}

func listIssue(user, repo string) ([]github.Issue, error) {
	url := base + user + "/" + repo + "/issues"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []github.Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func getIssue(user, repo string, id int) (*IssueResponse, error) {
	url := base + user + "/" + repo + "/issues/" + strconv.Itoa(id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var result IssueResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

type IssueRequest struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
	State string `json:"state,omitempty"`
}

type IssueResponse struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	State string `json:"state"`
}

func openIssue(user, repo string, issue *IssueRequest) (int, error) {
	url := base + user + "/" + repo + "/issues"
	body, err := json.Marshal(issue)
	if err != nil {
		return 0, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	var result IssueResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}
	return result.Id, nil
}

func updateIssue(user, repo string, id int, issue *IssueRequest) error {
	url := base + user + "/" + repo + "/issues/" + strconv.Itoa(id)
	str, err := json.Marshal(issue)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(str))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	var result IssueResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}
	return nil
}
