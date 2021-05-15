"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.main = void 0;
const tslib_1 = require("tslib");
const core_1 = require("@actions/core");
const github = tslib_1.__importStar(require("@actions/github"));
const fs_extra_1 = tslib_1.__importDefault(require("fs-extra"));
const simple_git_1 = tslib_1.__importDefault(require("simple-git"));
const git = simple_git_1.default(process.cwd());
function main() {
    return tslib_1.__awaiter(this, void 0, void 0, function* () {
        const token = core_1.getInput("token", { required: true });
        const docsDir = core_1.getInput("out-dir");
        const outFile = core_1.getInput("out-file", { required: true });
        const content = core_1.getInput("content", { required: true });
        const username = github.context.actor || github.context.repo.owner;
        const repoName = github.context.repo.repo;
        const remoteOrigin = `https://${token}@github.com/${repoName}`;
        const outFilePath = docsDir + (/\.mdx?$/.test(outFile) ? outFile : `${outFile}.md`);
        core_1.debug('Entire output file path: ' + outFilePath);
        yield fs_extra_1.default.outputFile(outFilePath, content);
        if (username) {
            yield git.addConfig('user.email', `${username}@users.noreply.github.com`);
        }
        yield git
            .addConfig('user.name', username)
            .add('.')
            .commit('docs: update')
            .push(remoteOrigin);
    });
}
exports.main = main;
