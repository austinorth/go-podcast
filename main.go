package main

import (
	"fmt"
	"html/template"

	"go.uber.org/zap"
)

var (
	tpl          *template.Template
	globalLogger *zap.Logger
)

type tplData struct {
	NavBar string
}

type listTypeCheck struct {
	tplData
}

func main() {
	fmt.Println("vim-go")
}
