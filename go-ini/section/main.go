package main

import (
	"fmt"

	"gopkg.in/ini.v1"
)

func SectionInfo() {
	cfg, err := ini.Load("my.ini")
	if err != nil {
		fmt.Println("Fail to read file: ", err)
		return
	}

	sections := cfg.Sections()
	names := cfg.SectionStrings()

	fmt.Println("sections: ", sections)
	fmt.Println("names: ", names)
}

func NewSection() {
	cfg, err := ini.Load("my.ini")
	if err != nil {
		fmt.Println("Fail to read file: ", err)
		return
	}

	newSection := cfg.Section("new")

	fmt.Println("new section: ", newSection)
	fmt.Println("names: ", cfg.SectionStrings())
}

func ParentChildSection() {
	cfg, err := ini.Load("parent_child.ini")
	if err != nil {
		fmt.Println("Fail to read file: ", err)
		return
	}

	fmt.Println("Clone url from package.sub:", cfg.Section("package.sub").Key("CLONE_URL").String())
}

func main() {
	// SectionInfo()

	// NewSection()

	ParentChildSection()
}