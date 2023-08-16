/*
 * Copyright 2022 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package repository

import (
	"context"
	"github.com/google/go-github/v53/github"
	"github.com/xanzy/go-gitlab"
	"golang.org/x/oauth2"
	"io/ioutil"
)

type GitHubApi struct {
	Client *github.Client
}

func NewClient(token string) *GitHubApi {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	return &GitHubApi{
		Client: github.NewClient(tc),
	}
}

func (g *GitHubApi) GetFile(owner, repoName, filePath, ref string) (*File, error) {
	fileContent, _, err := g.Client.Repositories.DownloadContents(context.Background(), owner, repoName, filePath, &github.RepositoryContentGetOptions{Ref: ref})
	if err != nil {
		return nil, err
	}
	defer fileContent.Close()

	content, err := ioutil.ReadAll(fileContent)
	if err != nil {
		return nil, err
	}

	return &File{
		Name:    filePath,
		Content: content,
	}, nil
}

func (g *GitHubApi) PushFilesToRepository(files map[string][]byte, owner, repoName, branch, commitMessage string) error {
	// Get a reference to the default branch
	ref, _, err := g.Client.Git.GetRef(context.Background(), owner, repoName, "refs/heads/"+branch)
	if err != nil {
		return err
	}

	// Create a new Tree object for the file to be pushed
	var treeEntries []*github.TreeEntry
	for filePath, content := range files {
		treeEntries = append(treeEntries, &github.TreeEntry{
			Path:    github.String(filePath),
			Content: github.String(string(content)),
			Mode:    github.String("100644"),
		})
	}
	newTree, _, err := g.Client.Git.CreateTree(context.Background(), owner, repoName, *ref.Object.SHA, treeEntries)
	if err != nil {
		return err
	}

	// Create a new commit object, using the new tree as its foundation
	newCommit, _, err := g.Client.Git.CreateCommit(context.Background(), owner, repoName, &github.Commit{
		Message: github.String(commitMessage),
		Tree:    newTree,
	})
	if err != nil {
		return err
	}

	// Update branch references to point to new submissions
	_, _, err = g.Client.Git.UpdateRef(context.Background(), owner, repoName, &github.Reference{
		Ref: github.String("refs/heads/" + branch),
		Object: &github.GitObject{
			SHA:  newCommit.SHA,
			Type: github.String("commit"),
		},
	}, true)
	if err != nil {
		return err
	}

	return nil
}

func (gl *GitLabApi) InitClient(token string) error {
	client, err := gitlab.NewClient(token)
	if err != nil {
		return err
	}
	gl.Client = client

	return nil
}