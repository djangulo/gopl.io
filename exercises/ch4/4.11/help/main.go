//Package help displays help messages for all the actions
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	readPkg   = "read"
	newPkg    = "new"
	updatePkg = "update"
)

var (
	validActions = map[string]struct{}{
		"read":   struct{}{},
		"update": struct{}{},
		"new":    struct{}{},
	}
)

func main() {
	args := os.Args
	action := args[len(args)-1]
	err := validateAction(action)
	if err != nil {
		log.Fatal(err)
	}

	var stdout, stderr bytes.Buffer

	// "github.com/djangulo/gopl.io/exercises/ch4/4.11/common"
	basePath := filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "djangulo", "gopl.io", "exercises", "ch4", "4.11", "bin")
	switch action {
	case readPkg:
		cmd := exec.Command(filepath.Join(basePath, "read"), "-h")
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		cmd.Run()
	case updatePkg:
		cmd := exec.Command(filepath.Join(basePath, "update"), "-h")
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		cmd.Run()
	case newPkg:
		cmd := exec.Command(filepath.Join(basePath, "new"), "-h")
		cmd.Env = os.Environ()
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		cmd.Run()
		// new.Usage()
	}
	fmt.Println(string(stdout.Bytes()))
	fmt.Println(string(stderr.Bytes()))

}

func validateAction(action string) error {
	if _, ok := validActions[action]; !ok {
		return fmt.Errorf("invalid action %s", action)
	}
	return nil
}
