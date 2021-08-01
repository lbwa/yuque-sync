import { getInput, info, startGroup, endGroup, setFailed } from '@actions/core'
import * as github from '@actions/github'
import fs from 'fs-extra'
import { $ } from 'zx'
import isString from 'lodash/isString'

const enum Input {
  TOKEN = 'token',
  OUT_DIR = 'out-dir',
  CLIENT_PAYLOAD = 'client-payload'
}

type ClientPayload = { id: number; title: string; post: string; path: string }

export async function main() {
  const token = getInput(Input.TOKEN, { required: true })
  const outDir = getInput(Input.OUT_DIR)
  const rawClientPayload: ClientPayload | string = getInput(
    Input.CLIENT_PAYLOAD,
    { required: true }
  )
  const clientPayload = isString(rawClientPayload)
    ? (JSON.parse(rawClientPayload) as ClientPayload)
    : rawClientPayload
  const outFile = clientPayload.title
  const fileContent = clientPayload.post

  if (!outFile || !fileContent) {
    setFailed(
      `Couldn't find available title or content, we got title ${outFile} and content ${fileContent.slice(
        0,
        10
      )}`
    )
  }

  // https://docs.github.com/en/actions/reference/context-and-expression-syntax-for-github-actions#github-context
  // https://github.com/actions/checkout/blob/25a956c84d5dd820d28caab9f86b8d183aeeff3d/src/input-helper.ts#L22
  const username = github.context.actor || github.context.repo.owner
  const repoName = github.context.repo.repo
  const remoteOrigin = `https://${username}:${token}@github.com/${username}/${repoName}.git`
  const outFilePath =
    (outDir.endsWith('/') ? outDir : `${outDir}/`) +
    (/\.mdx?$/.test(outFile) ? outFile : `${outFile}.md`)

  startGroup('Create local file')
  await fs.outputFile(outFilePath, fileContent)
  info(`New data is available in the ${username}/${repoName}/${outFilePath}`)
  endGroup()

  if (username) {
    await Promise.all([
      $`git config user.email ${username}@users.noreply.github.com`,
      $`git config user.name "${username}"`
    ])
  }

  const nameOnlyForLog =
    outFile.length > 24 ? `${outFile.slice(0, 21)}...` : outFile
  await $`git add .`
  await $`git commit -s -m \"docs: sync \`${nameOnlyForLog}\` from yuque.com\"`
  await $`git push ${remoteOrigin}`

  info(`New data has been uploaded to remote git.`)
}
