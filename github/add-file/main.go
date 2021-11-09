package main

import (
        "context"
        "fmt"
        "github.com/fluxcd/go-git-providers/github"
        "github.com/fluxcd/go-git-providers/gitprovider"
    	"io/ioutil"
    	"log"
    	"os"
    	"bufio"
    	"path/filepath"

)
const (
        githubDomain = "github.com"
	repoName = "arc-code"
)


func main() {

	iterate("/home/ubuntu/go_projects/src")
}

//iterate through all the files
func iterate(path string) {
    files := []gitprovider.CommitFile{}
    length:=len([]rune(path))
    filepath.Walk(path, func(path string, info os.FileInfo, err error) error{
        if err != nil {
            log.Fatalf(err.Error())
        }
        if info.IsDir()==false{
         fmt.Printf("File Name: %s\n", info.Name())
         //Encode the file to base64
         content := getContent(path)
         //send the file to the github repo
	 relativePath := path[length+1:]
	 fmt.Println(relativePath)
         files = append(files,gitprovider.CommitFile{Path : &relativePath, Content: &content})
        }
         return nil
    })
    fmt.Println(len(files))
    sendFile(files)
}

//get contents of the file
func getContent(filePath string) string {

    // Open file on disk.
    f, _ := os.Open(filePath)

    // Read entire file into byte slice.
    reader := bufio.NewReader(f)
    content, _ := ioutil.ReadAll(reader)

    return string(content)
}


func sendFile(files []gitprovider.CommitFile) {
        // Create a new client
        ctx := context.Background()
        githubToken := "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"
        c, err := github.NewClient(gitprovider.WithOAuth2Token(githubToken),)
        userRef := gitprovider.UserRef{
                Domain:    githubDomain,
                UserLogin: "chitti-intel",
        }

        userRepoRef := gitprovider.UserRepositoryRef{
                UserRef:        userRef,
                RepositoryName: repoName,
        }

        userRepo, err := c.UserRepositories().Get(ctx, userRepoRef)
        //Commit file to this repo
        _, err = userRepo.Commits().Create(ctx, "main", "Files added", files)

        fmt.Println(err)

}
