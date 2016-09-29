#!/bin/bash
# Create testwalk_windows.go
# Create Manifest test.manifest

# go get github.com/akavel/rsrc
rsrc -manifest test.manifest -o rsrc.syso

go build -ldflags="-H windowsgui"
