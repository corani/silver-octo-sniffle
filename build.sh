#!/bin/bash 

prog=$(realpath "$0")
root=$(dirname "$prog")

function print_usage {
    echo "Usage: $prog [options ...]"
    echo 
    echo "Available options:"
    echo "  -h          prints this help"
    echo "  -b          build program"
    echo "  -i          install tools"
    echo "  -t          run unit tests"
    echo "  -tr         record golden tests"
    echo "  -l          run golangci-lint"
}

function do_echo {
    echo "[CMD] $@"
    TIMEFORMAT="[INFO] took %3lR"
    time "$@"
    code=$?
    if [ $code -ne 0 ]; then
        rc=$code
        echo "[ERROR] return code $rc"
    fi
}

function go_install_tools {
    do_echo tools/install.sh
}

function go_build {
    if [ -z "${X_DEBUG:-}" ]; then
        build="release"
        opts=(-trimpath -ldflags '-s -w')
    else
        build="debug"
        opts=(-gcflags 'all=-N -l')
    fi

    echo "[INFO] building ${build} build"
    mkdir -p bin 
    do_echo go build "${opts[@]}" -o bin/compiler ./src
}

function go_unit_test {
    echo "[INFO] running unit tests..."

    mkdir -p gen

    # NOTE(daniel): -race is not supported on android/aarch64
    if [ $(uname -m) == "aarch64" ]; then
        opts=()
    else
        opts=(-race)
    fi

    do_echo go test -v -timeout 1m "${opts[@]}" \
        -cover -coverpkg=./... -coverprofile=gen/cover.out \
        ./... | tee gen/test.out
    sed -n -e '/^FAIL.*/p' -e '/^--- FAIL/p' gen/test.out > gen/summary.out

    coverage=$(go tool cover -func gen/cover.out | tail -1 | cut -d ')' -f2 | tr -d "[:space:]")
    echo "[INFO] coverage: ${coverage} of statements"

    count=$(wc -l gen/summary.out | cut -d' ' -f1)
    if [ $count -eq 0 ]; then
        echo "[INFO] all passed"
    else
        echo "[ERROR] failed"
        cat gen/summary.out
    fi
}

function go_golden_test_record {
    echo "[INFO] recording golden tests..."

    do_echo go test -v ./src/... -update
}

function run_golangci {
    PATH="${root}/tools/bin":$PATH

    command -v golangci-lint > /dev/null 2>&1
    if [ $? -eq 1 ]; then
        go_install_tools
    fi

    do_echo golangci-lint run               \
        --out-format colored-line-number    \
        --issues-exit-code 1                \
        ./...
}

if [ "$#" -eq "0" ]; then
    print_usage
    exit 0
fi

do_echo cd "$root"

while [ "$#" -gt "0" ]; do
  arg=$1
  shift

  case $arg in
    -h)
        print_usage
        exit 0
        ;;
    -b)
        go_build
        ;;
    -i)
        go_install_tools
        ;;
    -t)
        go_build
        go_unit_test
        ;;
    -tr)
        go_build
        go_golden_test_record
        ;;
    -l)
        run_golangci
        ;;
    *)
        echo "[ERROR] unrecognized argument '$arg'"
        print_usage
        exit 1
        ;;
  esac
done
