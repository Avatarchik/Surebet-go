#!/usr/bin/env bash

cd ..

set -e
packages=$(go list ./... | grep -v vendor)

echo "Static checks"

for d in $packages; do
    abs_path=$"$(go env GOPATH)/src/$d"
    commands=("gofmt -d -s" "go tool fix -diff")
    for c in "$commands"; do
        res=$(eval "$c $abs_path")
        if [ "$res" != "" ]; then
            printf "Code has to be reformatted:\n$res"
            exit 1
        fi
    done

    vet_flags=""
    if [[ $d == *"/surebetSearch/chrome" ]]; then
        vet_flags="-lostcancel=false"
    fi
    go vet $vet_flags $d

    staticcheck $d
done

echo "PASS"

echo "" > coverage.txt

for d in $packages; do
    go test -v -race -coverprofile=profile.out -covermode=atomic $d
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done