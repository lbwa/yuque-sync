# yuque-github-hook

> Only for private use.

Use [Repository Dispatch Event](https://docs.github.com/en/rest/reference/repos#create-a-repository-dispatch-event) to trigger Github Action.

## Usage

- clear building

```bash
go clean
```

- build

```bash
GOOS=linux GOARCH=amd64 go build -o main main.go
```
