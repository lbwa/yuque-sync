# YuQue Sync

Use [Repository Dispatch Event](https://docs.github.com/en/rest/reference/repos#create-a-repository-dispatch-event) to manually trigger Github Action.

## Github Action

```yml
- name: Generate local file
  uses: lbwa/yuque-sync@v1
  with:
    token: ${{ secrets.ACCESS_TOKEN }}
    out-dir: 'docs'
    out-file: ${{ github.event.client_payload.filename }}
    content: ${{ github.event.client_payload.post }}
```

|   name   |                 description                 | default  |
| :------: | :-----------------------------------------: | :------: |
|  token   | A repo scoped Github Personal Access Token  |   N/A    |
| out-dir  | Where should our documentations be place in | `'docs'` |
| out-file |       As generated markdown file name       |   N/A    |
| content  |     As generated markdown file content      |   N/A    |

## Serverless function

- clear building

```bash
go clean
```

- build

```bash
GOOS=linux GOARCH=amd64 go build -o main main.go
```

- use `ctrl+shift+b` to run vscode [build task](./.vscode/tasks.json)

## Notice

`YuQue repository -> settings -> advanced settings(in the SideBar) -> advanced settings(in the main content) -> enable automatic publish` item should **be disabled**, otherwise, YuQue webhook would never be triggered.

More details in [YuQue Webhook](https://www.yuque.com/yuque/developer/doc-webhook).

## FAQ

- Q: github rest api response 404?
- A: Please make sure [Personal access token](https://docs.github.com/en/github/authenticating-to-github/creating-a-personal-access-token) has **ALL ACCESSES** in `repo` scope, especially for private repo.

## References

- [create a repository dispatch event](https://docs.github.com/en/rest/reference/repos#create-a-repository-dispatch-event)
- [event that trigger workflow](https://docs.github.com/en/actions/reference/events-that-trigger-workflows#repository_dispatch)
