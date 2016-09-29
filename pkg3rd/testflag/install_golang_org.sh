#!/usr/bin/env bash

# All project named with [mirror] in github.com/golang/*
# for i in tour tools text sys review playground net mobile image exp debug crypto build arch; do
for i in tour tools text sys review playground net mobile image exp debug crypto build arch; do
	echo $i
	git clone https://github.com/golang/$i $GOPATH/src/golang.org/x/$i
done
