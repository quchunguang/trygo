trygo
=====
[![Build Status](https://travis-ci.org/quchunguang/trygo.svg?branch=master)](https://travis-ci.org/quchunguang/trygo)
[![Go Report Card](https://goreportcard.com/badge/github.com/quchunguang/trygo)](https://goreportcard.com/report/github.com/quchunguang/trygo)
[![GoDoc](https://godoc.org/github.com/quchunguang/trygo?status.svg)](https://godoc.org/github.com/quchunguang/trygo)

Simple samples of golang witch goals to covers full feature of golang.

Description
-----------

* testlang.go - Small demos about golang language specification.

* testlib.go - Small demos about golang standard packages.

* simtool - Simulate some tools of unix/linux

* eventframework - A simple event callback framework by golang.

* cat - A demo impletements the unix-like `cat` command.

* demos/.../ - Standalone demos.

Install
-------

```sh
sudo apt-get install libgl1-mesa-dev xorg-dev # dependence

go get github.com/quchunguang/trygo
go get github.com/quchunguang/trygo/demos/...
```

Usage
-----

Demos can be run separatly by testing. For example, we want to run DemoDefine(), we can,

```sh
go test -test.run TestDemoDefine
```

For one of standalone demo named testXXX, run `testXXX`.
