package main

import (
	_ "embed"
	"html/template"
)

//go:embed index.html
var pageTemplate string

func renderTemplate() *template.Template {
	funcMap := template.FuncMap{
	    "percent": func(x int, y int) int {
	        return (x / y) * 100
	    },
	}

	t := template.Must(
		template.New("page").Funcs(funcMap).Parse(pageTemplate),
	)

	return t
}
