package main

import (
	//"os"
	"fmt"
	"io/ioutil"
	github "github.com/google/go-github/github"
	oauth2 "golang.org/x/oauth2"
)

func githubCreateRepo(reponame string) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "5669659d4d29dd93c1a32e02d00ead59fe3e4198"},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	input := &github.Repository{Name: github.String(reponame)}
	_, _, err := client.Repositories.Create("", input)
	if err != nil {
		fmt.Println("Repositories.Create returned error: %v", err)
		//return true
	}
}

func deleteRepo(reponame string) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "5669659d4d29dd93c1a32e02d00ead59fe3e4198"},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	//input := &github.Repository{Name: github.String(reponame)}
	_, err := client.Repositories.Delete("chanhlv93", reponame)
	if err != nil {
		fmt.Println("Repositories.Create returned error: %v", err)
		//return true
	}
	//fmt.Println(res)
}

func listAll() {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "5669659d4d29dd93c1a32e02d00ead59fe3e4198"},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)
	repos, _, err := client.Repositories.List("", nil)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range repos {
		fmt.Println(*v.Name)
		var namerepo = *v.Name
		if len(namerepo) == 36 {
			go deleteRepo(namerepo)	
		}
		
	}
}

func main() {
	listAll()
}