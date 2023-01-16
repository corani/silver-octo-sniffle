#!/bin/bash
dir=$(dirname $(realpath "$0"))
export GOBIN="${dir}/bin"
cd "${dir}"

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

echo "[INFO] found tools:"
tools=$(cat ${dir}/tools.go | grep "_" | cut -d'"' -f2)
for tool in ${tools}; do
    echo "[INFO] - ${tool}"
done

do_echo go mod tidy 
do_echo go install ${tools}
