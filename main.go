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

type customStatus struct {
	state       string
	targetURL   string
	context     string
	description string
}

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	cs := customStatus{}
	ctx := context.Background()

	app := kingpin.New(filepath.Base(os.Args[0]), `gh-actions-custom-status - Create custom GitHub status.`)
	app.HelpFlag.Short('h')
	app.Flag("state", "The state of the status. Can be one of error, failure, pending, or success.").
		Required().
		StringVar(&cs.state)
	app.Flag("target_url", "The target URL to associate with this status.").
		StringVar(&cs.targetURL)
	app.Flag("context", "A string label to differentiate this status from the status of other systems.").
		StringVar(&cs.context)
	app.Flag("description", "A short description of the status.").
		StringVar(&cs.description)
	kingpin.MustParse(app.Parse(os.Args[1:]))

	envVars := []string{"GITHUB_TOKEN", "GITHUB_REPOSITORY", "GITHUB_REF"}
	for _, r := range envVars {
		if os.Getenv(r) == "" {
			log.Fatalln("env var not set correctly")
		}
	}

	// Get ref.
	ref := os.Getenv("GITHUB_REF")

	// Separate Owner and Repo.
	githubRepo := strings.Split(os.Getenv("GITHUB_REPOSITORY"), "/")

	// Setup GitHub client.
	ghToken := os.Getenv("GITHUB_TOKEN")
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: ghToken})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// Create GitHub status.
	rs := github.RepoStatus{
		TargetURL:   &cs.targetURL,
		Context:     &cs.context,
		Description: &cs.description,
		State:       &cs.state,
	}
	_, _, err := client.Repositories.CreateStatus(ctx, githubRepo[0], githubRepo[1], ref, &rs)
	if err != nil {
		log.Fatalln("could not create status")
	}
}
