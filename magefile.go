//go:build mage

package main

import (
	"os"
	"os/exec"
	"path/filepath"

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

func GenCert() error {
	goPath, err := exec.LookPath("go")
	if err != nil {
		return err
	}

	goDir := filepath.Dir(filepath.Dir(goPath))
	scriptPath := filepath.Join(goDir, "share/go/src/crypto/tls/generate_cert.go")

	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		return err
	}

	tlsDir := "tls"
	if err := os.MkdirAll(tlsDir, 0755); err != nil {
		return err
	}

	cmd := exec.Command("go", "run", scriptPath, "--rsa-bits=2048", "--host=localhost")
	cmd.Dir = tlsDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
