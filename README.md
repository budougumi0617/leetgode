leetgode
====

[![PkgGoDev](https://pkg.go.dev/badge/budougumi0617/leetgode)][godev]
![test](https://github.com/budougumi0617/leetgode/workflows/test/badge.svg?branch=master)

[godev]:https://pkg.go.dev/github.com/budougumi0617/regexponce

## Description
LeetCode CLI for Gophers.

## VS.


## Requirement
The leetgode CLI needs the authorization to execute some sub commands. Specifically, it needs `LEETCODE_SESSION`, and `csrftoken`.

1. Open chrome and paste the link below to the chrome linkbar.
    - `chrome://settings/cookies/detail?site=leetcode.com`
1. Copy the contents of `LEETCODE_SESSION`, and `csrftoken`.
1. Export below environment values by the use of `LEETCODE_SESSION`, and `csrftoken`.
```bash
export LEETCODE_SESSION=${LEETCODE_SESSION}
export LEETCODE_TOKEN=${csrftoken}
```


## Usage

```
leetgode help
Usage: leetgode is leetcode cli for gophers.

SubCommands:
 exec     Submit solution
 generate Generate the skeleton code with the test file by id
 help     Help shows usages
 list     List problems
 pick     Pick a problem by id
 test     Test solution
```

## Install
You can download binary from [release page](https://github.com/budougumi0617/leetgode/releases) and place it in $PATH directory.

### go get
```
go get -u https://github.com/budougumi0617/leetgode/cmd/leetgode
```

### macOS
If you want to install on macOS, you can use Homebrew.
```
brew install budougumi0617/tap/leetgode
```

## Contribution
1. Fork ([https://github.com/budougumi0617/leetgode/fork](https://github.com/budougumi0617/leetgode/fork))
2. Create a feature branch
3. Commit your changes
4. Rebase your local changes against the master branch
5. Run test suite with the `go test ./...` command and confirm that it passes
6. Run `gofmt -s`
7. Create new Pull Request

## License

[MIT](https://github.com/budougumi0617/leetgode/blob/master/LICENSE)

## Author

[Yoichiro Shimizu(@budougumi0617)](https://github.com/budougumi0617)
