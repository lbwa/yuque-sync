package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"yuque-github-hook/yuque"

	"github.com/google/go-github/v35/github"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
	"golang.org/x/oauth2"
)

var (
	GITHUB_OWNER         = os.Getenv("GITHUB_OWNER")
	GITHUB_REPO          = os.Getenv("GITHUB_REPO")
	GITHUB_ACCESS_TOKEN  = os.Getenv("GITHUB_ACCESS_TOKEN")
	GITHUB_WEBHOOK_EVENT = os.Getenv("GITHUB_WEBHOOK_EVENT")
)

func main() {

	// usage: https://cloud.tencent.com/document/product/583/18032
	// Make the handler available for Remote Procedure Call by Cloud Function
	cloudfunction.Start(dispatchGithubAction)
}

// payload from cloud function, should has a custom structure
// YuQue Webhook, see https://www.yuque.com/yuque/developer/doc-webhook#4da6e742
type CloudEvent struct {
	// 下文后面的 `json:xxx` 是为了将代码结构体中的字段与 json 数据解耦
	Data       yuque.DocDetailSerializer `json:"data"`
	Path       string                    `json:"path,omitempty"` // 文档的完整访问路径（不包括域名）
	ActionType string                    `json:"action_type"`    // 值有 publish - 发布、 update - 更新、 delete - 删除
	Publish    bool                      `json:"publish"`        // 文档是否为第一次发布，第一次发布时为 true
}

// Inspired by https://github.com/google/go-github/blob/a19996a59629e9dc2b32dc2fb8628040e6e38459/github/repos_test.go#L2213
// github v3 rest api: https://docs.github.com/en/rest
func dispatchGithubAction(ctx context.Context, event CloudEvent) error {
	fmt.Printf("Current github owner: %v\n", GITHUB_OWNER)
	fmt.Printf("Current github repo: %v\n", GITHUB_REPO)
	fmt.Printf("Current YuQue workspace: %v\n", event.Path)
	fmt.Printf("Current YuQue documentation: %v\n", event.Data.Title)

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: GITHUB_ACCESS_TOKEN})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// create a repository dispatch event
	// https://docs.github.com/en/rest/reference/repos#create-a-repository-dispatch-event
	repo, response, err := client.Repositories.Dispatch(ctx, GITHUB_OWNER, GITHUB_REPO, github.DispatchRequestOptions{
		// EventType is a custom webhook event name.(required)
		EventType: GITHUB_WEBHOOK_EVENT,
	})

	if err != nil {
		fmt.Printf("Repositories.Dispatch returned error: %v", err)
		return err
	}

	if response.StatusCode != 204 {
		// https://gobyexample.com/json
		// https://golang.org/pkg/encoding/json/#Marshal
		messageBytes, _ := json.Marshal(response)
		return errors.New(string(messageBytes))
	}

	fmt.Printf("Operation successfully: %v", repo.URL)
	return nil
}
