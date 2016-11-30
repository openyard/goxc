paragon
=======

An attempt to create, develop, manage and maintain software-products in the 21st century.

- product lifecycle
- trunk based development
- mono repository
- continous integration, continous deployment
- release often
- short feedback loops

---

## git commit style

- separate subject from body with a blank line
- limit the subject line to 50 characters
- capitalize the subject line
- do not end the subject line with a period
- use the imperative mood in the subject line
- wrap the body at 72 characters
- use the body to explain what and why vs. how

further information: http://chris.beams.io/posts/git-commit/

### structure

```
<type>(<scope>): <subject>
<BLANK LINE>
<body>
<BLANK LINE>
<footer>
```

### types

- feat (feature)
- fix (bugfix)
- docs (documentation)
- test (unit tests)
- style (formatting, typos etc.)
- refactor (refactoring)
- chore (general

### scopes

- module name
- client/customer name
- general

---

## git workflow

see also:

- http://nvie.com/posts/a-successful-git-branching-model/
- http://www.atlassian.com/git/tutorials/comparing-workflows/feature-branch-workflow
- http://www.atlassian.com/git/tutorials/comparing-workflows/gitflow-workflow

### branch naming convention

- main branches: master, develop
- feature branches: feature-*
- release branches: release-*
- hotfix branches: hotfix-*

### develop

- create new feature branch

```
$git checkout -b feature-xyz develop
$git push origin feature-xyz
```

- finish feature branch

```
$ git checkout develop
Switched to branch 'develop'
$ git merge --no-ff feature-xyz
Updating ea1b82a..05e9557
(Summary of changes)
$ git branch -d feature-xyz
Deleted branch feature-xyz (was 05e9557).
$ git push origin develop
```

> The --no-ff flag causes the merge to always create a new commit object, even if the merge could be performed with a fast-forward. This avoids losing information about the historical existence of a feature branch and groups together all commits that together added the feature.
> Yes, it will create a few more (empty) commit objects, but the gain is much bigger than the cost.

### release

- create new release branch

```
$ git checkout -b release-1.2 develop
Switched to a new branch "release-1.2"
$ ./bump-version.sh 1.2
Files modified successfully, version bumped to 1.2.
$ git commit -a -m "Bumped version number to 1.2"
[release-1.2 74d9424] Bumped version number to 1.2
1 files changed, 1 insertions(+), 1 deletions(-)
```

- finish release branch

```
$ git checkout master
Switched to branch 'master'
$ git merge --no-ff release-1.2
Merge made by recursive.
(Summary of changes)
$ git tag -a 1.2
```

```
$ git checkout develop
Switched to branch 'develop'
$ git merge --no-ff release-1.2
Merge made by recursive.
(Summary of changes)
```

```
$ git branch -d release-1.2
Deleted branch release-1.2 (was ff452fe).
```

### hotfix

- create new hotfix branch

```
$ git checkout -b hotfix-1.2.1 master
Switched to a new branch "hotfix-1.2.1"
$ ./bump-version.sh 1.2.1
Files modified successfully, version bumped to 1.2.1.
$ git commit -a -m "Bumped version number to 1.2.1"
[hotfix-1.2.1 41e61bb] Bumped version number to 1.2.1
1 files changed, 1 insertions(+), 1 deletions(-)
```

```
$ git commit -m "Fixed severe production problem"
[hotfix-1.2.1 abbe5d6] Fixed severe production problem
5 files changed, 32 insertions(+), 17 deletions(-)
```

- finsih hotfix branch

```
$ git checkout master
Switched to branch 'master'
$ git merge --no-ff hotfix-1.2.1
Merge made by recursive.
(Summary of changes)
$ git tag -a 1.2.1
```

```
$ git checkout develop
Switched to branch 'develop'
$ git merge --no-ff hotfix-1.2.1
Merge made by recursive.
(Summary of changes)
```

> When a release branch currently exists, the hotfix changes need to be merged into that release branch, instead of develop

```
$ git branch -d hotfix-1.2.1
Deleted branch hotfix-1.2.1 (was abbe5d6).
```
