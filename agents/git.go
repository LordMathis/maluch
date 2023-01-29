package agents

import (
	"context"
	"os"

	"github.com/go-git/go-git/v5"
)

type GitAgent struct {
	BaseAgent
	workdir string
	repo    string
}

const (
	Clone Action = "clone"
)

func (g *GitAgent) CloneRepo(ctx context.Context) error {
	_, err := git.PlainClone(g.workdir, false, &git.CloneOptions{
		URL:      g.repo,
		Progress: os.Stdout,
	})

	return err
}
