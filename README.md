[![codecov](https://codecov.io/gh/corani/silver-octo-sniffle/branch/master/graph/badge.svg?token=MZ7S8SGD7X)](https://codecov.io/gh/corani/silver-octo-sniffle)

# Silver Octo Sniffle

This started out as playing around a bit with generating LLVM IR in Go and has since devolved
into an attempt to implement a full [Oberon](http://oberon07.com/) compiler.

The name is just the random name that GitHub picked, if this results into something decent, I'll 
eventually rename the repo to something more memorable.

## Usage 

This project generates LLVM IR and subsequently compiles it with Clang. You'll need to have some
version of it installed. I'm using version 11 on Linux (X64) and version 15 on Android (ARM64) to
test it, presumably older versions will work as well as I'm not doing anything fancy. 

The CLI is not actively tested at the moment, I'm only running the golden tests (see below). But
the following should work:

```bash
$ ./build.sh -b

$ bin/compiler -src=test.obr
$ bin/test
```

## Testing 

To make testing easier, I'm using golden tests in the [test](test/) folder. They're saved as
Markdown files for easy viewing on GitHub and contain the source code with the expected tokens,
AST, IR and runtime output. The tests are executed together with the unit tests:

```bash
$ ./build.sh -t
```

To add a new test, simply copy one of the existing files and change the source code section. By
running:

```bash
$ ./build.sh -tr
```

you can record the output. Check the output in the new test file and make any necessary code
changes until it's correct, then move on. 

## CI/CD 

I'm running `golangci-lint` as a GitHub Action to statically check the code. This can be run
locally as well:

```bash
$ ./build.sh -l
```

There's also a GitHub Action that will run the tests and upload the test coverage to CodeCov.

Renovate is setup to keep the dependencies up-to-date.

Once this is in a usable state, I'll add a GoReleaser Action to create releases.
