
package main

import (
	"context"
	"fmt"

	"os"
	"github.com/fluxcd/go-git-providers/github"
	"github.com/fluxcd/go-git-providers/gitprovider"
)
const (
	githubDomain = "github.com"
	repoName = "arc-code"
)
func main() {
	// Create a new client
	ctx := context.Background()
	githubToken := "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" 
	c, err := github.NewClient(gitprovider.WithOAuth2Token(githubToken),)
	//checkErr(err)
        fmt.Println(err)
	// Get public information about the flux repository.
        userRef := gitprovider.UserRef{
		Domain:    githubDomain,
		UserLogin: "chitti-intel",
	}
	repos, err := c.UserRepositories().List(ctx, userRef)
	for _, repo := range repos {
		fmt.Fprintf(os.Stderr, "repo: %s\n", repo.Repository().GetRepository())
	}
        fmt.Fprintf(os.Stderr, "repos, len: %d\n", len(repos))
	//checkErr(err)
	
	desc := "This repo contains azure arc and git golang code"
	userRepoRef := gitprovider.UserRepositoryRef{
		UserRef:        userRef,
		RepositoryName: repoName,
	}
	userRepoInfo := gitprovider.RepositoryInfo{
		Description: &desc,
		Visibility:  gitprovider.RepositoryVisibilityVar(gitprovider.RepositoryVisibilityPublic),
	}

	// Check that the repository doesn't exist
	//_, err = c.UserRepositories().Get(ctx, userRepoRef)
	//Expect(errors.Is(err, gitprovider.ErrNotFound)).To(BeTrue())

	// Create it
	userRepo, err := c.UserRepositories().Create(ctx, userRepoRef, userRepoInfo, &gitprovider.RepositoryCreateOptions{
		AutoInit:        gitprovider.BoolVar(true),
		LicenseTemplate: gitprovider.LicenseTemplateVar(gitprovider.LicenseTemplateApache2),
	})
	//Expect(err).ToNot(HaveOccurred())


        fmt.Println(userRepo)
        fmt.Println(err)

}
