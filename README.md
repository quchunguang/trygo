trygo
=====

Simple samples of golang witch goals to covers full feature of golang.

Description
-----------

* testlang.go - Small demos about golang language specification.

* testlib.go - Small demos about golang standard packages.

* tools/simtool - Simulate some tools of unix/linux

* tools/eventframework - A simple event callback framework by golang.

* tools/cat - A demo impletements the unix-like `cat` command.

* standalone/... - Standalone demos.

Install
-------

```sh
sudo apt-get install libgl1-mesa-dev xorg-dev # dependence

go get github.com/quchunguang/trygo
go get github.com/quchunguang/trygo/standalone/...
```

Usage
-----

Demos can be run separatly by testing. For example, we want to run DemoDefine(), we can,

```sh
go test -test.run TestDemoDefine
```

For one of standalone demo named testXXX, run `testXXX`.
