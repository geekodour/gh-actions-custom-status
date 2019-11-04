package main

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/go-github/v28/github"
	"golang.org/x/oauth2"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	rs := github.RepoStatus{
		State:       github.String(""),
		TargetURL:   github.String(""),
		Context:     github.String("default"),
		Description: github.String(""),
	}
	ctx := context.Background()

	app := kingpin.New(filepath.Base(os.Args[0]), `gh-actions-custom-status - Create custom GitHub status.`)
	app.HelpFlag.Short('h')
	app.Flag("state", "The state of the status. Can be one of error, failure, pending, or success.").
		Required().
		StringVar(rs.State)
	app.Flag("target_url", "The target URL to associate with this status.").
		StringVar(rs.TargetURL)
	app.Flag("context", "A string label to differentiate this status from the status of other systems.").
		StringVar(rs.Context)
	app.Flag("description", "A short description of the status.").
		StringVar(rs.Description)
	kingpin.MustParse(app.Parse(os.Args[1:]))

	envVars := []string{"GITHUB_TOKEN", "GITHUB_REPOSITORY", "LAST_COMMIT_SHA"}
	for _, r := range envVars {
		if os.Getenv(r) == "" {
			log.Fatalln("env var not set correctly")
		}
	}

	// Get ref.
	ref := os.Getenv("LAST_COMMIT_SHA")

	// Separate Owner and Repo.
	githubRepo := strings.Split(os.Getenv("GITHUB_REPOSITORY"), "/")

	// Setup GitHub client.
	ghToken := os.Getenv("GITHUB_TOKEN")
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: ghToken})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// Create GitHub status.
	_, _, err := client.Repositories.CreateStatus(ctx, githubRepo[0], githubRepo[1], ref, &rs)
	if err != nil {
		log.Fatalf("%v : could not create status", err)
	}
}
