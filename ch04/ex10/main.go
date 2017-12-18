package main

import (
	"log"
	"os"

	"fmt"

	"time"

	"github.com/akito0107/gopl/ch04/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues: \n", result.TotalCount)

	now := time.Now()
	years := now.AddDate(-1, 0, 0)
	month := now.AddDate(0, -1, 0)
	issues := map[string][]*github.Issue{}

	for _, item := range result.Items {
		if item.CreatedAt.Before(years) {
			issues["1年以上"] = append(issues["1年以上"], item)
			continue
		}
		if item.CreatedAt.Before(month) {
			issues["1年未満"] = append(issues["1年未満"], item)
			continue
		}
		issues["1ヶ月未満"] = append(issues["1ヶ月未満"], item)
	}

	for k, v := range issues {
		fmt.Printf("~%s~\n", k)
		for _, issue := range v {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				issue.Number, issue.User.Login, issue.Title)
		}
	}
}
