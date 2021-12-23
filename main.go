package main

import (
	"text/template"
)

var packageTemplate = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.

package version

const SDKVersion = "{{ .Version}}"
`))

func main() {
	// ver := flag.String("version", "", "version")
	// flag.Parse()
	// fmt.Println(*ver)
	// fmt.Println(version.SDKVersion)

	// packageTemplate.Execute(f, struct {
	// 	Version string
	// }{
	// 	Version: ver,
	// })
}
