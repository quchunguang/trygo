package main

var genInfo = `
This source file was generated from template by "runtemplate"".
The command line is in the comment of ""testruntemplate.go" as "go generate"" command.
I'm {{.Author}}, and the version of the demo is {{.Version}}.

Following is always useful:
TemplateFile = {{.TemplateFile}}
OutFile      = {{.OutFile}}`
