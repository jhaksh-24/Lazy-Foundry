package tui

import (
	"bytes"
	"fmt"
	"github.com/jhaksh-24/Lazy-Foundry/internal/forge"
)

func executeForge(command string, args ...string) (string, error) {
	var output bytes.Buffer
	var errOutput bytes.Buffer
	
	var err error
	
	switch command {
	case "build":
		err = forge.Build(args...)
	case "test":
		err = forge.Test(args...)
	case "init":
		err = forge.Init(args...)
	case "coverage":
		err = forge.Coverage(args...)
	case "create":
		if len(args) > 0 {
			err = forge.Create(args[0], args[1:]...)
		} else {
			err = fmt.Errorf("contract name required")
		}
	case "script":
		if len(args) > 0 {
			err = forge.Script(args[0], args[1:]...)
		} else {
			err = fmt.Errorf("script path required")
		}
	case "install":
		if len(args) > 0 {
			err = forge.Install(args[0], args[1:]...)
		} else {
			err = fmt.Errorf("package name required")
		}
	default:
		return "", fmt.Errorf("unknown forge command: %s", command)
	}
	
	if err != nil {
		return errOutput.String(), err
	}
	
	return output.String(), nil
}
