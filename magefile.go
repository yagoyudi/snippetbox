//go:build mage

package main

import (
	"github.com/carolynvs/magex/pkg"
	"github.com/magefile/mage/sh"
)

var Default = Run

func Build() error {
	return sh.RunV("go", "build", "-o", "bin/snippetbox", "./cmd/snippetbox")
}

func Run() error {
	return sh.RunV("go", "run", "./cmd/snippetbox")
}

func Test() error {
	return sh.RunV("go", "test", "-v", "./...")
}

func Clean() error {
	return sh.RunV("rm", "-rf", "bin")
}

func EnsureMage() error {
	return pkg.EnsureMage("")
}
