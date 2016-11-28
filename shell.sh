#!/bin/sh

gen() {
    ./consul-template \
        -once \
        -consul localhost:8500 \
        -template "sherlock.ctmpl:sherlock-services.yaml"
}


until
    cmd=$1
    if [ -z "$cmd" ]; then
        onChange
    fi
    shift 1
    $cmd "$@"
    [ "$?" -ne 127 ]
do
    onChange
    exit
done
