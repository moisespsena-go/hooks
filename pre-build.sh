#!/usr/bin/env bash
this="$0"
find "${this%.sh}/" -name "*.sh" | sort | while read l
do
    echo "PRE-BUILD RUN '$l'"
    sh "$l" || exit $?
done