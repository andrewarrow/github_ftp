package main

import "github.com/andrewarrow/paradise_ftp/paradise"

type GitHubFiles struct {
}

func (ghf GitHubFiles) GetFiles() []string {
	files := make([]string, 3)

	return files
}

func NewGitHubFileManager() *paradise.FileManager {
	fm := paradise.FileManager{}
	fm.FileSystem = GitHubFiles{}
	return &fm
}
