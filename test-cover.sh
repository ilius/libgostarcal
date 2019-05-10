#!/bin/sh
set -e # Exit immediately if a command exits with a non-zero status
COVER_MODE=set # set or atomic

export RELEASE_STAGE=ci

OUT_FILE=$1
if [ -z $OUT_FILE ] ; then
	OUT_FILE="coverage.txt"
fi
TMP_FILE="/tmp/coverage-$USER.tmp"

echo "mode: $COVER_MODE" > "$OUT_FILE"
go list ./... | while read PKG ; do
	[ -f $GOPATH/src/$PKG/.notest ] && continue
	ls $GOPATH/src/$PKG/*_test.go >/dev/null 2>/dev/null || continue
	go test -tags test -v -covermode=$COVER_MODE -coverprofile=$TMP_FILE "$PKG"
	tail -n +2 $TMP_FILE >> "$OUT_FILE"
done

rm $TMP_FILE 2>/dev/null
