#!/bin/bash

if [ -f docs/docs.go ]; then
    cp docs/docs.go{,.old}
fi

swag init

if [ -f docs/docs.go.old ]; then
    if [ $(cat docs/docs.go|grep -v '^//'|md5) == $(cat docs/docs.go.old|grep -v '^//'|md5) ]; then
        mv docs/docs.go{.old,}
    else
        rm docs/docs.go.old
    fi
fi
