# YuQue Sync

[![GitHub Marketplace](https://img.shields.io/badge/Marketplace-YuQue%20Sync-blue.svg?colorA=24292e&colorB=0366d6&style=flat&longCache=true&logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAA4AAAAOCAYAAAAfSC3RAAAABHNCSVQICAgIfAhkiAAAAAlwSFlzAAAM6wAADOsB5dZE0gAAABl0RVh0U29mdHdhcmUAd3d3Lmlua3NjYXBlLm9yZ5vuPBoAAAERSURBVCiRhZG/SsMxFEZPfsVJ61jbxaF0cRQRcRJ9hlYn30IHN/+9iquDCOIsblIrOjqKgy5aKoJQj4O3EEtbPwhJbr6Te28CmdSKeqzeqr0YbfVIrTBKakvtOl5dtTkK+v4HfA9PEyBFCY9AGVgCBLaBp1jPAyfAJ/AAdIEG0dNAiyP7+K1qIfMdonZic6+WJoBJvQlvuwDqcXadUuqPA1NKAlexbRTAIMvMOCjTbMwl1LtI/6KWJ5Q6rT6Ht1MA58AX8Apcqqt5r2qhrgAXQC3CZ6i1+KMd9TRu3MvA3aH/fFPnBodb6oe6HM8+lYHrGdRXW8M9bMZtPXUji69lmf5Cmamq7quNLFZXD9Rq7v0Bpc1o/tp0fisAAAAASUVORK5CYII=)](https://github.com/marketplace/actions/yuque-sync)

Use [Repository Dispatch Event](https://docs.github.com/en/rest/reference/repos#create-a-repository-dispatch-event) to manually trigger Github Action.

## Github Action

```yml
- name: Generate local file
  uses: lbwa/yuque-sync@v2.0.3
  with:
    token: ${{ secrets.ACCESS_TOKEN }}
    out-dir: 'docs'
    client-payload: ${{toJson(github.event.client_payload)}}
```

|      name      |                       description                       | default  |
| :------------: | :-----------------------------------------------------: | :------: |
|     token      |       A repo scoped Github Personal Access Token        |   N/A    |
|    out-dir     |       Where should our documentations be place in       | `'docs'` |
| client-payload | A request payload from Github Repository Dispatch Event |   N/A    |

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
