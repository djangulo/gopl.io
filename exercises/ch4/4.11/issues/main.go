// Package github provides a Go API for the GitHub issue tracker.
package main

import (
	"flag"
	"time"
	"path/filepath"
	"os"
	"os/exec"
)

var (
	owner, repo string
	flagSet = flag.NewFlagSet(os.Args[0], ExitOnError)
	Usage       func()
	basePath string
)

func init() {
	const (
		ownerUsage = "owner of the repo"
		repoUsage  = "repo to create the issue"
	)
	basePath = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "djangulo", "gopl.io", "exercises", "ch4", "4.11", "issuses")
	if err := os.Mkdir(basePath, os.ModeDir); os.IsNotExist(err) {
		os.MkdirAll(filepath.Join(basePath, "bin"), os.ModeDir)
		for _, word := range []string{"read", "update", "new"} {
			err := exec.Command("go", "build", "-o", filepath.Join(basePath, "bin", word), filepath.Join(basePath, word)).Run()
			if err != nil {
				panic(err)
			}
		}
	}
	flag.StringVar(&owner, "owner", "", ownerUsage)
	flag.StringVar(&owner, "o", "", ownerUsage+" (shorthand)")
	flag.StringVar(&owner, "repo", "", repoUsage)
	flag.StringVar(&owner, "r", "", repoUsage+" (shorthand)")

	usageStr = `Usage of %s
	issues ACTION [options]

	Where ACTION is one of
		N, new		create new issues
		R, read		read issues
		U, update	update existing issues

	run
		issues help ACTION
	to see detailed usage
	`

	Usage = func() {
		fmt.Fprintf(
			flagSet.Output(),
			usageStr,
			os.Args[0],
		)
	}
	flagSet.Usage = Usage
}


func main() {
	flagSet.Parse()
	action := flagSet.Arg(0)
	err := validateAction(action)
	if err != nil {
		panic(err)
	}


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
	action = strings.Lower(action)
	valid := map[string]struct{}{
		"new":    struct{}{},
		"N":      struct{}{},
		"read":   struct{}{},
		"R":      struct{}{},
		"update": struct{}{},
		"U":      struct{}{},
	}
	if _, ok valid[action]; !ok {
		return fmt.Errorf("'%s' is not a valid action", action)
	}
	return nil
}

