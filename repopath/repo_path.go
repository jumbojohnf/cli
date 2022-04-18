package repopath

import "github.com/funcgql/cli/shell"

type RepoPath struct {
	Path string
}

func GitRepoPath() (RepoPath, error) {
	output, err := shell.Execute("git", "rev-parse", "--show-toplevel")
	if err != nil {
		return RepoPath{}, err
	}
	return RepoPath{Path: output.Combined}, nil
}
