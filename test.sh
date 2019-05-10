#!/bin/sh
set -e # Exit immediately if a command exits with a non-zero status

export RELEASE_STAGE=ci

go list ./... | while read PKG ; do
	[ -f $GOPATH/src/$PKG/.notest ] && continue
	ls $GOPATH/src/$PKG/*_test.go >/dev/null 2>/dev/null || continue
	go test -tags test -v "$PKG"
done
