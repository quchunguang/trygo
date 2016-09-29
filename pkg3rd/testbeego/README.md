Create Reference
================

To create beego project,

```sh
go get github.com/astaxie/beego # install library
go get github.com/beego/bee     # install command `bee`

bee new testbeego # create project file tree from template

cd testbeego
bee run watchall # develop run

go build         # build
./testbeego      # run
```
