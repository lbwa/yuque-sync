package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"regexp"
	"strings"
	"yuque-sync/model/yuque"

	"github.com/google/go-github/v35/github"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
	"github.com/tencentyun/scf-go-lib/events"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

const (
	YUQUE_HOST = "https://www.yuque.com"
)

var (
	GITHUB_OWNER    = os.Getenv("GITHUB_OWNER")
	GITHUB_REPO     = os.Getenv("GITHUB_REPO")
	GITHUB_PAT      = os.Getenv("GITHUB_PAT")      // Github personal access token
	GITHUB_RD_EVENT = os.Getenv("GITHUB_RD_EVENT") // repository_dispatch event type
)

func main() {

	// usage: https://cloud.tencent.com/document/product/583/18032
	// Make the handler available for Remote Procedure Call by Cloud Function
	cloudfunction.Start(dispatchGithubAction)
}

// payload from cloud function, should has a custom structure
// YuQue Webhook, see https://www.yuque.com/yuque/developer/doc-webhook#4da6e742
type YuQueEvent struct {
	Data yuque.DocDetailSerializer `json:"data"`
}

// should use `github.event.client_payload.post` to retrieve `Post` field in the repository action file(*.yml)
type GithubClientPayload struct {
	Id    yuque.YuQueId `json:"id"`
	Title string        `json:"title"`
	Post  string        `json:"post"`
	Path  string        `json:"path"`
}

// Inspired by https://github.com/google/go-github/blob/a19996a59629e9dc2b32dc2fb8628040e6e38459/github/repos_test.go#L2213
// github v3 rest api: https://docs.github.com/en/rest
// based on tencent cloud api gateway event, see https://github.com/tencentyun/scf-go-lib/blob/ccd4bf6de8cb891d5b58e49d6e03000337f9f817/events/apigw.go
func dispatchGithubAction(ctx context.Context, request events.APIGatewayRequest) (string, error) {
	// Debug level enabled by default in development
	// Info level enabled by default in production
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()

	client := createGithubClient(ctx, GITHUB_PAT)

	// regexp syntax, https://github.com/google/re2/wiki/Syntax
	isAuthorizedMethod, unauthorizedMethodErr := regexp.MatchString(
		// ignore letter case
		"(?i)"+http.MethodPost,
		request.Method,
	)
	if !isAuthorizedMethod || unauthorizedMethodErr != nil {
		return "", errors.New(`unauthorized method`)
	}

	sugar.Debug("Github owner: ", GITHUB_OWNER)
	sugar.Debug("Github repo: ", GITHUB_REPO)

	var yuQueData *YuQueEvent
	json.Unmarshal([]byte(request.Body), &yuQueData)

	postMeta := yuQueData.Data
	post := postMeta.Body

	docPathParts := createDocPathParts(yuQueData)
	docUrl := strings.Join(append([]string{YUQUE_HOST}, docPathParts...), "/")

	sugar.Debug("YuQue post action: ", postMeta.ActionType)
	sugar.Debug("YuQue post title: ", postMeta.Title)
	sugar.Debug("YuQue post URL: ", docUrl)

	stringifiedBodyBytes, _ := json.MarshalIndent(post, "", "  ")
	sugar.Debug("YuQue post body: ", string(stringifiedBodyBytes))

	clientPayload, err := serializeClientPayload(postMeta.Id, postMeta.Title, post, strings.Join(docPathParts, "/"))
	if err != nil {
		sugar.Debug("Got error when marshal GithubClientPayload: ", err)
		return "", err
	}
	// create a repository dispatch event
	// https://docs.github.com/en/rest/reference/repos#create-a-repository-dispatch-event
	repo, response, err := client.Repositories.Dispatch(
		ctx,
		GITHUB_OWNER,
		GITHUB_REPO,
		github.DispatchRequestOptions{
			// EventType is a custom webhook event name.(required)
			EventType:     GITHUB_RD_EVENT,
			ClientPayload: &clientPayload,
		},
	)

	if err != nil {
		sugar.Debug("Repositories.Dispatch returned error: ", err)
		return "", err
	}

	if response.StatusCode != http.StatusNoContent {
		// https://gobyexample.com/json
		// https://golang.org/pkg/encoding/json/#Marshal
		messageBytes, _ := json.Marshal(response)
		return "", errors.New(string(messageBytes))
	}

	sugar.Debug("Operation successfully: ", repo.HTMLURL)
	return http.StatusText(http.StatusOK), nil
}

func createGithubClient(ctx context.Context, pat string) *github.Client {
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: pat})
	tokenClient := oauth2.NewClient(ctx, tokenSource)
	return github.NewClient(tokenClient)
}

func createDocPathParts(yuQueData *YuQueEvent) (docPathParts []string) {
	postMeta := yuQueData.Data
	user := postMeta.User

	if postMeta.Path != "" {
		docPathParts = append(docPathParts, postMeta.Path)
	} else {
		docPathParts = append(docPathParts, []string{
			user.Login,         // username
			postMeta.Book.Slug, // repository slug
			postMeta.Slug,      // post slug
		}...)
	}
	return
}

func serializeClientPayload(id yuque.YuQueId, title, post, filepath string) (msg json.RawMessage, err error) {
	rawBytes, err := json.Marshal(GithubClientPayload{
		Id:    id,
		Title: title,
		Post:  post,
		Path:  filepath,
	})
	if err != nil {
		return []byte{}, err
	}
	return json.RawMessage(rawBytes), nil
}
