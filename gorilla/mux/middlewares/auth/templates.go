package main

import (
	"html/template"
	"log"
)

var (
	ptTemplate *template.Template
)

func init() {
	var err error
	ptTemplate, err = template.New("").ParseGlob("./tpls/*.tpl")
	if err != nil {
		log.Fatalf("load templates failed:%v", err)
	}
}
